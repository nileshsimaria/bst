package bst

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   []int
		error string
	}{
		{
			name:  "already-asc-sorted",
			input: []int{1, 2, 3, 4, 5},
			exp:   []int{1, 2, 3, 4, 5},
		}, {
			name:  "unsorted-1",
			input: []int{1, 20, 30, 4, 15},
			exp:   []int{1, 4, 15, 20, 30},
		}, {
			name:  "unsorted-2",
			input: []int{7, 3, 5, 6, 4, 2, 1},
			exp:   []int{1, 2, 3, 4, 5, 6, 7},
		}, {
			name:  "already-dec-sorted",
			input: []int{5, 4, 3, 2, 1},
			exp:   []int{1, 2, 3, 4, 5},
		}, {
			name:  "empty",
			input: []int{},
			exp:   []int{},
		}, {
			name:  "dup",
			input: []int{1, 10, 5, 3, 5},
			error: "Found duplicate: data=5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bstTree := NewBST()
			c := make(chan int)

			err := bstTree.AddBulk(test.input...)
			if test.error != "" {
				if err == nil {
					t.Error("expected error but error is nil")
				} else if err.Error() != test.error {
					t.Errorf("wrong error: got: %s, wanted: %s", err.Error(), test.error)
				}
				t.Skip()
			}

			func() {
				go bstTree.Walk(c)
			}()

			got := make([]int, 0)
			for d := range c {
				got = append(got, d)
			}
			if eq := reflect.DeepEqual(got, test.exp); !eq {
				t.Errorf("wanted: %v, got: %v", test.exp, got)
			}
		})
	}
}
