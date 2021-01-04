package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

// Product cool
type Product struct{}

//NewProduct cool
func NewProduct() Product {
	return Product{}
}

//Get cool
// @Summary 獲取單一產品
// @Produce json
// @Param id path int true "產品ID"
// @Success 200 {object} model.Product "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/products/{id} [get]
func (a Product) Get(c *gin.Context) {
	param := service.ProductRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	product, err := svc.GetProduct(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetProductFail)
		return
	}

	response.ToResponse(product)
	return
}

/*
func (a Product) List(c *gin.Context) {
	param := service.ProductListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	products, totalRows, err := svc.GetProductList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetProductList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetProductsFail)
		return
	}

	response.ToResponseList(products, totalRows)
	return
}


*/

//Create cool
// @Summary 建立產品
// @Produce json
// @Param name body string true "產品名稱"
// @Param stock body int true "存貨數量"
// @Param saleprice body int true "產品售價"
// @Param supplier body string true "供應商"
// @Success 200 {object} model.Product "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/products [post]
func (a Product) Create(c *gin.Context) {
	param := service.CreateProductRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateProduct(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateProductFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//Update cool
// @Summary 更新產品資料
// @Produce json
// @Param id body int true "產品ID"
// @Param name body string true "產品名稱"
// @Param stock body int true "存貨數量"
// @Param saleprice body int true "產品售價"
// @Param supplier body string true "供應商"
// @Success 200 {object} model.Product "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/products/{id} [put]
func (a Product) Update(c *gin.Context) {
	param := service.UpdateProductRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateProduct(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateProductFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//Delete cool
// @Summary 删除產品資料
// @Produce  json
// @Param id path int true "產品ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "内部錯誤"
// @Router /api/v1/products/{id} [delete]
func (a Product) Delete(c *gin.Context) {
	param := service.DeleteProductRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteProduct(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteProduct err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteProductFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
