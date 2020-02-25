package detector

import "github.com/Cuprumbur/weather-service/model"

type UseCase interface {
	FindDetector(ID int) (d *model.Detector, err error)
	FindAllDetectors() (detectors []*model.Detector, err error)
	Update(d *model.Detector) (err error)
	Delete(id int) (err error)
	Store(d *model.Detector) (id int64, err error)
}