package chartypes

import (
	"strconv"
	"testing"
)

const (
	M1  rune = '一'
	M2  rune = '二'
	M3  rune = '三'
	M4  rune = '四'
	M5  rune = '五'
	M6  rune = '六'
	M7  rune = '七'
	M8  rune = '八'
	M9  rune = '九'
	M10 rune = '十'
	M11 rune = '百'
	M12 rune = '千'
	M13 rune = '万'
	M14 rune = '億'
	M15 rune = '兆'
	H1a rune = '一'
	H1Z rune = '龠'
	H2  rune = '々'
	H3  rune = '〆'
	H4  rune = 'ヵ'
	H5  rune = 'ヶ'
	I1a rune = 'ぁ'
	I1Z rune = 'ん'
	K1a rune = 'ァ'
	K1Z rune = 'ヴ'
	K2  rune = 'ー'
	K3a rune = 'ｱ'
	K3Z rune = 'ﾝ'
	K4  rune = 'ﾞ'
	K5  rune = 'ｰ'
	A1a rune = 'a'
	A1Z rune = 'z'
	A2a rune = 'A'
	A2Z rune = 'Z'
	A3a rune = 'ａ'
	A3Z rune = 'ｚ'
	A4a rune = 'Ａ'
	A4Z rune = 'Ｚ'
	N1a rune = '0'
	N1Z rune = '9'
	N2a rune = '０'
	N2Z rune = '９'
)

func TestCharTypeAt(t *testing.T) {
	single := [][2]rune{
		{M1, Mchar},
		{M2, Mchar},
		{M3, Mchar},
		{M4, Mchar},
		{M5, Mchar},
		{M6, Mchar},
		{M7, Mchar},
		{M8, Mchar},
		{M9, Mchar},
		{M10, Mchar},
		{M11, Mchar},
		{M12, Mchar},
		{M13, Mchar},
		{M14, Mchar},
		{M15, Mchar},
		{K4, Kchar},
		{K5, Kchar},
		{H2, Hchar},
		{H3, Hchar},
		{H4, Hchar},
		{H5, Hchar},
		{K2, Kchar},
	}

	for _, test := range single {
		if p := CharTypeAt(test[0]); p != test[1] {
			t.Fatal(string(p), string(test[1]), "U+"+strconv.FormatInt(int64(test[0]), 16), string(test[0]))
		}
	}

	ranges := [][3]rune{
		{H1a, H1Z, Hchar},
		{I1a, I1Z, Ichar},
		{K1a, K1Z, Kchar},
		{K3a, K3Z, Kchar},
		{A1a, A1Z, Achar},
		{A2a, A2Z, Achar},
		{A3a, A3Z, Achar},
		{A4a, A4Z, Achar},
		{N1a, N1Z, Nchar},
		{N2a, N2Z, Nchar},
	}

	for _, test := range ranges {
	runes:
		for r := test[0]; r <= test[1]; r++ {
			switch r {
			case
				M1, M2, M3, M4, M5, M6, M7, M8, M9,
				M10, M11, M12, M13, M14, M15,
				K2, K4, K5,
				H2, H3,
				H4, H5:
				continue runes
			}

			if p := CharTypeAt(r); p != test[2] {
				t.Fatal(string(p), string(test[2]), "U+"+strconv.FormatInt(int64(r), 16), string(r))
			}
		}
	}

	for r := rune(0); r < 0x110000; r++ {
		switch r {
		case M1, M2, M3, M4, M5, M6, M7, M8, M9,
			M10, M11, M12, M13, M14, M15,
			K2, K4, K5,
			H2, H3,
			H4, H5:
			continue
		}

		c := CharTypeAt(r)

		switch {
		case H1a <= r && r <= H1Z:
			if c != Hchar {
				t.Fatal(string(c), string(Hchar), "U+"+strconv.FormatInt(int64(r), 16))
			}
		case I1a <= r && r <= I1Z:
			if c != Ichar {
				t.Fatal(string(c), string(Ichar), "U+"+strconv.FormatInt(int64(r), 16))
			}

		case K1a <= r && r <= K1Z:
			if c != Kchar {
				t.Fatal(string(c), string(Kchar), "U+"+strconv.FormatInt(int64(r), 16))
			}

		case K3a <= r && r <= K3Z:
			if c != Kchar {
				t.Fatal(string(c), string(Kchar), "U+"+strconv.FormatInt(int64(r), 16))
			}

		case A1a <= r && r <= A1Z:
			if c != Achar {
				t.Fatal(string(c), string(Achar), "U+"+strconv.FormatInt(int64(r), 16))
			}
		case A2a <= r && r <= A2Z:
			if c != Achar {
				t.Fatal(string(c), string(Achar), "U+"+strconv.FormatInt(int64(r), 16))
			}
		case A3a <= r && r <= A3Z:
			if c != Achar {
				t.Fatal(string(c), string(Achar), "U+"+strconv.FormatInt(int64(r), 16))
			}

		case A4a <= r && r <= A4Z:
			if c != Achar {
				t.Fatal(string(c), string(Achar), "U+"+strconv.FormatInt(int64(r), 16))
			}

		case N1a <= r && r <= N1Z:
			if c != Nchar {
				t.Fatal(string(c), string(Nchar), "U+"+strconv.FormatInt(int64(r), 16))
			}

		case N2a <= r && r <= N2Z:
			if c != Nchar {
				t.Fatal(string(c), string(Nchar), "U+"+strconv.FormatInt(int64(r), 16))
			}
		default:
			if c != Ochar {
				t.Fatal(string(c), string(Ochar), "U+"+strconv.FormatInt(int64(r), 16))
			}
		}
	}
}

func BenchmarkCharTypeAt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for r := rune(0); r < 0x10FFFF; r++ {
			CharTypeAt(r)
		}
	}
}
