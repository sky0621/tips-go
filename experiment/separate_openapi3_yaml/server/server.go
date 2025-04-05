package server

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/experiment/separate_openapi3_yaml/api"
)

var _ api.ServerInterface = (*Server)(nil)

func New() *Server {
	return &Server{}
}

type Server struct {
}

func (s Server) PostSkillrecords(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetSkills(ctx echo.Context, params api.GetSkillsParams) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PostSkills(ctx echo.Context, params api.PostSkillsParams) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetSkillsBySkillId(ctx echo.Context, bySkillId api.SkillId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetSkilltags(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PostSkilltags(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetSkilltagsBySkillTagId(ctx echo.Context, bySkillTagId api.SkillTagId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsers(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PostUsers(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) DeleteUsersByUserId(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdActivities(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdActivities(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdAppeals(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdAppeals(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdAttribute(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdAttribute(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdCareergroups(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PostUsersByUserIdCareergroups(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId api.UserId, byCareerGroupId api.CareerGroupId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId api.UserId, byCareerGroupId api.CareerGroupId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId api.UserId, byCareerGroupId api.CareerGroupId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PostUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx echo.Context, byUserId api.UserId, byCareerGroupId api.CareerGroupId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx echo.Context, byUserId api.UserId, byCareerGroupId api.CareerGroupId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdCareergroupsByCareerGroupIdCareersByCareerId(ctx echo.Context, byUserId api.UserId, byCareerGroupId api.CareerGroupId, byCareerId api.CareerId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdNotes(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PostUsersByUserIdNotes(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) DeleteUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId api.UserId, byNoteId api.NoteId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId api.UserId, byNoteId api.NoteId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdNotesByNoteIdItems(ctx echo.Context, byUserId api.UserId, byNoteId api.NoteId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdQualifications(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdQualifications(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdSkills(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersByUserIdSolutions(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) PutUsersByUserIdSolutions(ctx echo.Context, byUserId api.UserId) error {
	//TODO implement me
	panic("implement me")
}
