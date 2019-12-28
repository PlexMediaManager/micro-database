package database

import (
    format "fmt"
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

type DatabaseOrder struct {
    Column                  string
    Order                   string
}

type DatabaseQuery struct {
    Output                  interface{}
    Query                   interface{}
    Parameters              interface{}
    All                     bool
    Multiple                bool
    Count                   bool
    Counted                 int64
    Order                   DatabaseOrder
    Limit                   int64
    Offset                  int64
    Columns                 []string
}

// Get record from the database
func GetRecord(query *DatabaseQuery) error {
    var err error
    var builder *gorm.DB

    if query.All {
        if query.Columns != nil && len(query.Columns) > 0 {
            err = database.Select(query.Columns).Find(query.Output).Error
        } else {
            err = database.Find(query.Output).Error
        }
        if err != nil {
            return errors.DatabaseUndefinedError.ToErrorWithArguments(err, reflect.TypeOf(query.Output), err)
        }
        return nil
    }

    if query.Parameters == nil {
        builder = database.Where(query.Query)
    } else {
        builder = database.Where(query.Query, query.Parameters)
    }

    if query.Order.Column != "" {
        builder = builder.Order(format.Sprintf("%s %s", query.Order.Column, query.Order.Order))
    }

    if query.Limit > 0 {
        builder = builder.Limit(query.Limit)
    }

    if query.Offset > 0 {
        builder = builder.Offset(query.Offset)
    }

    if query.Columns != nil && len(query.Columns) > 0 {
        builder = builder.Select(query.Columns)
    }

    if query.Multiple {
        builder = builder.Find(query.Output)
    } else {
        builder = builder.First(query.Output)
    }

    if query.Count {
        builder = builder.Count(&query.Counted)
    }

    err = builder.Error

    if err != nil {
        return errors.DatabaseUndefinedError.ToErrorWithArguments(err, reflect.TypeOf(query.Output), err)
    }
    return nil
}

// Count total number of records on the table
func GetCount(model interface{}) (uint64, error) {
    var count uint64
    err := database.Model(model).Count(&count).Error
    if err != nil {
        return 0, errors.DatabaseCountError.ToErrorWithArguments(err, reflect.TypeOf(model), err)
    }
    return count, nil
}

// Check if record already exists in the database
func RecordExists(model Model) (bool, error) {
    var count uint64
    if model == nil {
        return false, errors.DatabaseModelNull.ToError(nil)
    }

    err := database.Table(model.TableName()).Where(model).Count(&count).Error
    if err != nil {
        return false, errors.DatabaseCountError.ToErrorWithArguments(err, reflect.TypeOf(model), err)
    }
    return count > 0, nil
}

type CreationParameters struct {
    Model               Model
    CreateWithID        bool
}

// Create database record
func CreateRecord(parameters CreationParameters) error {
    var err error
    if parameters.Model == nil {
        return errors.DatabaseModelNull.ToError(nil)
    }

    transaction := database.Begin()
    if transaction.NewRecord(parameters.Model) || parameters.CreateWithID {
        if beforeCreateError := parameters.Model.BeforeCreate(transaction); beforeCreateError != nil {
            transaction.Rollback()
            return errors.DatabaseModelCreationError.ToErrorWithArguments(beforeCreateError, parameters.Model)
        } else {
            err = transaction.Create(parameters.Model).Error
            transaction.Commit()
        }
    } else {
        transaction.Rollback()
        return errors.DatabaseModelCreationWithPrimaryKey.ToErrorWithArguments(nil, parameters.Model.TableName())
    }

    if err != nil {
        return errors.DatabaseModelCreationError.ToErrorWithArguments(err, parameters.Model.TableName(), err)
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