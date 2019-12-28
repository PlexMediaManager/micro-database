package repository

import (
    "github.com/plexmediamanager/micro-database/database"
    "github.com/plexmediamanager/micro-database/models"
)

type PermissionRepository struct {}

// Create new permission
func (*PermissionRepository) Create(model *models.Permission) (*models.Permission, error) {
    return model, database.CreateRecord(database.CreationParameters {
        Model:        model,
        CreateWithID: false,
    })
}

// Update existing permission
func (*PermissionRepository) Update(model *models.Permission, fields UpdatedFields) (*models.Permission, error) {
    return model, database.UpdateRecord(model, fields)
}

// Soft delete existing permission
func (*PermissionRepository) SoftDelete(model *models.Permission) (*models.Permission, error) {
    return model, database.SoftDeleteRecord(model)
}

// Force delete existing permission
func (*PermissionRepository) ForceDelete(model *models.Permission) (*models.Permission, error) {
    return model, database.ForceDeleteRecord(model)
}

// Find single permission by ID
func (repository *PermissionRepository) FindOneByID(id uint64) (*models.Permission, error) {
    return repository.findOne(models.Permission{ ID: id })
}

// Find many permissions by ID
func (repository *PermissionRepository) FindManyByID(id ...uint64) ([]*models.Permission, error) {
    return repository.findMany(models.KeyPermissionID + " IN (?)", id)
}

// Find single permission by Name
func (repository *PermissionRepository) FindOneByName(name string) (*models.Permission, error) {
    return repository.findOne(models.Permission { Name: name })
}

// Find single permission
func (*PermissionRepository) findOne(query interface{}) (*models.Permission, error) {
    record := &models.Permission{}
    err := database.GetRecord(&database.DatabaseQuery{
        Output:     record,
        Query:      query,
    })
    return record, err
}

// Find multiple permissions
func (*PermissionRepository) findMany(query interface{}, parameters interface{}) ([]*models.Permission, error) {
    var records []*models.Permission
    err := database.GetRecord(&database.DatabaseQuery{
        Output:     &records,
        Query:      query,
        Parameters: parameters,
        Multiple:   true,
    })
    return records, err
}