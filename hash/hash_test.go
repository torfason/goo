package hash

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type T struct {
	B string
	A int
	C bool
}

func TestHash(t *testing.T) {
	assert := assert.New(t)

	t.Run("hash example struct", func(*testing.T) {

		// Hash an example struct
		x := T{"hey", 7, true}
		actual := HashStruct(x)
		expected := "21f9a3e55776be4a6abb64f02371f36946ab20affce4b35eeed90fda89df7518"
		assert.Equal(expected, actual, x)

		// This should match the HashString of the json representaion of the stuct
		assert.Equal(HashString(`{"B":"hey","A":7,"C":true}`), expected)

	})

	t.Run("hash known good values", func(*testing.T) {

		// https://www.movable-type.co.uk/scripts/sha256.html
		tests := []struct {
			input    string
			expected string
		}{
			{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
			{"forty-two", "5cc47001f7c1334db3c568ddbb1c8ee51812aa8e75582ca616b1ec31bf2ddc16"},
			{"þýði", "7b245c95013b74c75b92519618b447d4d0f1d8740663f206fd6f0f6949e4118d"},
			{`{"B":"hey","A":7,"C":true}`, "21f9a3e55776be4a6abb64f02371f36946ab20affce4b35eeed90fda89df7518"},
		}

		for _, test := range tests {
			actual := HashString(test.input)
			expected := test.expected
			assert.Equal(expected, actual, test.input)
		}

	})

	t.Run("hash string to different length outputs", func(*testing.T) {

		tests := []struct {
			input    string
			bits     int
			expected string
		}{
			{"abc", 256, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
			{"abc", 128, "ba7816bf8f01cfea414140de5dae2223"},
			{"abc", 64, "ba7816bf8f01cfea"},
			{"abc", 32, "ba7816bf"},
			{"abc", 512, strings.Join([]string{
				"ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a",
				"2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"}, "")},
		}

		// Default size should be 256
		assert.Equal(tests[0].expected, HashString(tests[0].input))

		// Different sizes
		assert.Equal(tests[0].expected, HashString256(tests[0].input))
		assert.Equal(tests[1].expected, HashString128(tests[1].input))
		assert.Equal(tests[2].expected, HashString64(tests[2].input))
		assert.Equal(tests[3].expected, HashString32(tests[3].input))

		// 512 is a different algorithm
		assert.Equal(tests[4].expected, HashString512(tests[4].input))

	})

	t.Run("hash struct to different length outputs", func(*testing.T) {

		input := T{"hey", 7, true}
		tests := []struct {
			bits     int
			expected string
		}{
			{256, "21f9a3e55776be4a6abb64f02371f36946ab20affce4b35eeed90fda89df7518"},
			{128, "21f9a3e55776be4a6abb64f02371f369"},
			{64, "21f9a3e55776be4a"},
			{32, "21f9a3e5"},
			{512, strings.Join([]string{
				"27cf07949eb36b42416318ec1717a3dbdceaf0554ff753ed811f28cb4bc0e941",
				"ffb79ffd910ccbed16524def9957c24fbfe651165d16c46a56ba6283895a0119"}, "")},
		}

		// Default size should be 256
		assert.Equal(tests[0].expected, HashStruct(input))

		// Different sizes
		assert.Equal(tests[0].expected, HashStruct256(input))
		assert.Equal(tests[1].expected, HashStruct128(input))
		assert.Equal(tests[2].expected, HashStruct64(input))
		assert.Equal(tests[3].expected, HashStruct32(input))

		// 512 is a different algorithm
		assert.Equal(tests[4].expected, HashStruct512(input))

	})

	// Hashing a string as struct will first encapsulate it in a json object
	t.Run("hashing string as string not equal to hashing it as struct", func(t *testing.T) {

		s := "supercalifragilisticexpialidocious"
		hashAsString := HashString(s)
		hashAsStruct := HashStruct(s)
		assert.NotEqual(hashAsString, hashAsStruct)

	})
}

func TestSomeJsonIssuesRelatedToHashing(t *testing.T) {
	assert := assert.New(t)

	x := T{"hey", 7, true}
	rBytes, err := json.Marshal(x)
	rExpected := `{"B":"hey","A":7,"C":true}`
	assert.Nil(err)
	assert.Equal(rExpected, string(rBytes), "json encoding structs should preserve encountered order")

	y := make(map[string]int)
	y["b"] = 6
	y["e"] = 3
	y["gw"] = 4
	y["bs"] = 9
	y["e"] = 2
	y["wb"] = 8
	y["v"] = 5
	y["a"] = 1
	sBytes, err := json.Marshal(y)
	sExpected := `{"a":1,"b":6,"bs":9,"e":2,"gw":4,"v":5,"wb":8}`
	assert.Nil(err)
	assert.Equal(sExpected, string(sBytes), "json encoding maps should preserve alphabetical order")

}
