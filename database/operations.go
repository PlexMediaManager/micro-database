package database

import (
    "github.com/jinzhu/gorm"
    "github.com/plexmediamanager/micro-database/errors"
    "reflect"
)

type Model interface {
    PrimaryKey()                                            uint64
    BeforeCreate(*gorm.DB)                                  error
    BeforeUpdate(*gorm.DB)                                  error
    BeforeDelete(*gorm.DB)                                  error
    TableName()                                             string
    UpdatedFields(updateID bool, updatedFields ...string)   map[string]interface{}
}

// Get single record from the database
func GetSingleRecord(output interface{}, query interface{}) error {
    err := database.Where(query).First(output).Error
    if err != nil {
        return errors.DatabaseModelNotFound.ToErrorWithArguments(err, output.(Model).TableName())
    }
    return nil
}

// Get multiple records from the database
func GetMultipleRecords(output interface{}, query interface{}, parameters interface{}) error {
    var err error
    if parameters != nil {
        err = database.Where(query, parameters).Find(output).Error
    } else {
        err = database.Where(query).Find(output).Error
    }
    if err != nil {
        return errors.DatabaseUndefinedError.ToErrorWithArguments(err, reflect.TypeOf(output), err)
    }
    return nil
}

// Get all records from the database
func GetAllRecords(output interface{}) error {
    err := database.Find(output).Error
    if err != nil {
        return errors.DatabaseUndefinedError.ToErrorWithArguments(err, reflect.TypeOf(output), err)
    }
    return nil
}

// Create database record
func CreateRecord(model Model) error {
    var err error
    if model == nil {
        return errors.DatabaseModelNull.ToError(nil)
    }
    transaction := database.Begin()
    if transaction.NewRecord(model) {
        if creationError := model.BeforeCreate(transaction); creationError == nil {
            err = transaction.Create(model).Error
            transaction.Commit()
        } else {
            transaction.Rollback()
            return errors.DatabaseModelCreationError.ToErrorWithArguments(creationError, model)
        }
    } else {
        transaction.Rollback()
        return errors.DatabaseModelCreationWithPrimaryKey.ToErrorWithArguments(nil, model.TableName())
    }
    if err != nil {
        return errors.DatabaseModelCreationError.ToErrorWithArguments(err, model)
    }
    return nil
}

// Update database record
func UpdateRecord(model Model, fieldValueMap map[string]interface{}) error {
    var err error
    if model == nil {
        return errors.DatabaseModelNull.ToError(nil)
    }
    if model.PrimaryKey() > 0 {
        if database.First(model, model.PrimaryKey()).RecordNotFound() {
            return errors.DatabaseModelNotFound.ToError(nil)
        }
        err = database.Model(model).Updates(fieldValueMap).Error
    } else {
        return errors.DatabaseModelMissingRequiredField.ToErrorWithArguments(nil, "ID", model.TableName())
    }
    if err != nil {
        return errors.DatabaseModelUpdateError.ToErrorWithArguments(err, model)
    }
    return nil
}

// Soft Delete model from the database
func SoftDeleteRecord(model Model) error {
    var err error
    if model == nil {
        return errors.DatabaseModelNull.ToError(nil)
    }
    if model.PrimaryKey() > 0 {
        err = database.Delete(model).Error
    } else {
        return errors.DatabaseModelMissingRequiredField.ToErrorWithArguments(nil, "ID", model.TableName())
    }
    if err != nil {
        return errors.DatabaseModelDeletionError.ToErrorWithArguments(err, model)
    }
    return nil
}

// Force Delete model from the database
func ForceDeleteRecord(model Model) error {
    var err error
    if model == nil {
        return errors.DatabaseModelNull.ToError(nil)
    }
    if model.PrimaryKey() > 0 {
        err = database.Unscoped().Delete(model).Error
    } else {
        return errors.DatabaseModelMissingRequiredField.ToErrorWithArguments(nil, "ID", model.TableName())
    }
    if err != nil {
        return errors.DatabaseModelDeletionError.ToErrorWithArguments(err, model)
    }
    return nil
}