package service

import (
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/request"
	"github.com/ankit-kumar02/simple-crud/resource"
)

type BureauService interface {
	Create(bureau request.CreateBureauRequest)
	Update(bureau request.UpdateBureauRequest)
	Delete(bureauId string)
	FindById(bureauId string) resource.BureauResponse
	FindAll() []model.Bureau
}
