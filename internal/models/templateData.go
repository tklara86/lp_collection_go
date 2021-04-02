package models

import "github.com/tklara86/lp_collection_go/internal/forms"

// TemplateData holds data set from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	FlashMessage string
	Warning string
	Error string
	Form *forms.Form
}

