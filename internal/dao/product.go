package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

type Product struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Stock     uint32 `json:"stock"`
	Saleprice uint32 `json:"saleprice"`
	Supplier  string `json:"supplier"`
}

func (d *Dao) CreateProduct(param *Product) (*model.Product, error) {
	product := model.Product{
		Name:      param.Name,
		Stock:     param.Stock,
		Saleprice: param.Saleprice,
		Supplier:  param.Supplier,
	}
	return product.Create(d.engine)
}

func (d *Dao) UpdateProduct(param *Product) error {
	product := model.Product{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{}
	if param.Name != "" {
		values["name"] = param.Name
	}

	values["stock"] = param.Stock

	values["saleprice"] = param.Saleprice

	if param.Supplier != "" {
		values["supplier"] = param.Supplier
	}

	return product.Update(d.engine, values)
}

func (d *Dao) GetProduct(id uint32) (model.Product, error) {
	product := model.Product{Model: &model.Model{ID: id}}
	return product.Get(d.engine)
}

func (d *Dao) DeleteProduct(id uint32) error {
	product := model.Product{Model: &model.Model{ID: id}}
	return product.Delete(d.engine)
}

// func (d *Dao) CountProductListByTagID(id uint32, state uint8) (int, error) {
// 	product := model.Product{State: state}
// 	return product.CountByTagID(d.engine, id)
// }

// func (d *Dao) GetProductListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ProductRow, error) {
// 	product := model.Product{State: state}
// 	return product.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
// }
