package lookRepository

type Look struct {
	id     string
	Colour string
	Brand  string
}

func NewLook(id string, colour string, brand string) *Look {
	return &Look{
		id:     id,
		Colour: colour,
		Brand:  brand,
	}
}

func (look *Look) Id() string {
	return look.id
}
