package booking

import (
	"fmt"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

type CargoRoutingDTOAssembler struct {
}

func NewCargoRoutingDTOAssembler() *CargoRoutingDTOAssembler {
	return &CargoRoutingDTOAssembler{}
}

func (a CargoRoutingDTOAssembler) ToDTO(cargo *domain.Cargo) *CargoRoutingDTO {
	return &CargoRoutingDTO{
		TrackingID:               string(cargo.TrackingID),
		OriginUNLocode:           string(cargo.OriginUNLocode),
		FinalDestinationUNLocode: string(cargo.RouteSpecification.DestinationUNLocode),
		ArrivalDeadline:          cargo.RouteSpecification.ArrivalDeadline,
		Misrouted:                cargo.Delivery.Misdirected,
		Legs:                     NewLegDTOAssembler().ToDTOs(cargo.Itinerary.Legs),
	}
}

func (a CargoRoutingDTOAssembler) ToDTOs(cargos []domain.Cargo) []CargoRoutingDTO {
	var dtos []CargoRoutingDTO
	for _, cargo := range cargos {
		dtos = append(dtos, *a.ToDTO(&cargo))
	}
	return dtos
}

type LegDTOAssembler struct {
}

func NewLegDTOAssembler() *LegDTOAssembler {
	return &LegDTOAssembler{}
}

func (a LegDTOAssembler) ToDTO(leg *domain.Leg) *LegDTO {
	return &LegDTO{
		VoyageNubmer: string(leg.VoyageNumber),
		FromUNLocode: string(leg.LoadUNLocode),
		ToUNLocode:   string(leg.UnloadUNLocode),
		LoadTime:     leg.LoadTime,
		UnloadTime:   leg.UnloadTime,
	}
}

func (a LegDTOAssembler) ToDTOs(legs []domain.Leg) []LegDTO {
	var dtos []LegDTO
	for _, leg := range legs {
		dtos = append(dtos, *a.ToDTO(&leg))
	}
	return dtos
}

type RouteCandidateDTOAssembler struct {
}

func NewRouteCandidateDTOAssembler() *RouteCandidateDTOAssembler {
	return &RouteCandidateDTOAssembler{}
}

func (a RouteCandidateDTOAssembler) ToDTO(itinerary *domain.Itinerary) *RouteCandidateDTO {
	return &RouteCandidateDTO{
		Legs: NewLegDTOAssembler().ToDTOs(itinerary.Legs),
	}
}

func (a RouteCandidateDTOAssembler) ToDTOs(itineraries []domain.Itinerary) []RouteCandidateDTO {
	var dtos []RouteCandidateDTO
	for _, itinerary := range itineraries {
		dtos = append(dtos, *a.ToDTO(&itinerary))
	}
	return dtos
}

func (a RouteCandidateDTOAssembler) FromDTO(routeCandidateDTO *RouteCandidateDTO, voyageRepository domain.VoyageRepository, locationRepository domain.LocationRepository) (*domain.Itinerary, error) {
	var legs []domain.Leg
	for _, dto := range routeCandidateDTO.Legs {
		voyage, err := voyageRepository.Find(domain.VoyageNumber(dto.VoyageNubmer))
		if err != nil {
			return nil, fmt.Errorf("find voyage: %s", err)
		}

		from, err := locationRepository.Find(domain.UNLocode(dto.FromUNLocode))
		if err != nil {
			return nil, fmt.Errorf("find from location: %s", err)
		}

		to, err := locationRepository.Find(domain.UNLocode(dto.ToUNLocode))
		if err != nil {
			return nil, fmt.Errorf("find to location: %s", err)
		}

		legs = append(legs, *domain.NewLeg(voyage.VoyageNumber, from.UNLocode, to.UNLocode, dto.LoadTime, dto.UnloadTime))
	}
	return domain.NewItinerary(legs), nil
}

type LocationDTOAssembler struct {
}

func NewLocationDTOAssembler() *LocationDTOAssembler {
	return &LocationDTOAssembler{}
}

func (a LocationDTOAssembler) ToDTO(location *domain.Location) *LocationDTO {
	return &LocationDTO{
		UNLocode: string(location.UNLocode),
		Name:     location.Name,
	}
}

func (a LocationDTOAssembler) ToDTOs(locations []domain.Location) []LocationDTO {
	var dtos []LocationDTO
	for _, location := range locations {
		dtos = append(dtos, *a.ToDTO(&location))
	}
	return dtos
}
