package findLooksService_test

import (
	"testing"

	"github.com/juansgt/generics/services"
	"github.com/juansgt/model-test/v3/dataAccess/lookRepository"
	"github.com/juansgt/model-test/v3/services/findLooksService"
	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFindLooks_correctCalling_returnExpectedValues(t *testing.T) {

	// Arrange

	assert := assert.New(t)
	var expectedLook *lookRepository.Look = lookRepository.NewLook("1", "Dress", "Bash")
	var findLooksQueryService services.IQueryServiceNoInput[[]lookRepository.Look] = findLooksService.NewFindLooksQueryService(lookRepository.NewLookRepositoryMock())

	// Act

	looks := findLooksQueryService.Execute()

	// Assert

	assert.Equal(expectedLook.Id(), looks[0].Id())
	assert.Equal(expectedLook.Name, looks[0].Name)
	assert.Equal(expectedLook.Brand, looks[0].Brand)
}
