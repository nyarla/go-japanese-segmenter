package segmenter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/nyarla/go-japanese-segmenter/dicts/tinyseg"
)

type dummyWriter struct{}

func (dummy *dummyWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (dummy *dummyWriter) Reset() error {
	return nil
}

func ExampleSegmenter() {
	src := strings.NewReader("今日は良い天気ですね")
	dst := new(strings.Builder)
	dict := BiasCalculatorFunc(tinyseg.CalculateBias)
	seg := New(dst, src)

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

func BenchmarkSegmenter(b *testing.B) {
	msg := "今日は良い天気ですね"
	src := strings.NewReader(msg)
	dst := new(dummyWriter)
	dict := BiasCalculatorFunc(tinyseg.CalculateBias)
	seg := New(dst, src)

	b.ReportAllocs()
	b.ResetTimer()

	for idx := 0; idx < b.N; idx++ {
	loop:
		for {
			err := seg.Segment(dict)

			if err != nil && err != io.EOF {
				b.Fail()
			}

			if err == io.EOF {
				break loop
			}

			dst.Reset()
		}

		src.Reset(msg)
		seg.Reset(src)
	}
}

func BenchmarkSegmentTextInMemory(b *testing.B) {
	src, err := ioutil.ReadFile("timemachineu8j.txt")
	if err != nil {
		b.Fail()
	}

	r := bytes.NewReader(src)
	dst := new(dummyWriter)
	dict := BiasCalculatorFunc(tinyseg.CalculateBias)

	seg := New(dst, r)

	b.ResetTimer()

	for idx := 0; idx < b.N; idx++ {
	loop:
		for {
			errS := seg.Segment(dict)

			if errS != nil && errS != io.EOF {
				b.Fail()
			}

			if errS == io.EOF {
				break loop
			}

			dst.Reset()
		}

		r.Reset(src)
		seg.Reset(r)
	}
}

func BenchmarkSegmentTextInBufIO(b *testing.B) {
	src, err := os.Open("timemachineu8j.txt")
	if err != nil {
		b.Fail()
	}

	r := bufio.NewReader(src)
	dst := new(dummyWriter)
	dict := BiasCalculatorFunc(tinyseg.CalculateBias)

	seg := New(dst, r)

	b.ResetTimer()

	for idx := 0; idx < b.N; idx++ {
	loop:
		for {
			errS := seg.Segment(dict)

			if errS != nil && errS != io.EOF {
				b.Fail()
			}

			if errS == io.EOF {
				break loop
			}

			dst.Reset()
		}

		src.Close()
		src, err = os.Open("timemachineu8j.txt")
		if err != nil {
			b.Fail()
		}

		r.Reset(src)
		seg.Reset(r)
	}
}
