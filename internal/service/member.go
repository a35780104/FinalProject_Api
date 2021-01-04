package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/dao"
)

type MemberRequest struct {
	ID uint32 `form:"id"  binding:"required,gte=1"`
}

type MemberListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateMemberRequest struct {
	Name    string `form:"name" binding:"required,min=2,max=100"`
	Phone   string `form:"phone" binding:"required,min=2,max=100"`
	Address string `form:"address" binding:"required,min=2,max=100"`
	Email   string `form:"email" binding:"required,min=2,max=100"`
}

type UpdateMemberRequest struct {
	ID      uint32 `form:"id" binding:"required,gte=1"`
	Name    string `form:"name" binding:"required,min=2,max=100"`
	Phone   string `form:"phone" binding:"required,min=2,max=100"`
	Address string `form:"address" binding:"required,min=2,max=100"`
	Email   string `form:"email" binding:"required,min=2,max=100"`
}

type DeleteMemberRequest struct {
	ID uint32 `form:"id"`
}

type Member struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func (svc *Service) GetMember(param *MemberRequest) (*Member, error) {

	member, err := svc.dao.GetMember(param.ID)

	if err != nil {
		return nil, err
	}
	/*
			memberTag, err := svc.dao.GetMemberTagByAID(member.ID)
			if err != nil {
				return nil, err
			}

		tag, err := svc.dao.GetTag(memberTag.TagID, model.STATE_OPEN)
		if err != nil {
			return nil, err
		}
	*/
	return &Member{
		ID:      member.ID,
		Name:    member.Name,
		Phone:   member.Phone,
		Address: member.Address,
		Email:   member.Email,
	}, nil
}

/*
func (svc *Service) GetMemberList(param *MemberListRequest, pager *app.Pager) ([]*Member, int, error) {
	memberCount, err := svc.dao.CountMemberListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	members, err := svc.dao.GetMemberListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var memberList []*Member
	for _, member := range members {
		memberList = append(memberList, &Member{
			ID:            member.MemberID,
			Title:         member.MemberTitle,
			Desc:          member.MemberDesc,
			Content:       member.Content,
			CoverImageUrl: member.CoverImageUrl,
			Tag:           &model.Tag{Model: &model.Model{ID: member.TagID}, Name: member.TagName},
		})
	}

	return memberList, memberCount, nil
}
*/
func (svc *Service) CreateMember(param *CreateMemberRequest) error {
	// member, err := svc.dao.CreateMember(&dao.Member{
	_, err := svc.dao.CreateMember(&dao.Member{
		Name:    param.Name,
		Phone:   param.Phone,
		Address: param.Address,
		Email:   param.Email,
	})
	if err != nil {
		return err
	}
	/*
		err = svc.dao.CreateMemberTag(member.ID, param.TagID, param.CreatedBy)
		if err != nil {
			return err
		}
	*/
	return nil
}

func (svc *Service) UpdateMember(param *UpdateMemberRequest) error {
	err := svc.dao.UpdateMember(&dao.Member{
		ID:      param.ID,
		Name:    param.Name,
		Phone:   param.Phone,
		Address: param.Address,
		Email:   param.Email,
	})
	if err != nil {
		return err
	}
	/*
		err = svc.dao.UpdateMemberTag(param.ID, param.TagID, param.ModifiedBy)
		if err != nil {
			return err
		}
	*/
	return nil
}

func (svc *Service) DeleteMember(param *DeleteMemberRequest) error {
	err := svc.dao.DeleteMember(param.ID)
	if err != nil {
		return err
	}
	/*
		err = svc.dao.DeleteMemberTag(param.ID)
		if err != nil {
			return err
		}
	*/
	return nil
}
