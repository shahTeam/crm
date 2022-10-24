package repository_test

import (
	"context"
	"log"
	"teacher/domain/subject"
	"teacher/domain/teacher"
	"teacher/pkg/idgen"
	"teacher/repository"
	"testing"

	"github.com/shahTeam/crmconnect/postgres"
	"github.com/stretchr/testify/require"
)

var testPostgresCfg = postgres.Config{
	PostgresHost: "localhost",
	PostgresPort: "5432",
	PostgresUser: "shahzod",
	PostgresPassword: "7",
	PostgresDB: "crm_test",
	PostgresMigrationPath: "file:///home/shahzod/go/src/github.com/CRM/teacher_service/repository/migrations",
}

func TestPostgres_CreatSubject(t *testing.T) {
	p, err := repository.NewPostgres(testPostgresCfg)
	require.NoError(t, err)
 
	t.Cleanup(cleanUp(p))  
	
	subjectFactory := subject.NewFactory(idgen.Generator{})
	teacherFactory := teacher.NewFactory(idgen.Generator{})
t.Run("create subject", func(t *testing.T){
	t.Cleanup(cleanUp(p))
	sub, err := subjectFactory.NewSubject(
		"Math",
		"description",
	)
	require.NoError(t, err)

	err = p.CreateSubject(context.Background(), sub)
	require.NoError(t, err)
})

t.Run("create teacher", func(t *testing.T){
	t.Cleanup(cleanUp(p))
	sub, err := subjectFactory.NewSubject(
		"Math",
		"description",
	)
	require.NoError(t, err)

	err = p.CreateSubject(context.Background(), sub)
	require.NoError(t, err)

	tch, err := teacherFactory.NewFactory(
		"Shahzod",
		"Ibrohimov",
		"shahzod@Ibrohimov.com",
		"+998913084678",
		"1234",
        sub.ID(),
	)
	require.NoError(t, err)

	err = p.CreateTeacher(context.Background(), tch)
	require.NoError(t, err)
})
	
}

func cleanUp(p *repository.Postgres) func() {
	return func() {
		if err :=  p.CleanUp(context.Background()); err != nil {
			log.Panicln("failed to cleanup db, should be done manually", err)
		}
	}
}











































































































































































