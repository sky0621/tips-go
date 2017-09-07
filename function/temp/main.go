package main

func main() {
}

func mnd(selectDbFunc func(whereStr string) []string) (outFn func(ioTarget []string)) {

	selectResult := selectDbFunc("Taro")

	return func(ioTarget []string) {
		return selectResult
	}
}

type Option struct {
	element interface{}
	err     error
}

func NewOption(element interface{}, err error) Option {
	return Option{
		element: element,
		err:     err,
	}
}

func (o Option) Val() interface{} {
	if o.err != nil {
		return Unit{}
	}
	if o.element == nil {
		return Unit{}
	}
	return o.element
}

type Unit struct {
}

type Condition struct {
	wheres interface{}
}

type IoMonad func(condition Condition) Option
