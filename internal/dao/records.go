package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

type Records struct {
	ID        uint32 `json:"id"`
	Product   string `json:"product"`
	Saleprice uint32 `json:"saleprice"`
	Salesum   uint32 `json:"salesum"`
	Purchaser string `json:"purchaser"`
}

func (d *Dao) CreateRecords(param *Records) (*model.Records, error) {
	records := model.Records{
		Product:   param.Product,
		Saleprice: param.Saleprice,
		Salesum:   param.Salesum,
		Purchaser: param.Purchaser,
	}
	return records.Create(d.engine)
}

func (d *Dao) UpdateRecords(param *Records) error {
	records := model.Records{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{}
	if param.Product != "" {
		values["product"] = param.Product
	}

	values["paleprice"] = param.Saleprice

	values["palesum"] = param.Salesum

	if param.Purchaser != "" {
		values["purchaser"] = param.Purchaser
	}

	return records.Update(d.engine, values)
}

func (d *Dao) GetRecords(id uint32) (model.Records, error) {
	records := model.Records{Model: &model.Model{ID: id}}
	return records.Get(d.engine)
}

func (d *Dao) DeleteRecords(id uint32) error {
	records := model.Records{Model: &model.Model{ID: id}}
	return records.Delete(d.engine)
}

// func (d *Dao) CountRecordsListByTagID(id uint32, state uint8) (int, error) {
// 	records := model.Records{State: state}
// 	return records.CountByTagID(d.engine, id)
// }

// func (d *Dao) GetRecordsListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.RecordsRow, error) {
// 	records := model.Records{State: state}
// 	return records.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
// }
