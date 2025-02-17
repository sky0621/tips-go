# connect

## ref

https://connectrpc.com/docs/go/getting-started/

## curl

```
curl --header "Content-Type: application/json" --data '{}' http://localhost:8080/todo.v1.TodoService/ListTodo
{"todoList":[{"id":1,"msg":"Hello","completed":true},{"id":2,"msg":"Hey"}]}
```

```
curl --header "Content-Type: application/json" --data '{"id": 3}' http://localhost:8080/todo.v1.TodoService/GetTodo
{"todo":{"id":3,"msg":"Hello  3"}}

curl --header "Content-Type: application/json" --data '{"id": 429}' http://localhost:8080/todo.v1.TodoService/GetTodo
{"code":"resource_exhausted","message":"resource exhausted"}

curl --header "Content-Type: application/json" --data '{"id": 999}' http://localhost:8080/todo.v1.TodoService/GetTodo
{"code":"unknown","message":"unexpected id"}
```

## buf

### init

```
buf config init
```

### generate

```
buf generate
```
