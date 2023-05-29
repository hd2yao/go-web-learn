package database

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDBPing(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			DBPing()
		})
	}
}

func TestMysqlDemoCode(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MysqlDemoCode()
		})
	}
}
