package feistel


import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Unit tests should contain more examples
func TestXor(t *testing.T) {
	actual_result := xor(&[]int32{116,101},&[]int32{115,116})
	assert.Equal(t, []int32{7,17}, *actual_result)
}

func TestSplitString(t *testing.T) {
	evenWord := "tests"
	_, _, err := splitString(evenWord)
	expected_error := fmt.Errorf("Not able(for now) to split an odd lenght string: %v", evenWord)
	if assert.Error(t, err) {
		assert.Equal(t, expected_error, err)
	}
	test := "test"
	l, r, err := splitString(test)
	require.NoError(t, err)
	fmt.Printf("%v %v\n",l,r)
	assert.Equal(t, []int32{116,101}, *l)
	assert.Equal(t, []int32{115,116}, *r)
}