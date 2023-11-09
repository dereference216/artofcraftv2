package utils

import (
	"sync"

	tls_client "github.com/bogdanfinn/tls-client"
)

type Account struct {
	Email      string
	Password   string
	AuthToken  string
	Client     *tls_client.HttpClient
	Config     *Config
	XApiKey    string
	Etag       string
	ClaimLink  string
	ThreadLock *sync.Mutex
}

type Config struct {
	CapSolverKey string `json:"CAPkey"`
	Debug        bool   `json:"Debug"`
	Threads      int    `json:"Threads"`
	RequestRetry int    `json:"RequestRetry"`
}
