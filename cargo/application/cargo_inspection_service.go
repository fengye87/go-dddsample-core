package application

import (
	"fmt"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

//go:generate mockgen -destination=mock/cargo_inspection_service.go -package=mock . CargoInspectionService

type CargoInspectionService interface {
	InspectCargo(trackingID domain.TrackingID) error
}

type CargoInspectionServiceImpl struct {
	eventListener           EventListener
	cargoRepository         domain.CargoRepository
	handlingEventRepository domain.HandlingEventRepository
}

var _ CargoInspectionService = &CargoInspectionServiceImpl{}

func NewCargoInspectionService(eventListener EventListener, cargoRepository domain.CargoRepository, handlingEventRepository domain.HandlingEventRepository) *CargoInspectionServiceImpl {
	return &CargoInspectionServiceImpl{
		eventListener:           eventListener,
		cargoRepository:         cargoRepository,
		handlingEventRepository: handlingEventRepository,
	}
}

func (s CargoInspectionServiceImpl) InspectCargo(trackingID domain.TrackingID) error {
	cargo, err := s.cargoRepository.Find(trackingID)
	if err != nil {
		return fmt.Errorf("find cargo: %s", err)
	}

	handlingHistory, err := s.handlingEventRepository.LookupHandlingHistoryOfCargo(trackingID)
	if err != nil {
		return fmt.Errorf("lookup handling history of cargo: %s", err)
	}

	cargo.DeriveDeliveryProgress(handlingHistory)

	if cargo.Delivery.Misdirected {
		s.eventListener.CargoWasMisdirected(cargo)
	}

	if cargo.Delivery.IsUnloadedAtDestination {
		s.eventListener.CargoHasArrived(cargo)
	}

	if err := s.cargoRepository.Store(cargo); err != nil {
		return fmt.Errorf("store cargo: %s", err)
	}
	return nil
}
