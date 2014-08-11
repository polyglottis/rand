package rand

import (
	"strings"
	"testing"
)

var lengths = []int{0, 1, 2, 3, 5, 8, 13, 21, 34}

func TestId(t *testing.T) {
	for _, l := range lengths {
		id, err := Id(l)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("id:", id)
		assertSameLength(t, id, l)
	}
}

func assertSameLength(t *testing.T, id string, l int) {
	if len(id) != l {
		t.Errorf("Expected string of length %i but got %i characters: %s", l, len(id), id)
	}
}

func TestIdPrefix(t *testing.T) {
	for _, l := range lengths {
		for _, prefix := range []string{"", "test", "prefix"} {
			id, err := IdPrefix(prefix, l)

			if len(prefix) > l {
				if err == nil {
					t.Error("Weird behaviour")
				}
			} else {
				if err != nil {
					t.Error(err)
				}
				assertSameLength(t, id, l)
				if !strings.HasPrefix(id, prefix) {
					t.Errorf("%s should start with %s", id, prefix)
				}
			}
		}
	}
}
