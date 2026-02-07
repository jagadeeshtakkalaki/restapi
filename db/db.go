package db

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection using Postgres.
// If dsn is empty, it will use defaults: host=localhost user=admin password=admin dbname=gin_db port=5432 sslmode=disable
func Init(dsn string) error {
    if dsn == "" {
        host := os.Getenv("DB_HOST")
        if host == "" {
            host = "localhost"
        }
        port := os.Getenv("DB_PORT")
        if port == "" {
            port = "5432"
        }
        user := os.Getenv("DB_USER")
        if user == "" {
            user = "admin"
        }
        password := os.Getenv("DB_PASSWORD")
        if password == "" {
            password = "admin"
        }
        dbname := os.Getenv("DB_NAME")
        if dbname == "" {
            dbname = "mydb"
        }
        dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
    }

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    sqlDB, err := DB.DB()
    if err != nil {
        return err
    }
    sqlDB.SetMaxOpenConns(10)
    log.Printf("database initialized (postgres)")
    return nil
}
