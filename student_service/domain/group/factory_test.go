package group

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	//"github.com/shahTeam/crmconnect/id"
)

func TestFactory_NewGroup(t *testing.T) {
	type args struct {
		name          string
		mainTeacherID uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Group
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				name: "shahzod",
				mainTeacherID: testMainTeacherID,
			},
			want: Group{
				id: testGroupID,
				name: "shahzod",
				mainTeacherID: testMainTeacherID,
			},
			wantErr: false,
		},
	}

	f := NewFactory(testIDGenarator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewGroup(tt.args.name, tt.args.mainTeacherID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factory.NewGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testGroupID = uuid.New()
	testMainTeacherID = uuid.New()
)

type testIDGenarator struct{}

func (t testIDGenarator) GenerateUUID() uuid.UUID {
	return testGroupID
}
