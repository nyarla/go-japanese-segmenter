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

func CharTypeAt(r rune) rune {
	if 0xFF9F <= r {
		return Ochar
	}

	if 0xFF70 <= r {
		return Kchar
	}

	if 0xFF5B <= r {
		return Ochar
	}

	if 0xFF41 <= r {
		return Achar
	}

	if 0xFF3B <= r {
		return Ochar
	}

	if 0xFF21 <= r {
		return Achar
	}

	if 0xFF1A <= r {
		return Ochar
	}

	if 0xFF10 <= r {
		return Nchar
	}

	if 0x9FA1 <= r {
		return Ochar
	}

	if 0x4E00 <= r {
		if r == 0x767E {
			return Mchar
		}

		if r == 0x56DB {
			return Mchar
		}

		if r == 0x5343 {
			return Mchar
		}

		if r == 0x5341 {
			return Mchar
		}

		if r == 0x516D {
			return Mchar
		}

		if r == 0x516B {
			return Mchar
		}

		if r == 0x5146 {
			return Mchar
		}

		if r == 0x5104 {
			return Mchar
		}

		if r == 0x4E94 {
			return Mchar
		}

		if r == 0x4E8C {
			return Mchar
		}

		if r == 0x4E5D {
			return Mchar
		}

		if r == 0x4E09 {
			return Mchar
		}

		if r == 0x4E07 {
			return Mchar
		}

		if r == 0x4E03 {
			return Mchar
		}

		if r == 0x4E00 {
			return Mchar
		}

		return Hchar
	}

	if r == 0x30FC {
		return Kchar
	}

	if 0x30F7 <= r {
		return Ochar
	}

	if 0x30F5 <= r {
		return Hchar
	}

	if 0x30A1 <= r {
		return Kchar
	}

	if 0x3094 <= r {
		return Ochar
	}

	if 0x3041 <= r {
		return Ichar
	}

	if 0x3007 <= r {
		return Ochar
	}

	if 0x3005 <= r {
		return Hchar
	}

	if 0x7B <= r {
		return Ochar
	}

	if 0x61 <= r {
		return Achar
	}

	if 0x5B <= r {
		return Ochar
	}

	if 0x41 <= r {
		return Achar
	}

	if 0x3A <= r {
		return Ochar
	}

	if 0x30 <= r {
		return Nchar
	}

	return Ochar
}
