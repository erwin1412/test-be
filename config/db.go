package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "1234"
	}
	if dbname == "" {
		dbname = "cars"
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS cars (
		id SERIAL PRIMARY KEY,
		car_name CHAR(50) NOT NULL,
		day_rate DOUBLE PRECISION NOT NULL,
		month_rate DOUBLE PRECISION NOT NULL,
		image CHAR(256) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		car_id INT NOT NULL REFERENCES cars(id),
		order_date DATE NOT NULL,
		pickup_date DATE NOT NULL,
		dropoff_date DATE NOT NULL,
		pickup_location CHAR(50) NOT NULL,
		dropoff_location CHAR(50) NOT NULL
	);
	`)
	fmt.Println("Database connected!")
}
