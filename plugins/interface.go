package plugins

// Plugin é a interface para todos os plugins do interpretador.
type Plugin interface {
	Execute(method string, args []interface{}) (map[string]interface{}, error)
}