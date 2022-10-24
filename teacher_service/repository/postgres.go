package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"teacher/domain/subject"
	"teacher/domain/teacher"
	"teacher/pkg/errs"

	"github.com/jmoiron/sqlx"

	"github.com/shahTeam/crmconnect/postgres"
	"golang.org/x/crypto/bcrypt"
)

func NewPostgres(cfg postgres.Config) (*Postgres, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return nil, err
	}
	return &Postgres{
		db: db,
	}, nil
}

type Postgres struct {
	db *sqlx.DB
}

const (
	teachersTableName = "teachers"
	subjectsTableName = "subjects"
)

func (p Postgres) CreateTeacher(ctx context.Context, t teacher.Teacher) error {
	return p.createTeacher(ctx, toRepositoryTeacher(t))
}

func (p Postgres) createTeacher(ctx context.Context, t Teacher) error {
	query := `INSERT INTO teachers VALUES($1, $2, $3, $4, $5, $6, $7)`

	_, err := p.db.ExecContext(ctx, query,
		t.ID, t.FirstName, t.LastName, t.Email, t.PhoneNumber, t.Password, t.SubjectID)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) GetTeacher(ctx context.Context, by teacher.By) (teacher.Teacher, error) {
	t, err := p.getTeacher(ctx, by)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return teacher.Teacher{}, errs.NotFound
		}
		return teacher.Teacher{}, err
	}
	return teacher.UnmarshalTeacher(teacher.UnmarshalTeacherArgs(t))
}

func (p Postgres) getTeacher(ctx context.Context, by teacher.By) (Teacher, error) {
	var (
		query string
		arg   interface{}
	)
	switch b := by.(type) {
	case teacher.ByID:
		query = `SELECT * FROM teachers WHERE id = $1`
		arg = b.ID
	case teacher.ByEmail:
		query = `SELECT * FROM teachers WHERE email = $1`
		arg = b.Email
	case teacher.ByPhoneNumber:
		query = `SELECT * FROM teachers WHERE phone_number = $1`
		arg = b.PhoneNumber
	}
	var t Teacher
	if err := p.db.GetContext(ctx, &t, query, arg); err != nil {
		return Teacher{}, err
	}
	return t, nil
}

func (p Postgres) GetSubject(ctx context.Context, sub subject.By) (subject.Subject, error) {
	s, err := p.getSubject(ctx, sub)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return subject.Subject{}, errs.NotFound
		}
		return subject.Subject{}, err
	}
	return subject.UnmarshalSubject(subject.UnmarshalSubjectArgs(s))
}

func (p Postgres) UpdateTeacher(ctx context.Context, t teacher.Teacher) error {
	return p.updateTeacher(ctx, toRepositoryTeacher(t))
}

func (p Postgres) updateTeacher(ctx context.Context, t Teacher) error {
	bp, err := bcrypt.GenerateFromPassword([]byte(t.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	query := `
	   UPDATE teachers
	        SET first_name = $1, last_name = $2, email = $3, phone_number = $4, password = $5, subject_id = $6
	   WHERE id = $7
	`
	_, err = p.db.ExecContext(
		ctx, query,
		t.FirstName, t.LastName, t.Email, t.PhoneNumber, string(bp), t.SubjectID, t.ID,
	)
	return err
}

func (p Postgres) DeleteTeacher(ctx context.Context, by teacher.By) error {
	return p.deleteTeacher(ctx, by)
}

func (p Postgres) deleteTeacher(ctx context.Context, by teacher.By) error {
	var (
		query string
		arg   interface{}
	)
	switch b := by.(type) {
	case teacher.ByID:
		query = `DELTE * FROM teachers WHERE id = $1`
		arg = b.ID
	case teacher.ByEmail:
		query = `DELETE * FROM teachers WHERE email = $1`
		arg = b.Email
	case teacher.ByPhoneNumber:
		query = `DELETE * FROM teachers WHERE phone_number = $1`
		arg = b.PhoneNumber
	}

	_, err := p.db.ExecContext(ctx, query, arg)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) getSubject(ctx context.Context, by subject.By) (Subject, error) {
	var (
		query string
		arg   interface{}
	)

	switch b := by.(type) {
	case subject.ByID:
		query = `SELECT * FROM subjects WHERE id = $1`
		arg = b.ID
	case subject.ByName:
		query = `SELECT * FROM subjects WHERE name = $1`
		arg = b.Name
	}

	var s Subject
	if err := p.db.GetContext(ctx, &s, query, arg); err != nil {
		return Subject{}, err
	}
	return s, nil
}

func (p Postgres) UpdateSubject(ctx context.Context, sub subject.Subject) error {
	return p.updateSubject(ctx, toRepositorySubject(sub))
}

func (p Postgres) updateSubject(ctx context.Context, sub Subject) error {
	query := `
	    UPDATE subjects
		    SET name = $1, description = $2
	    WHERE id = $3
	`
	_, err := p.db.ExecContext(ctx, query, sub.Name, sub.Description, sub.ID)

	return err
}

func (p Postgres) DeleteSubject(ctx context.Context, sub subject.By) error {
     return p.deleteSubject(ctx, sub)
}

func (p Postgres) deleteSubject(ctx context.Context, sub subject.By) error {
	var (
		query string
		arg interface{}
	)

	switch s := sub.(type) {
		case subject.ByID:
			query = `DELETE FROM subjects WHERE id = $1`
			arg = s.ID
		case subject.ByName:
            query = `DELETE FROM subjects WHERE id = $2`
			arg = s.Name
	}

	_, err := p.db.ExecContext(ctx, query, arg)

	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) CreateSubject(ctx context.Context, sub subject.Subject) error {
	return p.createSubject(ctx, toRepositorySubject(sub))
}

func (p Postgres) createSubject(ctx context.Context, s Subject) error {
	query := `INSERT INTO subjects VALUES($1, $2, $3)`

	_, err := p.db.ExecContext(ctx, query, s.ID, s.Name, s.Description)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) ListTeachers(ctx context.Context, page, limit int32) ([]teacher.Teacher, int, error) {
     repoTeachers, count, err := p.listTeachers(ctx, page, limit)
	 if err != nil {
		return nil, 0, err
	 }

	 teachers, err := toDomainTeachers(repoTeachers)
	 if err != nil {
		return nil, 0, err
	 }
	 return teachers, count, nil
}

func (p Postgres) listTeachers(ctx context.Context, page, limit int32) ([]Teacher, int, error) {
	count, err := p.count(ctx, subjectsTableName)
	if err != nil {
		return nil, 0, err
	}

	query := `
	    SELECT * FROM teachers 
		OFFSET $1 LIMIT $2
	`
	offset := (page-1)*limit
	teachers := make([]Teacher, 0)
	if err := p.db.SelectContext(ctx, &teachers, query, offset, limit); err != nil {
		return nil, 0, err
	}

	return teachers, count, nil
}

func (p Postgres) ListSubjects(ctx context.Context, page, limit int32) ([]subject.Subject, int, error) {
	repoSubjects, count, err := p.listSubjects(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	subjects, err := toDomainSubjects(repoSubjects)
	if err != nil {
		return nil, 0, err
	}
	return subjects, count, nil
}

func (p Postgres) listSubjects(ctx context.Context, page, limit int32) ([]Subject, int, error) {
	count, err := p.count(ctx, subjectsTableName)
	if err != nil {
		return nil, 0, err
	}

	query := `
	    SELECT * FROM subjects 
		OFFSET $1 LIMIT $2
	`
	offset := (page-1)*limit
	subjects := make([]Subject, 0)
	if err := p.db.SelectContext(ctx, &subjects, query, offset, limit); err != nil {
		return nil, 0, err
	}
	return subjects, count, nil
}

func (p Postgres) count(ctx context.Context, table string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)

	var count int

	if err := p.db.GetContext(ctx, &count, query); err != nil {
		return count, err
	}
	return count, nil
}

func (p Postgres) CleanUp(ctx context.Context) error {
	query := `DELETE FROM teachers`
	_, err := p.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `DELETE FROM subjects`
	_, err = p.db.ExecContext(ctx, query)
	return err
}
