package apartment

import (
	"fmt"
	"strings"
)

type Apartment struct {
	View           string   `json:"view"`
	Сost           int      `json:"cost"`
	Area           float64  `json:"area"`
	CardinalPoints string   `json:"cardinalPoints"`
	Floor          int      `json:"floor"`
	WindowsNum     int      `json:"windowsNum"`
	IsRamadable    bool     `json:"isRamadable"`
	Links          []string `json:"links"`
	Plan           string   `json:"plan"`
}

func (a Apartment) String() string {
	var ir string
	if a.IsRamadable {
		ir = "возможна"
	} else {
		ir = "не возможна"
	}

	return fmt.Sprintf(
		"![](%s)\n- **Сылки:** %s\n- **Цена:** %d р.\n- **Площадь:** %.2f кв.м\n- **Цена за метр:** %.2f\n- **Вид:** %s\n- **Окна выходят на** %s\n- **Этаж:** %d\n- **Количество окон:** %d\n- **Перепланировка:** %s",
		a.Plan,
		strings.Join(a.Links, ", "),
		a.Сost,
		a.Area,
		float64(a.Сost)/a.Area,
		a.View,
		a.CardinalPoints,
		a.Floor,
		a.WindowsNum,
		ir,
	)
}

type Apartments []Apartment

func (as Apartments) String() string {
	s := []string{}
	for _, a := range as {
		s = append(s, a.String())
	}

	return strings.Join(s, "\n")
}

func (as Apartments) Load() {}
