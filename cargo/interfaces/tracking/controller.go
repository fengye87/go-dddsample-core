package tracking

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

type Controller struct {
	cargoRepository         domain.CargoRepository
	handlingEventRepository domain.HandlingEventRepository
}

func NewController(cargoRepository domain.CargoRepository, handlingEventRepository domain.HandlingEventRepository) *Controller {
	return &Controller{
		cargoRepository:         cargoRepository,
		handlingEventRepository: handlingEventRepository,
	}
}

func (c Controller) AddToRouter(router *mux.Router) {
	router.Path("/").Methods(http.MethodGet).HandlerFunc(c.Show)
}

func (c Controller) Show(w http.ResponseWriter, r *http.Request) {
	trackingID := r.URL.Query().Get("tracking_id")
	cargo, err := c.cargoRepository.Find(domain.TrackingID(trackingID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handlingHistory, err := c.handlingEventRepository.LookupHandlingHistoryOfCargo(cargo.TrackingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	events := handlingHistory.DistinctEventsByCompletionTime()
	dtos := NewHandlingEventDTOAssembler().ToDTOs(events)
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type HandlingEventDTO struct {
	Type           string    `json:"type"`
	VoyageNumber   string    `json:"voyage_number"`
	UNLocode       string    `json:"unlocode"`
	CompletionTime time.Time `json:"completion_time"`
}

type HandlingEventDTOAssembler struct {
}

func NewHandlingEventDTOAssembler() *HandlingEventDTOAssembler {
	return &HandlingEventDTOAssembler{}
}

func (a HandlingEventDTOAssembler) ToDTO(event *domain.HandlingEvent) *HandlingEventDTO {
	var typeStr string
	switch event.Type {
	case domain.HandlingEventTypeLoad:
		typeStr = "Load"
	case domain.HandlingEventTypeUnload:
		typeStr = "Unload"
	case domain.HandlingEventTypeReceive:
		typeStr = "Receive"
	case domain.HandlingEventTypeClaim:
		typeStr = "Claim"
	case domain.HandlingEventTypeCustoms:
		typeStr = "Customs"
	default:
		panic(nil)
	}
	return &HandlingEventDTO{
		Type:           typeStr,
		VoyageNumber:   string(event.VoyageNumber),
		UNLocode:       string(event.UNLocode),
		CompletionTime: event.CompletionTime,
	}
}

func (a HandlingEventDTOAssembler) ToDTOs(events []domain.HandlingEvent) []HandlingEventDTO {
	var dtos []HandlingEventDTO
	for _, event := range events {
		dtos = append(dtos, *a.ToDTO(&event))
	}
	return dtos
}
