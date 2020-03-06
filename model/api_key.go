package model

type ApiKey struct {
	ID         int
	Prefix     string
	HashKey    string `json:"-"`
	Scopes     []string
	DetectorID int
}
