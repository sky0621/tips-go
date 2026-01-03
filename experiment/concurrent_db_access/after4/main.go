package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		user  User
		group Group
		todos Todos
	)

	var wg sync.WaitGroup
	var errOnce sync.Once
	var firstErr error

	setErr := func(err error) {
		if err == nil {
			return
		}
		errOnce.Do(func() {
			firstErr = err
			cancel()
		})
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		u, err := getUser(ctx)
		if err != nil {
			setErr(err)
			return
		}
		user = u
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		g, err := getGroup(ctx)
		if err != nil {
			setErr(err)
			return
		}
		group = g
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		t, err := listTodo(ctx)
		if err != nil {
			setErr(err)
			return
		}
		todos = t
	}()

	wg.Wait()

	if firstErr != nil {
		fmt.Println(firstErr.Error())
	} else {
		result := Result{
			UserName:  user.Name,
			GroupName: group.Name,
			Contents:  todos.Contents(),
		}
		fmt.Println(result)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

type Result struct {
	UserName  string
	GroupName string
	Contents  []string
}

type User struct {
	ID   int
	Name string
}

func getUser(ctx context.Context) (User, error) {
	if err := sleepWithContext(ctx, 1*time.Second); err != nil {
		return User{}, err
	}
	user := User{ID: 1, Name: "John"}
	// 正常系
	return user, nil
	// 異常系
	// return User{}, errors.New("user not found")
}

type Group struct {
	ID   int
	Name string
}

func getGroup(ctx context.Context) (Group, error) {
	if err := sleepWithContext(ctx, 2*time.Second); err != nil {
		return Group{}, err
	}
	group := Group{ID: 1, Name: "GroupA"}
	// 正常系
	return group, nil
	// 異常系
	// return Group{}, errors.New("group does not exist")
}

type Todos []Todo

type Todo struct {
	ID      int
	Content string
}

func (t Todos) Contents() []string {
	contents := make([]string, 0, len(t))
	for _, todo := range t {
		contents = append(contents, todo.Content)
	}
	return contents
}

func listTodo(ctx context.Context) (Todos, error) {
	if err := sleepWithContext(ctx, 4*time.Second); err != nil {
		return nil, err
	}
	todos := Todos{
		{ID: 1, Content: "hello"},
		{ID: 2, Content: "bye"},
	}
	// 正常系
	return todos, nil
	// 異常系
	// return nil, errors.New("todos does not exist")
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
