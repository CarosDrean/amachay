package measure

import (
	"github.com/CarosDrean/amachay/kit/db"
	"github.com/CarosDrean/amachay/model"
)

var memory map[int] model.Measure

type Measure struct {
	db db.DB
}

func NewMeasure(db db.DB) *Measure {
	memory = make(map[int] model.Measure)

	return &Measure{
		db: db,
	}
}

func (m Measure) Create(measure model.Measure) error {
	memory[measure.ID] = measure
	return nil
}

func (m Measure) Update(ID int, measure model.Measure) error {
	memory[ID] = measure
	return nil
}

func (m Measure) Delete(ID int) error {
	delete(memory, ID)
	return nil
}

func (m Measure) Get(ID int) (model.Measure, error) {
	data := memory[ID]
	return data, nil
}

func (m Measure) GetAll() (model.Measures, error) {
	data := model.Measures{}

	for _, value := range memory {
		data = append(data, value)
	}
	return data, nil
}
