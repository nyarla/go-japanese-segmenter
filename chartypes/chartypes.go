package chartypes

const (
	Ochar rune = '0'
	Mchar rune = 'M'
	Hchar rune = 'H'
	Ichar rune = 'I'
	Kchar rune = 'K'
	Achar rune = 'A'
	Nchar rune = 'N'
	Bchar rune = 'B'
	Uchar rune = 'U'
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

func CharTypeAt(c rune) rune {
	switch c {
	case M1, M2, M3, M4, M5, M6, M7, M8, M9, M10, M11, M12, M13, M14, M15:
		return Mchar
	case H2, H3, H4, H5:
		return Hchar
	case K2, K4, K5:
		return Kchar
	}

	if N1a <= c && c <= N1Z {
		return Nchar
	}

	if A2a <= c && c <= A2Z {
		return Achar
	}

	if A1a <= c && c <= A1Z {
		return Achar
	}

	if I1a <= c && c <= I1Z {
		return Ichar
	}

	if H1a <= c && c <= H1Z {
		return Hchar
	}

	if K1a <= c && c <= K1Z {
		return Kchar
	}

	if N2a <= c && c <= N2Z {
		return Nchar
	}

	if A4a <= c && c <= A4Z {
		return Achar
	}

	if A3a <= c && c <= A3Z {
		return Achar
	}

	if K3a <= c && c <= K3Z {
		return Kchar
	}

	return Ochar
}
