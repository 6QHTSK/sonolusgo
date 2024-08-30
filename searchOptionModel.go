package sonolusgo

import (
	"fmt"
	"strconv"
)

type ServerForm struct {
	Type                string         `json:"type"`
	Title               string         `json:"title"`
	Icon                string         `json:"icon,omitempty"`
	Description         string         `json:"description,omitempty"`
	Help                string         `json:"help,omitempty"`
	RequireConfirmation bool           `json:"requireConfirmation"`
	Options             []ServerOption `json:"options"`
}

type ServerOption interface {
	GetQuery() string
	GetValueStr(queryResult string) string
}

type ServerTextOption struct {
	Query       string   `json:"query"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Required    bool     `json:"required"`
	Type        string   `json:"type"`
	Def         string   `json:"def"`
	PlaceHolder string   `json:"placeHolder"`
	Limit       int      `json:"limit"`
	Shortcuts   []string `json:"shortcuts"`
}

func (o ServerTextOption) GetQuery() string {
	return o.Query
}

func (o ServerTextOption) GetValueStr(queryResult string) string {
	return queryResult
}

func NewServerTextOption(query string, name string, description string, required bool, def string, placeHolder string, limit int, shortcuts []string) ServerTextOption {
	return ServerTextOption{
		Query:       query,
		Name:        name,
		Description: description,
		Required:    required,
		Type:        "text",
		Def:         def,
		PlaceHolder: placeHolder,
		Limit:       limit,
		Shortcuts:   shortcuts,
	}
}

type ServerSliderOption struct {
	Query       string  `json:"query"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Required    bool    `json:"required"`
	Type        string  `json:"type"`
	Def         float64 `json:"def"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
	Step        float64 `json:"step"`
	Unit        string  `json:"unit"`
}

func (o ServerSliderOption) GetQuery() string {
	return o.Query
}

func (o ServerSliderOption) GetValueStr(queryResult string) string {
	value, err := strconv.ParseFloat(queryResult, 64)
	if err != nil {
		return fmt.Sprintf("%f", o.Def)
	}
	if value < o.Min || value > o.Max {
		return fmt.Sprintf("%f", o.Def)
	}
	return queryResult
}

func NewServerSliderOption(query string, name string, description string, required bool, def float64, min float64, max float64, step float64, unit string) ServerSliderOption {
	return ServerSliderOption{
		Query:       query,
		Name:        name,
		Description: description,
		Required:    required,
		Type:        "slider",
		Def:         def,
		Min:         min,
		Max:         max,
		Step:        step,
		Unit:        unit,
	}
}

type ServerToggleOption struct {
	Query       string `json:"query"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	Def         int    `json:"def"`
}

func (o ServerToggleOption) GetQuery() string {
	return o.Query
}

func (o ServerToggleOption) GetValueStr(queryResult string) string {
	value, err := strconv.ParseFloat(queryResult, 64)
	if err != nil {
		return strconv.Itoa(o.Def)
	}
	if value != 0 && value != 1 {
		return strconv.Itoa(o.Def)
	}
	return queryResult
}

func NewServerToggleOption(query string, name string, description string, required bool, def int) ServerToggleOption {
	if def != 0 && def != 1 {
		panic("Toggle should be 0 or 1")
	}
	return ServerToggleOption{
		Query:       query,
		Name:        name,
		Description: description,
		Required:    required,
		Type:        "toggle",
		Def:         def,
	}
}

type ServerSelectOption struct {
	Query       string   `json:"query"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Required    bool     `json:"required"`
	Type        string   `json:"type"`
	Def         int      `json:"def"`
	Values      []string `json:"values,omitempty"`
}

func (o ServerSelectOption) GetQuery() string {
	return o.Query
}

func (o ServerSelectOption) GetValueStr(queryResult string) string {
	index, err := strconv.Atoi(queryResult)
	if err != nil || index < 0 {
		index = 0
	} else if index >= len(o.Values) {
		index = len(o.Values) - 1
	}
	return o.Values[index]
}

func NewServerSelectOption(query string, name string, description string, required bool, def int, values []string) ServerSelectOption {
	if def >= len(values) {
		panic("Def exceed values list")
	}
	return ServerSelectOption{
		Query:       query,
		Name:        name,
		Description: description,
		Required:    required,
		Type:        "select",
		Def:         def,
		Values:      values,
	}
}
