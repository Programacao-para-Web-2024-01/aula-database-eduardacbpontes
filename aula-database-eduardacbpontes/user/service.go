package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(user User) error {
	_, err := s.repo.Register(user)
	return err
}

func (s *Service) Login(username, password string) error {
	return s.repo.Login(username, password)
}
