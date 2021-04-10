package measure

import "github.com/CarosDrean/amachay/model"

type UseCase interface {
	Create(measure model.Measure) error
	Update(ID int, measure model.Measure) error
	Delete(ID int) error
	Get(ID int) (model.Measure, error)
	GetAll() (model.Measures, error)
}

type Storage interface {
	Create(measure model.Measure) error
	Update(ID int, measure model.Measure) error
	Delete(ID int) error
	Get(ID int) (model.Measure, error)
	GetAll() (model.Measures, error)
}
