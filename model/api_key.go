package model

import "encoding/json"

type ApiKey struct {
	ID         int
	HashKey    string `json:"-"`
	Scopes     []string
	DetectorID int
}

func (k *ApiKey) MarshalJSON() ([]byte, error) {
	runes := []rune(k.HashKey)
	prefix := string(runes[0:5])
	return json.Marshal(&struct {
		ID         int    `json:"id"`
		Prefix     string `json:"prefix"`
		Scopes     []string `json:"scopes"`
		DetectorID int    `json:"detector_id"`
	}{
		ID:         k.ID,
		Prefix:     prefix,
		Scopes:     k.Scopes,
		DetectorID: k.DetectorID,
	})
}
