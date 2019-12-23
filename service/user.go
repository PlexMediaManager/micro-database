package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
)

func userToStructure(result *proto.DatabaseResponse, err error) (*models.User, error) {
    var record *models.User
    return record, protoToStructure(&record, result, err)
}

func userArrayToStructure(result *proto.DatabaseResponse, err error) ([]*models.User, error) {
    var record []*models.User
    return record, protoToStructure(&record, result, err)
}

func GetUserService(client microClient.Client) proto.UserService {
    return proto.NewUserService(getServiceName(userServiceName), client)
}

func UserServiceFindOneByID (client microClient.Client, id uint64) (*models.User, error) {
    service := GetUserService(client)
    parameters := &proto.DatabaseID { Value: id }
    return userToStructure(service.FindOneByID(context.TODO(), parameters))
}

func UserServiceFindManyByID (client microClient.Client, id ...uint64) ([]*models.User, error) {
    service := GetUserService(client)
    parameters := &proto.DatabaseIDs { Values: id }
    return userArrayToStructure(service.FindManyByID(context.TODO(), parameters))
}

func UserServiceFindOneByUsername (client microClient.Client, username string) (*models.User, error) {
    service := GetUserService(client)
    parameters := &proto.DatabaseString { Value: username }
    return userToStructure(service.FindOneByUsername(context.TODO(), parameters))
}

func UserServiceFindManyByUsername (client microClient.Client, username ...string) ([]*models.User, error) {
    service := GetUserService(client)
    parameters := &proto.DatabaseStrings { Values: username }
    return userArrayToStructure(service.FindManyByUsername(context.TODO(), parameters))
}

func UserServiceFindOneByEmail (client microClient.Client, email string) (*models.User, error) {
    service := GetUserService(client)
    parameters := &proto.DatabaseString { Value: email }
    return userToStructure(service.FindOneByEmail(context.TODO(), parameters))
}

func UserServiceFindManyByEmail (client microClient.Client, email ...string) ([]*models.User, error) {
    service := GetUserService(client)
    parameters := &proto.DatabaseStrings { Values: email }
    return userArrayToStructure(service.FindManyByEmail(context.TODO(), parameters))
}
