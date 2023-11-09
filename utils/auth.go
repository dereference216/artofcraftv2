package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

func (Acc *Account) SignUP() error {
	Acc.Config.PrintDebug("Username", Acc.Email)
	Acc.Config.PrintDebug("Password", Acc.Password)
	var resp *http.Response
	var rawBytes []byte
	var data = strings.NewReader(`{"email":"` + Acc.Email + `","password":"` + Acc.Password + `","provider":"picsart"}`)
	req, err := http.NewRequest("POST", "https://api.picsart.com/users/signup.json", data)
	if err != nil {
		return fmt.Errorf("signup err:%s", err)
	}
	req.Header.Set("Host", "api.picsart.com")
	req.Header.Set("Is-Forced-Briteverify", "false")
	req.Header.Set("App", "com.picsart.studio")
	req.Header.Set("Is-Tablet", "0")
	req.Header.Set("Accept", "application/picsart-3.0+json")
	req.Header.Set("Country-Code", "US")
	req.Header.Set("User-Agent", "PicsArt-10.x")
	req.Header.Set("Platform", "android")
	req.Header.Set("Network", "WIFI")
	req.Header.Set("Os-Version", "9")
	req.Header.Set("Market", "google")
	req.Header.Set("Low-Memory", "0")
	req.Header.Set("Language-Code", "en")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return fmt.Errorf("SignUP err:%s", err)
	}
	for i := 0; i <= Acc.Config.RequestRetry; i++ {
		resp, err = (*Acc.Client).Do(req)
		if err != nil {
			if i == Acc.Config.RequestRetry {
				return fmt.Errorf("TimedOut (SignUP) (MaxRetry)")
			}
			Acc.Config.PrintWarn(fmt.Sprintf("TimeOut (SignUP) (%d)", i), err.Error())
			continue
		}
		break
	}
	rawBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error In ReadAll (SignUP): %s", err)
	}
	defer resp.Body.Close()
	if bytes.Contains(rawBytes, []byte(`success`)) {
		var cred struct {
			Key string `json:"key"`
		}
		err = json.Unmarshal(rawBytes, &cred)
		if err != nil {
			return fmt.Errorf("jsonUnmarshal(SignUP) err: %s", err)
		}
		Acc.XApiKey = cred.Key
		Acc.Config.PrintDebug("XApi-Key", Acc.XApiKey)
		return nil
	}
	Acc.Config.PrintNetworkError(resp.StatusCode, "SignUP", string(rawBytes))
	return fmt.Errorf("unknown error occured(SignUP)")
}

func (Acc *Account) Link() error {
	var rawBytes []byte
	var resp *http.Response
	req, err := http.NewRequest("GET", "https://api.picsart.com/discord/link", nil)
	if err != nil {
		return fmt.Errorf("Link err:%s", err)
	}
	req.Header.Set("Host", "api.picsart.com")
	req.Header.Set("Accept", "application/picsart-3.0+json")
	req.Header.Set("Platform", "android")
	req.Header.Set("Language-Code", "en")
	req.Header.Set("App", "com.picsart.studio")
	req.Header.Set("Os-Version", "9")
	req.Header.Set("Manufacturer", "Genymotion")
	req.Header.Set("Device-Model", "Fire 7")
	req.Header.Set("Market", "google")
	req.Header.Set("Is-Tablet", "0")
	req.Header.Set("Country-Code", "US")
	req.Header.Set("X-Api-Key", Acc.XApiKey)
	req.Header.Set("User-Agent", "PicsArt-10.x")
	for i := 0; i <= Acc.Config.RequestRetry; i++ {
		resp, err = (*Acc.Client).Do(req)
		if err != nil {
			if i == Acc.Config.RequestRetry {
				return fmt.Errorf("TimedOut (Link) (MaxRetry)")
			}
			Acc.Config.PrintWarn(fmt.Sprintf("TimeOut (Link) (%d)", i), err.Error())
			continue
		}
		break
	}
	defer resp.Body.Close()
	rawBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error In ReadAll (Link): %s", err)
	}
	defer resp.Body.Close()
	if bytes.Contains(rawBytes, []byte(`"status":"success",`)) {
		var prom struct {
			Response string `json:"response"`
		}
		err = json.Unmarshal(rawBytes, &prom)
		if err != nil {
			return fmt.Errorf("jsonUnmarshal(Link) err: %s", err)
		}
		Acc.ClaimLink = prom.Response
		Acc.Config.PrintGen("ClaimLink", Acc.ClaimLink)
		return nil
	}
	Acc.Config.PrintNetworkError(resp.StatusCode, "Link", string(rawBytes))
	return fmt.Errorf("unknown error occured(Link)")
}
