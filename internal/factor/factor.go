package factor

type Factor struct {
	Factor string             `json:"factor"`
	Value  map[string]float64 `json:"value"`
	Type   string             `json:"type"`
}

func (f Factor) ByValue(v string) (float64, bool) {
	for fv, s := range f.Value {
		if fv == v {
			return s, true
		}
	}
	return 1, false
}

type Factors []Factor

func (fs Factors) FactorList() (result []string) {
	for _, f := range fs {
		result = append(result, f.Factor)
	}
	return result
}

func (fs Factors) ByFactor(factor string) (Factor, bool) {
	for _, f := range fs {
		if f.Factor == factor {
			return f, true
		}
	}
	return Factor{}, false
}

func (fs Factors) ByValue(v string) (float64, bool) {
	for _, f := range fs {
		if s, ok := f.ByValue(v); ok {
			return s, true
		}
	}
	return 1, false
}

func (fs Factors) Load() {}
