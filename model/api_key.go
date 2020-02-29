package model

type ApiKey struct {
	KeyHash    string
	Prefix     string
	Scopes     []string
	DetectorID int
}

