package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"student/domain/group"
	"student/domain/student"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"

	"github.com/shahTeam/crmconnect/postgres"
)

const (
	studentsTableName = "students"
	groupsTableName   = "groups"
)

func NewPostgres(cfg postgres.Config) (*Postgres, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return &Postgres{}, err
	}
	return &Postgres{
		db: db.DB,
	}, nil
}

//Postgres implements service.Repository interface
type Postgres struct {
	db *sql.DB
}

func (p *Postgres) GetStudentsByGroup(ctx context.Context, groupID uuid.UUID) ([]student.Student, error) {
	repoStudents, err := p.getStudentsByGroup(ctx, groupID)
	if err != nil {
		return []student.Student{}, err
	}
	return toDomanStudents(repoStudents)
}

func (p *Postgres) getStudentsByGroup(ctx context.Context, groupID uuid.UUID) ([]Student, error) {
	query := `
	SELECT * FROM students WHERE group_id=$1
	`
	students := make([]Student, 0)
	rows, err := p.db.QueryContext(ctx, query, groupID)
	if err != nil {
		return []Student{}, err
	}
	for rows.Next() {
		var stu Student
		if err := rows.Scan(&stu.ID, &stu.FirstName, &stu.LastName, &stu.Email, &stu.PhoneNumber, &stu.Level, &stu.Password, &stu.GroupID); err != nil {
			return []Student{}, err
		}
		students = append(students, stu)
	}
	return students, nil
}

//ListGroups...
func (p *Postgres) ListGroup(ctx context.Context, page, limit int32) ([]group.Group, int, error) {
	repoGroups, count, err := p.listGroup(ctx, page, limit)
	if err != nil {
		return []group.Group{}, 0, err
	}

	groups, err := toDomanGroups(repoGroups)
	if err != nil {
		return []group.Group{}, 0, err
	}
	return groups, count, nil
}

func (p *Postgres) listGroup(ctx context.Context, page, limit int32) ([]Group, int, error) {
	count, err := p.count(ctx, groupsTableName)
	if err != nil {
		return []Group{}, 0, err
	}
	query := `SELECT * FROM groups offset $1 limit $2`
	offset := (page - 1) * limit
	groups := make([]Group, 0)
	rows, err := p.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return []Group{}, 0, err
	}
	for rows.Next() {
		var grp Group
		if err := rows.Scan(&grp.ID, &grp.Name, &grp.MainTeacherID); err != nil {
			return []Group{}, 0, err
		}
		groups = append(groups, grp)
	}
	return groups, count, nil
}

func (p *Postgres) GetGroup(ctx context.Context, by group.GroupBy) (group.Group, error){
	grr, err := p.getGroup(ctx, by)
	if err != nil {
		return group.Group{}, err
	}
	return group.UnmarshalGroup(group.UnmarshalGroupArgs(grr))
}

func (p *Postgres) getGroup(ctx context.Context, by group.GroupBy) (Group, error) {
	var (
		query string
		args interface{}
	)
	switch by := by.(type) {
	case group.ByID:
		query = `SELECT *FROM groups WHERE id = $1`
		args = by.ID
	case group.ByName:
		query = `SELECT *FROM groups WHERE name = $1`
		args = by.Name
	}
	 var gr Group
	 row := p.db.QueryRowContext(ctx, query, args)

     if err := row.Scan(&gr.ID, &gr.Name, gr.MainTeacherID); err != nil {
		return Group{}, err
	 }
	 return gr, nil
}

//Delete Group....
func (p *Postgres) DeleteGroup(ctx context.Context, by group.GroupBy) error {
	return p.DeleteGroup(ctx, by)
}

func (p *Postgres) deleteGroup(ctx context.Context, by group.GroupBy) error {
	var (
		query string
		args  interface{}
	)

	switch b := by.(type) {
	case group.ByID:
		query = `DELETE FROM groups WHERE id = $1`
		args = b.ID
	case group.ByName:
		query = `DELETE FROM groups WHERE name = $1`
		args = b.Name
	}
	_, err := p.db.ExecContext(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

//Update Group
func (p *Postgres) UpdateGroup(ctx context.Context, g group.Group) error {
	return p.updateGroup(ctx, g)
}

func (p *Postgres) updateGroup(ctx context.Context, g group.Group) error {
	query := `
	UPDATE groups SET name = $1, main_teacher = $2, id = $3
	`
	_, err := p.db.ExecContext(ctx, query, g.Name(), g.MainTeacherID(), g.ID())
	if err != nil {
		return err
	}
	return err
}

//List Students
func (p *Postgres) ListStudent(ctx context.Context, page, limit int32) ([]student.Student, int, error) {
	repoStudent, count, err := p.listStudent(ctx, page, limit)
	if err != nil {
		return []student.Student{}, 0, err
	}
	stu, err := toDomanStudents(repoStudent)
	if err != nil {
		return []student.Student{}, 0, err
	}
	return stu, count, nil
}

func (p *Postgres) listStudent(ctx context.Context, page, limit int32) ([]Student, int, error) {
	count, err := p.count(ctx, studentsTableName)
	if err != nil {
		return []Student{}, 0, err
	}
	query := `
	SELECT * FROM students offset $1 limit $2
	`
	offset := (page - 1) * limit
	students := make([]Student, 0)

	rows, err := p.db.QueryContext(ctx, query, offset, limit)

	for rows.Next() {
		var st Student
		if err := rows.Scan(&st.ID, &st.FirstName, &st.LastName, &st.Email, &st.PhoneNumber, &st.Level, &st.Password, &st.GroupID); err != nil {
			return []Student{}, 0, err
		}
		students = append(students, st)
	}
	return students, count, nil
}

//Delete Students
func (p *Postgres) DeleteStudent(ctx context.Context, by student.StudentBy) error {
	return p.deleteStudent(ctx, by)
}

func (p *Postgres) deleteStudent(ctx context.Context, by student.StudentBy) error {
	var (
		query string
		arg   interface{}
	)

	switch b := by.(type) {
	case student.ByEamil:
		query = `DELETE FROM students WHERE name = $1`
		arg = b.Email
	case student.ByID:
		query = `DELETE FROM students WHERE id = $1`
		arg = b.ID
	case student.ByPhoneNumber:
		query = `DELETE FROM students WHERE phone_number = $1`
		arg = b.ByPhoneNumber
	}

	if _, err := p.db.ExecContext(ctx, query, arg); err != nil {
		return err
	}
	return nil
}

//Update Student
func (p *Postgres) UpdateStudent(ctx context.Context, s student.Student) error {
	return p.updateStudent(ctx, toRepositoryStudents(s))
}

func (p *Postgres) updateStudent(ctx context.Context, s Student) error {
	bp, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.MinCost)
	query := `
	UPDATE students 
	SET first_name = $1. last_name = $2, email = $3, phone_number = $4, level = $5, password = $6, group_id = $7
	WHERE id = $8`
	if _, err = p.db.ExecContext(ctx, query, s.FirstName, s.LastName, s.Email, s.PhoneNumber, s.Level, string(bp), s.Password, s.GroupID, s.ID); err != nil {
		return err
	}
	return nil
}

//Get Student....
func (p *Postgres) GetStudent(ctx context.Context, by student.StudentBy) (student.Student, error) {
	s, err := p.getStudent(ctx, by)
	if err != nil {
		 if errors.Is(err, sql.ErrNoRows) {
             return student.Student{}, err
 		}
		return student.Student{}, err
	}
	return student.UnmarshalStudent(student.UnmarshalStudentArgs(s))
}

func (p *Postgres) getStudent(ctx context.Context, by student.StudentBy) (Student, error) {
	var (
		query string
		arg   interface{}
	)

	switch b := by.(type) {
	case student.ByID:
		query = `SELECT * FROM students WHERE id = $1`
		arg = b.ID
	case student.ByEamil:
		query = `SELECT * FROM students WHERE email = $1`
		arg = b.Email
	case student.ByPhoneNumber:
		query = `SELECT * FROM students WHERE phone_number = $1`
		arg = b.ByPhoneNumber
	}

	row := p.db.QueryRowContext(ctx, query, arg)
	var stud Student
	if err := row.Scan(&stud.ID, &stud.FirstName, &stud.LastName, &stud.Email, &stud.PhoneNumber, &stud.Level, &stud.Password, &stud.GroupID); err != nil {
		return Student{}, err
	}

	return stud, nil
}

//Create Student...
func (p *Postgres) CreateStudent(ctx context.Context, s student.Student) error {
	return p.createStudent(ctx, toRepositoryStudents(s))
}

func (p *Postgres) createStudent(ctx context.Context, s Student) error {
	bp, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO students 
	(id, first_name, last_name, email, phone_number, level, password, group_id)
	values($1, $2, $3, $4, $5, $6, $7, $8)
	`
	if _, err = p.db.ExecContext(ctx, query, s.ID, s.FirstName, s.LastName, s.Email, s.PhoneNumber, s.Level, string(bp), s.GroupID); err != nil {
		return err
	}
	return nil
}

//Create Group
func (p *Postgres) CreateGroup(ctx context.Context, g group.Group) error {
	return p.createGroup(ctx, toRepositoryGroup(g))
}

func (p *Postgres) createGroup(ctx context.Context, g Group) error {
	query := `
	INSERT INTO groups VALUES($1, $2, $3)
	`
	if _, err := p.db.ExecContext(ctx, query); err != nil {
		return err
	}
	return nil
}

func (p *Postgres) count(ctx context.Context, table string) (int, error) {
	query := fmt.Sprintf("select count(*) from %s", table)
	var count int
	row := p.db.QueryRowContext(ctx, query)
	row.Scan(&count)
	return count, nil
}

func (p *Postgres) cleanUp() func() {
	return func() {
		query := `DELETE FROM groups`
		_, err := p.db.Exec(query)
		if err != nil {
			log.Println(err)
		}
		query = `DELETE FROM students`
		if _, err := p.db.Exec(query); err != nil {
			log.Println(err)
		}
	}
}
