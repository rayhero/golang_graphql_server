// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package gorm

import (
	"fmt"

	"github.com/rayhero/gqlgen-todos/src/base/gorm/migration"

	//Imports the database dialect of choice
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "database/sql"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"

	"github.com/jinzhu/gorm"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	// Connection String details:
	// * username   - the user created inside the DB.
	// * passed     - the user password
	// * sqlIp 			- your SQL ip
	// * sqlPort    - your SQL port
	// * tableName  - the name of the database table
	dbConnectionString := "username:passwd@tcp(sqlIp:sqlPort)/tableName?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dbConnectionString)
	if err != nil {
		// logger
		fmt.Println(fmt.Errorf("[ORM] err: %v", err))
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	// db.LogMode(logMode)
	// Automigrate tables
	autoMigrate := true
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	fmt.Printf("[ORM] Database connection initialized.")
	return orm, err
}
