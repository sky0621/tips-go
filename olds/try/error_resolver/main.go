package main

import "fmt"

func main() {
	err1 := fmt.Errorf("some error: %w", ErrNotFound)
	err2 := fmt.Errorf("some error: %w", ErrPermissionDenied)
	err3 := fmt.Errorf("some error: %w", ErrUnknown)

	resolver := NewErrorResolverStrategy()
	fmt.Println(resolver.ResolveError(err1))
	fmt.Println(resolver.ResolveError(err2))
	fmt.Println(resolver.ResolveError(err3))
}
