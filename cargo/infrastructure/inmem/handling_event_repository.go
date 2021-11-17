package inmem

import "github.com/fengye87/dddsample-core/cargo/domain"

type HandlingEventRepository struct {
	handlingEvents []domain.HandlingEvent
}

var _ domain.HandlingEventRepository = &HandlingEventRepository{}

func NewHandlingEventRepository() *HandlingEventRepository {
	return &HandlingEventRepository{}
}

func (r *HandlingEventRepository) Store(event *domain.HandlingEvent) error {
	r.handlingEvents = append(r.handlingEvents, *event)
	return nil
}

func (r HandlingEventRepository) LookupHandlingHistoryOfCargo(trackingID domain.TrackingID) (*domain.HandlingHistory, error) {
	var events []domain.HandlingEvent
	for _, event := range r.handlingEvents {
		if event.TrackingID.SameValueAs(trackingID) {
			events = append(events, event)
		}
	}
	return &domain.HandlingHistory{
		HandlingEvents: events,
	}, nil
}
