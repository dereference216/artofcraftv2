package main

import (
	"artofcraft/utils"
	"encoding/json"
	"math/rand"
	"os"
	"sync"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func main() {
	var threadlock sync.Mutex
	var Config utils.Config
	var wg sync.WaitGroup
	config, err := os.ReadFile("./configs.json")
	if err != nil {
		utils.PrintError("Config", err.Error())
		return
	}
	if json.Unmarshal(config, &Config); err != nil {
		utils.PrintError("JsonUnarmshal (Config)", err.Error())
		return
	}
	proxies, err := utils.ReadFile("proxies.txt", &threadlock)
	if json.Unmarshal(config, &Config); err != nil {
		utils.PrintError("JsonUnarmshal (Config)", err.Error())
		return
	}
	if len(proxies) == 0 {
		utils.PrintError("proxies", "no proxies found")
		return
	}
	limiter := make(chan struct{}, Config.Threads)
	for {
		limiter <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				<-limiter
				wg.Done()
			}()
			jar := tls_client.NewCookieJar()
			options := []tls_client.HttpClientOption{
				tls_client.WithProxyUrl("http://" + proxies[rand.Intn(len(proxies))]),
				tls_client.WithNotFollowRedirects(),
				tls_client.WithClientProfile(profiles.Okhttp4Android9),
				tls_client.WithCookieJar(jar),
			}
			client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
			if err != nil {
				utils.PrintError("NEWHTTPCLIENT", err.Error())
				return
			}

			AccCreate := utils.Account{
				Email:    utils.RandomString(36),
				Password: utils.RandomPassword(),
				Client:   &client,
				Config:   &Config,
			}
			if err := AccCreate.SignUP(); err != nil {
				return
			}
			if err := AccCreate.Link(); err != nil {
				return
			}
			utils.Write("promos.txt", AccCreate.ClaimLink, &threadlock)
		}()
	}
}
