package segmenter

import (
	"io"
	"unicode/utf8"

	char "github.com/nyarla/go-japanese-segmenter/chartypes"
	dict "github.com/nyarla/go-japanese-segmenter/dictionary"
)

const (
	e1 = 0x110001
	e2 = 0x110002
	e3 = 0x110003
	e4 = 0x110004
	e5 = 0x110005
	e6 = 0x110006
)

type Segmenter interface {
	Segment(*dict.Dictionary) (string, error)
	Reset()
}

type segmenter struct {
	src                       string
	token                     string
	count, offset, start, end int
	p1, p2, p3                rune
	r1, r2, r3, r4, r5, r6    rune
	t1, t2, t3, t4, t5, t6    rune
	i4, i5, i6, last          int
}

func New(src string) Segmenter {
	this := new(segmenter)
	this.src = src
	this.last = len(this.src)

	this.Reset()

	return this
}

func (this *segmenter) Reset() {
	this.token = ""

	this.count = 0
	this.offset = 0
	this.start = 0
	this.end = 0

	this.p1 = char.Uchar
	this.p2 = char.Uchar
	this.p3 = char.Uchar

	this.r1 = e1
	this.r2 = e2
	this.r3 = e3
	this.r4 = 0
	this.r5 = 0
	this.r6 = 0

	this.t1 = char.Ochar
	this.t2 = char.Ochar
	this.t3 = char.Ochar
	this.t4 = 0
	this.t5 = 0
	this.t6 = 0

	this.i4 = 0
	this.i5 = 0
	this.i6 = 0
}

func (this *segmenter) Segment(dict *dict.Dictionary) (string, error) {
	for pos, r := range this.src[this.offset:] {
		idx := pos
		if this.count > 0 && this.offset != 0 {
			idx += this.offset
		}
		switch this.count {
		case 0:
			this.r4 = r
			this.t4 = char.CharTypeAt(this.r4)
			this.i4 = idx
		case 1:
			this.r5 = r
			this.t5 = char.CharTypeAt(this.r5)
			this.i5 = idx
		case 2:
			this.r6 = r
			this.t6 = char.CharTypeAt(this.r6)
			this.i6 = idx
		default:
			this.r1 = this.r2
			this.r2 = this.r3
			this.r3 = this.r4
			this.r4 = this.r5
			this.r5 = this.r6
			this.r6 = r

			this.t1 = this.t2
			this.t2 = this.t3
			this.t3 = this.t4
			this.t4 = this.t5
			this.t5 = this.t6
			this.t6 = char.CharTypeAt(this.r6)

			this.i4 = this.i5
			this.i5 = this.i6
			this.i6 = idx
		}

		this.count += 1

		if this.count < 3 {
			continue
		}

		bias := dict.CalculateBias(
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

		if bias > 0 {
			this.p1 = this.p2
			this.p2 = this.p3
			this.p3 = char.Bchar

			this.end = this.i4
			if this.end != 0 {
				this.token = this.src[this.start:this.end]
			}

			this.start = this.end

			if this.token != "" {
				this.offset = this.i6 + utf8.RuneLen(this.r6)
				return this.token, nil
			}
		} else {
			this.p1 = this.p2
			this.p2 = this.p3
			this.p3 = char.Ochar
		}
	}

	for this.r6 != e6 {
		this.r1 = this.r2
		this.r2 = this.r3
		this.r3 = this.r4
		this.r4 = this.r5
		this.r5 = this.r6

		switch this.r5 {
		case e4:
			this.r6 = e5
		case e5:
			this.r6 = e6
		default:
			this.r6 = e4
		}

		this.t1 = this.t2
		this.t2 = this.t3
		this.t3 = this.t4
		this.t4 = this.t5
		this.t5 = this.t6
		this.t6 = char.Ochar

		this.i4 = this.i5
		this.i5 = this.i6
		this.i6 = this.last

		bias := dict.CalculateBias(
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

		if bias > 0 {
			this.p1 = this.p2
			this.p2 = this.p3
			this.p3 = char.Bchar

			this.end = this.i4
			if this.end != 0 {
				this.token = this.src[this.start:this.end]
			}

			this.start = this.end
			return this.token, nil
		} else {
			this.p1 = this.p2
			this.p2 = this.p3
			this.p3 = char.Ochar
		}
	}

	this.token = this.src[this.start:this.last]

	if this.token != "" {
		return this.token, nil
	}

	return "", io.EOF
}
