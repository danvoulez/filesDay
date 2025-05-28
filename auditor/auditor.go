package auditor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"logline-runtime/registry"
)

// LogLine representa a estrutura canônica de uma LogLine.
type LogLine registry.LogLine

type Client struct {
	registryClient *registry.Client
	auditKey       []byte
}

func NewClient(regClient *registry.Client, auditKeyStr string) *Client {
	return &Client{
		registryClient: regClient,
		auditKey:       []byte(auditKeyStr),
	}
}

func (c *Client) StoreAuditLog(logLine LogLine) error {
	logLineCopy := logLine
	if logLineCopy.Details == nil {
		logLineCopy.Details = make(map[string]interface{})
	}
	delete(logLineCopy.Details, "signature")
	bytesToSign, err := json.Marshal(logLineCopy)
	if err != nil {
		return fmt.Errorf("falha ao serializar LogLine para assinatura HMAC: %w", err)
	}
	if len(c.auditKey) > 0 {
		h := hmac.New(sha256.New, c.auditKey)
		h.Write(bytesToSign)
		signature := h.Sum(nil)
		logLine.Details["signature"] = fmt.Sprintf("hmac-sha256:%x", signature)
	}
	regLogLine := registry.LogLine(logLine)
	return c.registryClient.StoreAuditLog(regLogLine)
}