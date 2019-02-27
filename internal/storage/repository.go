package storage

import (
	"database/sql"

	"todoservice/internal/listing"

	//Database driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	createTableStatement = "CREATE TABLE IF NOT EXISTS todo (id INTEGER PRIMARY KEY AUTOINCREMENT, userId INTEGER, description TEXT, complete INTEGER)"
	getTodoQuery         = "SELECT id, userId, description, complete from todo where id = ?"
	updateTodoQuery      = "UPDATE todo set userId = ?, description = ?, complete = ? where id = ?"
	insertTodoQuery      = "INSERT INTO todo(userId, description, complete) VALUES(?, ?, ?)"
	getTodoListQuery     = "SELECT id, userId, description, complete from todo where userId = ?"
)

type Storage struct {
	db *sql.DB
}

func New(host string) (*Storage, error) {
	s := new(Storage)
	var err error
	s.db, err = sql.Open("sqlite3", host)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening DB connection")
	}

	statement, err := s.db.Prepare(createTableStatement)
	if err != nil {
		return nil, errors.Wrap(err, "Error preparing statement")
	}
	_, err = statement.Exec()
	if err != nil {
		return nil, errors.Wrap(err, "Error creating table")
	}

	logrus.Info("Connection to database successful")
	return s, nil
}

func (s *Storage) Add(userID int64, desc string, complete int32) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "Error opening transaction")
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(insertTodoQuery)
	if err != nil {
		return 0, errors.Wrap(err, "Error preparing user sql statement")
	}

	res, err := stmt.Exec(userID, desc, complete)
	if err != nil {
		return 0, errors.Wrap(err, "Error executing todo sql statement")
	}

	todoID, err := res.LastInsertId()

	tx.Commit()
	return todoID, nil
}

func (s *Storage) Update(id int64, userID int64, desc string, complete int32) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "Error opening transaction")
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(updateTodoQuery)
	if err != nil {
		return errors.Wrap(err, "Error preparing todo sql statement")
	}

	_, err = stmt.Exec(userID, desc, complete, id)
	if err != nil {
		return errors.Wrap(err, "Error executing todo sql statement")
	}

	tx.Commit()
	return nil
}

func (s *Storage) Get(id int64) (listing.Todo, error) {
	var todo listing.Todo
	var complete int32
	err := s.db.QueryRow(getTodoQuery, id).Scan(&todo.ID, &todo.UserID, &todo.Description, &complete)

	todo.Complete = complete == 1

	if err == sql.ErrNoRows {
		//Return empty Todo if no rows returned
		return listing.Todo{}, nil
	} else if err != nil {
		return listing.Todo{}, errors.Wrap(err, "Error querying database")
	}

	return todo, nil
}

func (s *Storage) GetList(userID int64) ([]listing.Todo, error) {
	var todoList []listing.Todo
	todoRows, err := s.db.Query(getTodoListQuery, userID)
	defer todoRows.Close()

	if err != nil {
		return nil, errors.Wrap(err, "Error querying database")
	}
	for todoRows.Next() {
		var todo listing.Todo
		var complete int32

		todoRows.Scan(&todo.ID, &todo.UserID, &todo.Description, complete)
		todo.Complete = complete == 1
		todoList = append(todoList, todo)
	}
	err = todoRows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "Error scanning results")
	}
	return todoList, nil
}
