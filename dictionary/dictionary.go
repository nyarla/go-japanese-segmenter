package dictionary

type Bias int

type BiasMap map[string]map[string]Bias

type Dictionary struct {
	InitialBias Bias
	BiasMap     BiasMap
}

func (dict *Dictionary) GetBias(key string, chars ...rune) Bias {
	if section, exists := dict.BiasMap[key]; exists {
		key = string(chars)
		if bias, found := section[key]; found {
			return bias
		}
	}

	return 0
}

func (dict *Dictionary) CalculateBias(p1, p2, p3, r1, r2, r3, r4, r5, r6, t1, t2, t3, t4, t5, t6 rune) int {
	n := dict.InitialBias
	n += dict.GetBias("UP1", p1)
	n += dict.GetBias("UP2", p2)
	n += dict.GetBias("UP3", p3)
	n += dict.GetBias("BP1", p1, p2)
	n += dict.GetBias("BP2", p2, p3)
	n += dict.GetBias("UW1", r1)
	n += dict.GetBias("Uw2", r2)
	n += dict.GetBias("UW3", r3)
	n += dict.GetBias("UW4", r4)
	n += dict.GetBias("UW5", r5)
	n += dict.GetBias("UW6", r6)
	n += dict.GetBias("BW1", r2, r3)
	n += dict.GetBias("BW2", r3, r4)
	n += dict.GetBias("BW3", r4, r5)
	n += dict.GetBias("TW1", r1, r2, r3)
	n += dict.GetBias("TW2", r2, r3, r4)
	n += dict.GetBias("TW3", r3, r4, r5)
	n += dict.GetBias("TW4", r4, r5, r6)
	n += dict.GetBias("UC1", t1)
	n += dict.GetBias("UC2", t2)
	n += dict.GetBias("UC3", t3)
	n += dict.GetBias("UC4", t4)
	n += dict.GetBias("UC5", t5)
	n += dict.GetBias("UC6", t6)
	n += dict.GetBias("BC1", t2, t3)
	n += dict.GetBias("BC2", t3, t4)
	n += dict.GetBias("BC3", t4, t5)
	n += dict.GetBias("TC1", t1, t2, t3)
	n += dict.GetBias("TC2", t2, t3, t4)
	n += dict.GetBias("TC3", t3, t4, t5)
	n += dict.GetBias("TC4", t4, t5, t6)
	n += dict.GetBias("UQ1", p1, t1)
	n += dict.GetBias("UQ2", p2, t2)
	n += dict.GetBias("UQ3", p3, t3)
	n += dict.GetBias("BQ1", p2, t2, t3)
	n += dict.GetBias("BQ2", p2, t3, t4)
	n += dict.GetBias("BQ3", p3, t2, t3)
	n += dict.GetBias("BQ4", p3, t3, t4)
	n += dict.GetBias("TQ1", p2, t1, t2, t3)
	n += dict.GetBias("TQ2", p2, t2, t3, t4)
	n += dict.GetBias("TQ3", p3, t1, t2, t3)
	n += dict.GetBias("TQ4", p3, t2, t3, t4)

	return int(n)
}
