# Forum REST Api

A simple REST Api for a forum like Reddit built in Go

# To Do

- [x] Relational data modelling

- [x] Create models
- [x] Implement validations for models
- [x] Declare repository interfaces
- [x] Implement repositories for SQLite3
    - [x] Implement custom errors
- [x] Test repositories

- [x] Declare service interfaces
- [x] Implement services

- [ ] Implement controllers

- [x] Implement router
- [ ] Implement auth
- [ ] Implement logging

- [x] Dockerize

# New Objectives

- [ ] Implement error handling for all errors in all layers
- [ ] Fix Update methods for all services
- [ ] Log all errors that lead to a InternalServerError
- [ ] Improve repository documentation
- [ ] Improve service documentation
- [ ] Test Services
- [ ] Test Controllers
- [ ] Reorganize util package

# Required environment variables

* SQLITE_DB_FILE_NAME
* SQLITE_DB_FOLDER_NAME
* PORT
