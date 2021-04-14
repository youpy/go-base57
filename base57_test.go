package base57_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/youpy/go-base57"
)

func Test_base57_Encode(t *testing.T) {

	type args struct {
		u uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid UUID string",
			args: args{u: uuid.MustParse("3b1f8b40-222c-4a6e-b77e-779d5a94e21c")},
			want: "CXc85b4rqinB7s5J52TRYb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := base57.New()
			if got := b.Encode(tt.args.u); got != tt.want {
				t.Errorf("base57.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base57_Decode(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name:    "valid UUID string",
			args:    args{u: "CXc85b4rqinB7s5J52TRYb"},
			want:    uuid.MustParse("3b1f8b40-222c-4a6e-b77e-779d5a94e21c"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := base57.New()
			got, err := b.Decode(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("base57.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base57.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
