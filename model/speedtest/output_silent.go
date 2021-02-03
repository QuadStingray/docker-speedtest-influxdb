package speedtest

import (
	"github.com/m-lab/ndt7-client-go/spec"
)

type SilentOutput struct {
}

// OnStarting handles the start event
func (h SilentOutput) OnStarting(test spec.TestKind) error {
	return nil
}

// OnError handles the error event
func (h SilentOutput) OnError(test spec.TestKind, err error) error {
	return nil
}

// OnConnected handles the connected event
func (h SilentOutput) OnConnected(test spec.TestKind, fqdn string) error {
	return nil
}

// OnDownloadEvent handles an event emitted by the download test
func (h SilentOutput) OnDownloadEvent(m *spec.Measurement) error {
	return h.onSpeedEvent(m)
}

// OnUploadEvent handles an event emitted during the upload test
func (h SilentOutput) OnUploadEvent(m *spec.Measurement) error {
	return h.onSpeedEvent(m)
}

func (h SilentOutput) onSpeedEvent(m *spec.Measurement) error {
	return nil
}

// OnComplete handles the complete event
func (h SilentOutput) OnComplete(test spec.TestKind) error {
	return nil
}

// OnSummary handles the summary event.
func (h SilentOutput) OnSummary(s *Summary) error {
	return nil
}
