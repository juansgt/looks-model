package dataAccess

type LookRepositoryMock struct {
}

func (*LookRepositoryMock) FindLooks() []Look {
	var looks []Look
	var look *Look = &Look{
		id:    "1",
		Name:  "Dress",
		Brand: "Bash",
	}
	looks = make([]Look, 0, 3)

	looks = append(looks, *look)

	return looks
}
