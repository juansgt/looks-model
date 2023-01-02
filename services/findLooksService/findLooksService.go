package findLooksService

import (
	"github.com/juansgt/generics/services"
	"github.com/juansgt/model-test/v2/dataAccess"
)

type FindLooksQueryService struct {
	lookRepository dataAccess.ILookRepository
}

func NewFindLooksQueryService(lookRepository dataAccess.ILookRepository) services.IQueryServiceNoInput[[]dataAccess.Look] {
	return &FindLooksQueryService{
		lookRepository: lookRepository,
	}
}

func (findLooksQueryService *FindLooksQueryService) Execute() []dataAccess.Look {
	return findLooksQueryService.lookRepository.FindLooks()
}
