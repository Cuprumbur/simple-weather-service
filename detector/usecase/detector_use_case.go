package detector

import (
	"github.com/Cuprumbur/weather-service/detector"
	"github.com/Cuprumbur/weather-service/model"
)

type detectorUseCase struct {
	r detector.Repository
}

func NewDetectorUseCase(r detector.Repository) detector.UseCase {
	return &detectorUseCase{r}
}

func (u *detectorUseCase) FindDetector(id int) (d *model.Detector, err error) {
	return u.r.FindDetector(id);
}

func (u *detectorUseCase) FindAllDetectors() (detectors []*model.Detector, err error) {
	return u.r.FindAllDetectors()
}

func (u *detectorUseCase) Update(d *model.Detector) (err error) {
	return u.r.Update(d)
}

func (u *detectorUseCase) Delete(id int) (err error) {
	return u.r.Delete(id)
}

func (u *detectorUseCase) Store(d *model.Detector) (id int64, err error) {
	return u.r.Store(d)
}
