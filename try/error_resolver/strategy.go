package main

type ErrorResolverStrategy struct {
	ErrorResolvers []ErrorResolver
}

func NewErrorResolverStrategy() ErrorResolverStrategy {
	return ErrorResolverStrategy{
		ErrorResolvers: []ErrorResolver{
			NotFoundResolver{},
			PermissionDeniedResolver{},
			OtherResolver{},
		},
	}
}

func (s ErrorResolverStrategy) ResolveError(e error) error {
	for _, x := range s.ErrorResolvers {
		e = x.ResolveError(e)
	}
	return e
}
