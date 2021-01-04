package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Member struct {
	*Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type MemberSwagger struct {
	List  []*Member
	Pager *app.Pager
}

func (a Member) TableName() string {
	return "se_member"
}

func (a Member) Create(db *gorm.DB) (*Member, error) {

	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Member) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ?", a.ID).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a Member) Get(db *gorm.DB) (Member, error) {
	var member Member

	db = db.Where("id = ?", a.ID)

	err := db.First(&member).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return member, err
	}

	return member, nil
}

func (a Member) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", a.ID).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

/*
type MemberRow struct {
	MembereID     uint32
	TagID         uint32
	TagName       string
	MemberTitle  string
	MemberDesc   string
	CoverImageUrl string
	Content       string
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
*/
