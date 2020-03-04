package data

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ORM is the object holding the active database connection interface
var ORM *gorm.DB

func init() {
	log.Println("greefdata initializing")
	var err error
	ORM, err = gorm.Open("sqlite3", "greefdb.sqlite3")

	if err != nil {
		log.Println(err)
	}
	if os.Getenv("GREEF_DB_LOGMODE") == "" {
		os.Setenv(`GREEF_DB_LOGMODE`, "ON")
	}
	ORM.LogMode(os.Getenv("GREEF_DB_LOGMODE") == "ON")
}

// MigrateDataSchema auto-migrates structs
func MigrateDataSchema() {
	ORM.AutoMigrate(EnvironmentalParameters{})
	ORM.AutoMigrate(Livestock{})
	ORM.AutoMigrate(SystemDetails{})
	// ORM.AutoMigrate(DeviceState{})
}

func Service() *gorm.DB {
	return ORM
}
