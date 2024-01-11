package service

import "github.com/ankit-kumar02/simple-crud/request"

type BureauService interface {
	Create(bureau request.CreateBureauRequest)
	Update(bureau request.UpdateBureauRequest)
	Delete(bureauId string)
	FindById(bureauId string)
	FindAll()
}
