package measure

import (
	"fmt"
	"github.com/CarosDrean/amachay/model"
)

type Measure struct {
	storage Storage
}

func NewMeasure(storage Storage) *Measure{
	return &Measure{
		storage: storage,
	}
}

func (m Measure) Create(measure model.Measure) error {
	if err := m.validateDependencies(); err != nil {
		return err
	}

	err := m.storage.Create(measure)
	if err != nil {
		return err
	}

	return nil
}

func (m Measure) Update(ID int, measure model.Measure) error {
	if err := m.validateDependencies(); err != nil {
		return err
	}

	err := m.storage.Update(ID, measure)
	if err != nil {
		return err
	}

	return nil
}

func (m Measure) Delete(ID int) error {
	if err := m.validateDependencies(); err != nil {
		return err
	}

	err := m.storage.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (m Measure) Get(ID int) (model.Measure, error) {
	if err := m.validateDependencies(); err != nil {
		return model.Measure{}, err
	}

	data, err := m.storage.Get(ID)
	if err != nil {
		return model.Measure{}, err
	}

	return data, nil
}

func (m Measure) GetAll() (model.Measures, error) {
	if err := m.validateDependencies(); err != nil {
		return nil, err
	}

	data, err := m.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m Measure) validateDependencies() error {
	if m.storage == nil {
		return fmt.Errorf("storage is nil")
	}

	return nil
}
