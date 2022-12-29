package dataAccess

type Look struct {
	id    string
	Name  string
	Brand string
}

func NewLook(id string, name string, brand string) *Look {
	return &Look{
		id:    id,
		Name:  name,
		Brand: brand,
	}
}

func (look *Look) Id() string {
	return look.id
}
