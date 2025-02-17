package main

import (
	"connectrpc.com/connect"
	"context"
	todov1 "github.com/sky0621/tips-go/library_selection/communication/grpc/connect/gen/todo/v1"
	"github.com/sky0621/tips-go/library_selection/communication/grpc/connect/gen/todo/v1/todov1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type TodoServer struct {
}

func (t TodoServer) CreateTodo(ctx context.Context, c *connect.Request[todov1.CreateTodoRequest]) (*connect.Response[todov1.CreateTodoResponse], error) {
	// TODO implement me
	panic("implement me")
}

func (t TodoServer) ListTodo(ctx context.Context, c *connect.Request[todov1.ListTodoRequest]) (*connect.Response[todov1.ListTodoResponse], error) {
	res := connect.NewResponse(&todov1.ListTodoResponse{TodoList: []*todov1.Todo{
		{Id: 1, Msg: "Hello", Completed: true},
		{Id: 2, Msg: "Hey", Completed: false},
	}})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
}

func main() {
	svr := &TodoServer{}
	mux := http.NewServeMux()
	path, handler := todov1connect.NewTodoServiceHandler(svr)
	mux.Handle(path, handler)
	err := http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
	if err != nil {
		log.Fatal(err)
	}
}
