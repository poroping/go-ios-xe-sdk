package config

import "net/http"

type Config struct {
	Password  string
	Username  string
	UserAgent string
	Insecure  bool
	Host      string
	Schema    string
	HTTPCon   *http.Client
}
