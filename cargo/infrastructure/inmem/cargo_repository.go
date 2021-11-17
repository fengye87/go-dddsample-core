package inmem

import (
	"fmt"
	"strconv"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

type CargoRepsitory struct {
	cargos []domain.Cargo
}

var _ domain.CargoRepository = &CargoRepsitory{}

func NewCargoRepository() *CargoRepsitory {
	return &CargoRepsitory{}
}

func (r CargoRepsitory) Find(trackingID domain.TrackingID) (*domain.Cargo, error) {
	for _, cargo := range r.cargos {
		if cargo.TrackingID.SameValueAs(trackingID) {
			return &cargo, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (r CargoRepsitory) FindAll() ([]domain.Cargo, error) {
	return r.cargos, nil
}

func (r *CargoRepsitory) Store(cargo *domain.Cargo) error {
	for i, c := range r.cargos {
		if c.TrackingID.SameValueAs(cargo.TrackingID) {
			r.cargos[i] = *cargo
			return nil
		}
	}
	r.cargos = append(r.cargos, *cargo)
	return nil
}

var trackingID int

func (r CargoRepsitory) NextTrackingID() (domain.TrackingID, error) {
	trackingID++
	return domain.TrackingID(strconv.Itoa(trackingID)), nil
}
