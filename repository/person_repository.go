package repository

import "gin-basics/model"

type PersonRepository interface {
	FindAll() []model.Person
	FindByID(id int) model.Person
	Save(person model.Person) model.Person
	Delete(id int)
}
