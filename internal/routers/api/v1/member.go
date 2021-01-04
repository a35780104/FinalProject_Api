package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

// Member cool
type Member struct{}

//NewMember cool
func NewMember() Member {
	return Member{}
}

// @Summary 獲取一個用戶
// @Produce json
// @Param id path int true "用戶ID"
// @Success 200 {object} model.Member "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/members/{id} [get]
func (a Member) Get(c *gin.Context) {

	param := service.MemberRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	fmt.Print("clear")
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	member, err := svc.GetMember(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetMember err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetMemberFail)
		return
	}

	response.ToResponse(member)
	return
}

// func (a Member) List(c *gin.Context) {
// 	param := service.MemberListRequest{}
// 	response := app.NewResponse(c)
// 	valid, errs := app.BindAndValid(c, &param)
// 	if !valid {
// 		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
// 		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
// 		return
// 	}

// 	svc := service.New(c.Request.Context())
// 	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
// 	members, totalRows, err := svc.GetMemberList(&param, &pager)
// 	if err != nil {
// 		global.Logger.Errorf(c, "svc.GetMemberList err: %v", err)
// 		response.ToErrorResponse(errcode.ErrorGetMembersFail)
// 		return
// 	}

// 	response.ToResponseList(members, totalRows)
// 	return
// }

// @Summary 註冊用戶
// @Produce json
// @Param name body string true "用戶名稱"
// @Param phone body string true "連絡電話"
// @Param address body string true "聯絡地址"
// @Param email body string true "電子信箱"
// @Success 200 {object} model.Member "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/members [post]
func (a Member) Create(c *gin.Context) {
	param := service.CreateMemberRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.CreateMember(&param)

	if err != nil {
		global.Logger.Errorf(c, "svc.CreateMember err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateMemberFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新用戶資料
// @Produce json
// @Param id path int true "用戶ID"
// @Param name body string true "用戶名稱"
// @Param phone body string true "連絡電話"
// @Param address body string true "聯絡地址"
// @Param email body string true "電子信箱"
// @Success 200 {object} model.Member "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/members/{id} [put]
func (a Member) Update(c *gin.Context) {
	param := service.UpdateMemberRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateMember(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateMember err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateMemberFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除用戶
// @Produce  json
// @Param id path int true "用戶ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/members/{id} [delete]
func (a Member) Delete(c *gin.Context) {
	param := service.DeleteMemberRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteMember(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteMember err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteMemberFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
