package booking

import (
	"fmt"
	"time"

	"github.com/fengye87/dddsample-core/cargo/application"
	"github.com/fengye87/dddsample-core/cargo/domain"
)

type BookingServiceFacadeImpl struct {
	bookingService     application.BookingService
	locationRepository domain.LocationRepository
	cargoRepository    domain.CargoRepository
	voyageRepository   domain.VoyageRepository
}

var _ BookingServiceFacade = &BookingServiceFacadeImpl{}

func NewBookingServiceFacade(bookingService application.BookingService, locationRepository domain.LocationRepository, cargoRepository domain.CargoRepository, voyageRepository domain.VoyageRepository) *BookingServiceFacadeImpl {
	return &BookingServiceFacadeImpl{
		bookingService:     bookingService,
		locationRepository: locationRepository,
		cargoRepository:    cargoRepository,
		voyageRepository:   voyageRepository,
	}
}

func (f BookingServiceFacadeImpl) ListShippingLocations() ([]LocationDTO, error) {
	locations, err := f.locationRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("find all locations: %s", err)
	}
	return NewLocationDTOAssembler().ToDTOs(locations), nil
}

func (f BookingServiceFacadeImpl) BookNewCargo(origin string, destination string, arrivalDeadline time.Time) (string, error) {
	trackingID, err := f.bookingService.BookNewCargo(domain.UNLocode(origin), domain.UNLocode(destination), arrivalDeadline)
	if err != nil {
		return "", fmt.Errorf("book new cargo: %s", err)
	}
	return string(trackingID), nil
}

func (f BookingServiceFacadeImpl) LoadCargoForRouting(trackingID string) (*CargoRoutingDTO, error) {
	cargo, err := f.cargoRepository.Find(domain.TrackingID(trackingID))
	if err != nil {
		return nil, fmt.Errorf("find cargo: %s", err)
	}
	return NewCargoRoutingDTOAssembler().ToDTO(cargo), nil
}

func (f BookingServiceFacadeImpl) AssignCargoToRoute(trackingID string, route *RouteCandidateDTO) error {
	itinerary, err := NewRouteCandidateDTOAssembler().FromDTO(route, f.voyageRepository, f.locationRepository)
	if err != nil {
		return err
	}
	return f.bookingService.AssignCargoToRoute(itinerary, domain.TrackingID(trackingID))
}

func (f BookingServiceFacadeImpl) ChangeDestination(trackingID string, destinationUNLocode string) error {
	return f.bookingService.ChangeDestination(domain.TrackingID(trackingID), domain.UNLocode(destinationUNLocode))
}

func (f BookingServiceFacadeImpl) ListAllCargos() ([]CargoRoutingDTO, error) {
	cargos, err := f.cargoRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("find all cargos: %s", err)
	}
	return NewCargoRoutingDTOAssembler().ToDTOs(cargos), nil
}

func (f BookingServiceFacadeImpl) RequestPossibleRoutesForCargo(trackingID string) ([]RouteCandidateDTO, error) {
	itineraries, err := f.bookingService.RequestPossibleRoutesForCargo(domain.TrackingID(trackingID))
	if err != nil {
		return nil, fmt.Errorf("request possible routes for cargo: %s", err)
	}
	return NewRouteCandidateDTOAssembler().ToDTOs(itineraries), nil
}
