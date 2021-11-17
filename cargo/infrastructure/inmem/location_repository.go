package inmem

import (
	"fmt"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

type LocationRepository struct {
	locations []domain.Location
}

var _ domain.LocationRepository = &LocationRepository{}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{
		locations: []domain.Location{{
			UNLocode: "CNHKG",
			Name:     "Hongkong",
		}, {
			UNLocode: "AUMEL",
			Name:     "Melbourne",
		}, {
			UNLocode: "SESTO",
			Name:     "Stockholm",
		}, {
			UNLocode: "FIHEL",
			Name:     "Helsinki",
		}, {
			UNLocode: "USCHI",
			Name:     "Chicago",
		}, {
			UNLocode: "JNTKO",
			Name:     "Tokyo",
		}, {
			UNLocode: "DEHAM",
			Name:     "Hamburg",
		}, {
			UNLocode: "CNSHA",
			Name:     "Shanghai",
		}, {
			UNLocode: "NLRTM",
			Name:     "Rotterdam",
		}, {
			UNLocode: "SEGOT",
			Name:     "Göteborg",
		}, {
			UNLocode: "CNHGH",
			Name:     "Hangzhou",
		}, {
			UNLocode: "USNYC",
			Name:     "New York",
		}, {
			UNLocode: "USDAL",
			Name:     "Dallas",
		}},
	}
}

func (r LocationRepository) Find(unLocode domain.UNLocode) (*domain.Location, error) {
	for _, location := range r.locations {
		if location.UNLocode.SameValueAs(unLocode) {
			return &location, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (r LocationRepository) FindAll() ([]domain.Location, error) {
	return r.locations, nil
}
