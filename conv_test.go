package imperial

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type example struct {
	Input    string
	Expected float64
}

var examples = []example{
	{Input: `1 1/2"`, Expected: 0.0381},
	{Input: `1'`, Expected: 0.3048},
	{Input: `1 1/2'`, Expected: 0.4572},
	{Input: `1 1/2' 2"`, Expected: 0.4572 + 0.0508},
	{Input: `1 1/2' 2 1/2"`, Expected: 0.4572 + 0.0635},
	{Input: `1' 2 1/2"`, Expected: 0.3683},
	{Input: `2 1/2"`, Expected: 0.0635},
	{Input: `32'`, Expected: 9.7536},
}

func TestParse(t *testing.T) {
	for _, eg := range examples {
		parsed, err := Parse(eg.Input)
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, eg.Expected, parsed)
	}
}

func TestEdges(t *testing.T) {
	fail(t, `1`)
	fail(t, `2 2`)
	fail(t, `1' 3`)
	fail(t, `2' 2\" 2`)
	fail(t, `2' 2\" 2"`)
	fail(t, `2' 2'`)
	fail(t, `2\" 2"`)
	fail(t, `2\" 2 2"`)
	fail(t, `2 2"`)
	fail(t, `2' k"`)
	fail(t, `2' 2/2/2"`)
	fail(t, `2 2/2/2' 2"`)
	fail(t, `2 2/2' 2""`)
	fail(t, `2 2/2'' 2"`)
	fail(t, `2' 2/2' 2"`)

	pass(t, `1'`)
	pass(t, `1' 3"`)
	pass(t, `1 1/2'`)
	pass(t, `1 1/2' 2 2/3"`)
	pass(t, `2 2/3"`)
	pass(t, `2"`)
	pass(t, ` 2 2 / 3"`)
	pass(t, `2 "`)
	pass(t, `2' 2"`)
}

func fail(t *testing.T, test string) {
	_, err := Parse(test)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "could not parse")
}
func pass(t *testing.T, test string) {
	_, err := Parse(test)
	require.Nil(t, err)
}
