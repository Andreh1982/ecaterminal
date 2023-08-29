package appcontext

import (
	"context"
)

type ContextKey string

const (
	AppContextKey               ContextKey = "appContext"
	defaultBackgroundContextKey ContextKey = "ctx"
)

type Context interface {
	Done()
	Context() context.Context
}

func New(ctx context.Context) Context {
	return &appContext{
		defaultBackgroundContext: ctx,
	}
}

func NewBackground() Context {
	ctx := context.Background()

	return &appContext{
		defaultBackgroundContext: ctx,
	}
}

type appContext struct {
	defaultBackgroundContext context.Context
}

func (appContext *appContext) Context() context.Context {
	return appContext.defaultBackgroundContext
}

func (appContext *appContext) Done() {
	appContext.defaultBackgroundContext = nil
}
