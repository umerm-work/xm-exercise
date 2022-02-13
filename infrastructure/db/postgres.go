package db

import (
	"bytes"
	"fmt"
	"github.com/umerm-work/xm-exercise/config"
	"github.com/umerm-work/xm-exercise/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os/exec"
)

// CONN global gorm db connection.
var CONN *gorm.DB

// Connect set the gorm db connection.
func Connect(env config.Env) *gorm.DB {
	if CONN != nil {
		return CONN
	}

	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)
	if CONN, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("failed to open the database connection. %w", err))
	}
	createPgDb(env)
	CONN.AutoMigrate(&models.Company{})

	// Add table suffix when creating tables
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	return CONN
}
func createPgDb(env config.Env) {
	cmd := exec.Command("createdb", "-p", env.DBPort, "-h", env.DBHost, "-U", env.DBUser, "-w", env.DBPassword, "-e", env.DBName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Output: %q\n", out.String())
}
