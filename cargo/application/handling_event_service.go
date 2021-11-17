package application

import (
	"fmt"
	"time"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

//go:generate mockgen -destination=mock/handling_event_service.go -package=mock . HandlingEventService

type HandlingEventService interface {
	RegisterHandlingEvent(completionTime time.Time, trackingID domain.TrackingID, voyageNumber domain.VoyageNumber, unLocode domain.UNLocode, typ domain.HandlingEventType) error
}

type HandlingEventServiceImpl struct {
	eventListener           EventListener
	handlingEventRepository domain.HandlingEventRepository
	handlingEventFactory    domain.HandlingEventFactory
}

var _ HandlingEventService = &HandlingEventServiceImpl{}

func NewHandlingEventService(eventListener EventListener, handlingEventRepository domain.HandlingEventRepository, handlingEventFactory domain.HandlingEventFactory) *HandlingEventServiceImpl {
	return &HandlingEventServiceImpl{
		eventListener:           eventListener,
		handlingEventRepository: handlingEventRepository,
		handlingEventFactory:    handlingEventFactory,
	}
}

func (s HandlingEventServiceImpl) RegisterHandlingEvent(completionTime time.Time, trackingID domain.TrackingID, voyageNumber domain.VoyageNumber, unLocode domain.UNLocode, tp domain.HandlingEventType) error {
	registrationTime := time.Now()
	event, err := s.handlingEventFactory.CreateHandlingEvent(registrationTime, completionTime, trackingID, voyageNumber, unLocode, tp)
	if err != nil {
		return fmt.Errorf("create handling event: %s", err)
	}

	if err := s.handlingEventRepository.Store(event); err != nil {
		return fmt.Errorf("store handling event: %s", err)
	}

	s.eventListener.CargoWasHandled(event)
	return nil
}
