package eventServiceAnnotation

import "github.com/MarcGrol/golangAnnotations/annotation"

const (
	TypeEventService    = "EventService"
	TypeEventOperation  = "EventOperation"
	ParamSelf           = "self"
	ParamAsync          = "async"
	ParamTopic          = "topic"
	ParamProcess        = "process"
	ParamNoTest         = "notest"
	ParamProducesEvents = "producesevents"
)

func Get() []annotation.AnnotationDescriptor {
	return []annotation.AnnotationDescriptor{
		{
			Name:       TypeEventService,
			ParamNames: []string{ParamSelf, ParamAsync, ParamNoTest},
			Validator:  validateEventServiceAnnotation,
		},
		{
			Name:       TypeEventOperation,
			ParamNames: []string{ParamTopic, ParamProcess},
			Validator:  validateEventOperationAnnotation,
		}}
}

func validateEventServiceAnnotation(annot annotation.Annotation) bool {
	if annot.Name == TypeEventService {
		return true
	}
	return false
}

func validateEventOperationAnnotation(annot annotation.Annotation) bool {
	if annot.Name == TypeEventOperation {
		_, ok := annot.Attributes[ParamTopic]
		return ok
	}
	return false
}
