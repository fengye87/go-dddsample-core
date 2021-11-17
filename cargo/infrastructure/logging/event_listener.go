package logging

import (
	"log"

	"github.com/fengye87/dddsample-core/cargo/application"
	"github.com/fengye87/dddsample-core/cargo/domain"
)

type EventListener struct {
}

func NewEventListener() *EventListener {
	return &EventListener{}
}

var _ application.EventListener = &EventListener{}

func (l EventListener) CargoWasHandled(event *domain.HandlingEvent) {
	log.Printf("[CargoWasHandled] %#v\n", *event)
}

func (l EventListener) CargoWasMisdirected(cargo *domain.Cargo) {
	log.Printf("[CargoWasMisdirected] %#v\n", *cargo)
}

func (l EventListener) CargoHasArrived(cargo *domain.Cargo) {
	log.Printf("[CargoHasArrived] %#v\n", *cargo)
}

func (l EventListener) ReceivedHandlingEventRegistrationAttempt(attempt *application.HandlingEventRegistrationAttempt) {
	log.Printf("[ReceivedHandlingEventRegistrationAttempt] %#v\n", *attempt)
}
