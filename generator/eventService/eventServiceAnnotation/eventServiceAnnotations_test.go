package eventServiceAnnotation

import (
	"testing"

	"github.com/MarcGrol/golangAnnotations/annotation"
	"github.com/stretchr/testify/assert"
)

func TestCorrectEventServiceAnnotation(t *testing.T) {
	registry := annotation.NewRegistry(Get())

	ann, ok := registry.ResolveAnnotationByName([]string{`// @EventService( Self = "caregiverService", process = "myprocess", async="true", admin="true", producesEvents="x,y" )`}, "EventService")
	assert.True(t, ok)
	{
		self, ok := ann.Attributes[ParamSelf]
		assert.True(t, ok)
		assert.Equal(t, "caregiverService", self)
	}
	{
		async, ok := ann.Attributes[ParamAsync]
		assert.True(t, ok)
		assert.Equal(t, "true", async)
	}
	{
		process, ok := ann.Attributes[ParamProcess]
		assert.True(t, ok)
		assert.Equal(t, "myprocess", process)
	}
	{
		producesEvents, ok := ann.Attributes[ParamProducesEvents]
		assert.True(t, ok)
		assert.Equal(t, "x,y", producesEvents)
	}
}

func TestCorrectEventOperationAnnotation(t *testing.T) {
	registry := annotation.NewRegistry(Get())

	_, ok := registry.ResolveAnnotation(`// @EventOperation( topic = "order" )`)
	assert.True(t, ok)
}
