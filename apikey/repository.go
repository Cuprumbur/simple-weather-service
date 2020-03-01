package apikey

import "github.com/Cuprumbur/weather-service/model"

type Repository interface {
	FindApiKey(id int) (key *model.ApiKey, err error)
	FindApiKeys(detectorID int) (keys []*model.ApiKey, err error)
	FindAllApiKeys() (keys []*model.ApiKey, err error)
	UpdateScopes(id int, scopes []string) (err error)
	Delete(id int) (err error)
	Store(key model.ApiKey) (err error)
}
