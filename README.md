# Forum REST Api

A simple REST Api for a forum like Reddit built in Go

# How to run locally

This project requires go and gcc installed in your system

First, clone this repository
```bash
git clone https://github.com/AlejandroJorge/Forum-REST-api.git
cd Forum-REST-api
```

Then, create a .env file on project folder with the following environment variables:
- SQLITE_DB_FOLDER_NAME
- SQLITE_DB_FILE_NAME
- PORT
- AUTH_SECRET

## Build natively

This will build and test the project
```bash
make
```

To run it then simply:
```bash
./build/server
```

## Build docker image

This will build the docker image
```bash
docker build -t forum-rest-api:latest .
```

To run it then simply:
```bash
docker run forum-rest-api:latest
```

# Development Roadmap

Where is development going right now

## Code quality and reliability improvement

- [x] Reorganize testing
- [x] Check exported functions / types / constants
- [x] Custom errors for data / repository layer (include default error)
- [x] Logging for data / repository layer errors
- [x] Remove unnecesary transactions
- [x] Custom errors for domain / service layer (include default error)
- [x] Logging for domain / service layer errors
- [x] Expand update methods for services (decompress)
- [x] Documentation in comments for repositories (errors, optional args, etc)
- [x] Documentation in comments for services (errors, optional args, etc)
- [x] Refactor controllers with error handling
- [x] Implement auth for every necessary operation
- [ ] Remake tests for repositories
- [ ] Repository level validation for tests
- [ ] Make tests for services
- [ ] Make tests for controllers / routes

## Next features

- [ ] Auth expiration
- [ ] CORS fix
- [ ] Pagination for profiles
- [ ] Pagination for posts
- [ ] Pagination for comments
- [ ] Searching for posts
- [ ] Multiple subforums
- [ ] Different auth levels
- [ ] Admin dashboard

