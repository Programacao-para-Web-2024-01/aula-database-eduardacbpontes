package student

import "aula-database/subject"

type StudentService struct {
	repository *StudentRepository
	subjectSvc *subject.Service
}

func NewStudentService(repository *StudentRepository, subjectSvc *subject.Service) *StudentService {
	return &StudentService{repository: repository, subjectSvc: subjectSvc}
}

func (svc *StudentService) List() ([]Student, error) {
	return svc.repository.List()
}

func (svc *StudentService) Get(id int) (*Student, error) {
	s, err := svc.repository.Get(id)
	if err != nil {
		return nil, err
	}

	s.Subjects, err = svc.subjectSvc.GetByStudentID(id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (svc *StudentService) Create(s Student) (*Student, error) {
	newId, err := svc.repository.Create(s)
	if err != nil {
		return nil, err
	}

	s.Id = newId

	return &s, nil
}

func (svc *StudentService) Update(s Student) error {
	id := int(s.Id)

	_, err := svc.Get(id)
	if err != nil {
		return err
	}

	return svc.repository.Update(id, s)
}

func (svc *StudentService) Delete(id int) error {
	_, err := svc.Get(id)
	if err != nil {
		return err
	}

	return svc.repository.Delete(id)
}

func (svc *StudentService) AssociateSubjects(studentID int, subjectIDs []int) error {
	return svc.repository.AssociateSubjects(studentID, subjectIDs)
}
