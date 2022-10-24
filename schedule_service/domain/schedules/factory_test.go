package schedules

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestFactory_NewSchedule(t *testing.T) {
	type args struct {
		groupID      uuid.UUID
		subjectID    uuid.UUID
		teacherID    uuid.UUID
		weekday      time.Weekday
		lessonNumber int32
	}
	tests := []struct {
		name    string
		args    args
		want    Schedule
		wantErr bool
	}{
		{
		  name: "should pass",
		  args: args{
			groupID: testGroupID,
			subjectID: testSubjectID,
			teacherID: testTeacherID,
			weekday: time.Monday,
			lessonNumber: 1,
		  },	
		  want: Schedule{
			id: testScheduleID,
			groupID: testGroupID,
			subjectID: testSubjectID,
			teacherID: testTeacherID,
			weekday: time.Monday,
			lessonNumber: 1,
		  },
		  wantErr: false,
		},
		{
			name: "Invalid weekday",
			args: args{
				groupID: testGroupID,
				subjectID: testSubjectID,
				teacherID: testTeacherID,
				weekday: time.Saturday,
				lessonNumber: 1,
			},
			want: Schedule{},
			wantErr: true,
		},
		{
			name: "invalid lesson number",
			args: args{
				groupID: testGroupID,
				subjectID: testSubjectID,
				teacherID: testTeacherID,
				weekday: time.Monday,
				lessonNumber: 7,
			},
			want: Schedule{},
			wantErr: true,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewSchedule(tt.args.groupID, tt.args.subjectID, tt.args.teacherID, tt.args.weekday, tt.args.lessonNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factory.NewSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.NewSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testTeacherID = uuid.New()
	testGroupID = uuid.New()
	testSubjectID = uuid.New()
	testScheduleID = uuid.New()
)

type testIDGenerator struct{}

func (t testIDGenerator) GenerateUUID() uuid.UUID {
	return testScheduleID
}


