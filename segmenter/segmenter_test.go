package segmenter

import (
	"fmt"
	"io"
	"testing"

	"github.com/nyarla/go-japanese-segmenter/defaults"
)

func ExampleSegmenter() {
	src := "今日は良い天気ですね"
	seg := New(src)

	for {
		token, err := seg.Segment(defaults.Dictionary)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}

		fmt.Println(token)
	}

	// Output:
	// 今日
	// は
	// 良い
	// 天気
	// です
	// ね
}

func BenchmarkSegmenter(b *testing.B) {
	src := "今日は良い天気ですね"
	seg := New(src)
	num := 0

	b.ReportAllocs()
	b.ResetTimer()

	for idx := 0; idx < b.N; idx++ {
	loop:
		for {
			_, err := seg.Segment(defaults.Dictionary)

			if err != nil && err != io.EOF {
				b.Fail()
			}

			if err == io.EOF {
				break loop
			}

			num += 1
		}

		if num != 6 {
			b.Fatal(num)
		}

		seg.Reset()
		num = 0
	}
}
