package domain

type Location struct {
	UNLocode UNLocode
	Name     string
}

func NewLocation(unLocode UNLocode, name string) *Location {
	return &Location{
		UNLocode: unLocode,
		Name:     name,
	}
}

func (l Location) SameIdentityAs(other *Location) bool {
	return other.UNLocode.SameValueAs(l.UNLocode)
}

type UNLocode string

func (u UNLocode) SameValueAs(other UNLocode) bool {
	return other == u
}

//go:generate mockgen -destination=mock/location_repository.go -package=mock . LocationRepository

type LocationRepository interface {
	Find(unLocode UNLocode) (*Location, error)
	FindAll() ([]Location, error)
}
