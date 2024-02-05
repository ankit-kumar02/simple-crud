package repository

import (
	"errors"
	"fmt"

	"github.com/ankit-kumar02/simple-crud/helper"
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/request"
	"gorm.io/gorm"
)

type BureauRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) BureauRepository {
	return &BureauRepositoryImpl{Db: Db}
}

func (t *BureauRepositoryImpl) Save(burea model.Bureau) {
	result := t.Db.Create(&burea)
	helper.ErrorPanic(result.Error)

}

func (t *BureauRepositoryImpl) Update(bureau model.Bureau) {
	var updateTag = request.UpdateBureauRequest{
		DiffNumM2M8: bureau.DiffNumM2M8,
		CountPlLoan: bureau.CountPlLoan,
	}
	result := t.Db.Model(&bureau).Where("bureau_reference_number = ?", bureau.BureauReferenceNumber).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t *BureauRepositoryImpl) Delete(bureauId string) {
	result := t.Db.Unscoped().Where("bureau_reference_number = ?", bureauId).Delete(&model.Bureau{})
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// If the error is not due to record not found, panic
		helper.ErrorPanic(result.Error)
	}

}

func (t *BureauRepositoryImpl) FindById(bureauId string) (model.Bureau, error) {
	var tag model.Bureau
	result := t.Db.Find(&tag, "bureau_reference_number = ?", bureauId)
	fmt.Println("repo", result)
	if result.RowsAffected == 0 {
		fmt.Println("user not found")
		return tag, errors.New("Bureau is not found")
	} else {
		fmt.Println("user  found")
		return tag, nil
	}
}
func (t *BureauRepositoryImpl) FindAll() []model.Bureau {
	var tags []model.Bureau
	results := t.Db.Find(&tags)
	helper.ErrorPanic(results.Error)
	return tags
}
