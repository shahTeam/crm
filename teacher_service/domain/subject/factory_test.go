package subject

import (
	"reflect"
	"teacher/pkg/idgen"
	"testing"

	"github.com/google/uuid"
)	

func TestFactory_NewSubject(t *testing.T) {
	type fields struct {
		idGenerator idgen.IGenerator
	}
	type args struct {
		name        string
		description string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Subject
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				name: "Math",
				description: "study math",
			},
			want: Subject{
				id: testSubjectID,
				name: "Math",
				description: "study math",
			},
			wantErr: false,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewSubject(tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factory.NewSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.NewSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}


var testSubjectID = uuid.New()

type testIDGenerator struct {}

func (g testIDGenerator) GeneratorUUID() uuid.UUID {
	return testSubjectID
}