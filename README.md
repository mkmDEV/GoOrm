# GoOrm

Followed [Creating a Go API using an ORM - Tutorial](https://www.youtube.com/watch?v=VAGodyl84OY "Creating a Go API using an ORM - Tutorial
") to implement a REST API backend with ORM SQLite DB written in [GO](https://golang.org/ "GO Lang")

### Install [mux router](https://github.com/gorilla/mux "mux router") and [gorm](https://gorm.io/ "gorm") before run this application
```BASH
$ go get -u github.com/gorilla/mux
$ go get -u github.com/jinzhu/gorm
```

### Run this app in `terminal` with
```BASH
$ go run *.go
```

### Endpoints at `http://localhost:8081/`
use [Postman](https://www.getpostman.com/ "Postman") to test them
- [ ] `GET /` print *"Hello World"*
- [ ] `GET /users` get all users
- [ ] `POST /user/{name}/{email}` create a new user
  <details><summary>Request sample</summary>
  
  ```
    http://localhost:8081/user/elliot/no@email.com
  ```
  
  </details>

- [ ] `DELETE /user/{name}` delete a user
  <details><summary>Request sample</summary>
  
  ```
    http://localhost:8081/user/elliot
  ```
  
  </details>
  
- [ ] `PUT /user/{name}/{email}` update a user
  <details><summary>Request sample</summary>
  
  ```
  http://localhost:8081/user/elliot/another.new@email.com
  ```
  </details>
