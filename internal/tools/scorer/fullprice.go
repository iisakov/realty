package scorer

import (
	"fmt"
	"realty/internal/developer"
)

func (s Scorer) FullPrice(ds developer.Developers) {
	for _, d := range ds {
		for _, r := range d.Residentials {
			for _, a := range r.Apartments {
				area := a.Area
				cost := a.Cost
				cPa := float64(a.Cost) / a.Area
				divCPa := 0.0
				if cPa > s.AvgAreaCost {
					divCPa = cPa - s.AvgAreaCost
				}
				initialPayment20 := float64(a.Cost) * 0.20
				initialPayment30 := float64(a.Cost) * 0.30

				fmt.Printf(
					"\nПлощадь: %.2f,\nЦена: %d,\nЗа метр: %.2f,\nПереплата за квадрат: %.2f,\nПервоначальный взнос 20: %.2f + %.2f Переплата - %.2f,\nПервоначальный взнос 30: %.2f + %.2f Переплата - %.2f\n",
					area,
					cost,
					cPa,
					divCPa,
					initialPayment20,
					divCPa*area,
					initialPayment20+divCPa*area,
					initialPayment30,
					divCPa*area,
					initialPayment30+divCPa*area,
				)
			}
		}
	}
}
