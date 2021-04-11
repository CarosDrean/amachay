package model

type Measure struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Measures []Measure
