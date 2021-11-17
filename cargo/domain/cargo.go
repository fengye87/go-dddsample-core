package domain

import (
	"reflect"
	"time"
)

type Cargo struct {
	TrackingID         TrackingID
	OriginUNLocode     UNLocode
	RouteSpecification *RouteSpecification
	Itinerary          *Itinerary
	Delivery           *Delivery
}

func NewCargo(trackingID TrackingID, routeSpecification *RouteSpecification) *Cargo {
	itinerary := NewItinerary(nil)
	handlingHistory := NewHandlingHistory(nil)
	return &Cargo{
		TrackingID:         trackingID,
		OriginUNLocode:     routeSpecification.OriginUNLocode,
		RouteSpecification: routeSpecification,
		Itinerary:          itinerary,
		Delivery:           deliveryDerivedFrom(routeSpecification, itinerary, handlingHistory),
	}
}

func (c *Cargo) SpecifyNewRoute(routeSpecification *RouteSpecification) {
	c.RouteSpecification = routeSpecification
	c.Delivery = c.Delivery.updateOnRouting(c.RouteSpecification, c.Itinerary)
}

func (c *Cargo) AssignToRoute(itinerary *Itinerary) {
	c.Itinerary = itinerary
	c.Delivery = c.Delivery.updateOnRouting(c.RouteSpecification, c.Itinerary)
}

func (c *Cargo) DeriveDeliveryProgress(handlingHistory *HandlingHistory) {
	c.Delivery = deliveryDerivedFrom(c.RouteSpecification, c.Itinerary, handlingHistory)
}

func (c Cargo) SameIdentityAs(other *Cargo) bool {
	return other.TrackingID.SameValueAs(c.TrackingID)
}

type TrackingID string

func (id TrackingID) SameValueAs(other TrackingID) bool {
	return other == id
}

type RouteSpecification struct {
	OriginUNLocode      UNLocode
	DestinationUNLocode UNLocode
	ArrivalDeadline     time.Time
}

func NewRouteSpecification(originUNLocode UNLocode, destinationUNLocode UNLocode, arrivalDeadline time.Time) *RouteSpecification {
	return &RouteSpecification{
		OriginUNLocode:      originUNLocode,
		DestinationUNLocode: destinationUNLocode,
		ArrivalDeadline:     arrivalDeadline,
	}
}

func (s RouteSpecification) IsSatisfiedBy(itinerary *Itinerary) bool {
	return s.OriginUNLocode.SameValueAs(itinerary.InitialDepartureUNLocode()) && s.DestinationUNLocode.SameValueAs(itinerary.FinalArrivalUNLocode()) && s.ArrivalDeadline.After(itinerary.FinalArrivalDate())
}

func (s RouteSpecification) SameValueAs(other *RouteSpecification) bool {
	return reflect.DeepEqual(other, &s)
}

type Itinerary struct {
	Legs []Leg
}

func NewItinerary(legs []Leg) *Itinerary {
	return &Itinerary{
		Legs: legs,
	}
}

func (it Itinerary) IsExpected(event *HandlingEvent) bool {
	if len(it.Legs) == 0 {
		return true
	}

	switch event.Type {
	case HandlingEventTypeReceive:
		return it.InitialDepartureUNLocode().SameValueAs(event.UNLocode)
	case HandlingEventTypeLoad:
		for _, leg := range it.Legs {
			if leg.LoadUNLocode.SameValueAs(event.UNLocode) && leg.VoyageNumber.SameValueAs(event.VoyageNumber) {
				return true
			}
		}
		return false
	case HandlingEventTypeUnload:
		for _, leg := range it.Legs {
			if leg.UnloadUNLocode.SameValueAs(event.UNLocode) && leg.VoyageNumber.SameValueAs(event.VoyageNumber) {
				return true
			}
		}
		return false
	case HandlingEventTypeClaim:
		return it.FinalArrivalUNLocode().SameValueAs(event.UNLocode)
	default:
		return true
	}
}

func (it Itinerary) InitialDepartureUNLocode() UNLocode {
	if len(it.Legs) == 0 {
		return UNLocode("")
	}
	return it.Legs[0].LoadUNLocode
}

func (it Itinerary) FinalArrivalUNLocode() UNLocode {
	if len(it.Legs) == 0 {
		return UNLocode("")
	}
	return it.Legs[len(it.Legs)-1].UnloadUNLocode
}

func (it Itinerary) FinalArrivalDate() time.Time {
	if len(it.Legs) == 0 {
		return time.Time{}
	}
	return it.Legs[len(it.Legs)-1].UnloadTime
}

func (it Itinerary) SameValueAs(other *Itinerary) bool {
	return reflect.DeepEqual(other, &it)
}

type Leg struct {
	VoyageNumber   VoyageNumber
	LoadUNLocode   UNLocode
	UnloadUNLocode UNLocode
	LoadTime       time.Time
	UnloadTime     time.Time
}

func NewLeg(voyageNumber VoyageNumber, loadUNLocode UNLocode, unloadUNLocode UNLocode, loadTime time.Time, unloadTime time.Time) *Leg {
	return &Leg{
		VoyageNumber:   voyageNumber,
		LoadUNLocode:   loadUNLocode,
		UnloadUNLocode: unloadUNLocode,
		LoadTime:       loadTime,
		UnloadTime:     unloadTime,
	}
}

func (l Leg) SameValueAs(other *Leg) bool {
	return reflect.DeepEqual(other, &l)
}

type Delivery struct {
	TransportStatus         TransportStatus
	LastKnownUNLocode       UNLocode
	CurrentVoyageNumber     VoyageNumber
	Misdirected             bool
	ETA                     time.Time
	NextExpectedActivity    *HandlingActivity
	IsUnloadedAtDestination bool
	RoutingStatus           RoutingStatus
	CalculatedAt            time.Time
	LastEvent               *HandlingEvent
}

func newDelivery(lastEvent *HandlingEvent, itinerary *Itinerary, routeSpecification *RouteSpecification) *Delivery {
	d := &Delivery{
		CalculatedAt: time.Now(),
		LastEvent:    lastEvent,
	}

	d.Misdirected = d.calculateMisdirectionStatus(itinerary)
	d.RoutingStatus = d.calculateRoutingStatus(itinerary, routeSpecification)
	d.TransportStatus = d.calculateTransportStatus()
	d.LastKnownUNLocode = d.calculateLastKnownUNLocode()
	d.CurrentVoyageNumber = d.calculateCurrentVoyageNumber()
	d.ETA = d.calculateETA(itinerary)
	d.NextExpectedActivity = d.calculateNextExpectedActivity(routeSpecification, itinerary)
	d.IsUnloadedAtDestination = d.calculateUnloadedAtDestination(routeSpecification)
	return d
}

func deliveryDerivedFrom(routeSpecification *RouteSpecification, itinerary *Itinerary, handlingHistory *HandlingHistory) *Delivery {
	var lastEvent *HandlingEvent
	if handlingHistory != nil {
		lastEvent = handlingHistory.MostRecentlyCompletedEvent()
	}
	return newDelivery(lastEvent, itinerary, routeSpecification)
}

func (d Delivery) updateOnRouting(routeSpecification *RouteSpecification, itinerary *Itinerary) *Delivery {
	return newDelivery(d.LastEvent, itinerary, routeSpecification)
}

func (d Delivery) calculateTransportStatus() TransportStatus {
	if d.LastEvent == nil {
		return TransportStatusNotReceived
	}

	switch d.LastEvent.Type {
	case HandlingEventTypeLoad:
		return TransportStatusOnboardCarrier
	case HandlingEventTypeUnload, HandlingEventTypeReceive, HandlingEventTypeCustoms:
		return TransportStatusInPort
	case HandlingEventTypeClaim:
		return TransportStatusClaimed
	default:
		return TransportStatusUnknown
	}
}

func (d Delivery) calculateLastKnownUNLocode() UNLocode {
	if d.LastEvent == nil {
		return UNLocode("")
	}
	return d.LastEvent.UNLocode
}

func (d Delivery) calculateCurrentVoyageNumber() VoyageNumber {
	if d.LastEvent == nil {
		return VoyageNumber("")
	}
	return d.LastEvent.VoyageNumber
}

func (d Delivery) calculateMisdirectionStatus(itinerary *Itinerary) bool {
	if d.LastEvent == nil {
		return false
	}
	return !itinerary.IsExpected(d.LastEvent)
}

func (d Delivery) calculateETA(itinerary *Itinerary) time.Time {
	if d.onTrack() {
		return itinerary.FinalArrivalDate()
	}
	return time.Time{}
}

func (d Delivery) calculateNextExpectedActivity(routeSpecification *RouteSpecification, itinerary *Itinerary) *HandlingActivity {
	if !d.onTrack() {
		return nil
	}

	if d.LastEvent == nil {
		return NewHandlingActivity(HandlingEventTypeReceive, routeSpecification.OriginUNLocode, VoyageNumber(""))
	}

	switch d.LastEvent.Type {
	case HandlingEventTypeLoad:
		for _, leg := range itinerary.Legs {
			if leg.LoadUNLocode.SameValueAs(d.LastEvent.UNLocode) {
				return NewHandlingActivity(HandlingEventTypeUnload, leg.UnloadUNLocode, leg.VoyageNumber)
			}
		}
		return nil
	case HandlingEventTypeUnload:
		for i := 0; i < len(itinerary.Legs); i++ {
			leg := itinerary.Legs[i]
			if leg.UnloadUNLocode.SameValueAs(d.LastEvent.UNLocode) {
				if i+1 >= len(itinerary.Legs) {
					return NewHandlingActivity(HandlingEventTypeClaim, leg.UnloadUNLocode, VoyageNumber(""))
				}
				nextLeg := itinerary.Legs[i+1]
				return NewHandlingActivity(HandlingEventTypeLoad, nextLeg.LoadUNLocode, nextLeg.VoyageNumber)
			}
		}
		return nil
	case HandlingEventTypeReceive:
		firstLeg := itinerary.Legs[0]
		return NewHandlingActivity(HandlingEventTypeLoad, firstLeg.LoadUNLocode, firstLeg.VoyageNumber)
	default:
		return nil
	}
}

func (d *Delivery) calculateRoutingStatus(itinerary *Itinerary, routeSpecification *RouteSpecification) RoutingStatus {
	if itinerary == nil {
		return RoutingStatusNotRouted
	}
	if routeSpecification.IsSatisfiedBy(itinerary) {
		return RoutingStatusRouted
	}
	return RoutingStatusMisrouted
}

func (d Delivery) calculateUnloadedAtDestination(routeSpecification *RouteSpecification) bool {
	if d.LastEvent == nil {
		return false
	}
	return d.LastEvent.Type.SameValueAs(HandlingEventTypeUnload) && routeSpecification.DestinationUNLocode.SameValueAs(d.LastEvent.UNLocode)
}

func (d Delivery) onTrack() bool {
	return d.RoutingStatus == RoutingStatusRouted && !d.Misdirected
}

type TransportStatus int

const (
	TransportStatusUnset TransportStatus = iota
	TransportStatusNotReceived
	TransportStatusInPort
	TransportStatusOnboardCarrier
	TransportStatusClaimed
	TransportStatusUnknown
)

func (s TransportStatus) SameValueAs(other TransportStatus) bool {
	return other == s
}

type HandlingActivity struct {
	Type         HandlingEventType
	UNLocode     UNLocode
	VoyageNumber VoyageNumber
}

func NewHandlingActivity(tp HandlingEventType, unLocode UNLocode, voyageNumber VoyageNumber) *HandlingActivity {
	return &HandlingActivity{
		Type:         tp,
		UNLocode:     unLocode,
		VoyageNumber: voyageNumber,
	}
}

func (a HandlingActivity) SameValueAs(other *HandlingActivity) bool {
	return reflect.DeepEqual(other, &a)
}

type RoutingStatus int

const (
	RoutingStatusUnset RoutingStatus = iota
	RoutingStatusNotRouted
	RoutingStatusRouted
	RoutingStatusMisrouted
)

func (s RoutingStatus) SameValueAs(other RoutingStatus) bool {
	return other == s
}

//go:generate mockgen -destination=mock/cargo_repository.go -package=mock . CargoRepository

type CargoRepository interface {
	Find(trackingID TrackingID) (*Cargo, error)
	FindAll() ([]Cargo, error)
	Store(cargo *Cargo) error
	NextTrackingID() (TrackingID, error)
}
