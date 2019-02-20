package chartypes

import (
	"testing"
)

func TestCharTypeAt(t *testing.T) {
	// Mchar group
	if CharTypeAt(M1) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M2) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M3) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M4) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M5) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M6) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M7) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M8) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M9) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M10) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M11) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M12) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M13) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M14) != Mchar {
		t.Fail()
	}
	if CharTypeAt(M15) != Mchar {
		t.Fail()
	}

	// Hchar group
	if CharTypeAt(H1Z) != Hchar {
		t.Fail()
	}
	if CharTypeAt(H2) != Hchar {
		t.Fail()
	}
	if CharTypeAt(H3) != Hchar {
		t.Fail()
	}
	if CharTypeAt(H4) != Hchar {
		t.Fail()
	}
	if CharTypeAt(H5) != Hchar {
		t.Fail()
	}

	// Ichar group
	if CharTypeAt(I1a) != Ichar {
		t.Fail()
	}
	if CharTypeAt(I1Z) != Ichar {
		t.Fail()
	}

	// Kchar group
	if CharTypeAt(K1a) != Kchar {
		t.Fail()
	}
	if CharTypeAt(K1Z) != Kchar {
		t.Fail()
	}
	if CharTypeAt(K2) != Kchar {
		t.Fail()
	}
	if CharTypeAt(K3a) != Kchar {
		t.Fail()
	}
	if CharTypeAt(K3Z) != Kchar {
		t.Fail()
	}
	if CharTypeAt(K4) != Kchar {
		t.Fail()
	}
	if CharTypeAt(K5) != Kchar {
		t.Fail()
	}

	// Achar group
	if CharTypeAt(A1a) != Achar {
		t.Fail()
	}
	if CharTypeAt(A1Z) != Achar {
		t.Fail()
	}
	if CharTypeAt(A2a) != Achar {
		t.Fail()
	}
	if CharTypeAt(A2Z) != Achar {
		t.Fail()
	}
	if CharTypeAt(A3a) != Achar {
		t.Fail()
	}
	if CharTypeAt(A3Z) != Achar {
		t.Fail()
	}
	if CharTypeAt(A4a) != Achar {
		t.Fail()
	}
	if CharTypeAt(A4Z) != Achar {
		t.Fail()
	}

	// Nchar group
	if CharTypeAt(N1a) != Nchar {
		t.Fail()
	}
	if CharTypeAt(N1Z) != Nchar {
		t.Fail()
	}
	if CharTypeAt(N2a) != Nchar {
		t.Fail()
	}
	if CharTypeAt(N2Z) != Nchar {
		t.Fail()
	}

	// Ochar group
	if CharTypeAt('ω') != Ochar {
		t.Fail()
	}
}

func BenchmarkCharTypeAt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Mchar group
		if CharTypeAt(M1) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M2) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M3) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M4) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M5) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M6) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M7) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M8) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M9) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M10) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M11) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M12) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M13) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M14) != Mchar {
			b.Fail()
		}
		if CharTypeAt(M15) != Mchar {
			b.Fail()
		}

		// Hchar group
		if CharTypeAt(H1Z) != Hchar {
			b.Fail()
		}
		if CharTypeAt(H2) != Hchar {
			b.Fail()
		}
		if CharTypeAt(H3) != Hchar {
			b.Fail()
		}
		if CharTypeAt(H4) != Hchar {
			b.Fail()
		}
		if CharTypeAt(H5) != Hchar {
			b.Fail()
		}

		// Ichar group
		if CharTypeAt(I1a) != Ichar {
			b.Fail()
		}
		if CharTypeAt(I1Z) != Ichar {
			b.Fail()
		}

		// Kchar group
		if CharTypeAt(K1a) != Kchar {
			b.Fail()
		}
		if CharTypeAt(K1Z) != Kchar {
			b.Fail()
		}
		if CharTypeAt(K2) != Kchar {
			b.Fail()
		}
		if CharTypeAt(K3a) != Kchar {
			b.Fail()
		}
		if CharTypeAt(K3Z) != Kchar {
			b.Fail()
		}
		if CharTypeAt(K4) != Kchar {
			b.Fail()
		}
		if CharTypeAt(K5) != Kchar {
			b.Fail()
		}

		// Achar group
		if CharTypeAt(A1a) != Achar {
			b.Fail()
		}
		if CharTypeAt(A1Z) != Achar {
			b.Fail()
		}
		if CharTypeAt(A2a) != Achar {
			b.Fail()
		}
		if CharTypeAt(A2Z) != Achar {
			b.Fail()
		}
		if CharTypeAt(A3a) != Achar {
			b.Fail()
		}
		if CharTypeAt(A3Z) != Achar {
			b.Fail()
		}
		if CharTypeAt(A4a) != Achar {
			b.Fail()
		}
		if CharTypeAt(A4Z) != Achar {
			b.Fail()
		}

		// Nchar group
		if CharTypeAt(N1a) != Nchar {
			b.Fail()
		}
		if CharTypeAt(N1Z) != Nchar {
			b.Fail()
		}
		if CharTypeAt(N2a) != Nchar {
			b.Fail()
		}
		if CharTypeAt(N2Z) != Nchar {
			b.Fail()
		}

		// Ochar group
		if CharTypeAt('ω') != Ochar {
			b.Fail()
		}
	}
}
