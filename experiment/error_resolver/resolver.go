package main

type ErrorResolver interface {
	ResolveError(e error) error
}
