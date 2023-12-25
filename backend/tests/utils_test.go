package tests

import (
	"fmt"
	"testing"

	"github.com/SanjaySinghRajpoot/backend/utils"
	"github.com/stretchr/testify/assert"
)

func Test_HashPassword(t *testing.T) {

	hashedPassword, err := utils.HashPassword("test")
	assert.Nil(t, err)
	assert.NotNil(t, hashedPassword)

	// If passwords are not same then error will be returned
	error := utils.VerifyPassword(hashedPassword, "test")

	assert.Nil(t, error)
}

func Test_ValidPassword(t *testing.T) {

	checkPass := utils.ValidPassword("Uu@123rrrr")

	fmt.Println(checkPass)

	assert.Equal(t, true, checkPass)
}
