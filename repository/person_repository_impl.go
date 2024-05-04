package repository

import (
	"gin-basics/model"
	"gorm.io/gorm"
)

type PersonRepositoryImpl struct {
	Db *gorm.DB
}

func (p PersonRepositoryImpl) FindAll() []model.Person {
	var persons []model.Person
	if results := p.Db.Find(&persons); results.Error != nil {
		panic(results.Error)
	}
	return persons
}

func (p PersonRepositoryImpl) FindByID(id int) model.Person {
	var person model.Person
	if result := p.Db.First(&person, id); result.Error != nil {
		panic(result.Error)
	}
	return person
}

func (p PersonRepositoryImpl) Save(person model.Person) model.Person {
	if result := p.Db.Create(&person); result.Error != nil {
		panic(result.Error)
	}
	return person
}

func (p PersonRepositoryImpl) Delete(id int) {
	var person model.Person
	if result := p.Db.Where("id = ?", id).Delete(person); result.Error != nil {
		panic(result.Error)
	}
}

func NewPersonRepositoryImpl(db *gorm.DB) PersonRepository {
	return &PersonRepositoryImpl{Db: db}
}
