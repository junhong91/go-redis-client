package album

import "github.com/junhong91/clean-architecture-go-redis/pkg/entity"

//Reader interface
type Reader interface {
	Find(id entity.ID) (*entity.Album, error)
}

//Writer album writer
type Writer interface {
	Store(a *entity.Album) (entity.ID, error)
	Delete(id entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
