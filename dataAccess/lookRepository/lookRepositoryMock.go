package lookRepository

type LookRepositoryMock struct {
}

func NewLookRepositoryMock() *LookRepositoryMock {
	return new(LookRepositoryMock)
}

func (*LookRepositoryMock) FindLooks() []Look {
	var looks []Look
	var look *Look = &Look{
		id:     "1",
		Colour: "Dress",
		Brand:  "Bash",
	}
	looks = make([]Look, 0, 3)

	looks = append(looks, *look)

	return looks
}
