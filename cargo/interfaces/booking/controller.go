package booking

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Controller struct {
	bookingServiceFacade BookingServiceFacade
}

func NewController(bookingServiceFacade BookingServiceFacade) *Controller {
	return &Controller{
		bookingServiceFacade: bookingServiceFacade,
	}
}

func (c Controller) AddToRouter(router *mux.Router) {
	router.Path("/new").Methods(http.MethodGet).HandlerFunc(c.New)
	router.Path("/").Methods(http.MethodPost).HandlerFunc(c.Create)
	router.Path("/").Methods(http.MethodGet).HandlerFunc(c.List)
	router.Path("/{tracking_id}").Methods(http.MethodGet).HandlerFunc(c.Show)
	router.Path("/{tracking_id}/select_itinerary").Methods(http.MethodGet).HandlerFunc(c.SelectItinerary)
	router.Path("/{tracking_id}/assign_itinerary").Methods(http.MethodPost).HandlerFunc(c.AssignItinerary)
	router.Path("/{tracking_id}/pick_new_destination").Methods(http.MethodGet).HandlerFunc(c.PickNewDestination)
	router.Path("/{tracking_id}/change_destination").Methods(http.MethodPost).HandlerFunc(c.ChangeDestination)
}

func (c Controller) New(w http.ResponseWriter, r *http.Request) {
	locations, err := c.bookingServiceFacade.ListShippingLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	var command RegistrationCommand
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	trackingID, err := c.bookingServiceFacade.BookNewCargo(command.OriginUNLocode, command.DestinationUNLocode, command.ArrivalDeadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(trackingID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type RegistrationCommand struct {
	OriginUNLocode      string    `json:"origin_unlocode"`
	DestinationUNLocode string    `json:"destination_unlocode"`
	ArrivalDeadline     time.Time `json:"arrival_deadline"`
}

func (c Controller) List(w http.ResponseWriter, r *http.Request) {
	cargos, err := c.bookingServiceFacade.ListAllCargos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cargos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackingID := vars["tracking_id"]
	cargoRouting, err := c.bookingServiceFacade.LoadCargoForRouting(trackingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cargoRouting); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) SelectItinerary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackingID := vars["tracking_id"]
	cargo, err := c.bookingServiceFacade.LoadCargoForRouting(trackingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	routeCandidates, err := c.bookingServiceFacade.RequestPossibleRoutesForCargo(trackingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode([]interface{}{cargo, routeCandidates}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) AssignItinerary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackingID := vars["tracking_id"]

	var command RouteAssignmentCommand
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var legDTOs []LegDTO
	for _, leg := range command.Legs {
		legDTOs = append(legDTOs, LegDTO{
			VoyageNubmer: leg.VoyageNumber,
			FromUNLocode: leg.FromUNLocode,
			ToUNLocode:   leg.ToUNLocode,
			LoadTime:     leg.FromTime,
			UnloadTime:   leg.ToTime,
		})
	}

	selectedRoute := RouteCandidateDTO{
		Legs: legDTOs,
	}
	if err := c.bookingServiceFacade.AssignCargoToRoute(trackingID, &selectedRoute); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type RouteAssignmentCommand struct {
	Legs []LegCommand `json:"legs"`
}

type LegCommand struct {
	VoyageNumber string    `json:"voyage_number"`
	FromUNLocode string    `json:"from_unlocode"`
	ToUNLocode   string    `json:"to_unlocode"`
	FromTime     time.Time `json:"from_time"`
	ToTime       time.Time `json:"to_time"`
}

func (c Controller) PickNewDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackingID := vars["tracking_id"]
	cargo, err := c.bookingServiceFacade.LoadCargoForRouting(trackingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locations, err := c.bookingServiceFacade.ListShippingLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode([]interface{}{cargo, locations}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Controller) ChangeDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackingID := vars["tracking_id"]
	unLocode := r.URL.Query().Get("unlocode")
	if err := c.bookingServiceFacade.ChangeDestination(trackingID, unLocode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//go:generate mockgen -destination=mock/booking_service_facade.go -package=mock . BookingServiceFacade

type BookingServiceFacade interface {
	BookNewCargo(originUNLocode string, destinationUNLocode string, arrivalDeadline time.Time) (string, error)
	LoadCargoForRouting(trackingID string) (*CargoRoutingDTO, error)
	AssignCargoToRoute(trackingID string, route *RouteCandidateDTO) error
	ChangeDestination(trackingID string, newDestinationUNLocode string) error
	RequestPossibleRoutesForCargo(trackingID string) ([]RouteCandidateDTO, error)
	ListShippingLocations() ([]LocationDTO, error)
	ListAllCargos() ([]CargoRoutingDTO, error)
}

type CargoRoutingDTO struct {
	TrackingID               string    `json:"tracking_id"`
	OriginUNLocode           string    `json:"origin_unlocode"`
	FinalDestinationUNLocode string    `json:"final_destination_unlocode"`
	ArrivalDeadline          time.Time `json:"arrival_deadline"`
	Misrouted                bool      `json:"misrouted"`
	Legs                     []LegDTO  `json:"legs"`
}

type LegDTO struct {
	VoyageNubmer string    `json:"voyage_number"`
	FromUNLocode string    `json:"from_unlocode"`
	ToUNLocode   string    `json:"to_unlocode"`
	LoadTime     time.Time `json:"load_time"`
	UnloadTime   time.Time `json:"unload_time"`
}

type RouteCandidateDTO struct {
	Legs []LegDTO `json:"legs"`
}

type LocationDTO struct {
	UNLocode string `json:"unlocode"`
	Name     string `json:"name"`
}
