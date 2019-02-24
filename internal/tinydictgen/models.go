package tinydictgen

import (
	"sort"
)

type JSONData map[string]map[string]int64

func (json JSONData) Items() []*Item {
	var list = make(List, 0)

	for section, data := range json {
	biases:
		for chars, bias := range data {
			item := NewItem(section, chars, bias)

			for idx, target := range list {
				if merged, ok := Merge(target, item); ok {
					sort.Sort(merged.List)
					list[idx] = merged
					continue biases
				}
			}

			list = append(list, item)
		}
	}

	sort.Sort(list)
	return list
}

type List []*Item

func (list List) Len() int {
	return len(list)
}

func (list List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list List) Less(i, j int) bool {
	iV := []rune(list[i].Var)
	jV := []rune(list[j].Var)

	if iV[0] < jV[0] {
		return true
	}

	if iV[1] < jV[1] {
		return true
	}

	if list[i].Rune < list[j].Rune {
		return true
	}

	if list[i].Bias < list[j].Bias {
		return true
	}

	if len(list[i].List) < len(list[j].List) {
		return true
	}

	return false
}

type Item struct {
	Depth int
	Var   string
	Rune  rune
	Bias  int64
	List  List
}

func NewItem(section string, chars string, bias int64) *Item {
	var targets []string
	switch section {
	case "UP1":
		targets = []string{"p1"}
	case "UP2":
		targets = []string{"p2"}
	case "UP3":
		targets = []string{"p3"}

	case "UW1":
		targets = []string{"r1"}
	case "UW2":
		targets = []string{"r2"}
	case "UW3":
		targets = []string{"r3"}
	case "UW4":
		targets = []string{"r4"}
	case "UW5":
		targets = []string{"r5"}
	case "UW6":
		targets = []string{"r6"}

	case "UC1":
		targets = []string{"t1"}
	case "UC2":
		targets = []string{"t2"}
	case "UC3":
		targets = []string{"t3"}
	case "UC4":
		targets = []string{"t4"}
	case "UC5":
		targets = []string{"t5"}
	case "UC6":
		targets = []string{"t6"}

	case "BP1":
		targets = []string{"p1", "p2"}
	case "BP2":
		targets = []string{"p2", "p3"}

	case "BW1":
		targets = []string{"r2", "r3"}
	case "BW2":
		targets = []string{"r3", "r4"}
	case "BW3":
		targets = []string{"r4", "r5"}

	case "BC1":
		targets = []string{"t2", "t3"}
	case "BC2":
		targets = []string{"t3", "t4"}
	case "BC3":
		targets = []string{"t4", "t5"}

	case "UQ1":
		targets = []string{"p1", "t1"}
	case "UQ2":
		targets = []string{"p2", "t2"}
	case "UQ3":
		targets = []string{"p3", "t3"}

	case "TW1":
		targets = []string{"r1", "r2", "r3"}
	case "TW2":
		targets = []string{"r2", "r3", "r4"}
	case "TW3":
		targets = []string{"r3", "r4", "r5"}
	case "TW4":
		targets = []string{"r4", "r5", "r6"}

	case "TC1":
		targets = []string{"t1", "t2", "t3"}
	case "TC2":
		targets = []string{"t2", "t3", "t4"}
	case "TC3":
		targets = []string{"t3", "t4", "t5"}
	case "TC4":
		targets = []string{"t4", "t5", "t6"}

	case "BQ1":
		targets = []string{"p2", "t2", "t3"}
	case "BQ2":
		targets = []string{"p2", "t3", "t4"}
	case "BQ3":
		targets = []string{"p3", "t2", "t3"}
	case "BQ4":
		targets = []string{"p3", "t3", "t4"}

	case "TQ1":
		targets = []string{"p2", "t1", "t2", "t3"}
	case "TQ2":
		targets = []string{"p2", "t2", "t3", "t4"}
	case "TQ3":
		targets = []string{"p3", "t1", "t2", "t3"}
	case "TQ4":
		targets = []string{"p3", "t2", "t3", "t4"}
	}

	var runes = []rune(chars)
	if len(targets) < len(runes) {
		var tmp []rune

		if runes[0] == 'B' || runes[0] == 'Ｂ' {
			tmp = []rune{0x110001}
		}

		if runes[0] == 'E' || runes[0] == 'Ｅ' {
			tmp = []rune{0x110004}
		}

		if len(tmp) != 0 {
			if len(runes) >= 3 {
				tmp = append(tmp, runes[2:]...)
			}

			runes = tmp
		}
	}

	root := new(Item)
	root.Var = targets[0]
	root.Rune = runes[0]
	root.Depth = 0

	if len(targets) == 1 {
		root.Bias = bias
		return root
	}

	var current = root

	for depth := 1; depth < len(targets) && depth < len(runes); depth++ {
		nest := new(Item)
		nest.Var = targets[depth]
		nest.Rune = runes[depth]
		nest.Depth = depth

		if depth == len(targets)-1 {
			nest.Bias = bias
		}

		current.List = List{nest}
		current = nest
	}

	return root
}

func Merge(dst, src *Item) (*Item, bool) {
	switch {
	case src == nil && dst != nil:
		return dst, true
	case src != nil && dst == nil:
		return src, true
	case src == nil && dst == nil:
		return nil, false
	}

	if dst.Depth != src.Depth {
		return nil, false
	}

	if dst.Var != src.Var {
		return nil, false
	}

	if dst.Rune != src.Rune {
		return nil, false
	}

	switch {
	case len(dst.List) == 0 && len(src.List) != 0:
		dst.List = src.List
		sort.Sort(dst.List)
		return dst, true
	case len(dst.List) != 0 && len(src.List) == 0:
		dst.Bias = src.Bias
		return dst, true
	}

src:
	for _, srcItem := range src.List {
		for dstIdx, dstItem := range dst.List {
			if merged, ok := Merge(dstItem, srcItem); ok {
				sort.Sort(merged.List)
				dst.List[dstIdx] = merged
				continue src
			}
		}

		dst.List = append(dst.List, srcItem)
	}

	sort.Sort(dst.List)
	return dst, true
}
