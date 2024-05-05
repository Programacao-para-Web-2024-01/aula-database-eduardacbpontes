package subject

type Service struct {
	repo *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) GetByStudentID(studentID int) ([]Subject, error) {
	return s.repo.GetByStudentID(studentID)
}
