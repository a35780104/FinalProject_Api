package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/dao"
)

type ProductRequest struct {
	ID uint32 `form:"id"  binding:"required,gte=1"`
}

type ProductListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateProductRequest struct {
	Name      string `form:"name"  binding:"required,min=2,max=100"`
	Stock     uint32 `form:"stock"  binding:"required"`
	Saleprice uint32 `form:"saleprice"  binding:"required"`
	Supplier  string `form:"supplier"  binding:"required,min=2,max=100"`
}

type UpdateProductRequest struct {
	ID        uint32 `form:"id"  binding:"required,gte=1"`
	Name      string `form:"name"  binding:"required,min=2,max=100"`
	Stock     uint32 `form:"stock"`
	Saleprice uint32 `form:"saleprice"`
	Supplier  string `form:"supplier"  binding:"required,min=2,max=100"`
}

type DeleteProductRequest struct {
	ID uint32 `form:"id"`
}

type Product struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Stock     uint32 `json:"stock"`
	Saleprice uint32 `json:"saleprice"`
	Supplier  string `json:"supplier"`
}

func (svc *Service) GetProduct(param *ProductRequest) (*Product, error) {
	product, err := svc.dao.GetProduct(param.ID)
	if err != nil {
		return nil, err
	}
	/*
			productTag, err := svc.dao.GetProductTagByAID(product.ID)
			if err != nil {
				return nil, err
			}

		tag, err := svc.dao.GetTag(productTag.TagID, model.STATE_OPEN)
		if err != nil {
			return nil, err
		}
	*/
	return &Product{
		ID:        product.ID,
		Name:      product.Name,
		Stock:     product.Stock,
		Saleprice: product.Saleprice,
		Supplier:  product.Supplier,
	}, nil
}

/*
func (svc *Service) GetProductList(param *ProductListRequest, pager *app.Pager) ([]*Product, int, error) {
	productCount, err := svc.dao.CountProductListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	products, err := svc.dao.GetProductListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var productList []*Product
	for _, product := range products {
		productList = append(productList, &Product{
			ID:            product.ProductID,
			Title:         product.ProductTitle,
			Desc:          product.ProductDesc,
			Content:       product.Content,
			CoverImageUrl: product.CoverImageUrl,
			Tag:           &model.Tag{Model: &model.Model{ID: product.TagID}, Name: product.TagName},
		})
	}

	return productList, productCount, nil
}
*/
func (svc *Service) CreateProduct(param *CreateProductRequest) error {
	// product, err := svc.dao.CreateProduct(&dao.Product{
	_, err := svc.dao.CreateProduct(&dao.Product{
		Name:      param.Name,
		Stock:     param.Stock,
		Saleprice: param.Saleprice,
		Supplier:  param.Supplier,
	})
	if err != nil {
		return err
	}
	/*
		err = svc.dao.CreateProductTag(product.ID, param.TagID, param.CreatedBy)
		if err != nil {
			return err
		}
	*/
	return nil
}

func (svc *Service) UpdateProduct(param *UpdateProductRequest) error {
	err := svc.dao.UpdateProduct(&dao.Product{
		ID:        param.ID,
		Name:      param.Name,
		Stock:     param.Stock,
		Saleprice: param.Saleprice,
		Supplier:  param.Supplier,
	})
	if err != nil {
		return err
	}
	/*
		err = svc.dao.UpdateProductTag(param.ID, param.TagID, param.ModifiedBy)
		if err != nil {
			return err
		}
	*/
	return nil
}

func (svc *Service) DeleteProduct(param *DeleteProductRequest) error {
	err := svc.dao.DeleteProduct(param.ID)
	if err != nil {
		return err
	}
	/*
		err = svc.dao.DeleteProductTag(param.ID)
		if err != nil {
			return err
		}
	*/
	return nil
}
