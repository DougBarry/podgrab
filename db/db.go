package db

import (
	"fmt"
	"log"
	"os"
	"path"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

// DB is
var DB *gorm.DB

// Init is used to Initialize Database
func Init() (*gorm.DB, error) {
	// github.com/mattn/go-sqlite3
	configPath := os.Getenv("CONFIG")
	dbPath := path.Join(configPath, "podgrab.db")

	dbDSN := "file:" + dbPath + "?_sync=0&_journal_mode=WAL&cache=shared&_cache_size=50000&temp_store=memory&mmap_size = 30000000000"
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	log.Println(dbPath)
	db, err := gorm.Open(sqlite.Open(dbDSN), gormConfig)
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}

	localDB, _ := db.DB()
	localDB.SetMaxIdleConns(20)
	//db.LogMode(true)
	DB = db
	return DB, nil
}

// Migrate Database
func Migrate() {
	DB.AutoMigrate(&Podcast{}, &PodcastItem{}, &Setting{}, &Migration{}, &JobLock{}, &Tag{})
	RunMigrations()
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
