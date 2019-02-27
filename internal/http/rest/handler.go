package rest

import (
	"encoding/json"
	"net/http"
	"os"

	"todoservice/internal/adding"
	"todoservice/internal/listing"
	"todoservice/internal/updating"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

const (
	//ContentType header for type of content in response
	ContentType = "Content-Type"
	//ApplicationJSON value for content type in response
	ApplicationJSON = "application/json"
	//InternalServerErrorMessage message for an unexpected error
	InternalServerErrorMessage = "Internal Server Error"
	//BadRequestMessage message for a bad request sent from client
	BadRequestMessage = "Bad Request"
)

// Handler initializes routes
func Handler(addService adding.Service, listingService listing.Service, updateService updating.Service) http.Handler {
	router := httprouter.New()

	logrus.Info("Initalizing routes")
	router.POST("/todo-service/todo", addTodo(addService))
	router.PUT("/todo-service/todo/:id", updateTodo(updateService))
	router.GET("/todo-service/todo/:id", getTodo(listingService))
	router.GET("/todo-service/todo", getTodoList(listingService))

	logrus.Info("Adding Middleware")
	handler := handlers.CombinedLoggingHandler(os.Stdout, router)
	return handler
}

func addTodo(s adding.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var todo adding.Todo
		err := decoder.Decode(&todo)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusBadRequest, BadRequestMessage)
			return
		}

		todoID, err := s.AddTodo(todo)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusInternalServerError, InternalServerErrorMessage)
			return
		}

		w.Header().Set(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todoID)
	}
}

func updateTodo(s updating.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		id := ps.ByName("id")

		var todo updating.Todo

		err := decoder.Decode(&todo)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusBadRequest, BadRequestMessage)
			return
		}

		err = s.UpdateTodo(id, todo)

		w.Header().Set(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusNoContent)
	}
}

func getTodo(s listing.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		todo, err := s.GetTodo(id)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusInternalServerError, InternalServerErrorMessage)
			return
		}

		w.Header().Set(ContentType, ApplicationJSON)
		json.NewEncoder(w).Encode(todo)
	}
}

func getTodoList(s listing.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		queryValues := r.URL.Query()
		userID := queryValues.Get("userId")
		if userID == "" {
			returnError(w, http.StatusBadRequest, BadRequestMessage)
			return
		}
		todoList, err := s.GetTodoList(userID)
		if err != nil {
			logrus.Error(err)
			returnError(w, http.StatusInternalServerError, InternalServerErrorMessage)
			return
		}

		w.Header().Set(ContentType, ApplicationJSON)
		json.NewEncoder(w).Encode(todoList)
	}
}

func returnError(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(httpCode)

	json.NewEncoder(w).Encode(message)
}
