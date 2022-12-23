package dataAccess

type Look struct {
	id    string
	name  string
	brand string
}

type ILookRepository interface {
	FindLooks() Look
}

type LookRepositoryMock struct {
}

func (*LookRepositoryMock) FindLooks() []Look {
	var looks []Look
	var look *Look = &Look{
		id:    "1",
		name:  "Dress",
		brand: "Bash",
	}
	looks = make([]Look, 0, 3)

	looks = append(looks, *look)

	return looks
}
