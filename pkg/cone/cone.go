/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

cone.go - small library to control fallback mode of C1NS Hostend Infrastucture

*/

package cone 

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type COneError struct {
	Message string `json:"message"`
}

func (e COneError) Error() string {
	return e.Message
}

type CloudOneNS struct {
	apiKey         string
	cloudOneRegion string
	accountId      string
	awsRegion      string
}

func NewCloudOneNS(apiKey string, cloudOneRegion string, accountId string, awsRegion string) *CloudOneNS {
	return &CloudOneNS{
		apiKey:         apiKey,
		cloudOneRegion: cloudOneRegion,
		accountId:      accountId,
		awsRegion:      awsRegion,
	}
}

//go:generate  enum -package cone -type Status --names success,fail,in-progress
type COneNSBypassStatus struct {
	AccountID          string    `json:"accountId"`
	Action             Action    `json:"action"`
	AwsRegion          string    `json:"awsRegion"`
	Error              string    `json:"error"`
	InitiateByCustomer bool      `json:"initiateByCustomer"`
	InitiatorAccountID string    `json:"initiatorAccountId"`
	Status             string    `json:"status"`
	UpdateTime         time.Time `json:"updateTime"`
}

func (c *CloudOneNS) GetInspectionBypassStatus(ctx context.Context) (*COneNSBypassStatus, error) {
	uri := fmt.Sprintf("https://network.%s.cloudone.trendmicro.com/api/nsaas/inspection-bypass?accountId=%s&awsRegion=%s",
		c.cloudOneRegion, c.accountId, c.awsRegion)
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "ApiKey "+c.apiKey)
	req.Header.Set("Api-Version", "v1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request: %w", err)
	}
	if resp.StatusCode == 200 {
		var response COneNSBypassStatus
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}
		return &response, nil
	}
	var cOneError COneError
	if err := json.NewDecoder(resp.Body).Decode(&cOneError); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}
	return nil, cOneError
}

//go:generate  enum -package cone -type Action --names bypass,inspect
type COneNSBypassRequest struct {
	AccountID string `json:"accountId"`
	Action    Action `json:"action"`
	AwsRegion string `json:"awsRegion"`
}

func (c *CloudOneNS) SetInspectionBypass(ctx context.Context, action Action) error {
	uri := fmt.Sprintf("https://network.%s.cloudone.trendmicro.com/api/nsaas/inspection-bypass",
		c.cloudOneRegion)
	request := COneNSBypassRequest{
		AccountID: c.accountId,
		Action:    action,
		AwsRegion: c.awsRegion,
	}
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&request); err != nil {
		return fmt.Errorf("json encode: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", uri, &body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "ApiKey "+c.apiKey)
	req.Header.Set("Api-Version", "v1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request: %w", err)
	}
	if resp.StatusCode == 202 {
		return nil
	}
	var cOneError COneError
	if err := json.NewDecoder(resp.Body).Decode(&cOneError); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}
	return cOneError
}
