package infra

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
	builder.WriteString(fmt.Sprintf("UpdatedAt: %s", u.UpdatedAt))
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
	builder.WriteString(fmt.Sprintf("UpdatedAt: %v", p.UpdatedAt))
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
	builder.WriteString(fmt.Sprintf("UpdatedAt: %v", c.UpdatedAt))
	builder.WriteString("}")
	return builder.String()
}

func (u ListUserWithPostAndCommentsRow) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString(fmt.Sprintf("UserID: %d, ", u.UserID))
	builder.WriteString(fmt.Sprintf("UserName: %s, ", u.UserName))
	builder.WriteString(fmt.Sprintf("UserCreatedAt: %v, ", u.UserCreatedAt))
	builder.WriteString(fmt.Sprintf("UserUpdatedAt: %v, ", u.UserUpdatedAt))
	builder.WriteString(fmt.Sprintf("PostID: %d, ", u.PostID))
	builder.WriteString(fmt.Sprintf("PostTitle: %s, ", u.PostTitle))
	builder.WriteString(fmt.Sprintf("PostContent: %v, ", u.PostContent))
	builder.WriteString(fmt.Sprintf("PostCreatedAt: %v, ", u.PostCreatedAt))
	builder.WriteString(fmt.Sprintf("PostUpdatedAt: %v, ", u.PostUpdatedAt))
	builder.WriteString(fmt.Sprintf("CommentID: %d, ", u.CommentID))
	builder.WriteString(fmt.Sprintf("CommentContent: %v, ", u.CommentContent))
	builder.WriteString(fmt.Sprintf("CommentCreatedAt: %v, ", u.CommentCreatedAt))
	builder.WriteString(fmt.Sprintf("CommentUpdatedAt: %v", u.CommentUpdatedAt))
	builder.WriteString("}")
	return builder.String()
}

func (u ListUsersWithPostAndCommentCountRow) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString(fmt.Sprintf("UserID: %d, ", u.UserID))
	builder.WriteString(fmt.Sprintf("UserName: %s, ", u.UserName))
	builder.WriteString(fmt.Sprintf("PostCount: %d, ", u.PostCount))
	builder.WriteString(fmt.Sprintf("CommentCount: %d", u.CommentCount))
	builder.WriteString("}")
	return builder.String()
}

func (u ListUsersWithRecentPostAndCommentCountRow) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString(fmt.Sprintf("UserID: %d, ", u.UserID))
	builder.WriteString(fmt.Sprintf("UserName: %s, ", u.UserName))
	builder.WriteString(fmt.Sprintf("RecentPosts: %d, ", u.RecentPosts))
	builder.WriteString(fmt.Sprintf("RecentComments: %d", u.RecentComments))
	builder.WriteString("}")
	return builder.String()
}

func (c ListRecentCommentByPostsRow) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString(fmt.Sprintf("PostID: %d, ", c.PostID))
	builder.WriteString(fmt.Sprintf("PostTitle: %s, ", c.PostTitle))
	builder.WriteString(fmt.Sprintf("AuthorName: %s, ", c.AuthorName))
	builder.WriteString(fmt.Sprintf("LatestCommentContent: %v, ", c.LatestCommentContent))
	builder.WriteString(fmt.Sprintf("LatestCommenterName: %v, ", c.LatestCommenterName))
	builder.WriteString(fmt.Sprintf("LatestCommentCreatedAt: %v, ", c.LatestCommentCreatedAt))
	builder.WriteString("}")
	return builder.String()
}
