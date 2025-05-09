// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package infra

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateClassBatch(ctx context.Context) (sql.Result, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (int64, error)
	CreateDepartmentsBatch(ctx context.Context) (sql.Result, error)
	CreateEmployeesBatch(ctx context.Context) (sql.Result, error)
	CreateGrade(ctx context.Context) (int64, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (int64, error)
	CreateSchool(ctx context.Context) (int64, error)
	CreateStudentBatch(ctx context.Context) (sql.Result, error)
	CreateUser(ctx context.Context, name string) (int64, error)
	CreateUsersBatch(ctx context.Context, arg CreateUsersBatchParams) (sql.Result, error)
	DeleteComment(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetComment(ctx context.Context, id int64) (Comment, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetSchoolByID(ctx context.Context, schoolID int64) (School, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListClassByGradeID(ctx context.Context, gradeID sql.NullInt64) ([]Class, error)
	ListClassInGradeIDs(ctx context.Context, gradeids []sql.NullInt64) ([]Class, error)
	ListCommentsByPost(ctx context.Context, postID sql.NullInt64) ([]Comment, error)
	ListEmployeesOrderByDepartmentIdAsc(ctx context.Context) ([]Employee, error)
	ListEmployeesOrderByDepartmentIdDesc(ctx context.Context) ([]Employee, error)
	ListEmployeesOrderByJoinDateAsc(ctx context.Context) ([]Employee, error)
	ListEmployeesOrderByJoinDateDesc(ctx context.Context) ([]Employee, error)
	ListEmployeesOrderBySalaryAsc(ctx context.Context) ([]Employee, error)
	ListEmployeesOrderBySalaryDesc(ctx context.Context) ([]Employee, error)
	// これは機能しない。
	ListEmployeesOrderByXXXX(ctx context.Context, dollar_1 interface{}) ([]Employee, error)
	ListGradeBySchoolID(ctx context.Context, schoolID sql.NullInt64) ([]Grade, error)
	ListPostsByLikeTitle(ctx context.Context, title string) ([]Post, error)
	ListPostsByUser(ctx context.Context, userID sql.NullInt64) ([]Post, error)
	ListRecentCommentByPosts(ctx context.Context) ([]ListRecentCommentByPostsRow, error)
	ListStudentByClassID(ctx context.Context, classID sql.NullInt64) ([]Student, error)
	ListStudentInClassIDs(ctx context.Context, classids []sql.NullInt64) ([]Student, error)
	ListStudentsWithClassWithGradeWithSchool(ctx context.Context, schoolID int64) ([]ListStudentsWithClassWithGradeWithSchoolRow, error)
	ListUserWithPostAndComments(ctx context.Context) ([]ListUserWithPostAndCommentsRow, error)
	ListUsers(ctx context.Context) ([]User, error)
	// https://docs.sqlc.dev/en/latest/howto/select.html#mysql-and-sqlite
	ListUsersByIDs(ctx context.Context, ids []int64) ([]User, error)
	ListUsersWithPostAndCommentCount(ctx context.Context) ([]ListUsersWithPostAndCommentCountRow, error)
	ListUsersWithRecentPostAndCommentCount(ctx context.Context) ([]ListUsersWithRecentPostAndCommentCountRow, error)
	ListWithComplexQuery(ctx context.Context, arg ListWithComplexQueryParams) ([]ListWithComplexQueryRow, error)
	MaxUsersID(ctx context.Context) (interface{}, error)
	OrgListUsersByIDs(ctx context.Context, id int64) ([]User, error)
	Relations(ctx context.Context) ([]RelationsRow, error)
	SearchEmployees(ctx context.Context) ([]Employee, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) error
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error
}

var _ Querier = (*Queries)(nil)
