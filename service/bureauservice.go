package service

import (
	"github.com/ankit-kumar02/simple-crud/helper"
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/repository"
	"github.com/ankit-kumar02/simple-crud/request"
)

type BureauServiceImpl struct {
	BureauRepository repository.BureauRepository
}

func NewBureauServiceImple(bureaurepository repository.BureauRepository) BureauService {
	return &BureauServiceImpl{
		BureauRepository: bureaurepository,
	}
}
func (b BureauServiceImpl) Create(bureau request.CreateBureauRequest) {
	tagModel := model.Bureau{
		BureauReferenceNumber: bureau.BureauReferenceNumber,
	}
	b.BureauRepository.Save(tagModel)
}

func (b BureauServiceImpl) Update(bureauId request.UpdateBureauRequest) {
	tagData, err := b.BureauRepository.FindById(bureauId)
	helper.ErrorPanic(err)
	tagData.BureauReferenceNumber = bureauId.BureauReferenceNumber
	b.BureauRepository.Update(tagData)
}

func (t BureauServiceImpl) Delete(tagId string) {
	t.BureauRepository.Delete(tagId)
}

func (t BureauServiceImpl) FindById(tagId int) {
	tagData, err := t.BureauRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t BureauServiceImpl) FindAll() []response.TagsResponse {
	result := t.BureauRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}
