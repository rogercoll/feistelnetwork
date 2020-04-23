package feistel


import (
	"log"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Example_run() {
	ceasar := func(text []rune) []rune {
		//for quick test we set the same shift = 3
		result := make([]rune, len(text))
		for i,t := range text {
			s := int(t) + 3
			if s > 'z' {
				s = s - 26
			} else if s < 'a' {
				s = s + 26
			}
			result[i] = rune(s)
		}
		return result
	}
	encrypted, err := Run("hell", 1, ceasar)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encrypted)
	// Output:
    // [7 10 108 108]
}

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