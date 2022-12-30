package findLooksService

import (
	"testing"

	"github.com/juansgt/model-test/dataAccess"
	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFindLooks_correctCalling_returnExpectedValues(t *testing.T) {

	// Arrange

	assert := assert.New(t)
	var expectedLook *dataAccess.Look = dataAccess.NewLook("1", "Dress", "Bash")
	findLooksQueryService := NewFindLooksQueryService(new(dataAccess.LookRepositoryMock))

	// Act

	looks := findLooksQueryService.Execute()

	// Assert

	assert.Equal(expectedLook.Id(), looks[0].Id())
	assert.Equal(expectedLook.Name, looks[0].Name)
	assert.Equal(expectedLook.Brand, looks[0].Brand)
}
