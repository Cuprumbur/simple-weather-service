package apikey

import (
	"errors"
	"math/rand"
	"time"

	apikey "github.com/Cuprumbur/weather-service/api-key"
	detector "github.com/Cuprumbur/weather-service/detector"
	"github.com/Cuprumbur/weather-service/model"
	"golang.org/x/crypto/bcrypt"
)

type apiKeyUseCase struct {
	detectorRepo detector.Repository
	apikeyRepo   apikey.Repository
}

func init() {
	rand.Seed(time.Now().Unix())
}

func NewApiKeyUseCase(detectorRepo detector.Repository, apikeyRepo apikey.Repository) apikey.UseCase {
	return &apiKeyUseCase{detectorRepo, apikeyRepo}
}

func (u *apiKeyUseCase) CheckAccess(apiKey string, detectorID int, scope string) (err error) {

	return
}

func (u *apiKeyUseCase) FindApiKeys(idDetector int) (keys []*model.ApiKey, err error) {
	return
}

func (u *apiKeyUseCase) FindAllApiKeys() (keys []*model.ApiKey, err error) {
	return
}

func (u *apiKeyUseCase) UpdateScopes(keyHash string, scopes []string) (err error) {
	return
}

func (u *apiKeyUseCase) Delete(keyHash string) (err error) {
	return
}

func (u *apiKeyUseCase) CreateApiKey(detectorID int, scopes []string) (string, error) {
	if scopes == nil || len(scopes) == 0 {

		return "", errors.New("scopes cannot be empty")
	}

	d, err := u.detectorRepo.FindDetector(detectorID)
	if err != nil {
		return "", err
	}

	if d == nil {
		return "", errors.New("detector does not exist")
	}

	key := generateKey()
	keyHash, err := calcHash(key)
	if err != nil {
		return "", err
	}

	modelKey := model.ApiKey{
		KeyHash:    keyHash,
		Scopes:     scopes,
		DetectorID: detectorID,
	}
	prefix, err := getPrefix(key)
	if err != nil {
		return "", err
	}

	err = u.apikeyRepo.Store(modelKey)
	if err != nil {
		return "", err
	}

	apikey := prefix + "." + key

	return apikey, nil
}

func calcHash(apiKey string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateKey() string {
	b := make([]rune, 25)
	length := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}

	return string(b)
}

func getPrefix(key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	runes := []rune(key)

	prefix := string(runes[0:5])
	return prefix, nil
}
