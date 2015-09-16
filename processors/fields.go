package processors

import "loggo"

type FieldsProcessor struct {
	fields map[string]interface{}
}

func NewFieldsProcessor(fields map[string]interface{}) *FieldsProcessor {
	return &FieldsProcessor{
		fields: fields,
	}
}

func (p *FieldsProcessor) Process(entry *loggo.Entry) {
	for key, value := range p.fields {
		entry.Fields[key] = value
	}
}
