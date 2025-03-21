// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// CreatedAt 作成日時
type CreatedAt = time.Time

// Todo defines model for Todo.
type Todo struct {
	// Completed 完了フラグ
	Completed *TodoCompleted `json:"completed,omitempty"`

	// Content TODOの内容
	Content *TodoContent `json:"content,omitempty"`

	// CreatedAt 作成日時
	CreatedAt CreatedAt `json:"createdAt"`

	// TodoID TODOを一意に識別するキー
	TodoID TodoID `json:"id"`

	// Title TODOのタイトル
	Title TodoTitle `json:"title"`
}

// TodoCompleted 完了フラグ
type TodoCompleted = bool

// TodoContent TODOの内容
type TodoContent = string

// TodoCreateInput 作成するTODOの入力情報
type TodoCreateInput struct {
	// Content TODOの内容
	Content *TodoContent `json:"content,omitempty"`

	// Title TODOのタイトル
	Title TodoTitle `json:"title"`
}

// TodoID TODOを一意に識別するキー
type TodoID = int

// TodoTitle TODOのタイトル
type TodoTitle = string

// GetTodosParams defines parameters for GetTodos.
type GetTodosParams struct {
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// PostTodosJSONRequestBody defines body for PostTodos for application/json ContentType.
type PostTodosJSONRequestBody = TodoCreateInput

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 全TODO返却
	// (GET /todos)
	GetTodos(ctx echo.Context, params GetTodosParams) error
	// TODO登録
	// (POST /todos)
	PostTodos(ctx echo.Context) error
	// 指定TODO返却
	// (GET /todos/{todoId})
	GetTodosTodoId(ctx echo.Context, todoId TodoID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTodos converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodos(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTodosParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTodos(ctx, params)
	return err
}

// PostTodos converts echo context to params.
func (w *ServerInterfaceWrapper) PostTodos(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTodos(ctx)
	return err
}

// GetTodosTodoId converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodosTodoId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "todoId" -------------
	var todoId TodoID

	err = runtime.BindStyledParameterWithOptions("simple", "todoId", ctx.Param("todoId"), &todoId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter todoId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTodosTodoId(ctx, todoId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/todos", wrapper.GetTodos)
	router.POST(baseURL+"/todos", wrapper.PostTodos)
	router.GET(baseURL+"/todos/:todoId", wrapper.GetTodosTodoId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6yVz0/bSBTH/xX0do8mP5abbyxIq5zgkBviMMQvYVDsMeMJIoosEWfDIja7rVBRm7ZS",
	"e4CqFSJqqx5a2oo/ZnAK/0U1tuM4iQOk7S2237wfn+93XhpQYqbNLLSEA3oDnNImmiT4ucSRCDQWhXow",
	"0ClxagvKLNDh6uvz/sHD/uPTftcDDXCXmHYVQQfQoMy4SQToYBCB84KaCBqIuq0+O4JTqwKuBkVmMJWX",
	"VKsrZdDXGvA7xzLo8Ft22FA26iarosN2CpZdE+BqDbA5s5ELikGz6lAVBRrq4c5UcbCrQSk55m0nhzxc",
	"Dei9ChUMcF0NOG7XKFe9ramDyZrrMRu2sYUlAe56RGcpOdEofr/XubrYl61j2XojvbdJBQSvYZxyg7Eq",
	"EgvijJZAK0XO4sryimz2/P223/s0oqd/cdj/v3XTPZHekWx2pfev3POm6ZlUaIplwhyDeu1T//BZv9X2",
	"X74DbULQuNm75QxDXQ0EFarvu88Ug8BxccLjk5qE8xWMKei8o6uPe/2/H8jm2fX5E//gNELlncvWlyTP",
	"fOJ6UEsMQVJLYAU5aLA7X2HzFjHV26Do8qB8cTBbqnjSu5TeiWwdyNZZuoSTsqnxqVVmKUmZweYWVwsQ",
	"I41fzTnId4JWd5A7YXg+k8vkVJ/MRovYFHRYyOQyC0pVIjYDQbOCGSz4VcEUf/jt17L5SjZ7EdHry0f+",
	"fx+SnlPuICpa6QB/oSgGCVUJTkwUyJ1gkVCVbruGvA4aRCSr1KQKd2gBVd0ku9SsmaDnczkNTGpFTxOS",
	"uO66MoljM8sJrflHLjfmUGLbVVoKmstuOWqeRqIUFWg69zFl4OCwOuGc1EOFbuEUGNipmSbh9fCbehui",
	"U8ls5oipnv3W/XzTeT+d8CpzYsTqkqAj/mRGfabRZ1rqo7dRbTJ3gnz+l5ZPAxyt+TG0ilkILPgQmjnb",
	"EMFacKe6ut/5x+89lc1j6XVk88WM3o6WTrrD1cUaGlwMQkcBajPACP6sftbrP0I8BdIY/TAi6W3Xdb8H",
	"AAD//9AzODm7CAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
