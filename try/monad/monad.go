package monad

type Monad interface {
	Unit(t interface{}) Monad
	FlatMap()
}
