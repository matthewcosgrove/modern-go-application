// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package todogen

import (
	"context"
	"emperror.dev/errors"
	"fmt"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo"
)

// MarkedAsDoneHandler handles MarkedAsDone events.
type MarkedAsDoneHandler interface {
	// MarkedAsDone handles a(n) MarkedAsDone event.
	MarkedAsDone(ctx context.Context, event todo.MarkedAsDone) error
}

// MarkedAsDoneEventHandler handles MarkedAsDone events.
type MarkedAsDoneEventHandler struct {
	handler MarkedAsDoneHandler
	name    string
}

// NewMarkedAsDoneEventHandler returns a new MarkedAsDoneEventHandler instance.
func NewMarkedAsDoneEventHandler(handler MarkedAsDoneHandler, name string) MarkedAsDoneEventHandler {
	return MarkedAsDoneEventHandler{
		handler: handler,
		name:    name,
	}
}

// HandlerName returns the name of the event handler.
func (h MarkedAsDoneEventHandler) HandlerName() string {
	return h.name
}

// NewEvent returns a new empty event used for serialization.
func (h MarkedAsDoneEventHandler) NewEvent() interface{} {
	return &todo.MarkedAsDone{}
}

// Handle handles an event.
func (h MarkedAsDoneEventHandler) Handle(ctx context.Context, event interface{}) error {
	e, ok := event.(*todo.MarkedAsDone)
	if !ok {
		return errors.NewWithDetails("unexpected event type", "type", fmt.Sprintf("%T", event))
	}

	return h.handler.MarkedAsDone(ctx, *e)
}
