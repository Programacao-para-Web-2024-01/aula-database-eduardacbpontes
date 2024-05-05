package student

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(database *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: database,
	}
}

func (repo *StudentRepository) List() ([]Student, error) {
	data, err := repo.db.Query(`
		SELECT st.id,
		       st.name,
		       st.age,
		       st.email,
		       st.phone, 
		       IF(COUNT(su.id) > 0, JSON_ARRAYAGG(su.name), NULL) subjects
		FROM students st
         	LEFT JOIN students_subjects ss ON st.id = ss.student_id
         	LEFT JOIN subjects su ON ss.subject_id = su.id
		GROUP BY st.id`,
	)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var allStudents []Student

	for data.Next() {
		var s Student
		var subs []byte
		err = data.Scan(
			&s.Id,
			&s.Name,
			&s.Age,
			&s.Email,
			&s.Phone,
			&subs,
		)
		if err != nil {
			return nil, err
		}

		if subs != nil {
			err = json.Unmarshal(subs, &s.SubjectsName)
			if err != nil {
				return nil, err
			}
		}

		allStudents = append(allStudents, s)
	}

	return allStudents, nil
}

func (repo *StudentRepository) Get(id int) (*Student, error) {
	data := repo.db.QueryRow(`
		SELECT id, name, age, email, phone
		FROM students
		WHERE id = ?`, id)

	var s Student
	err := data.Scan(&s.Id, &s.Name, &s.Age, &s.Email, &s.Phone)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (repo *StudentRepository) Create(s Student) (int64, error) {
	res, err := repo.db.Exec(`INSERT INTO students(name, age, email, phone)
					  VALUES (?, ?, ?, ?)`,
		s.Name, s.Age, s.Email, s.Phone)

	if err != nil {
		return 0, err
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (repo *StudentRepository) Update(id int, s Student) error {
	_, err := repo.db.Exec(`UPDATE students
						SET name=?,
						    age=?,
						    email=?,
						    phone=?
						WHERE id=?`,
		s.Name, s.Age, s.Email, s.Phone, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo *StudentRepository) Delete(id int) error {
	_, err := repo.db.Exec(`DELETE
							FROM students
							WHERE id = ?`, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo *StudentRepository) AssociateSubjects(studentID int, subjectIDs []int) error {
	var b strings.Builder
	var p []any
	for i, subjectID := range subjectIDs {
		p = append(p, studentID, subjectID)
		b.WriteString("(?, ?)")
		if i < len(subjectIDs)-1 {
			b.WriteString(",")
		}
	}

	const q = `
			INSERT INTO students_subjects(student_id, subject_id)
			VALUES %s;`

	fmt.Println(fmt.Sprintf(q, b.String()))
	_, err := repo.db.Exec(fmt.Sprintf(q, b.String()), p...)
	if err != nil {
		return err
	}

	return nil
}
