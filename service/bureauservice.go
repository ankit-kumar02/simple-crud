package service

import (
	"encoding/json"

	"github.com/ankit-kumar02/simple-crud/helper"
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/repository"
	"github.com/ankit-kumar02/simple-crud/request"
	"github.com/ankit-kumar02/simple-crud/resource"
)

type BureauServiceImpl struct {
	BureauRepository repository.BureauRepository
}

func NewBureauServiceImple(bureaurepository repository.BureauRepository) BureauService {
	return &BureauServiceImpl{BureauRepository: bureaurepository}

}
func (b BureauServiceImpl) Create(bureau request.CreateBureauRequest) {
	data := bureau.NewBureauVariable

	// Convert data to JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic("Error marshaling JSON")
	}
	tagModel := model.Bureau{
		UserID:                bureau.UserID,
		BureauReferenceNumber: helper.GenerateBureauReference(),
		UserReferenceNumber:   helper.GenrateUserReference(),
		DiffNumM2M8:           bureau.DiffNumM2M8,
		CountPlLoan:           bureau.CountPlLoan,
		NewBureauVariable:     string(jsonData),
	}
	b.BureauRepository.Save(tagModel)

}

func (b BureauServiceImpl) Update(bureau request.UpdateBureauRequest) {
	tagData, err := b.BureauRepository.FindById(bureau.BureauReferenceNumber)

	helper.ErrorPanic(err)
	tagData.BureauReferenceNumber = bureau.BureauReferenceNumber
	b.BureauRepository.Update(tagData)
}

func (t BureauServiceImpl) Delete(tagId string) {
	t.BureauRepository.Delete(tagId)
}

func (t BureauServiceImpl) FindById(tagId string) (tagResponse resource.BureauResponse) {
	tagData, err := t.BureauRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagResponse = resource.BureauResponse{
		UserID:                tagData.UserID,
		BureauReferenceNumber: tagData.BureauReferenceNumber,
		UserReferenceNumber:   tagData.UserReferenceNumber,
		DiffNumM2M8:           tagData.DiffNumM2M8,
		CountPlLoan:           tagData.CountPlLoan,
	}
	return tagResponse
}

func (t BureauServiceImpl) FindAll() []model.Bureau {
	result := t.BureauRepository.FindAll()

	var tags []model.Bureau
	for _, value := range result {
		tag := model.Bureau{
			UserID:                value.UserID,
			BureauReferenceNumber: value.BureauReferenceNumber,
			CountPlLoan:           value.CountPlLoan,
			DiffNumM2M8:           value.DiffNumM2M8,
			NewBureauVariable:     value.NewBureauVariable,
		}
		tags = append(tags, tag)
	}
	return tags
}
