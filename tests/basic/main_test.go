package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), input %d, actual = %d ", input, output, actual)
	// }

	assert.Equal(t, AddOne(1), 3, "AddOne(1) should be 2")
	assert.NotEqual(t, 2, 3)
	assert.Nil(t, nil, nil)
}

func TestRequire(t *testing.T) {
	require.Equal(t, AddOne(1), 4, "AddOne(1) should be 2")
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, AddOne(1), 4, "AddOne(1) should be 2")
	fmt.Println("Executing")
}

func TestAddOne2(t *testing.T) {
	require.Equal(t, AddOne2(1), 4, "AddOne(1) should be 3")
}
