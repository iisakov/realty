package scorer

type Scorable interface {
	Params() map[string]any
}

type Scorer struct {
	MaxCost           int     `json:"maxCost"`
	MaxInitialPayment int     `json:"maxInitialPayment"`
	AvgAreaCost       float64 `json:"avgAreaCost"`
	MinArea           float64 `json:"minArea"`
}

func (s Scorer) Load() {}
