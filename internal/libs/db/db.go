package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"angmorning.com/internal/config"
)

func InitDb() *sql.DB {
	dbPort, err := strconv.Atoi(config.DbPort)
	if err != nil {
		return nil
	}

	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DbHost, dbPort, config.DbUser, config.DbPassword, config.DbName))

	if err != nil {
		fmt.Printf("Failed to connect to the database: %v", err)
		return nil
	}

	if err := conn.Ping(); err != nil {
		fmt.Printf("Failed to connect to the database: %v", err)
		return nil
	}
	fmt.Println("Successfully connected!")

	return conn
}
