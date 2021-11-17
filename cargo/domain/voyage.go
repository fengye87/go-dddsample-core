package domain

import (
	"reflect"
	"time"
)

type Voyage struct {
	VoyageNumber VoyageNumber
	Schedule     *Schedule
}

func NewVoyage(voyageNumber VoyageNumber, schedule *Schedule) *Voyage {
	return &Voyage{
		VoyageNumber: voyageNumber,
		Schedule:     schedule,
	}
}

func (v Voyage) SameIdentityAs(other *Voyage) bool {
	return other.VoyageNumber.SameValueAs(v.VoyageNumber)
}

type VoyageNumber string

func (n VoyageNumber) SameValueAs(other VoyageNumber) bool {
	return other == n
}

type Schedule struct {
	CarrierMovements []CarrierMovement
}

func NewSchedule(carrierMovements []CarrierMovement) *Schedule {
	return &Schedule{
		CarrierMovements: carrierMovements,
	}
}

func (s Schedule) SameValueAs(other *Schedule) bool {
	return reflect.DeepEqual(other, &s)
}

type CarrierMovement struct {
	DepartureUNLocode UNLocode
	ArrivalUNLocode   UNLocode
	DepartureTime     time.Time
	ArrivalTime       time.Time
}

func NewCarrierMovement(departureUNLocode UNLocode, arrivalUNLocode UNLocode, departureTime time.Time, arrivalTime time.Time) *CarrierMovement {
	return &CarrierMovement{
		DepartureUNLocode: departureUNLocode,
		ArrivalUNLocode:   arrivalUNLocode,
		DepartureTime:     departureTime,
		ArrivalTime:       arrivalTime,
	}
}

func (m CarrierMovement) SameValueAs(other *CarrierMovement) bool {
	return reflect.DeepEqual(other, &m)
}

type VoyageBuilder struct {
	CarrierMovements []CarrierMovement
	VoyageNumber     VoyageNumber
	DepatureUNLocode UNLocode
}

func NewVoyageBuilder(voyageNumber VoyageNumber, depatureUNLocode UNLocode) *VoyageBuilder {
	return &VoyageBuilder{
		VoyageNumber:     voyageNumber,
		DepatureUNLocode: depatureUNLocode,
	}
}

func (b *VoyageBuilder) AddMovement(arrivalUNLocode UNLocode, depatureTime time.Time, arrivalTime time.Time) *VoyageBuilder {
	b.CarrierMovements = append(b.CarrierMovements, *NewCarrierMovement(b.DepatureUNLocode, arrivalUNLocode, depatureTime, arrivalTime))
	b.DepatureUNLocode = arrivalUNLocode
	return b
}

func (b VoyageBuilder) Build() *Voyage {
	return NewVoyage(b.VoyageNumber, NewSchedule(b.CarrierMovements))
}

//go:generate mockgen -destination=mock/voyage_repository.go -package=mock . VoyageRepository

type VoyageRepository interface {
	Find(voyageNumber VoyageNumber) (*Voyage, error)
	FindAll() ([]Voyage, error)
}
