package detector

import "github.com/Cuprumbur/weather-service/model"

type Repository interface {
	FindDetector(id int) (d *model.Detector, err error)
	FindAllDetectors() (detectors []*model.Detector, err error)
	Update(d *model.Detector) (err error)
	Delete(id int) (err error)
	Store(d *model.Detector) (id int64, err error)
}
