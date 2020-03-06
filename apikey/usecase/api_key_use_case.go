package apikey

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	apikey "github.com/Cuprumbur/weather-service/apikey"
	detector "github.com/Cuprumbur/weather-service/detector"
	"github.com/Cuprumbur/weather-service/model"
	"github.com/google/uuid"
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

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	strID := id.String()
	prefix, err := subString(strID, ApiKeyPrefixLength)
	if err != nil {
		return "", err
	}

	hash, err := hash(strID)
	if err != nil {
		return "", err
	}

	modelKey := model.ApiKey{
		HashKey:    hash,
		Prefix:     prefix,
		Scopes:     scopes,
		DetectorID: detectorID,
	}

	err = u.apikeyRepo.Store(modelKey)
	if err != nil {
		return "", err
	}

	return strID, nil
}

var ApiKeyPrefixLength = 5

func subString(str string, length int) (string, error) {
	asRune := []rune(str)
	if len(asRune) < length {
		return "", errors.New("Length of str cannot be shorter then length of sub-string")
	}
	return string(asRune[:length]), nil
}

func hash(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
