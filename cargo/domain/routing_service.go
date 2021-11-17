package domain

//go:generate mockgen -destination=mock/routing_service.go -package=mock . RoutingService

type RoutingService interface {
	FetchRoutesForSpecification(routeSpecification *RouteSpecification) ([]Itinerary, error)
}
