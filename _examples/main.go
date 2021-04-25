package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/nyarla/go-japanese-segmenter/dicts/tinyseg"
	"github.com/nyarla/go-japanese-segmenter/segmenter"
)

func main() {
	src := strings.NewReader("今日は良い天気ですね")
	dst := new(strings.Builder)
	dict := segmenter.BiasCalculatorFunc(tinyseg.CalculateBias)
	seg := segmenter.New(dst, src)

	for {
		err := seg.Segment(dict)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}

		fmt.Println(dst.String())
		dst.Reset()
	}

	fmt.Println(dst.String())
	dst.Reset()

	// Output:
	// 今日
	// は
	// 良い
	// 天気
	// ですね
}
