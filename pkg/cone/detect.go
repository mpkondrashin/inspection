package cone

import (
	"context"
	"errors"
	"sync"
)

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

func DetectAWSRegions(ctx context.Context, accountID, apiKey, cloudOneRegion string) (awsRegion []string) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, region := range AWSRegions {
		wg.Add(1)
		go func(r string) {
			defer wg.Done()
			c := NewCloudOneNS(apiKey, cloudOneRegion, accountID, r)
			status, err := c.GetInspectionBypassStatus(ctx)
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

var CloudOneRegions = []string{
	"us-1",
	"in-1",
	"gb-1",
	"jp-1",
	"de-1",
	"au-1",
	"ca-1",
	"sg-1",
	"trend-us-1",
}

var ErrUndetectedCloudOneRegion = errors.New("undetected Cloud One region")

func DetectCloudOneRegion(ctx context.Context, accountID, apiKey, awsRegion string) (cloudOneRegion string, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	err = ErrUndetectedCloudOneRegion
	for _, c1r := range CloudOneRegions {
		wg.Add(1)
		go func(c1r string) {
			//log.Println("test ", r)
			defer wg.Done()
			c := NewCloudOneNS(apiKey, c1r, accountID, awsRegion)
			if _, err := c.GetInspectionBypassStatus(ctx); err != nil {
				return
			}
			cancel()
			cloudOneRegion = c1r
			err = nil
		}(c1r)
	}
	wg.Wait()
	return
}

var ErrUndetectedRegions = errors.New("undetected regions")

func DetectBothRegions(ctx context.Context, accountID, apiKey string) (awsRegion, cloudOneRegion string, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	err = ErrUndetectedCloudOneRegion
	for _, c1r := range CloudOneRegions {
		for _, ar := range AWSRegions {
			wg.Add(1)
			go func(c1r, ar string) {
				defer wg.Done()
				c := NewCloudOneNS(apiKey, c1r, accountID, ar)
				if _, err := c.GetInspectionBypassStatus(ctx); err != nil {
					return
				}
				cancel()
				awsRegion = ar
				cloudOneRegion = c1r
				err = nil
			}(c1r, ar)
		}
	}
	wg.Wait()
	return
}
