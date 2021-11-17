package dummy

import (
	"fmt"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

type RoutingService struct {
	voyageRepository domain.VoyageRepository
}

var _ domain.RoutingService = &RoutingService{}

func NewRoutingService(voyageRepository domain.VoyageRepository) *RoutingService {
	return &RoutingService{
		voyageRepository: voyageRepository,
	}
}

func (s RoutingService) FetchRoutesForSpecification(routeSpecification *domain.RouteSpecification) ([]domain.Itinerary, error) {
	voyages, err := s.voyageRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("find all voyages: %s", err)
	}

	var itineraries []domain.Itinerary
	for _, voyage := range voyages {
		itinerary := voyageItinerary(&voyage)
		if routeSpecification.IsSatisfiedBy(itinerary) {
			itineraries = append(itineraries, *itinerary)
		}
	}
	return itineraries, nil
}

func voyageItinerary(voyage *domain.Voyage) *domain.Itinerary {
	var legs []domain.Leg
	for _, movement := range voyage.Schedule.CarrierMovements {
		legs = append(legs, domain.Leg{
			VoyageNumber:   voyage.VoyageNumber,
			LoadUNLocode:   movement.DepartureUNLocode,
			UnloadUNLocode: movement.ArrivalUNLocode,
			LoadTime:       movement.DepartureTime,
			UnloadTime:     movement.ArrivalTime,
		})
	}
	return domain.NewItinerary(legs)
}
