package application

import (
	"fmt"
	"time"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

//go:generate mockgen -destination=mock/booking_service.go -package=mock . BookingService

type BookingService interface {
	BookNewCargo(originUNLocode domain.UNLocode, destinationUNLocode domain.UNLocode, arrivalDeadline time.Time) (domain.TrackingID, error)
	RequestPossibleRoutesForCargo(trackingID domain.TrackingID) ([]domain.Itinerary, error)
	AssignCargoToRoute(itinerary *domain.Itinerary, trackingID domain.TrackingID) error
	ChangeDestination(trackingID domain.TrackingID, newDestinationUNLocode domain.UNLocode) error
}

type BookingServiceImpl struct {
	cargoRepository    domain.CargoRepository
	locationRepository domain.LocationRepository
	routingService     domain.RoutingService
}

var _ BookingService = &BookingServiceImpl{}

func NewBookingService(cargoRepository domain.CargoRepository, locationRepository domain.LocationRepository, routingService domain.RoutingService) *BookingServiceImpl {
	return &BookingServiceImpl{
		cargoRepository:    cargoRepository,
		locationRepository: locationRepository,
		routingService:     routingService,
	}
}

func (s BookingServiceImpl) BookNewCargo(originUNLocode domain.UNLocode, destinationUNLocode domain.UNLocode, arrivalDeadline time.Time) (domain.TrackingID, error) {
	trackingID, err := s.cargoRepository.NextTrackingID()
	if err != nil {
		return "", fmt.Errorf("find cargo: %s", err)
	}

	origin, err := s.locationRepository.Find(originUNLocode)
	if err != nil {
		return "", fmt.Errorf("find origin: %s", err)
	}

	destination, err := s.locationRepository.Find(destinationUNLocode)
	if err != nil {
		return "", fmt.Errorf("find destination: %s", err)
	}

	routeSpecification := domain.NewRouteSpecification(origin.UNLocode, destination.UNLocode, arrivalDeadline)
	cargo := domain.NewCargo(trackingID, routeSpecification)
	if err := s.cargoRepository.Store(cargo); err != nil {
		return "", fmt.Errorf("store cargo: %s", err)
	}
	return trackingID, nil
}

func (s BookingServiceImpl) RequestPossibleRoutesForCargo(trackingID domain.TrackingID) ([]domain.Itinerary, error) {
	cargo, err := s.cargoRepository.Find(trackingID)
	if err != nil {
		return nil, fmt.Errorf("find cargo: %s", err)
	}
	return s.routingService.FetchRoutesForSpecification(cargo.RouteSpecification)
}

func (s BookingServiceImpl) AssignCargoToRoute(itinerary *domain.Itinerary, trackingID domain.TrackingID) error {
	cargo, err := s.cargoRepository.Find(trackingID)
	if err != nil {
		return fmt.Errorf("find cargo: %s", err)
	}

	cargo.AssignToRoute(itinerary)
	if err := s.cargoRepository.Store(cargo); err != nil {
		return fmt.Errorf("store cargo: %s", err)
	}
	return nil
}

func (s BookingServiceImpl) ChangeDestination(trackingID domain.TrackingID, newDestinationUNLocode domain.UNLocode) error {
	cargo, err := s.cargoRepository.Find(trackingID)
	if err != nil {
		return fmt.Errorf("find cargo: %s", err)
	}

	newDestination, err := s.locationRepository.Find(newDestinationUNLocode)
	if err != nil {
		return fmt.Errorf("find new destination: %s", err)
	}

	routeSpecification := domain.NewRouteSpecification(cargo.OriginUNLocode, newDestination.UNLocode, cargo.RouteSpecification.ArrivalDeadline)
	cargo.SpecifyNewRoute(routeSpecification)
	if err := s.cargoRepository.Store(cargo); err != nil {
		return fmt.Errorf("store cargo: %s", err)
	}
	return nil
}
