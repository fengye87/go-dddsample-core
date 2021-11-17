package handling

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/fengye87/dddsample-core/cargo/application"
	"github.com/fengye87/dddsample-core/cargo/domain"
)

type Controller struct {
	eventListener application.EventListener
}

func NewController(eventListener application.EventListener) *Controller {
	return &Controller{
		eventListener: eventListener,
	}
}

func (c Controller) AddToRouter(router *mux.Router) {
	router.Path("/submit_report").Methods(http.MethodPost).HandlerFunc(c.SubmitReport)
}

func (c Controller) SubmitReport(w http.ResponseWriter, r *http.Request) {
	var report HandlingReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, trackingID := range report.TrackingIDs {
		eventType, err := parseHandlingEventType(report.Type)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		attempt := application.NewHandlingEventRegistrationAttempt(time.Now(), report.CompletionTime, domain.TrackingID(trackingID), domain.VoyageNumber(report.VoyageNumber), eventType, domain.UNLocode(report.UNLocode))
		c.eventListener.ReceivedHandlingEventRegistrationAttempt(attempt)
	}
}

type HandlingReport struct {
	CompletionTime time.Time `json:"completion_time"`
	TrackingIDs    []string  `json:"tracking_ids"`
	Type           string    `json:"type"`
	UNLocode       string    `json:"unlocode"`
	VoyageNumber   string    `json:"voyage_number"`
}

func parseHandlingEventType(s string) (domain.HandlingEventType, error) {
	switch s {
	case "Load":
		return domain.HandlingEventTypeLoad, nil
	case "Unload":
		return domain.HandlingEventTypeUnload, nil
	case "Receive":
		return domain.HandlingEventTypeReceive, nil
	case "Claim":
		return domain.HandlingEventTypeClaim, nil
	case "Customs":
		return domain.HandlingEventTypeCustoms, nil
	default:
		return 0, fmt.Errorf("invalid handling event type")
	}
}
