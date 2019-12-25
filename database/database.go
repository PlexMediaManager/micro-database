package database

import (
    format "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/plexmediamanager/micro-database/errors"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/service"
    "github.com/plexmediamanager/service/helpers"
)

type DatabaseClient struct {
    databaseHost        string
    databasePort        uint64
    databaseUsername    string
    databasePassword    string
    databaseName        string
}

var (
    application         *service.Application
    database            *gorm.DB
)

// Initialize database client
func Initialize(app *service.Application) *DatabaseClient {
    application = app
    return &DatabaseClient {
        databaseHost:       helpers.GetEnvironmentVariableAsString("DATABASE_HOST", "localhost"),
        databasePort:       uint64(helpers.GetEnvironmentVariableAsInteger("DATABASE_PORT", 3306)),
        databaseUsername:   helpers.GetEnvironmentVariableAsString("DATABASE_USERNAME", "forge"),
        databasePassword:   helpers.GetEnvironmentVariableAsString("DATABASE_PASSWORD", "secret"),
        databaseName:       helpers.GetEnvironmentVariableAsString("DATABASE_NAME", "app"),
    }
}

// Connect to the database
func (client *DatabaseClient) Connect() error {
    var err error
    database, err = gorm.Open("mysql", client.buildDatabaseDSN())
    if err != nil {
        return errors.DatabaseConnectionError.ToError(err)
    }
    client.autoMigrate()
    return nil
}

// Get database connection instance
func Instance() *gorm.DB {
    return database
}

// Disconnect from the database
func Disconnect() error {
    return database.Close()
}

// Build database DSN connection string
func (client *DatabaseClient) buildDatabaseDSN() string {
    return format.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
        client.databaseUsername,
        client.databasePassword,
        client.databaseHost,
        client.databasePort,
        client.databaseName,
    )
}

// Perform automatic migration
func (client *DatabaseClient) autoMigrate() {
    database.AutoMigrate(
        &models.Genre{},
        &models.Language{},
        &models.Movie{},
        &models.Permission{},
        &models.User{},
    )
}