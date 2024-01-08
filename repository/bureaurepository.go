package repository

import (
	"errors"

	"github.com/ankit-kumar02/simple-crud/helper"
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/request"
	"gorm.io/gorm"
)

type BureauRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) BureauRepository {
	return &BureauRepositoryImpl{DB: Db}
}

func (t BureauRepositoryImpl) Save(burea model.Bureau) {
	result := t.DB.Create(&burea)
	helper.ErrorPanic(result.Error)

}

func (t BureauRepositoryImpl) Update(bureau model.Bureau) {
	var updateTag = request.UpdateBureauRequest{
		UserID:                bureau.UserID,
		BureauReferenceNumber: bureau.BureauReferenceNumber,
	}
	result := t.DB.Model(&bureau).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t BureauRepositoryImpl) Delete(bureauId int) {
	var tags model.Bureau
	result := t.DB.Where("id = ?", bureauId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

func (t BureauRepositoryImpl) FindById(tagsId int) (model.Bureau, error) {
	var tag model.Bureau
	result := t.DB.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("Bureau is not found")
	}
}

func (t BureauRepositoryImpl) FindAll() []model.Bureau {
	var tags []model.Bureau
	results := t.DB.Find(&tags)
	helper.ErrorPanic(results.Error)
	return tags
}
