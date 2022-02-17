package temblate_go

// langTemplates maps a language (2-char string) to messageTemplates
type langTemplates map[string]string

// MessageTemplate maps a key to a string message; the string is intended to contain a go template
type messageTemplates map[string]langTemplates
