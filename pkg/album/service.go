package album

import "github.com/junhong91/clean-architecture-go-redis/pkg/entity"

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Find a bookmark
func (s *Service) Find(id entity.ID) (*entity.Album, error) {
	return s.repo.Find(id)
}

//Store a bookmark
func (s *Service) Store(a *entity.Album) (entity.ID, error) {
	return s.repo.Store(a)
}

//Delete a bookmark
func (s *Service) Delete(id entity.ID) error {
	_, err := s.Find(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
