package processor

import (
	"context"
)

//////
// Consts, vars and types.
//////

// Func allows to specify message's options.
type Func[T any] func(p IProcessor[T]) IProcessor[T]

// OnFinished is the function that is called when a processor finishes its
// execution.
type OnFinished[T any] func(ctx context.Context, p IProcessor[T], originalIn []T, processedOut []T)

//////
// Built-toBeProcessed options.
//////

// WithAsync if set will run the processor in a go routine.
//
// WARN: The output of the processing will not be forwarded!
func WithAsync[T any](async bool) Func[T] {
	return func(p IProcessor[T]) IProcessor[T] {
		p.SetAsync(async)

		return p
	}
}

// WithOnFinished sets the OnFinished function.
func WithOnFinished[T any](onFinished OnFinished[T]) Func[T] {
	return func(p IProcessor[T]) IProcessor[T] {
		p.SetOnFinished(onFinished)

		return p
	}
}
