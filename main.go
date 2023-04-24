package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config = gorm.Config{
	PrepareStmt:            true,
	SkipDefaultTransaction: true,
}

func main() {
	arg := strings.ToLower(os.Args[1])
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
	dsn := "postgresql://postgres:senha123@localhost:5432/postgres?sslmode=disable"

	start := time.Now()
	db, err := gorm.Open(postgres.Open(dsn), &config)

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Printf("==> Open connection: %s\n", time.Since(start))
	executeBenchmark(db)
}

func mysqlBench() {
	dsn := "root:senha123@tcp(localhost:3306)/teste?parseTime=true"

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
