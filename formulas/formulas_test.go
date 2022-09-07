package formulas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormulas_Sum(t *testing.T) {
	// act
	result := Sum(1, 2)

	// assert
	assert.Equal(t, 3, result)
}
