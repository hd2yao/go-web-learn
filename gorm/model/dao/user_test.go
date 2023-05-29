package dao

import (
	"testing"

	"go-web/gorm/model/dao/table"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		user *table.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{user: &table.User{UserName: "Kevin", Password: "123456"}},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := CreateUser(test.args.user); (err != nil) != test.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsers, err := GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, user := range gotUsers {
				t.Logf("user: %v", user)
			}

		})
	}
}

func TestUpdateUserNameById(t *testing.T) {
	type args struct {
		userName string
		userId   int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				userName: "Klein",
				userId:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserNameById(tt.args.userName, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserNameById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
