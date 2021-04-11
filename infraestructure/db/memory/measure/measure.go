package measure

import (
	"fmt"
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
	_, exist := memory[measure.ID]
	if exist {
		return fmt.Errorf("ID exist")
	}

	memory[measure.ID] = measure
	return nil
}

func (m Measure) Update(ID int, measure model.Measure) error {
	_, exist := memory[ID]
	if !exist {
		return fmt.Errorf("ID no exist")
	}

	memory[ID] = measure
	return nil
}

func (m Measure) Delete(ID int) error {
	_, exist := memory[ID]
	if !exist {
		return fmt.Errorf("ID no exist")
	}

	delete(memory, ID)
	return nil
}

func (m Measure) Get(ID int) (model.Measure, error) {
	_, exist := memory[ID]
	if !exist {
		return model.Measure{}, nil
	}

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
