package repository

import (
    "github.com/plexmediamanager/micro-database/database"
    "github.com/plexmediamanager/micro-database/models"
)

type UserRepository struct {}

// Create new user
func (*UserRepository) Create(model *models.User) (*models.User, error) {
    return model, database.CreateRecord(model)
}

// Update existing user
func (*UserRepository) Update(model *models.User, fields UpdatedFields) (*models.User, error) {
    return model, database.UpdateRecord(model, fields)
}

// Soft delete existing user
func (*UserRepository) SoftDelete(model *models.User) (*models.User, error) {
    return model, database.SoftDeleteRecord(model)
}

// Force delete existing user
func (*UserRepository) ForceDelete(model *models.User) (*models.User, error) {
    return model, database.ForceDeleteRecord(model)
}

// Find single user by ID
func (repository *UserRepository) FindOneByID(id uint64) (*models.User, error) {
    return repository.findOne(models.User{ ID: id })
}

// Find many users by ID
func (repository *UserRepository) FindManyByID(id ...uint64) ([]*models.User, error) {
    return repository.findMany(models.KeyUserID + " IN (?)", id)
}

// Find single user by Username
func (repository *UserRepository) FindOneByUsername(username string) (*models.User, error) {
    return repository.findOne(models.User{ Username: username })
}

// Find many users by Username
func (repository *UserRepository) FindManyByUsername(username ...string) ([]*models.User, error) {
    return repository.findMany(models.KeyUserUsername + " IN (?)", username)
}

// Find single user by Email
func (repository *UserRepository) FindOneByEmail(email string) (*models.User, error) {
    return repository.findOne(models.User{ Email: email })
}

// Find many users by E-Mail
func (repository *UserRepository) FindManyByEmail(email ...string) ([]*models.User, error) {
    return repository.findMany(models.KeyUserEmail + " IN (?)", email)
}

// Find single user
func (*UserRepository) findOne(query interface{}) (*models.User, error) {
    user := &models.User{}
    err := database.GetSingleRecord(user, query)
    return user, err
}

// Find multiple users
func (*UserRepository) findMany(query interface{}, parameters interface{}) ([]*models.User, error) {
    var users []*models.User
    err := database.GetMultipleRecords(&users, query, parameters)
    return users, err
}