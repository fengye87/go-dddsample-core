package domain

import (
	"fmt"
	"reflect"
	"time"
)

type HandlingEvent struct {
	Type             HandlingEventType
	VoyageNumber     VoyageNumber
	UNLocode         UNLocode
	CompletionTime   time.Time
	RegistrationTime time.Time
	TrackingID       TrackingID
}

func NewHandlingEvent(trackingID TrackingID, completionTime time.Time, registrationTime time.Time, tp HandlingEventType, unLocode UNLocode, voyageNumber VoyageNumber) (*HandlingEvent, error) {
	if tp.RequiresVoyage() && voyageNumber == VoyageNumber("") {
		return nil, fmt.Errorf("voyage number is required for this type of event")
	}
	if tp.ProhibitsVoyage() && voyageNumber != VoyageNumber("") {
		return nil, fmt.Errorf("voyage number is prohibited for this type of event")
	}
	return &HandlingEvent{
		VoyageNumber:     voyageNumber,
		CompletionTime:   completionTime,
		RegistrationTime: registrationTime,
		Type:             tp,
		UNLocode:         unLocode,
		TrackingID:       trackingID,
	}, nil
}

func (e HandlingEvent) SameEventAs(other *HandlingEvent) bool {
	return reflect.DeepEqual(other, &e)
}

type HandlingEventType int

const (
	HandlingEventTypeLoad HandlingEventType = iota
	HandlingEventTypeUnload
	HandlingEventTypeReceive
	HandlingEventTypeClaim
	HandlingEventTypeCustoms
)

func (t HandlingEventType) RequiresVoyage() bool {
	switch t {
	case HandlingEventTypeLoad, HandlingEventTypeUnload:
		return true
	case HandlingEventTypeReceive, HandlingEventTypeClaim, HandlingEventTypeCustoms:
		return false
	default:
		panic(t)
	}
}

func (t HandlingEventType) ProhibitsVoyage() bool {
	return !t.RequiresVoyage()
}

func (t HandlingEventType) SameValueAs(other HandlingEventType) bool {
	return other == t
}

type HandlingHistory struct {
	HandlingEvents []HandlingEvent
}

func NewHandlingHistory(handlingEvents []HandlingEvent) *HandlingHistory {
	return &HandlingHistory{
		HandlingEvents: handlingEvents,
	}
}

func (h HandlingHistory) DistinctEventsByCompletionTime() []HandlingEvent {
	// TODO: sort
	return h.HandlingEvents
}

func (h HandlingHistory) MostRecentlyCompletedEvent() *HandlingEvent {
	distinctEvents := h.DistinctEventsByCompletionTime()
	if len(distinctEvents) == 0 {
		return nil
	}
	return &distinctEvents[len(distinctEvents)-1]
}

func (h HandlingHistory) SameValueAs(other *HandlingHistory) bool {
	return reflect.DeepEqual(other, &h)
}

type HandlingEventFactory struct {
	cargoRepository    CargoRepository
	voyageRepository   VoyageRepository
	locationRepository LocationRepository
}

func NewHandlingEventFactory(cargoRepository CargoRepository, voyageRepository VoyageRepository, locationRepository LocationRepository) *HandlingEventFactory {
	return &HandlingEventFactory{
		cargoRepository:    cargoRepository,
		voyageRepository:   voyageRepository,
		locationRepository: locationRepository,
	}
}

func (f HandlingEventFactory) CreateHandlingEvent(registrationTime time.Time, completionTime time.Time, trackingID TrackingID, voyageNumber VoyageNumber, unLocode UNLocode, typ HandlingEventType) (*HandlingEvent, error) {
	cargo, err := f.findCargo(trackingID)
	if err != nil {
		return nil, fmt.Errorf("find cargo: %s", err)
	}

	voyage, err := f.findVoyage(voyageNumber)
	if err != nil {
		return nil, fmt.Errorf("find voyage: %s", err)
	}

	location, err := f.findLocation(unLocode)
	if err != nil {
		return nil, fmt.Errorf("find location: %s", err)
	}

	return NewHandlingEvent(cargo.TrackingID, completionTime, registrationTime, typ, location.UNLocode, voyage.VoyageNumber)
}

func (f HandlingEventFactory) findCargo(trackingID TrackingID) (*Cargo, error) {
	return f.cargoRepository.Find(trackingID)
}

func (f HandlingEventFactory) findVoyage(voyageNumber VoyageNumber) (*Voyage, error) {
	return f.voyageRepository.Find(voyageNumber)
}

func (f HandlingEventFactory) findLocation(unLocode UNLocode) (*Location, error) {
	return f.locationRepository.Find(unLocode)
}

//go:generate mockgen -destination=mock/handling_event_repository.go -package=mock . HandlingEventRepository

type HandlingEventRepository interface {
	Store(event *HandlingEvent) error
	LookupHandlingHistoryOfCargo(trachkingID TrackingID) (*HandlingHistory, error)
}
