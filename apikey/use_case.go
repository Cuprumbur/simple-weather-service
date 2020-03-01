package apikey

import (
	"github.com/Cuprumbur/weather-service/model"
)

type UseCase interface {
	FindApiKeys(detectorID int) (keys []*model.ApiKey, err error)
	FindAllApiKeys() (keys []*model.ApiKey, err error)
	UpdateScopes(id int, scopes []string) (err error)
	Delete(id int) (err error)
	CreateApiKey(detectorID int, scopes []string) (apiKey string, err error)
}
