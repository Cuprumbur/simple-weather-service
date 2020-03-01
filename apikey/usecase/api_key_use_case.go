package apikey

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	apikey "github.com/Cuprumbur/weather-service/apikey"
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

func (u *apiKeyUseCase) FindApiKeys(detectorID int) (keys []*model.ApiKey, err error) {
	return u.apikeyRepo.FindApiKeys(detectorID)
}

func (u *apiKeyUseCase) FindAllApiKeys() (keys []*model.ApiKey, err error) {
	return u.apikeyRepo.FindAllApiKeys()
}

func (u *apiKeyUseCase) UpdateScopes(id int, scopes []string) error {
	if scopes == nil || len(scopes) == 0 {
		return errors.New("scopes cannot be empty")
	}

	for i := range scopes {
		scopes[i] = strings.TrimSpace(scopes[i])
	}

	return u.apikeyRepo.UpdateScopes(id, scopes)
}

func (u *apiKeyUseCase) Delete(id int) error {
	return u.apikeyRepo.Delete(id)
}

func (u *apiKeyUseCase) CreateApiKey(detectorID int, scopes []string) (string, error) {
	if scopes == nil || len(scopes) == 0 {
		return "", errors.New("scopes cannot be empty")
	}

	for i := range scopes {
		scopes[i] = strings.TrimSpace(scopes[i])
	}

	d, err := u.detectorRepo.FindDetector(detectorID)
	if err != nil {
		return "", err
	}

	if d == nil {
		return "", errors.New("detector does not exist")
	}

	key := generateKey()
	k, err := combinePrefixHash(key)
	if err != nil {
		return "", err
	}

	modelKey := model.ApiKey{
		HashKey:    k,
		Scopes:     scopes,
		DetectorID: detectorID,
	}

	err = u.apikeyRepo.Store(modelKey)
	if err != nil {
		return "", err
	}

	return key, nil
}

var (
	letters      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lengthPrefix = 5
	lengthKey    = 30
)

func generateKey() string {
	b := make([]rune, lengthKey)
	length := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}

	return string(b)
}

func combinePrefixHash(key string) (string, error) {
	prefix, err := getPrefix(key, lengthPrefix)
	if err != nil {
		return "", err
	}

	hash, err := calcHash(key)
	if err != nil {
		return "", err
	}

	return prefix + "." + hash, nil
}

func getPrefix(key string, length int) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	runes := []rune(key)

	prefix := string(runes[0:length])
	return prefix, nil
}

func calcHash(apiKey string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
