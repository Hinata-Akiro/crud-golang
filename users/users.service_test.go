package users

import (
	"reflect"
	"testing"
	"time"
)

func Test_createNewUser(t *testing.T) {
	type args struct {
		userInput *User
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "should create a new user",
			args: args{
				userInput: &User{
					FirstName: "John",
					LastName:  "Doe",
					Email:     " john@example.com",
					PassWord:  "yhadyyudw36782",
				},
			},
			want: &User{
				FirstName: "John",
				LastName:  "Doe",
				Email:     " john@example.com",
				PassWord:  "yhadyyudw36782",
			},
			wantErr: false,
		},
		{
			name: "should not create a new user",
			args: args{
				userInput: &User{
					FirstName: "John",
					LastName:  "Doe",
					Email:     " john@example.com",
					PassWord:  "yhadyyudw36782",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createNewUser(tt.args.userInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("createNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createNewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkUserExists(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUserExists(tt.args.email); got != tt.want {
				t.Errorf("checkUserExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "should find a user",
            args: args{
                email: "jane@example.com",
				},
            want: &User{
				FirstName: "Jane",
                LastName:  "Doe",
                Email:     "jane@example.com",
                PassWord:  " 64738273hqh",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				},
            wantErr: false,
		},
		{
			name: "should not find a user",
            args: args{
                email: "jane@example.com",
                },
            want: nil,
            wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
