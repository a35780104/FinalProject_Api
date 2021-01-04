package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/dao"
)

type RecordsRequest struct {
	ID uint32 `form:"id"  binding:"required,gte=1"`
}

type RecordsListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateRecordsRequest struct {
	Product   string `form:"product"  binding:"required,min=2,max=100"`
	Saleprice uint32 `form:"saleprice"  binding:"required,min=1"`
	Salesum   uint32 `form:"salesum"  binding:"required,min=1"`
	Purchaser string `form:"purchaser"  binding:"required,min=2,max=100"`
}

type UpdateRecordsRequest struct {
	ID        uint32 `form:"id"  binding:"required,gte=1"`
	Product   string `form:"product"  binding:"required,min=2,max=100"`
	Saleprice uint32 `form:"saleprice"  binding:"required,min=1"`
	Salesum   uint32 `form:"salesum"  binding:"required,min=1"`
	Purchaser string `form:"purchaser"  binding:"required,min=2,max=100"`
}

type DeleteRecordsRequest struct {
	ID uint32 `form:"id"`
}

type Records struct {
	ID        uint32 `json:"id"`
	Product   string `json:"product"`
	Saleprice uint32 `json:"saleprice"`
	Salesum   uint32 `json:"salesum"`
	Purchaser string `json:"purchaser"`
}

func (svc *Service) GetRecords(param *RecordsRequest) (*Records, error) {
	records, err := svc.dao.GetRecords(param.ID)
	if err != nil {
		return nil, err
	}
	/*
			recordsTag, err := svc.dao.GetRecordsTagByAID(records.ID)
			if err != nil {
				return nil, err
			}

		tag, err := svc.dao.GetTag(recordsTag.TagID, model.STATE_OPEN)
		if err != nil {
			return nil, err
		}
	*/
	return &Records{
		ID:        records.ID,
		Product:   records.Product,
		Saleprice: records.Saleprice,
		Salesum:   records.Salesum,
		Purchaser: records.Purchaser,
	}, nil
}

/*
func (svc *Service) GetRecordsList(param *RecordsListRequest, pager *app.Pager) ([]*Records, int, error) {
	recordsCount, err := svc.dao.CountRecordsListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	recordss, err := svc.dao.GetRecordsListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var recordsList []*Records
	for _, records := range recordss {
		recordsList = append(recordsList, &Records{
			ID:            records.RecordsID,
			Title:         records.RecordsTitle,
			Desc:          records.RecordsDesc,
			Content:       records.Content,
			CoverImageUrl: records.CoverImageUrl,
			Tag:           &model.Tag{Model: &model.Model{ID: records.TagID}, Name: records.TagName},
		})
	}

	return recordsList, recordsCount, nil
}
*/
func (svc *Service) CreateRecords(param *CreateRecordsRequest) error {
	// records, err := svc.dao.CreateRecords(&dao.Records{
	_, err := svc.dao.CreateRecords(&dao.Records{
		Product:   param.Product,
		Saleprice: param.Saleprice,
		Salesum:   param.Salesum,
		Purchaser: param.Purchaser,
	})
	if err != nil {
		return err
	}
	/*
		err = svc.dao.CreateRecordsTag(records.ID, param.TagID, param.CreatedBy)
		if err != nil {
			return err
		}
	*/
	return nil
}

func (svc *Service) UpdateRecords(param *UpdateRecordsRequest) error {
	err := svc.dao.UpdateRecords(&dao.Records{
		ID:        param.ID,
		Product:   param.Product,
		Saleprice: param.Saleprice,
		Salesum:   param.Salesum,
		Purchaser: param.Purchaser,
	})
	if err != nil {
		return err
	}
	/*
		err = svc.dao.UpdateRecordsTag(param.ID, param.TagID, param.ModifiedBy)
		if err != nil {
			return err
		}
	*/
	return nil
}

func (svc *Service) DeleteRecords(param *DeleteRecordsRequest) error {
	err := svc.dao.DeleteRecords(param.ID)
	if err != nil {
		return err
	}
	/*
		err = svc.dao.DeleteRecordsTag(param.ID)
		if err != nil {
			return err
		}
	*/
	return nil
}
