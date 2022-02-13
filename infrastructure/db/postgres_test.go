package db

import (
	"github.com/umerm-work/xm-exercise/config"
	"testing"
)

func TestConnect(t *testing.T) {

	con := config.Env{
		DBName:     "postgres",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "123",
	}
	if got := Connect(con); got == nil {
		t.Error("DB connection test failed")
	}
}
