# todo-service
[![Build Status](https://travis-ci.org/davidseyboldblog/todo-service.svg?branch=master)](https://travis-ci.org/davidseyboldblog/todo-service)
```
Request:
POST /todo-service/todo
{
  "userId": "1",
  "description": "Thing to do",
  "complete": false
}

Response:
201 CREATED
{
  "id": 1
}
```

```
Request:
GET /todo-service/todo/1

Response:
200 OK
{
    "id": 1,
    "userId": 1,
    "description": "Thing to do",
    "complete": false
}
```

```
Request:
PUT /todo-service/todo/1
{
  "userId": "1",
  "description": "Thing to do",
  "complete": false
}

Response:
204 NO CONTENT
```

```
Request:
GET /todo-service/todo?userId=1

Response:
200 OK
{
    "todos": [
        {
            "id": 1,
            "userId": 1,
            "description": "Todo 1",
            "complete": false
        },
        {
            "id": 2,
            "userId": 1,
            "description": "Todo 2",
            "complete": false
        }
    ]
}
```
