package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

type Member struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func (d *Dao) CreateMember(param *Member) (*model.Member, error) {
	member := model.Member{
		Name:    param.Name,
		Phone:   param.Phone,
		Address: param.Address,
		Email:   param.Email,
	}
	return member.Create(d.engine)
}

func (d *Dao) UpdateMember(param *Member) error {
	member := model.Member{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{}
	if param.Name != "" {
		values["name"] = param.Name
	}
	if param.Phone != "" {
		values["phone"] = param.Phone
	}
	if param.Address != "" {
		values["address"] = param.Address
	}
	if param.Email != "" {
		values["email"] = param.Email
	}

	return member.Update(d.engine, values)
}

func (d *Dao) GetMember(id uint32) (model.Member, error) {

	member := model.Member{Model: &model.Model{ID: id}}
	return member.Get(d.engine)
}

func (d *Dao) DeleteMember(id uint32) error {
	member := model.Member{Model: &model.Model{ID: id}}
	return member.Delete(d.engine)
}

// func (d *Dao) CountMemberListByTagID(id uint32, state uint8) (int, error) {
// 	member := model.Member{State: state}
// 	return member.CountByTagID(d.engine, id)
// }

// func (d *Dao) GetMemberListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.MemberRow, error) {
// 	member := model.Member{State: state}
// 	return member.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
// }
