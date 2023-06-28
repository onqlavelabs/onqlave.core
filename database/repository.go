package database

type DbRepository[T any] interface {
	FindOneRepository[T]
	FindAllRepository[T]
	CreateRepository[T]
	UpdateRepository[T]
	DeleteRepository[T]
}

type CreateRepository[T any] interface {
	Insert(model *T) (*T, error)
}

type FindOneRepository[T any] interface {
	FindOne(query interface{}, args ...interface{}) (*T, error)
}

type FindAllRepository[T any] interface {
	FindAll(query interface{}, args ...interface{}) ([]T, error)
}

type UpdateRepository[T any] interface {
	Update(model *T) (*T, error)
}

type DeleteRepository[T any] interface {
	Delete(model *T) (*T, error)
}
