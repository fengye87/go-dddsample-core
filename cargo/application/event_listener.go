package application

import (
	"time"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

//go:generate mockgen -destination=mock/event_listener.go -package=mock . EventListener

type EventListener interface {
	CargoWasHandled(event *domain.HandlingEvent)
	CargoWasMisdirected(cargo *domain.Cargo)
	CargoHasArrived(cargo *domain.Cargo)
	ReceivedHandlingEventRegistrationAttempt(attempt *HandlingEventRegistrationAttempt)
}

type HandlingEventRegistrationAttempt struct {
	RegistrationTime time.Time
	CompletionTime   time.Time
	TrackingID       domain.TrackingID
	VoyageNumber     domain.VoyageNumber
	Type             domain.HandlingEventType
	UNLocode         domain.UNLocode
}

func NewHandlingEventRegistrationAttempt(registrationTime time.Time, completionTime time.Time, trackingID domain.TrackingID, voyageNumber domain.VoyageNumber, tp domain.HandlingEventType, unLocode domain.UNLocode) *HandlingEventRegistrationAttempt {
	return &HandlingEventRegistrationAttempt{
		RegistrationTime: registrationTime,
		CompletionTime:   completionTime,
		TrackingID:       trackingID,
		VoyageNumber:     voyageNumber,
		Type:             tp,
		UNLocode:         unLocode,
	}
}
