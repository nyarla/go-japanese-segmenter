package tinydictgen

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	var data = JSONData{
		"BQ4": map[string]int64{
			"BIH": 3761,
			"BIK": 1348,
		},
	}

	var dst = new(strings.Builder)

	err := Render(dst, "example", "-332", data)
	if err != nil {
		t.Fatal(err)
	}
}
