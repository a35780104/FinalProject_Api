package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Records struct {
	*Model
	Product   string `json:"product"`
	Saleprice uint32 `json:"saleprice"`
	Salesum   uint32 `json:"salesum"`
	Purchaser string `json:"purchaser"`
}

type RecordsSwagger struct {
	List  []*Records
	Pager *app.Pager
}

func (a Records) TableName() string {
	return "se_records"
}

func (a Records) Create(db *gorm.DB) (*Records, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Records) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ?", a.ID).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a Records) Get(db *gorm.DB) (Records, error) {
	var records Records
	db = db.Where("id = ?", a.ID)
	err := db.First(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return records, err
	}

	return records, nil
}

func (a Records) Delete(db *gorm.DB) error {
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
