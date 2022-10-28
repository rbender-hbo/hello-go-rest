package foo

import (
	log "github.com/sirupsen/logrus"
)

type FooRepository interface {
	Save(*Foo) error
	FindAll() ([]*Foo, error)
	FindById(int) (*Foo, error)
}

type InMemoryFooRepository struct {
	fooMap    map[int]*Foo
	nextId    int
	LatestFoo *Foo
}

func NewInMemoryFooRepository() *InMemoryFooRepository {
	repo := new(InMemoryFooRepository)
	repo.fooMap = make(map[int]*Foo)
	repo.nextId = 1
	return repo
}

func (fooRepo *InMemoryFooRepository) Save(foo *Foo) error {
	log.WithFields(log.Fields{"foo": foo}).Info("Save Foo")
	if foo.FooId >= fooRepo.nextId {
		fooRepo.nextId = foo.FooId + 1
	}
	if foo.FooId == 0 {
		foo.FooId = fooRepo.generateId()
	}
	fooRepo.fooMap[foo.FooId] = foo
	fooRepo.LatestFoo = foo

	return nil
}

func (fooRepo *InMemoryFooRepository) generateId() int {
	id := fooRepo.nextId
	log.WithFields(log.Fields{"id": id}).Debug("Generate new fooId")
	fooRepo.nextId++
	return id
}

func (fooRepo *InMemoryFooRepository) FindAll() ([]*Foo, error) {
	values := make([]*Foo, 0, len(fooRepo.fooMap))
	for _, value := range fooRepo.fooMap {
		values = append(values, value)
	}
	return values, nil
}

func (fooRepo *InMemoryFooRepository) FindById(fooId int) (*Foo, error) {
	foo := fooRepo.fooMap[fooId]
	return foo, nil
}
