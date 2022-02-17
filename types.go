package temblate_go

// LangTemplates maps a language (2-char string) to messageTemplates
type LangTemplates map[string]string

// MessageTemplate maps a key to a string message; the string is intended to contain a go template
type messageTemplates map[string]LangTemplates
