package main

import (
	"net/http"

	"todoservice/internal/adding"
	"todoservice/internal/http/rest"
	"todoservice/internal/listing"
	"todoservice/internal/storage"
	"todoservice/internal/updating"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Only log the Info severity or above.
	logrus.SetLevel(logrus.InfoLevel)

	storage, err := storage.New("./todo.db")
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	addService := adding.NewService(storage)
	updateService := updating.NewService(storage)
	getService := listing.NewService(storage)

	router := rest.Handler(addService, getService, updateService)

	err = http.ListenAndServe(":12001", router)
	if err != nil {
		logrus.Errorf("Server exited with error: %v", err)
	}
}
