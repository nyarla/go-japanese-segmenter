package segmenter

import (
	"io"
	"unicode/utf8"

	char "github.com/nyarla/go-japanese-segmenter/chartypes"
)

const (
	b1 = 0x110001
	b2 = 0x110002
	b3 = 0x110003
	e1 = 0x110004
)

type BiasCalculator interface {
	CalculateBias(p1, p2, p3, r1, r2, r3, r4, r5, r6, t1, t2, t3, t4, t5, t6 rune) int64
}

type BiasCalculatorFunc func(p1, p2, p3, r1, r2, r3, r4, r5, r6, t1, t2, t3, t4, t5, t6 rune) int64

func (calc BiasCalculatorFunc) CalculateBias(p1, p2, p3, r1, r2, r3, r4, r5, r6, t1, t2, t3, t4, t5, t6 rune) int64 {
	return calc(p1, p2, p3, r1, r2, r3, r4, r5, r6, t1, t2, t3, t4, t5, t6)
}

type Segmenter interface {
	Segment(BiasCalculator) error
	Reset(src io.RuneReader)
}

type segmenter struct {
	src                    io.RuneReader
	dst                    io.Writer
	buf                    []byte
	p1, p2, p3             rune
	r1, r2, r3, r4, r5, r6 rune
	t1, t2, t3, t4, t5, t6 rune
}

func New(dst io.Writer, src io.RuneReader) Segmenter {
	this := new(segmenter)
	this.dst = dst
	this.buf = make([]byte, 4)

	this.Reset(src)

	return this
}

func (this *segmenter) Reset(src io.RuneReader) {
	this.src = src

	this.buf[0] = 0
	this.buf[1] = 0
	this.buf[2] = 0
	this.buf[3] = 0

	this.p1 = char.Uchar
	this.p2 = char.Uchar
	this.p3 = char.Uchar

	this.r1 = b3
	this.r2 = b2
	this.r3 = b1
	this.r4 = 0
	this.r5 = 0
	this.r6 = 0

	this.t1 = char.Ochar
	this.t2 = char.Ochar
	this.t3 = char.Ochar
	this.t4 = 0
	this.t5 = 0
	this.t6 = 0
}

func (this *segmenter) Segment(calc BiasCalculator) error {
	for {
		var t rune
		r, _, err := this.src.ReadRune()

		if err != nil {
			if err != io.EOF {
				return err
			}

			if this.r6 < e1 {
				r = e1
			} else {
				r = this.r6 + 1
			}

			t = char.Ochar
		}

		t = char.CharTypeAt(r)

		switch {
		case this.r4 == 0:
			this.r4 = r
			this.t4 = t
			continue
		case this.r5 == 0:
			this.r5 = r
			this.t5 = t
			continue
		case this.r6 == 0:
			// do nothing
		default:
			this.r1, this.r2, this.r3, this.r4, this.r5 = this.r2, this.r3, this.r4, this.r5, this.r6
			this.t1, this.t2, this.t3, this.t4, this.t5 = this.t2, this.t3, this.t4, this.t5, this.t6
		}

		this.r6 = r
		this.t6 = t

		if this.r3 < b1 {
			size := utf8.EncodeRune(this.buf, this.r3)
			if _, err = this.dst.Write(this.buf[0:size]); err != nil {
				return err
			}
		}

		bias := calc.CalculateBias(
			this.p1,
			this.p2,
			this.p3,
			this.r1,
			this.r2,
			this.r3,
			this.r4,
			this.r5,
			this.r6,
			this.t1,
			this.t2,
			this.t3,
			this.t4,
			this.t5,
			this.t6,
		)

		this.p1, this.p2 = this.p2, this.p3

		if bias > 0 {
			this.p3 = char.Bchar

			if this.r3 < b1 {
				return nil
			}
		} else {
			this.p3 = char.Ochar
		}

		if this.r3 >= e1 {
			break
		}
	}

	return io.EOF
}
