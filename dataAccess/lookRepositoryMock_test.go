package dataAccess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFindLooks_correctCalling_returnExpectedValues(t *testing.T) {

	// Arrange

	assert := assert.New(t)
	var expectedLook *Look = &Look{
		id:    "1",
		Name:  "Dress",
		Brand: "Bash",
	}
	lookRepositoryMock := new(LookRepositoryMock)

	// Act

	looks := lookRepositoryMock.FindLooks()

	// Assert

	assert.Equal(expectedLook.id, looks[0].id)
	assert.Equal(expectedLook.Name, looks[0].Name)
	assert.Equal(expectedLook.Brand, looks[0].Brand)
}
