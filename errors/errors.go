package errors

import "github.com/plexmediamanager/service/errors"

const (
    ServiceID       errors.Service      =   2
)

var (
    // Library Errors Group
    DatabaseConnectionError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    1,
        },
        Message:    "Failed to connect to the database",
    }
    DatabaseMarshalError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    2,
        },
        Message:    "Failed to convert database model to bytes",
    }
    DatabaseUnmarshalError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    3,
        },
        Message:    "Failed to convert bytes to database model",
    }

    // Service Errors Group
    DatabaseModelNull = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    1,
        },
        Message:    "Provided model is not set and equals to NULL",
    }
    DatabaseModelCreationError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    2,
        },
        Message:    "Unable to create model '%v': %v",
    }
    DatabaseModelCreationWithPrimaryKey = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    3,
        },
        Message:    "Unable to create model on '%s' as specified primary key already exists",
    }
    DatabaseModelMissingRequiredField = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    4,
        },
        Message:    "Missing required field `%s` on the model %s",
    }
    DatabaseModelDeletionError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    5,
        },
        Message:    "Unable to delete model '%s' %s",
    }
    DatabaseModelNotFound = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    6,
        },
        Message:    "Unable to find model on '%s'",
    }
    DatabaseModelUpdateError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    7,
        },
        Message:    "Unable to update model '%s' %s",
    }
    DatabaseUndefinedError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    7,
        },
        Message:    "Failed to execute database query on '%s' with error: %v",
    }
    DatabaseCountError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    8,
        },
        Message:    "Failed to count rows on '%s' with error: %v",
    }
    DatabaseUserNotFoundError = errors.Error {
        Code:       errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeService,
            ErrorNumber:    9,
        },
        Message:    "Failed to fetch information for user with specified query",
    }
)