package main

// ここで構造体を指定のインタフェースに格納しておくことで、インタフェースの実装漏れがあった時にコンパイルエラーになって検知できる。
var _ SomeInterface = &Impl{}

type Impl struct{}

func (i *Impl) SomeMethod() {}
