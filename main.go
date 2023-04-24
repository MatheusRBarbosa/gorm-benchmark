package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var envs map[string]string
var config = gorm.Config{
	PrepareStmt:            true,
	SkipDefaultTransaction: true,
}

func main() {
	loadEnvs()
	arg := GetEnv("TARGET")
	if arg == "postgres" {
		postgresBench()
		return
	} else if arg == "mysql" {
		mysqlBench()
		return
	}

	return
}

func postgresBench() {
	dsn := GetEnv("POSTGRES_DSN")

	fmt.Printf("DSN: %s\n", dsn)

	start := time.Now()
	db, err := gorm.Open(postgres.Open(dsn), &config)

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Printf("==> Open connection: %s\n", time.Since(start))
	executeBenchmark(db)
}

func mysqlBench() {
	dsn := GetEnv("MYSQL_DSN")

	start := time.Now()
	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Printf("==> Open connection: %s\n", time.Since(start))
	executeBenchmark(db)
}

func executeBenchmark(db *gorm.DB) {
	user := &User{
		Name:     "Inserted",
		Email:    "m@m.com",
		Password: "asdf",
	}

	// Test Create performance
	start := time.Now()
	db.Create(&user)
	fmt.Printf("==> Create user: %s\n", time.Since(start))

	// Test Select by email performance
	start = time.Now()
	db.Where(&User{Email: user.Email}).First(&user)
	fmt.Printf("==> Select by EMAIL user: %s\n", time.Since(start))

	// Test Update performance
	start = time.Now()
	db.Model(&user).Where(&User{ID: user.ID}).Update("name", "updated")
	fmt.Printf("==> Update user: %s\n", time.Since(start))

	// Test delete performance
	start = time.Now()
	db.Unscoped().Delete(&User{}, user.ID)
	fmt.Printf("==> Delete User user: %s\n", time.Since(start))
}

func GetEnv(name string) string {
	return envs[name]
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs = map[string]string{
		"APP_ENV":      os.Getenv("APP_ENV"),
		"POSTGRES_DSN": os.Getenv("POSTGRES_DSN"),
		"MYSQL_DSN":    os.Getenv("MYSQL_DSN"),
		"TARGET":       os.Getenv("TARGET"),
	}
}
