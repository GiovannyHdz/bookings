package models

// TemplateDate holds data send from handlers to template
type TemplateDate struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string // Success messages
	Warning   string
	Error     string
}
