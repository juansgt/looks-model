package lookRepository_test

import (
	"testing"

	"github.com/juansgt/looks-model/v3/dataAccess/lookRepository"
	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFindLooks_correctCalling_returnExpectedValues(t *testing.T) {

	// Arrange

	assert := assert.New(t)
	var expectedLook *lookRepository.Look = lookRepository.NewLook("1", "Dress", "Bash")
	var lookRepository lookRepository.ILookRepository = lookRepository.NewLookRepositoryMock()

	// Act

	looks := lookRepository.FindLooks()

	// Assert

	assert.Equal(expectedLook.Id(), looks[0].Id())
	assert.Equal(expectedLook.Colour, looks[0].Colour)
	assert.Equal(expectedLook.Brand, looks[0].Brand)
}
