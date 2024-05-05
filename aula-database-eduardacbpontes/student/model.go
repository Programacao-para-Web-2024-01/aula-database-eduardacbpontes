package student

import "aula-database/subject"

type Student struct {
	Id           int64             `json:"id"`
	Name         string            `json:"name"`
	Age          int               `json:"age"`
	Email        string            `json:"email"`
	Phone        string            `json:"phone"`
	Subjects     []subject.Subject `json:"subjects,omitempty"`
	SubjectsName []string          `json:"subjects_name,omitempty"`
}
