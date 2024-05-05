package subject

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByStudentID(studentID int) ([]Subject, error) {
	rows, err := r.db.Query(`
			SELECT su.id,
			       su.name,
			       su.workload
			FROM subjects su
			         JOIN students_subjects ss ON su.id = ss.subject_id
			WHERE ss.student_id = ?`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []Subject
	for rows.Next() {
		var subject Subject

		err := rows.Scan(&subject.Id, &subject.Name, &subject.WordLoad)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}
