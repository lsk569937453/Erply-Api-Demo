package client

import (
	"crypto/tls"
	"erply-api/internal/common"
	"net/http"
	"time"

	"github.com/erply/api-go-wrapper/pkg/api/auth"

	"github.com/erply/api-go-wrapper/pkg/api"
)

var ApiClient *api.Client

func init() {
	const (
		username   = "lsk569937453@gmail.com"
		password   = "demo1234"
		clientCode = "113073"
		partnerKey = "113073"
	)

	apiClientInit, err := buildClient()
	common.Die(err)
	ApiClient = apiClientInit
}

func buildClient() (*api.Client, error) {
	username := "lsk569937453@gmail.com"
	password := "demo1234"
	clientCode := "113073"

	connectionTimeout := 60 * time.Second
	transport := &http.Transport{
		DisableKeepAlives:     true,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: connectionTimeout,
	}
	httpCl := &http.Client{Transport: transport}

	sessionKey, err := auth.VerifyUser(username, password, clientCode, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	apiClient, err := api.NewClient(sessionKey, clientCode, httpCl)
	if err != nil {
		panic(err)
	}

	return apiClient, nil
}
