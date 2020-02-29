package apikey

import (
	"errors"
	"testing"

	"github.com/Cuprumbur/weather-service/api-key/mocks"
	detectorMock "github.com/Cuprumbur/weather-service/detector/mocks"
	"github.com/Cuprumbur/weather-service/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

)

func TestCreatApiKey(t *testing.T) {
	t.Run("success creating api key", func(t *testing.T) {

		detectorRepo := &detectorMock.Repository{}
		detectorID := 1
		d := &model.Detector{
			ID: detectorID,
		}
		detectorRepo.On("FindDetector", detectorID).Return(d, nil)
		apikeyRepo := &mocks.Repository{}
		apikeyRepo.On("Store", mock.Anything).Return(nil)
		useCase := NewApiKeyUseCase(detectorRepo, apikeyRepo)

		scopes := []string{"data"}
		apiKey, err := useCase.CreateApiKey(detectorID, scopes)

		assert.NotEmpty(t, apiKey)
		assert.Nil(t, err)
	})

	t.Run("error failed", func(t *testing.T) {

		detectorRepo := &detectorMock.Repository{}
		errMsg := "not found"

		detectorRepo.On("FindDetector", mock.Anything).Return(nil, errors.New(errMsg))
		
		apikeyRepo := &mocks.Repository{}
		useCase := NewApiKeyUseCase(detectorRepo, apikeyRepo)

		detectorID := 1
		scopes := []string{"data"}
		apiKey, err := useCase.CreateApiKey(detectorID, scopes)

		assert.Empty(t, apiKey)
		assert.EqualError(t, err, errMsg)
	})


	t.Run("error failed Detector not found", func(t *testing.T) {

		detectorRepo := &detectorMock.Repository{}

		detectorRepo.On("FindDetector", mock.Anything).Return(nil, nil)
		
		apikeyRepo := &mocks.Repository{}
		useCase := NewApiKeyUseCase(detectorRepo, apikeyRepo)

		detectorID := 1
		scopes := []string{"data"}
		apiKey, err := useCase.CreateApiKey(detectorID, scopes)

		assert.Empty(t, apiKey)

		assert.EqualError(t, err, "detector does not exist")
	})

	t.Run("scopes cannot be empty", func(t *testing.T) {

		detectorRepo := &detectorMock.Repository{}
		apikeyRepo := &mocks.Repository{}
		u := NewApiKeyUseCase(detectorRepo, apikeyRepo)

		detectorID := 1
		emptyScopes := []string{}
		apiKey, err := u.CreateApiKey(detectorID, emptyScopes)

		assert.Empty(t, apiKey)
		assert.EqualError(t, err, "scopes cannot be empty")
	})
}
