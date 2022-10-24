package student

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	//"github.com/shahTeam/crmconnect/id"
)

func TestFactory_NewStudent(t *testing.T) {
	type args struct {
		firstName   string
		lastName    string
		email       string
		phoneNumber string
		level       int32
		password    string
		groupID     uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Student
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				firstName:   "shahzod",
				lastName:    "Ibrohimov",
				email:       "shahzodasturchi@gmail.com",
				phoneNumber: "+998913084678",
				level:       4,
				password:    "1234",
				groupID:     testGroupID,
			},
			want: Student{
				id:          testStudentID,
				firstName:   "shahzod",
				lastName:    "Ibrohimov",
				email:       "shahzodasturchi@gamil.com",
				phoneNumber: "+998913084678",
				level:       4,
				password:    "1234",
				groupID:     testGroupID,
			},
			wantErr: false,
		},
		// {
		// 	name: "with empty name",
		// 	args: args{
		// 		firstName:   "",
		// 		lastName:    "Ibrohimov",
		// 		email:       "shahzodasturchi@gmail.com",
		// 		phoneNumber: "+998913084678",
		// 		level:       4,
		// 		groupID:     testGroupID,
		// 	},
		// 	want:    Student{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "with empty surname",
		// 	args: args{
		// 		firstName:   "shahzod",
		// 		lastName:    "",
		// 		email:       "shahzodasturchi@gmail.com",
		// 		phoneNumber: "+998913084678",
		// 		level:       4,
		// 		groupID:     testGroupID,
		// 	},
		// 	want:    Student{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "with invalid email",
		// 	args: args{
		// 		firstName:   "shahzod",
		// 		lastName:    "Ibrohimov",
		// 		email:       "shahzodastu.com",
		// 		phoneNumber: "+998913084678",
		// 		level:       4,
		// 		groupID:     testGroupID,
		// 	},
		// 	want:    Student{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "with invalid phone number",
		// 	args: args{
		// 		firstName:   "shahzod",
		// 		lastName:    "Ibrohimov",
		// 		email:       "shahzodasturchi@gmail.com",
		// 		phoneNumber: "3084678",
		// 		level:       4,
		// 		groupID:     testGroupID,
		// 	},
		// 	want:    Student{},
		// 	wantErr: true,
		// },
		// {
		// 	name: "with invalid level",
		// 	args: args{
		// 		firstName:   "shahzod",
		// 		lastName:    "Ibrohimov",
		// 		email:       "shahzodasturchi@gmail.com",
		// 		phoneNumber: "+998913084678",
		// 		level:       7,
		// 		groupID:     testGroupID,
		// 	},
		// 	want:    Student{},
		// 	wantErr: true,
		// },
	}
	
    f := NewFactory(testIDGenarator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewStudent(tt.args.firstName, tt.args.lastName, tt.args.email, tt.args.phoneNumber, tt.args.level, tt.args.password, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factory.NewStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.NewStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testGroupID = uuid.New()
	testStudentID = uuid.New()
)

type testIDGenarator struct{}

func (g testIDGenarator) GenerateUUID() uuid.UUID {
	fmt.Println(testStudentID)
     return testStudentID
}
