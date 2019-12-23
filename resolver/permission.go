package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-database/proto"
    "github.com/plexmediamanager/micro-database/repository"
)

type PermissionService struct {}

func (service PermissionService) FindOneByID (_ context.Context, parameters *proto.DatabaseID, response *proto.DatabaseResponse) error {
    result, err := repository.Permission.FindOneByID(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service PermissionService) FindManyByID (_ context.Context, parameters *proto.DatabaseIDs, response *proto.DatabaseResponse) error {
    result, err := repository.Permission.FindManyByID(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}

func (service PermissionService) FindOneByName (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.Permission.FindOneByName(parameters.Value)
    return structureToBytesWithError(result, err, response)
}
