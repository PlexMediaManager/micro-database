package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
)

func permissionToStructure(result *proto.DatabaseResponse, err error) (*models.Permission, error) {
    var record *models.Permission
    return record, protoToStructure(&record, result, err)
}

func permissionArrayToStructure(result *proto.DatabaseResponse, err error) ([]*models.Permission, error) {
    var record []*models.Permission
    return record, protoToStructure(&record, result, err)
}

func GetPermissionService(client microClient.Client) proto.PermissionService {
    return proto.NewPermissionService(getServiceName(permissionServiceName), client)
}

func PermissionServiceFindOneByID (client microClient.Client, id uint64) (*models.Permission, error) {
    service := GetPermissionService(client)
    parameters := &proto.DatabaseID { Value: id }
    return permissionToStructure(service.FindOneByID(context.TODO(), parameters))
}

func PermissionServiceFindManyByID (client microClient.Client, id ...uint64) ([]*models.Permission, error) {
    service := GetPermissionService(client)
    parameters := &proto.DatabaseIDs { Values: id }
    return permissionArrayToStructure(service.FindManyByID(context.TODO(), parameters))
}

func PermissionServiceFindOneByName (client microClient.Client, name string) (*models.Permission, error) {
    service := GetPermissionService(client)
    parameters := &proto.DatabaseString { Value: name }
    return permissionToStructure(service.FindOneByName(context.TODO(), parameters))
}