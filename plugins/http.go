package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"logline-runtime/auditor" // Para gravação de LogLines de auditoria do plugin
)

// HTTPPlugin representa um plugin que faz requisições HTTP reais.
type HTTPPlugin struct {
	client        *http.Client
	auditorClient *auditor.Client // Cliente para gravar LogLines
}

// NewHTTPPlugin cria e retorna uma nova instância do HTTPPlugin.
func NewHTTPPlugin(auditorClient *auditor.Client) *HTTPPlugin {
	return &HTTPPlugin{
		client:        &http.Client{Timeout: 15 * time.Second},
		auditorClient: auditorClient,
	}
}

// Execute faz a requisição HTTP baseada em serviceMethod e args.
func (p *HTTPPlugin) Execute(serviceMethod string, args []interface{}) (map[string]interface{}, error) {
	log.Printf("HTTPPlugin: Executando %s com args: %v", serviceMethod, args)

	verb := "POST"
	url := serviceMethod
	parts := strings.SplitN(serviceMethod, " ", 2)
	if len(parts) == 2 {
		verb = strings.ToUpper(parts[0])
		url = parts[1]
	}

	var payload []byte
	if len(args) > 0 && args[0] != nil {
		payload, _ = json.Marshal(args[0])
	}

	req, err := http.NewRequest(verb, url, bytes.NewBuffer(payload))
	if err != nil {
		p.storePluginAuditLog("http_request_failed", "create_request", false, map[string]interface{}{"url": url, "verb": verb, "error": err.Error()})
		return map[string]interface{}{"ok": false, "error": err.Error()}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		p.storePluginAuditLog("http_request_failed", "do_request", false, map[string]interface{}{"url": url, "verb": verb, "error": err.Error()})
		return map[string]interface{}{"ok": false, "error": err.Error()}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		p.storePluginAuditLog("http_request_failed", "read_response_body", false, map[string]interface{}{"url": url, "verb": verb, "error": err.Error()})
		return map[string]interface{}{"ok": false, "error": err.Error()}, err
	}

	var responseData map[string]interface{}
	if len(bodyBytes) > 0 {
		if err := json.Unmarshal(bodyBytes, &responseData); err != nil {
			p.storePluginAuditLog("http_request_failed", "unmarshal_response_body", false, map[string]interface{}{"url": url, "verb": verb, "body": string(bodyBytes), "error": err.Error()})
			return map[string]interface{}{"ok": false, "error": fmt.Sprintf("Failed to parse JSON response: %v, Body: %s", err, string(bodyBytes))}, err
		}
	} else {
		responseData = map[string]interface{}{"message": "Empty response body"}
	}

	isOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	status := "valid"
	if !isOK {
		status = "failed"
	}

	p.storePluginAuditLog(
		"http_request_completed",
		"http_request",
		isOK,
		map[string]interface{}{
			"url":          url,
			"verb":         verb,
			"status_code":  resp.StatusCode,
			"response_data": responseData,
			"response_size": len(bodyBytes),
		},
	)

	return map[string]interface{}{"ok": isOK, "data": responseData, "status_code": resp.StatusCode, "error": nil}, nil
}

// storePluginAuditLog grava LogLines de auditoria do plugin.
func (p *HTTPPlugin) storePluginAuditLog(didSuffix, thisPrefix string, ok bool, details map[string]interface{}) {
	status := "valid"
	if !ok {
		status = "failed"
	}
	logLine := auditor.LogLine{
		Who:         "runtime@plugin.http",
		Did:         fmt.Sprintf("http_%s", didSuffix),
		This:        fmt.Sprintf("http_call_%s", thisPrefix),
		When:        time.Now().Format(time.RFC3339),
		ConfirmedBy: "plugin.http_executor",
		IfOk:        "request_handled",
		IfDoubt:     "review_http_log",
		IfNot:       "http_error_alert",
		Status:      status,
		Type:        "plugin_execution",
		Details:     details,
	}
	if err := p.auditorClient.StoreAuditLog(logLine); err != nil {
		log.Printf("ERRO CRÍTICO: Falha ao gravar LogLine de auditoria do HTTP Plugin: %v", err)
	}
}