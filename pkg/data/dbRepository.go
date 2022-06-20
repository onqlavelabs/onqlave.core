package data

type DbRepository[T any] interface {
	FindOne(query interface{}, args ...interface{}) (*T, error)
	FindAll(query interface{}, args ...interface{}) ([]T, error)
	Insert(model *T) (*T, error)
	Update(model *T) (*T, error)
	Delete(model *T) (*T, error)
}
