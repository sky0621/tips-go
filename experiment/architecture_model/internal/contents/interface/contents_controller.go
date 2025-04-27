package _interface

type ContentsController interface {
	PostContents() error
	GetContents() error
}

type contentsController struct {
}

func NewContentsController() ContentsController {
	return &contentsController{}
}
