package factor

type Factor struct {
	Type  string  `json:"type"`
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

type Factors []Factor

func (fs Factors) ByType(t string) (Factors, bool) {
	result, ok := Factors{}, false
	for _, f := range fs {
		if f.Type == t {
			result = append(result, f)
			ok = true
		}
	}
	return result, ok
}

func (fs Factors) ByLabel(l string) (float64, bool) {
	for _, f := range fs {
		if f.Label == l {
			return f.Score, true
		}
	}
	return 0, false
}

func (fs Factors) Load() {}
