package shared

import (
	"expvar"

	"github.com/thalesfsp/sypl"
)

// IMeta defines what an `Entity` must do.
type IMeta interface {
	// GetName returns the `Name` of the entity.
	GetName() string

	// GetLogger returns the `Logger` of the entity.
	GetLogger() sypl.ISypl

	// GetDescription returns the `Description` of the processor.
	GetDescription() string
}

// IMetrics defines how to interact with the metrics.
type IMetrics interface {
	// GetCounterCreated returns the `CounterCreated` metric.
	GetCounterCreated() *expvar.Int

	// GetCounterRunning returns the `CounterRunning` metric.
	GetCounterRunning() *expvar.Int

	// GetCounterFailed returns the `CounterFailed` metric.
	GetCounterFailed() *expvar.Int

	// GetCounterDone returns the `CounterDone` metric.
	GetCounterDone() *expvar.Int

	// GetStatus returns the `Status` metric.
	GetStatus() *expvar.String
}
