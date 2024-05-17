package pipeline

import (
	"context"

	"github.com/thalesfsp/etler/v2/internal/shared"
	"github.com/thalesfsp/etler/v2/task"
	"github.com/thalesfsp/status"
)

// IPipeline defines what a `Pipeline` must do.
type IPipeline[ProcessedData, ConvertedOut any] interface {
	shared.IMeta

	shared.IMetrics

	// GetPaused returns the Paused status.
	GetPaused() status.Status

	// SetPause the pipeline.
	SetPause(state bool)

	// GetOnFinished returns the `OnFinished` function.
	GetOnFinished() OnFinished[ProcessedData, ConvertedOut]

	// SetOnFinished sets the `OnFinished` function.
	SetOnFinished(onFinished OnFinished[ProcessedData, ConvertedOut])

	// Run the pipeline.
	Run(ctx context.Context, processedData []ProcessedData) ([]task.Task[ProcessedData, ConvertedOut], error)
}
