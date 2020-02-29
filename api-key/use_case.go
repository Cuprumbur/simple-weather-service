package apikey

import (
	"github.com/Cuprumbur/weather-service/model"
)

type UseCase interface {
	CheckAccess(apiKey string, detectorID int, scope string) (err error)
	FindApiKeys(idDetector int) (keys []*model.ApiKey, err error)
	FindAllApiKeys() (keys []*model.ApiKey, err error)
	UpdateScopes(keyHash string, scopes []string) (err error)
	Delete(keyHash string) (err error)
	CreateApiKey(detectorID int, scopes []string) (apiKey string, err error)
}
