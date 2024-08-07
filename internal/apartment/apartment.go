package apartment

import (
	"fmt"
	"strings"
)

type Apartment struct {
	Tags       []string `json:"tags"`
	View       string   `json:"view"`
	Cost       int      `json:"cost"`
	Area       float64  `json:"area"`
	Side       string   `json:"side"`
	Floor      int      `json:"floor"`
	WindowsNum int      `json:"windowsNum"`
	Ramadable  string   `json:"ramadable"`
	Links      []string `json:"links"`
	Plan       string   `json:"plan"`
}

func (a Apartment) String() string {
	return fmt.Sprintf(
		"![](%s)\n- **Сылки:** %s\n- **Цена:** %d р.\n- **Площадь:** %.2f кв.м\n- **Цена за метр:** %.2f\n- **Вид:** %s\n- **Окна выходят на** %s\n- **Этаж:** %d\n- **Количество окон:** %d\n- **Можно перепланировать:** %s\n- **Метки** %s",
		a.Plan,
		strings.Join(a.Links, ", "),
		a.Cost,
		a.Area,
		float64(a.Cost)/a.Area,
		a.View,
		a.Side,
		a.Floor,
		a.WindowsNum,
		a.Ramadable,
		strings.Join(a.Tags, ", "),
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
