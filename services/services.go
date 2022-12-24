package services

import "github.com/juansgt/model-test/dataAccess"

type FindLooksQueryService struct {
	lookRepository dataAccess.ILookRepository
}

func (findLooksQueryService *FindLooksQueryService) FindLooks() []dataAccess.Look {
	return findLooksQueryService.lookRepository.FindLooks()
}
