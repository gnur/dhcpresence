package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

const API_ENDPOINT = "http://modem.uranus.aawa.nl/ws/NeMo/Intf/lan:getMIBs"

type authResponse struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
		Username  string `json:"username"`
		Groups    string `json:"groups"`
	} `json:"data"`
}

type authRequest struct {
	Service    string `json:"service"`
	Method     string `json:"method"`
	Parameters struct {
		ApplicationName string `json:"applicationName"`
		Username        string `json:"username"`
		Password        string `json:"password"`
	} `json:"parameters"`
}

func getDevices() ([]Device, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("unable to create cookiejar: %w", err)
	}

	authReq := authRequest{
		Service: "sah.Device.Information",
		Method:  "createContext",
		Parameters: struct {
			ApplicationName string `json:"applicationName"`
			Username        string `json:"username"`
			Password        string `json:"password"`
		}{
			ApplicationName: "webui",
			Username:        "admin",
			Password:        `Y~X"F2w}t3`,
		},
	}
	authReqJSON, err := json.Marshal(authReq)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal request: %w", err)
	}

	body := bytes.NewReader(authReqJSON)
	req, err := http.NewRequest("POST", API_ENDPOINT, body)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Set("Authorization", "X-Sah-Login")
	req.Header.Set("Content-Type", "application/x-sah-ws-4-call+js; charset=utf-8")

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request: %w", err)
	}
	defer resp.Body.Close()

	//parse body to get contextID
	var authResp authResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return nil, fmt.Errorf("unable to parse response: %w", err)
	}

	devicesBody := bytes.NewReader([]byte(`{"service":"Devices","method":"get","parameters":{"expression":"not interface and not self and not voice"}}`))
	req, err = http.NewRequest("POST", API_ENDPOINT, devicesBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
		// handle err
	}
	req.Header.Set("X-Context", authResp.Data.ContextID)
	req.Header.Set("Cookie", resp.Header.Get("set-cookie"))

	resp, err = client.Do(req)
	if err != nil {
		// handle err
		return nil, fmt.Errorf("unable to get ips: %w", err)
	}
	defer resp.Body.Close()

	//parse body to get ips
	var ipsResp DeviceResp
	err = json.NewDecoder(resp.Body).Decode(&ipsResp)
	if err != nil {
		return nil, fmt.Errorf("unable to parse response: %w", err)
	}

	return ipsResp.Status, nil
}

type Payload struct {
	Service    string     `json:"service"`
	Method     string     `json:"method"`
	Parameters Parameters `json:"parameters"`
}
type Parameters struct {
}
