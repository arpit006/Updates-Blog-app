package datastore

import "context"

var CONTEXT context.Context = nil

func ContextFactory() context.Context {
	if CONTEXT != nil {
		return CONTEXT
	}
	CONTEXT := context.Background()
	return CONTEXT
}
