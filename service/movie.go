package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
)

func movieToStructure(result *proto.DatabaseResponse, err error) (*models.Movie, error) {
    var record *models.Movie
    return record, protoToStructure(&record, result, err)
}

func movieArrayToStructure(result *proto.DatabaseResponse, err error) ([]*models.Movie, error) {
    var record []*models.Movie
    return record, protoToStructure(&record, result, err)
}

func GetMovieService(client microClient.Client) proto.MovieService {
    return proto.NewMovieService(getServiceName(movieServiceName), client)
}

func MovieServiceFindDownloaded (client microClient.Client) ([]*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseEmpty {}
    return movieArrayToStructure(service.FindDownloaded(context.TODO(), parameters))
}