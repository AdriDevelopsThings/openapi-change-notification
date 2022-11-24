package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/adridevelopsthings/openapi-change-notification/config"
)

type HCaptchaSiteVerifyResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	Credit      bool     `json:"credit"`
	ErrorCodes  []string `json:"error-codes"`
}

func VerifyHCaptcha(token string) (bool, error) {
	secret := config.GetConfig().HCaptchaSecret
	if secret == "" {
		fmt.Printf("WARNING: HCaptcha token can't be verrified in this request, because secret is not set.\n")
		return true, nil
	}
	response, err := http.PostForm("https://hcaptcha.com/siteverify", url.Values{
		"secret":   {secret},
		"response": {token},
	})
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		fmt.Printf("Not 200 status code while request to hcaptcha: %d:\n", response.StatusCode)
		b, _ := io.ReadAll(response.Body)
		fmt.Printf("%s\n", string(b))
		return false, nil
	}
	var hcaptchaResponse HCaptchaSiteVerifyResponse
	d := json.NewDecoder(response.Body)
	err = d.Decode(&hcaptchaResponse)
	return hcaptchaResponse.Success, err
}
