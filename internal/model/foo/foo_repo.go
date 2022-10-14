package foo

import (
	log "github.com/sirupsen/logrus"
)

type FooRepository struct {
	fooMap map[int]*Foo
	nextId int
}

func NewFooRepository() *FooRepository {
	repo := new(FooRepository)
	repo.fooMap = make(map[int]*Foo)
	repo.nextId = 1
	return repo
}

func (fooRepo *FooRepository) Save(foo *Foo) {
	log.WithFields(log.Fields{"foo": foo}).Info("Save Foo")
	if foo.FooId >= fooRepo.nextId {
		fooRepo.nextId = foo.FooId + 1
	}
	if foo.FooId == 0 {
		foo.FooId = fooRepo.generateId()
	}
	fooRepo.fooMap[foo.FooId] = foo
}

func (fooRepo *FooRepository) generateId() int {
	id := fooRepo.nextId
	log.WithFields(log.Fields{"id": id}).Debug("Generate new fooId")
	fooRepo.nextId++
	return id
}

func (fooRepo *FooRepository) FindAll() []*Foo {
	values := make([]*Foo, 0, len(fooRepo.fooMap))
	for _, value := range fooRepo.fooMap {
		values = append(values, value)
	}
	return values
}

func (fooRepo *FooRepository) FindById(fooId int) (foo *Foo, ok bool) {
	foo, ok = fooRepo.fooMap[fooId]
	return foo, ok
}
