package tinydictgen

import (
	"reflect"
	"testing"
)

func TestModels(t *testing.T) {
	var A = NewItem("BQ4", "BIH", 3761)
	var B = NewItem("BQ4", "BIK", 1348)

	var C, ok = Merge(A, B)
	if !ok {
		t.Fail()
	}

	var D = &Item{
		Var:   "p3",
		Rune:  'B',
		Depth: 0,
		List: []*Item{
			&Item{
				Var:   "t3",
				Rune:  'I',
				Depth: 1,
				List: []*Item{
					&Item{
						Var:   "t4",
						Rune:  'H',
						Depth: 2,
						Bias:  3761,
					},
					&Item{
						Var:   "t4",
						Rune:  'K',
						Depth: 2,
						Bias:  1348,
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(C, D) {
		t.Fatalf("%+v != %+v", C, D)
	}

	var J = JSONData{
		"BQ4": map[string]int64{
			"BIH": 3761,
			"BIK": 1348,
		},
	}

	if !reflect.DeepEqual(J.Items()[0], D) {
		t.Fatalf("%+v != %+v", J.Items()[0], D)
	}
}
