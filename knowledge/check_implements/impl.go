package main

var _ SomeInterface = &Impl{}

type Impl struct{}

func (i *Impl) SomeMethod() {}
