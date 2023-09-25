package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/joho/godotenv"
    "log"
    "os"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
    // Load konfigurasi dari file .env
    if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }

    db, err := gorm.Open("mysql", getDatabaseURL())

    if err != nil {
        return nil, err
    }

    DB = db
    return db, nil
}

func getDatabaseURL() string {
    return os.Getenv("DB_USERNAME") + ":" +
        os.Getenv("DB_PASSWORD") + "@tcp(" +
        os.Getenv("DB_HOST") + ":" +
        os.Getenv("DB_PORT") + ")/" +
        os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
}

