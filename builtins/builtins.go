package builtins

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// BuiltinFn é o tipo de função nativa do interpretador.
type BuiltinFn func(args ...interface{}) (interface{}, error)

// NewBuiltinsMap retorna um mapa de todas as funções built-in disponíveis.
func NewBuiltinsMap() map[string]BuiltinFn {
	return map[string]BuiltinFn{
		"now":     Now,
		"len":     Len,
		"upper":   Upper,
		"add_days": AddDays,
		"sha256":  SHA256,
		"format":  Format,
	}
}

// Now retorna timestamp UTC ISO8601.
func Now(args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("função 'now' não aceita argumentos")
	}
	return time.Now().UTC().Format(time.RFC3339), nil
}

// Len retorna o comprimento de string ou slice.
func Len(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("função 'len' requer 1 argumento")
	}
	switch arg := args[0].(type) {
	case string:
		return len(arg), nil
	case []interface{}:
		return len(arg), nil
	case []string:
		return len(arg), nil
	default:
		return nil, fmt.Errorf("tipo de argumento não suportado para 'len': %T", args[0])
	}
}

// Upper converte string para maiúsculas.
func Upper(args ...interface{}) (interface{}, error) {
	if len(args) != 1 || reflect.TypeOf(args[0]).Kind() != reflect.String {
		return nil, fmt.Errorf("função 'upper' requer 1 argumento do tipo string")
	}
	return strings.ToUpper(args[0].(string)), nil
}

// AddDays adiciona dias a uma data ISO8601.
func AddDays(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("função 'add_days' requer 2 argumentos: data e número de dias")
	}
	dateStr, ok1 := args[0].(string)
	daysFloat, ok2 := args[1].(float64)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("argumentos inválidos para 'add_days'")
	}
	days := int(daysFloat)
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return nil, fmt.Errorf("falha ao parsear data '%s': %w", dateStr, err)
	}
	newDate := t.Add(time.Duration(days) * 24 * time.Hour)
	return newDate.Format(time.RFC3339), nil
}

// SHA256 calcula hash SHA256 de string.
func SHA256(args ...interface{}) (interface{}, error) {
	if len(args) != 1 || reflect.TypeOf(args[0]).Kind() != reflect.String {
		return nil, fmt.Errorf("função 'sha256' requer 1 argumento do tipo string")
	}
	hash := sha256.Sum256([]byte(args[0].(string)))
	return fmt.Sprintf("%x", hash), nil
}

// Format: "Olá, {}!".format("Mundo") => "Olá, Mundo!"
func Format(args ...interface{}) (interface{}, error) {
	if len(args) < 2 || reflect.TypeOf(args[0]).Kind() != reflect.String {
		return nil, fmt.Errorf("função 'format' requer pelo menos 2 argumentos: string de formato e valores")
	}
	formatString := args[0].(string)
	placeholders := args[1:]
	if len(placeholders) > 0 {
		return strings.Replace(formatString, "{}", fmt.Sprintf("%v", placeholders[0]), 1), nil
	}
	return formatString, nil
}