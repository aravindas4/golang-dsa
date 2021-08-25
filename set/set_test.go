package set

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	set := &Set{}
	set.New()

	t.Run("Test Initially", func(t *testing.T) {
		received := set.GetElements()
		expected := map[int]bool{}
		if !reflect.DeepEqual(expected, received) {
			t.Error("Test Failed: inputted ", nil, " expected ", expected, ", recieved ", received)
		}
	})

	t.Run("Test AddElement", func(t *testing.T) {
		set.AddElement(1)
		received := set.ContainsElement(1)
		expected := true
		if expected != received {
			t.Error("Test Failed: inputted ", 1, " expected ", expected, ", recieved ", received)
		}

		set.AddElement(1)
		received_count := set.CountElements()
		expected_count := 1
		if expected != received {
			t.Error("Test Failed: inputted ", nil, " expected ", expected_count, ", recieved ", received_count)
		}
	})

	t.Run("Test DeleteElement", func(t *testing.T) {
		set.DeleteElement(1)
		received := set.ContainsElement(1)
		expected := false
		if expected != received {
			t.Error("Test Failed: inputted ", 1, " expected ", expected, ", recieved ", received)
		}
	})
}
