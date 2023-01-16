package findLooksService

import "github.com/juansgt/model-test/v3/dataAccess/lookRepository"

type FindLooksQueryService struct {
	lookRepository lookRepository.ILookRepository
}

func NewFindLooksQueryService(lookRepository lookRepository.ILookRepository) *FindLooksQueryService {
	return &FindLooksQueryService{
		lookRepository: lookRepository,
	}
}

func (findLooksQueryService *FindLooksQueryService) Execute() []lookRepository.Look {
	return findLooksQueryService.lookRepository.FindLooks()
}
