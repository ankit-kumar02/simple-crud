package repository

import (
	"github.com/ankit-kumar02/simple-crud/model"
)

type BureauRepository interface {
	Save(bureau model.Bureau)
	Update(bureau model.Bureau)
	Delete(bureauId int)
	FindById(bureauId int) (bureau model.Bureau, err error)
	FindAll() []model.Bureau
}
