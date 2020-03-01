package apikey

import (
	"database/sql"
	"strings"

	"github.com/Cuprumbur/weather-service/apikey"
	"github.com/Cuprumbur/weather-service/model"

)

type mySqlApiKeyRepository struct {
	db *sql.DB
}

func NewMySqlApiKeyRepository(db *sql.DB) apikey.Repository {
	return &mySqlApiKeyRepository{db}
}

func (r *mySqlApiKeyRepository) FindApiKey(id int) (apikey *model.ApiKey, err error) {
	rows, err := r.db.Query("SELECT * FROM api_key WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var hashKey, scopes string
		var detectorID int
		err = rows.Scan(&id, &hashKey, &scopes, &detectorID)
		if err != nil {
			return nil, err
		}

		apikey = &model.ApiKey{
			ID:         id,
			HashKey:    hashKey,
			Scopes:     strings.Split(scopes, ","),
			DetectorID: detectorID,
		}
		return apikey, nil
	}

	return nil, nil
}

func (r *mySqlApiKeyRepository) FindApiKeys(detectorID int) (keys []*model.ApiKey, err error) {
	rows, err := r.db.Query("SELECT * FROM api_key WHERE detector_id = ?", detectorID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var hashKey, scopes string
		var detectorID int
		err = rows.Scan(&id, &hashKey, &scopes, &detectorID)
		if err != nil {
			return nil, err
		}

		apikey := &model.ApiKey{
			ID:         id,
			HashKey:    hashKey,
			Scopes:     strings.Split(scopes, ","),
			DetectorID: detectorID,
		}
		keys = append(keys, apikey)
	}

	return keys, nil
}

func (r *mySqlApiKeyRepository) FindAllApiKeys() (keys []*model.ApiKey, err error) {
	rows, err := r.db.Query("SELECT * FROM api_key")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var hashKey, scopes string
		var detectorID int
		err = rows.Scan(&id, &hashKey, &scopes, &detectorID)
		if err != nil {
			return nil, err
		}

		apikey := &model.ApiKey{
			ID:         id,
			HashKey:    hashKey,
			Scopes:     strings.Split(scopes, ","),
			DetectorID: detectorID,
		}
		keys = append(keys, apikey)
	}

	return keys, nil
}

func (r *mySqlApiKeyRepository) UpdateScopes(id int, scopes []string) error {
	upd, err := r.db.Prepare("UPDATE api_key SET scopes = ? WHERE  id = ?")
	if err != nil {
		return err
	}

	_, err = upd.Exec(strings.Join(scopes, ","), id)
	return err
}

func (r *mySqlApiKeyRepository) Delete(id int) error {
	del, err := r.db.Prepare("DELETE FROM api_key WHERE id=?")
	if err != nil {
		return err
	}
	_, err = del.Exec(id)

	return err
}

func (r *mySqlApiKeyRepository) Store(key model.ApiKey) error {

	ins, err := r.db.Prepare("INSERT INTO api_key(hash_key, scopes, detector_id) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = ins.Exec(key.HashKey, strings.Join(key.Scopes, ","), key.DetectorID)
	return err
}
