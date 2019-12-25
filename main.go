package main

import (
    "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/database"
    "github.com/plexmediamanager/micro-database/proto"
    "github.com/plexmediamanager/micro-database/resolver"
    "github.com/plexmediamanager/service"
    "github.com/plexmediamanager/service/log"
    "time"
)

func main() {
    application := service.CreateApplication()
    // Initialize Application Configuration
    err := application.InitializeConfiguration()
    if err != nil {
        log.Panic(err)
    }
    // Initialize Application Micro Service
    err = application.InitializeMicroService()
    if err != nil {
        log.Panic(err)
    }

    err = application.Service().Client().Init(
        client.PoolSize(10),
        client.Retries(30),
        client.RequestTimeout(1 * time.Second),
    )
    if err != nil {
        log.Panic(err)
    }

    err = database.Initialize().Connect()
    if err != nil {
        log.Panic(err)
    }

    registerResolvers(application)
    go application.StartMicroService()

    service.WaitForOSSignal(1)
}

func registerResolvers(application *service.Application) {
    server := application.Service().Server()

    if err := proto.RegisterGenreServiceHandler(server, new(resolver.GenreService)); err != nil {
        log.Panic(err)
    }

    if err := proto.RegisterMovieServiceHandler(server, new(resolver.MovieService)); err != nil {
        log.Panic(err)
    }

    if err := proto.RegisterUserServiceHandler(server, new(resolver.UserService)); err != nil {
        log.Panic(err)
    }

    if err := proto.RegisterPermissionServiceHandler(server, new(resolver.PermissionService)); err != nil {
        log.Panic(err)
    }
}