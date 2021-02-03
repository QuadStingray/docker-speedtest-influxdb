package speedtest

import (
	"context"
	"github.com/m-lab/ndt7-client-go"
	"github.com/m-lab/ndt7-client-go/spec"
)

type TestRunner struct {
	Client *ndt7.Client
	Output OutputType
}

func (r TestRunner) doRunTest(
	ctx context.Context, test spec.TestKind,
	start func(context.Context) (<-chan spec.Measurement, error),
	emitEvent func(m *spec.Measurement) error,
) int {
	ch, err := start(ctx)
	if err != nil {
		r.Output.OnError(test, err)
		return 1
	}
	err = r.Output.OnConnected(test, r.Client.FQDN)
	if err != nil {
		return 1
	}
	for ev := range ch {
		err = emitEvent(&ev)
		if err != nil {
			return 1
		}
	}
	return 0
}

func (r TestRunner) runTest(
	ctx context.Context, test spec.TestKind,
	start func(context.Context) (<-chan spec.Measurement, error),
	emitEvent func(m *spec.Measurement) error,
) int {
	err := r.Output.OnStarting(test)
	if err != nil {
		return 1
	}
	code := r.doRunTest(ctx, test, start, emitEvent)
	err = r.Output.OnComplete(test)
	if err != nil {
		return 1
	}
	return code
}

func (r TestRunner) RunDownload(ctx context.Context) int {
	return r.runTest(ctx, spec.TestDownload, r.Client.StartDownload,
		r.Output.OnDownloadEvent)
}

func (r TestRunner) RunUpload(ctx context.Context) int {
	return r.runTest(ctx, spec.TestUpload, r.Client.StartUpload,
		r.Output.OnUploadEvent)
}
