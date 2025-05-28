package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"logline-runtime/auditor"
)

type LLMPlugin struct {
	client            *http.Client
	auditorClient     *auditor.Client
	openAIApiKey      string
	anthropicApiKey   string
	googleGeminiApiKey string
}

func NewLLMPlugin(auditorClient *auditor.Client) *LLMPlugin {
	return &LLMPlugin{
		client:            &http.Client{Timeout: 60 * time.Second},
		auditorClient:     auditorClient,
		openAIApiKey:      os.Getenv("OPENAI_API_KEY"),
		anthropicApiKey:   os.Getenv("ANTHROPIC_API_KEY"),
		googleGeminiApiKey: os.Getenv("GEMINI_API_KEY"),
	}
}

func (p *LLMPlugin) Execute(method string, args []interface{}) (map[string]interface{}, error) {
	if method != "dispatch" {
		return nil, fmt.Errorf("método '%s' não suportado pelo plugin LLM. Use 'dispatch'", method)
	}
	if len(args) == 0 {
		return nil, fmt.Errorf("chamada 'llm.dispatch' requer um mapa de argumentos (prompt, model, etc.)")
	}
	params, ok := args[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("o primeiro argumento para 'llm.dispatch' deve ser um mapa de parâmetros")
	}
	prompt, _ := params["prompt"].(string)
	model, _ := params["model"].(string)
	temperature, _ := params["temperature"].(float64)
	maxTokens, _ := params["max_tokens"].(float64)

	if prompt == "" || model == "" {
		p.storePluginAuditLog("llm_dispatch_failed", "missing_params", false, map[string]interface{}{"error": "Prompt ou modelo ausente"})
		return nil, fmt.Errorf("prompt e modelo são obrigatórios para llm.dispatch")
	}

	llmURL := ""
	llmAPIKey := ""
	requestBody := map[string]interface{}{}
	responsePath := []string{}

	if strings.HasPrefix(model, "gpt") || strings.Has