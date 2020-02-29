package apikey

import "github.com/Cuprumbur/weather-service/model"

type Repository interface {
	FindApiKey(keyHash string) (key model.ApiKey, err error)
	FindApiKeys(idDetector int) (keys []model.ApiKey, err error)
	FindAllApiKeys() (keys []model.ApiKey, err error)
	UpdateScopes(keyHash string, scopes []string) (err error)
	Delete(keyHash string) (err error)
	Store(key model.ApiKey) (err error)
}