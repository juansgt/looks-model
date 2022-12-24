package services

import "github.com/juansgt/model-test/dataAccess"

type FindLooksQueryService struct {
	lookRepository dataAccess.ILookRepository
}

func (this *FindLooksQueryService) FindLooks() []dataAccess.Look {
	return this.lookRepository.FindLooks()
}
