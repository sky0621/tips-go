package main

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	todov1 "github.com/sky0621/tips-go/library_selection/communication/grpc/connect/gen/todo/v1"
	"github.com/sky0621/tips-go/library_selection/communication/grpc/connect/gen/todo/v1/todov1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/durationpb"
	"log"
	"net/http"
	"time"
)

type TodoServer struct {
}

func (t TodoServer) CreateTodo(ctx context.Context, c *connect.Request[todov1.CreateTodoRequest]) (*connect.Response[todov1.CreateTodoResponse], error) {
	// TODO implement me
	panic("implement me")
}

func (t TodoServer) GetTodo(ctx context.Context, c *connect.Request[todov1.GetTodoRequest]) (*connect.Response[todov1.GetTodoResponse], error) {
	todo, err := todoByID(c.Msg.GetId())
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&todov1.GetTodoResponse{Todo: todo})
	res.Header().Set("Todo-Version", "v1")
	return res, nil
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

func todoByID(id int32) (*todov1.Todo, error) {
	switch {
	case id >= 1 && id <= 10:
		return &todov1.Todo{
			Id:        id,
			Msg:       fmt.Sprintf("Hello %2d", id),
			Completed: id%2 == 0,
		}, nil
	case id == 400:
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid id"))
	case id == 1400:
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("invalid id"))
	case id == 2400:
		return nil, connect.NewError(connect.CodeOutOfRange, errors.New("invalid id"))
	case id == 401:
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	case id == 403:
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	case id == 404:
		return nil, connect.NewError(connect.CodeNotFound, errors.New("not found"))
	case id == 409:
		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("already exists"))
	case id == 1409:
		return nil, connect.NewError(connect.CodeAborted, errors.New("aborted"))
	case id == 429:
		return nil, connect.NewError(connect.CodeResourceExhausted, errors.New("resource exhausted"))
	case id == 499:
		return nil, connect.NewError(connect.CodeCanceled, errors.New("canceled"))
	case id == 500:
		return nil, connect.NewError(connect.CodeUnknown, errors.New("unknown error"))
	case id == 1500:
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	case id == 2500:
		return nil, connect.NewError(connect.CodeDataLoss, errors.New("data loss"))
	case id == 501:
		return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
	case id == 503:
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("unavailable"))
	case id == 504:
		ce := connect.NewError(connect.CodeDeadlineExceeded, errors.New("deadline exceeded"))
		retryInfo := &errdetails.RetryInfo{RetryDelay: durationpb.New(1 * time.Second)}
		time.Sleep(2 * time.Second)
		if detail, detailErr := connect.NewErrorDetail(retryInfo); detailErr != nil {
			ce.AddDetail(detail)
		}
		return nil, ce
	default:
		return nil, errors.New("unexpected id")
	}
}
