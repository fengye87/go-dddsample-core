package inmem

import (
	"fmt"
	"time"

	"github.com/fengye87/dddsample-core/cargo/domain"
)

type VoyageRepository struct {
	voyages []domain.Voyage
}

var _ domain.VoyageRepository = &VoyageRepository{}

func NewVoyageRepository() *VoyageRepository {

	return &VoyageRepository{
		voyages: []domain.Voyage{
			*domain.NewVoyageBuilder("V100", "CNHKG").
				AddMovement("JNTKO", mustParseTime("2009-03-03"), mustParseTime("2009-03-05")).
				AddMovement("USNYC", mustParseTime("2009-03-06"), mustParseTime("2009-03-09")).Build(),
			*domain.NewVoyageBuilder("V200", "JNTKO").
				AddMovement("USNYC", mustParseTime("2009-03-06"), mustParseTime("2009-03-08")).
				AddMovement("USCHI", mustParseTime("2009-03-10"), mustParseTime("2009-03-14")).
				AddMovement("SESTO", mustParseTime("2009-03-14"), mustParseTime("2009-03-16")).Build(),
			*domain.NewVoyageBuilder("V300", "JNTKO").
				AddMovement("NLRTM", mustParseTime("2009-03-08"), mustParseTime("2009-03-11")).
				AddMovement("DEHAM", mustParseTime("2009-03-11"), mustParseTime("2009-03-12")).
				AddMovement("AUMEL", mustParseTime("2009-03-14"), mustParseTime("2009-03-18")).
				AddMovement("JNTKO", mustParseTime("2009-03-19"), mustParseTime("2009-03-21")).Build(),
			*domain.NewVoyageBuilder("V400", "DEHAM").
				AddMovement("SESTO", mustParseTime("2009-03-14"), mustParseTime("2009-03-15")).
				AddMovement("FIHEL", mustParseTime("2009-03-15"), mustParseTime("2009-03-16")).
				AddMovement("DEHAM", mustParseTime("2009-03-20"), mustParseTime("2009-03-22")).Build(),
			*domain.NewVoyageBuilder("0100S", "CNHKG").
				AddMovement("CNHGH", mustParseTime("2008-10-01 12:00"), mustParseTime("2008-10-03 14:30")).
				AddMovement("JNTKO", mustParseTime("2008-10-03 21:00"), mustParseTime("2008-10-06 06:15")).
				AddMovement("AUMEL", mustParseTime("2008-10-06 11:00"), mustParseTime("2008-10-12 11:30")).
				AddMovement("USNYC", mustParseTime("2008-10-14 12:00"), mustParseTime("2008-10-23 23:10")).Build(),
			*domain.NewVoyageBuilder("0200T", "USNYC").
				AddMovement("USCHI", mustParseTime("2008-10-24 07:00"), mustParseTime("2008-10-24 17:45")).
				AddMovement("USDAL", mustParseTime("2008-10-24 21:25"), mustParseTime("2008-10-25 19:30")).Build(),
			*domain.NewVoyageBuilder("0300A", "USDAL").
				AddMovement("DEHAM", mustParseTime("2008-10-29 03:30"), mustParseTime("2008-10-31 14:00")).
				AddMovement("SESTO", mustParseTime("2008-11-01 15:20"), mustParseTime("2008-11-01 18:40")).
				AddMovement("FIHEL", mustParseTime("2008-11-02 09:00"), mustParseTime("2008-11-02 11:15")).Build(),
			*domain.NewVoyageBuilder("0301S", "USDAL").
				AddMovement("FIHEL", mustParseTime("2008-10-29 03:30"), mustParseTime("2008-11-05 15:45")).Build(),
			*domain.NewVoyageBuilder("0400S", "FIHEL").
				AddMovement("NLRTM", mustParseTime("2008-11-04 05:50"), mustParseTime("2008-11-06 14:10")).
				AddMovement("CNSHA", mustParseTime("2008-11-10 21:45"), mustParseTime("2008-11-22 16:40")).
				AddMovement("CNHKG", mustParseTime("2008-11-24 07:00"), mustParseTime("2008-11-28 13:37")).Build(),
		},
	}
}

func (r VoyageRepository) Find(voyageNumber domain.VoyageNumber) (*domain.Voyage, error) {
	for _, voyage := range r.voyages {
		if voyage.VoyageNumber.SameValueAs(voyageNumber) {
			return &voyage, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (r VoyageRepository) FindAll() ([]domain.Voyage, error) {
	return r.voyages, nil
}

func mustParseTime(value string) time.Time {
	if t, err := time.Parse("2006-01-02", value); err == nil {
		return t
	}
	if t, err := time.Parse("2006-01-02 15:04", value); err == nil {
		return t
	}
	panic(value)
}
