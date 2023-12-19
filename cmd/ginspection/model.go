package main

import (
	"crypto/sha1"
	"fmt"
	"inspection/pkg/cone"
)

type Model struct {
	password string
	config   Configuration
	hash     string
}

func (m *Model) CalculateHash() string {
	h := sha1.New()
	h.Write([]byte(m.password))
	h.Write([]byte(m.config.APIKey))
	h.Write([]byte(m.config.Region))
	h.Write([]byte(m.config.AccountID))
	h.Write([]byte(m.config.AWSRegion))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (m *Model) Load(fileName string) error {
	if err := m.config.Load(fileName, m.password); err != nil {
		return err
	}
	m.hash = m.CalculateHash()
	return nil
}

func (m *Model) Save(fileName string) error {
	if err := m.config.Save(fileName, m.password); err != nil {
		return err
	}
	m.hash = m.CalculateHash()
	return nil
}

func (m *Model) Changed() bool {
	return m.hash != m.CalculateHash()
}

func (m *Model) COne() *cone.CloudOneNS {
	return cone.NewCloudOneNS(m.config.apiKeyDecrypted, m.config.Region, m.config.AccountID, m.config.AWSRegion)
}
