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

- [ ] Reorganize testing
- [ ] Custom errors for data / repository layer (include default error)
- [ ] Logging for data / repository layer errors
- [ ] Repository level validation
- [ ] Remove unnecesary transactions
- [ ] Custom errors for domain / service layer (include default error)
- [ ] Logging for domain / service layer errors
- [ ] Validation of JSON requests in controllers (mandatory fields, format, etc)
- [ ] Expand update methods for services (decompress)
- [ ] Documentation in comments for repositories (errors, optional args, etc)
- [ ] Documentation in comments for services (errors, optional args, etc)
- [ ] Remake tests for repositories
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

