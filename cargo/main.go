package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fengye87/dddsample-core/cargo/application"
	"github.com/fengye87/dddsample-core/cargo/infrastructure/dummy"
	"github.com/fengye87/dddsample-core/cargo/infrastructure/inmem"
	"github.com/fengye87/dddsample-core/cargo/infrastructure/logging"
	"github.com/fengye87/dddsample-core/cargo/interfaces/booking"
	"github.com/fengye87/dddsample-core/cargo/interfaces/handling"
	"github.com/fengye87/dddsample-core/cargo/interfaces/tracking"
)

func main() {
	cargoRepository := inmem.NewCargoRepository()
	locationRepository := inmem.NewLocationRepository()
	voyageRepository := inmem.NewVoyageRepository()
	handlingEventRepository := inmem.NewHandlingEventRepository()
	routingService := dummy.NewRoutingService(voyageRepository)

	bookingService := application.NewBookingService(cargoRepository, locationRepository, routingService)
	eventListener := logging.NewEventListener()

	bookingServiceFacade := booking.NewBookingServiceFacade(bookingService, locationRepository, cargoRepository, voyageRepository)
	bookingController := booking.NewController(bookingServiceFacade)
	handlingController := handling.NewController(eventListener)
	trackingController := tracking.NewController(cargoRepository, handlingEventRepository)

	router := mux.NewRouter()
	bookingController.AddToRouter(router.PathPrefix("/booking/").Subrouter())
	handlingController.AddToRouter(router.PathPrefix("/handling/").Subrouter())
	trackingController.AddToRouter(router.PathPrefix("/tracking/").Subrouter())

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("error running server: %s", err)
	}
}
