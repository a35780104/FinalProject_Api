package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

//Records cool
type Records struct{}

//NewRecords cool
func NewRecords() Records {
	return Records{}
}

//Get cool
// @Summary 獲取交易紀錄
// @Produce json
// @Param id path int true "紀錄ID"
// @Success 200 {object} model.Records "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/recordss/{id} [get]
func (a Records) Get(c *gin.Context) {
	param := service.RecordsRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	records, err := svc.GetRecords(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetRecords err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetRecordsFail)
		return
	}

	response.ToResponse(records)
	return
}

/*
func (a Records) List(c *gin.Context) {
	param := service.RecordsListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	recordss, totalRows, err := svc.GetRecordsList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetRecordsList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetRecordssFail)
		return
	}

	response.ToResponseList(recordss, totalRows)
	return
}
*/
// @Summary 建立交易紀錄
// @Produce json
// @Param product body string true "產品名稱"
// @Param saleprice body int true "產品單價"
// @Param salesum body int true "產品數量"
// @Param purchaser body string true "購買人"
// @Success 200 {object} model.Records "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/recordss [post]
func (a Records) Create(c *gin.Context) {
	param := service.CreateRecordsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateRecords(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateRecords err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateRecordsFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//Update cool
// @Summary 更新交易紀錄
// @Produce json
// @Param id body int true "紀錄ID"
// @Param product body string true "產品名稱"
// @Param saleprice body int true "產品售價"
// @Param salesum body int true "產品數量"
// @Param purchaser body string true "購買人"
// @Success 200 {object} model.Records "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/recordss/{id} [put]
func (a Records) Update(c *gin.Context) {
	param := service.UpdateRecordsRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateRecords(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateRecords err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateRecordsFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//Delete cool
// @Summary 删除交易紀錄
// @Produce  json
// @Param id path int true "紀錄ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/recordss/{id} [delete]
func (a Records) Delete(c *gin.Context) {
	param := service.DeleteRecordsRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteRecords(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteRecords err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteRecordsFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
