package models

import (
	"fmt"
	"strings"
)

func (u User) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString(fmt.Sprintf("ID: %d, ", u.ID))
	builder.WriteString(fmt.Sprintf("Name: %s, ", u.Name))
	builder.WriteString(fmt.Sprintf("CreatedAt: %s, ", u.CreatedAt))
	builder.WriteString(fmt.Sprintf("UpdatedAt: %s, ", u.UpdatedAt))
	builder.WriteString("Posts: [")
	for idx, post := range u.Posts {
		builder.WriteString(fmt.Sprintf("[%d]{%s} ", idx, post.String()))
	}
	builder.WriteString("]")
	builder.WriteString("}")
	return builder.String()
}

func (p Post) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString(fmt.Sprintf("ID: %d, ", p.ID))
	builder.WriteString(fmt.Sprintf("Title: %s, ", p.Title))
	builder.WriteString(fmt.Sprintf("Content: %s, ", p.Content))
	builder.WriteString(fmt.Sprintf("UserID: %d, ", p.UserID))
	builder.WriteString(fmt.Sprintf("CreatedAt: %v, ", p.CreatedAt))
	builder.WriteString(fmt.Sprintf("UpdatedAt: %v, ", p.UpdatedAt))
	builder.WriteString("Comments: [")
	for idx, comment := range p.Comments {
		builder.WriteString(fmt.Sprintf("[%d]{%s} ", idx, comment.String()))
	}
	builder.WriteString("]")
	builder.WriteString("}")
	return builder.String()
}

func (c Comment) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString(fmt.Sprintf("ID: %d, ", c.ID))
	builder.WriteString(fmt.Sprintf("Content: %s, ", c.Content))
	builder.WriteString(fmt.Sprintf("UserID: %d, ", c.UserID))
	builder.WriteString(fmt.Sprintf("PostID: %d, ", c.PostID))
	builder.WriteString(fmt.Sprintf("CreatedAt: %v, ", c.CreatedAt))
	builder.WriteString(fmt.Sprintf("UpdatedAt: %v, ", c.UpdatedAt))
	builder.WriteString("}")
	return builder.String()
}
