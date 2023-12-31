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
	"io"
	"net/http"
	"sync"
	"time"
)

type COneError struct {
	Message string `json:"message"`
}

func (e COneError) Error() string {
	return e.Message
}

type CloudOneNS struct {
	APIKey         string
	CloudOneRegion string
	AccountId      string
	// AWSRegion      string
}

func NewCloudOneNS(apiKey string, cloudOneRegion string, accountId string /*, awsRegion string*/) *CloudOneNS {
	return &CloudOneNS{
		APIKey:         apiKey,
		CloudOneRegion: cloudOneRegion,
		AccountId:      accountId,
		//	AWSRegion:      awsRegion,
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
	Status             Status    `json:"status"`
	UpdateTime         time.Time `json:"updateTime"`
}

func (c *CloudOneNS) GetInspectionBypassStatus_(ctx context.Context, AWSRegion string) (*COneNSBypassStatus, error) {
	uri := fmt.Sprintf("https://network.%s.cloudone.trendmicro.com/api/nsaas/inspection-bypass?accountId=%s&awsRegion=%s",
		c.CloudOneRegion, c.AccountId /*c.*/, AWSRegion)
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "ApiKey "+c.APIKey)
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

/*
	func (c *CloudOneNS) SetInspectionBypass(ctx context.Context, action Action) error {
		uri := fmt.Sprintf("https://network.%s.cloudone.trendmicro.com/api/nsaas/inspection-bypass",
			c.CloudOneRegion)
		request := COneNSBypassRequest{
			AccountID: c.AccountId,
			Action:    action,
			AwsRegion: c.AWSRegion,
		}
		var body bytes.Buffer
		if err := json.NewEncoder(&body).Encode(&request); err != nil {
			return fmt.Errorf("json encode: %w", err)
		}
		//fmt.Println(body.String())
		req, err := http.NewRequestWithContext(ctx, "PUT", uri, &body)
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}
		req.Header.Set("Authorization", "ApiKey "+c.APIKey)
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
	}*/

type AccountInfo struct {
	ID           string    `json:"id"`
	Alias        string    `json:"alias"`
	Locale       string    `json:"locale"`
	Timezone     string    `json:"timezone"`
	Region       string    `json:"region"`
	State        string    `json:"state"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	Urn          string    `json:"urn"`
	MfaRequired  bool      `json:"mfaRequired"`
	Links        []struct {
		Rel    string `json:"rel"`
		Href   string `json:"href"`
		Method string `json:"method"`
	} `json:"links"`
}

/*
	func (c *CloudOneNS) GetAccountInfo(ctx context.Context) (*AccountInfo, error) {
		uri := fmt.Sprintf("https://accounts.cloudone.trendmicro.com/api/accounts/%s", c.AccountId)
		req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
		if err != nil {
			return nil, fmt.Errorf("create request: %w", err)
		}
		req.Header.Set("Authorization", "ApiKey "+c.APIKey)
		req.Header.Set("Api-Version", "v1")
		//log.Println(req.Header)
		//log.Println(uri)
		client := &http.Client{}
		resp, err := client.Do(req)
		//log.Println(resp, err)
		if err != nil {
			return nil, fmt.Errorf("HTTP request: %w", err)
		}
		if resp.StatusCode == 200 {
			var response AccountInfo
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
*/
func (c *CloudOneNS) request(ctx context.Context, method string, uri string, body io.Reader, response any) error {
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "ApiKey "+c.APIKey)
	req.Header.Set("Api-Version", "v1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request: %w", err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		if response != nil {
			if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
				return fmt.Errorf("decode error: %w", err)
			}
		}
		return nil
	}
	var cOneError COneError
	if err := json.NewDecoder(resp.Body).Decode(&cOneError); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}
	return cOneError
}

func (c *CloudOneNS) GetAccountInfo_(ctx context.Context) (*AccountInfo, error) {
	uri := fmt.Sprintf("https://accounts.cloudone.trendmicro.com/api/accounts/%s", c.AccountId)
	var response AccountInfo
	err := c.request(ctx, "GET", uri, nil, &response)
	return &response, err

}

func (c *CloudOneNS) SetInspectionBypass_(ctx context.Context, awsRegion string, action Action) error {
	uri := fmt.Sprintf("https://network.%s.cloudone.trendmicro.com/api/nsaas/inspection-bypass",
		c.CloudOneRegion)
	request := COneNSBypassRequest{
		AccountID: c.AccountId,
		Action:    action,
		AwsRegion:/*c.*/ awsRegion,
	}
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&request); err != nil {
		return fmt.Errorf("json encode: %w", err)
	}
	return c.request(ctx, "PUT", uri, &body, nil)
}

func (c *CloudOneNS) GetInspectionBypassStatus(ctx context.Context, AWSRegion string) (*COneNSBypassStatus, error) {

	uri := fmt.Sprintf("https://network.%s.cloudone.trendmicro.com/api/nsaas/inspection-bypass?accountId=%s&awsRegion=%s",
		c.CloudOneRegion, c.AccountId /*c.*/, AWSRegion)
	var response COneNSBypassStatus
	err := c.request(ctx, "GET", uri, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil

}

var AWSRegions = []string{
	"af-south-1",
	"ap-east-1",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-northeast-3",
	"ap-south-1",
	"ap-south-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-southeast-3",
	"ap-southeast-4",
	"ca-central-1",
	"eu-central-1",
	"eu-central-2",
	"eu-north-1",
	"eu-south-1",
	"eu-south-2",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"il-central-1",
	"me-central-1",
	"me-south-1",
	"sa-east-1",
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
}

func (c *CloudOneNS) DetectAWSRegions(ctx context.Context) (awsRegion []string) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, region := range AWSRegions {
		wg.Add(1)
		go func(r string) {
			defer wg.Done()
			status, err := c.GetInspectionBypassStatus(ctx, r)
			if err != nil || len(status.Error) > 0 {
				return
			}
			mu.Lock()
			awsRegion = append(awsRegion, r)
			mu.Unlock()
		}(region)
	}
	wg.Wait()
	return
}
