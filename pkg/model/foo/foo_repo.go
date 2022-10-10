package foo

type FooRepository struct {
	fooMap map[int]*Foo
}

func NewFooRepository() *FooRepository {
	repo := new(FooRepository)
	repo.fooMap = make(map[int]*Foo)
	return repo
}

func (fooRepo *FooRepository) Save(foo *Foo) {
	fooRepo.fooMap[foo.FooId] = foo
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
