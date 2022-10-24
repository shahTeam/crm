package teacher

import (
	"reflect"
	//"teacher/pkg/idgen"
	"testing"
	"fmt"

	"github.com/google/uuid"
)

func TestFactory_NewFactory(t *testing.T) {
	type args struct {
		firstName   string
		lastName    string
		email       string
		phoneNumber string
		password    string
		subjectID   uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Teacher
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "shuld pass",
			args: args{
				firstName:   "shahzod",
				lastName:    "Ibrohimov",
				email:       "shahzod@Ibrohimov.com",
				phoneNumber: "+998913084678",
				password:    "12344",
				subjectID:   testSubjectID,
			},
			want: Teacher{
				id:          testTeacherID,
				firstName:   "shahzod",
				lastName:    "Ibrohimov",
				email:       "shahzod@Ibrohimov.com",
				phoneNumber: "+998913084678",
				password:    "12344",
				subjectID:   testSubjectID,
			},
			wantErr: false,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewFactory(tt.args.firstName, tt.args.lastName, tt.args.email, tt.args.phoneNumber,tt.args.password, tt.args.subjectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factory.NewFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.NewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testSubjectID = uuid.New()
	testTeacherID = uuid.New()
	
)


type testIDGenerator struct{}

func (g testIDGenerator) GeneratorUUID() uuid.UUID {
	fmt.Println(testTeacherID)
	return testTeacherID
	
}
