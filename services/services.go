package services

import "github.com/juansgt/model-test/dataAccess"

type FindLooksQueryService struct {
	lookRepository dataAccess.ILookRepository
}

func NewFindLooksQueryService(lookRepository dataAccess.ILookRepository) *FindLooksQueryService {
	return &FindLooksQueryService{
		lookRepository: lookRepository,
	}
}

func (findLooksQueryService *FindLooksQueryService) FindLooks() []dataAccess.Look {
	return findLooksQueryService.lookRepository.FindLooks()
}
