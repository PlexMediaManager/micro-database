package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-database/proto"
    "github.com/plexmediamanager/micro-database/repository"
)

type UserService struct {}

func (service UserService) FindOneByID (_ context.Context, parameters *proto.DatabaseID, response *proto.DatabaseResponse) error {
    result, err := repository.User.FindOneByID(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service UserService) FindManyByID (_ context.Context, parameters *proto.DatabaseIDs, response *proto.DatabaseResponse) error {
    result, err := repository.User.FindManyByID(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}

func (service UserService) FindOneByUsername (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.User.FindOneByUsername(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service UserService) FindManyByUsername (_ context.Context, parameters *proto.DatabaseStrings, response *proto.DatabaseResponse) error {
    result, err := repository.User.FindManyByUsername(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}

func (service UserService) FindOneByEmail (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.User.FindOneByEmail(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service UserService) FindManyByEmail (_ context.Context, parameters *proto.DatabaseStrings, response *proto.DatabaseResponse) error {
    result, err := repository.User.FindManyByEmail(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}