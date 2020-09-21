package exporters

import (
	"context"
	"log"
	"net/http"

	"github.com/topfreegames/fluxcloud/pkg/msg"
)

//Exporter An exporter sends a formatted event to an upstream.
type Exporter interface {
	// Send a message through the exporter.
	Send(c context.Context, client *http.Client, message msg.Message) error

	// Return a new line as a string for the exporter.
	NewLine() string

	// Return a link formatted for the exporter.
	FormatLink(link string, name string) string

	// Returns the name of the exporter.
	Name() string

	// determines whether eventType should be reported on or not
	Excluded(eventType string) bool
}

//ExcludingExporter base type that exporters should extend to support the Excluded method
type ExcludingExporter struct {
	ExcludedTypes []string
}

//Excluded determines whether this event type should be reported
func (e *ExcludingExporter) Excluded(eventType string) bool {
	for _, eT := range e.ExcludedTypes {
		if eT == eventType {
			log.Printf("event type %v is configured to be ignored.", eventType)
			return true
		}
	}
	return false
}
