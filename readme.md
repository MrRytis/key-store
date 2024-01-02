# This is simple Go application for key value data storage.

### What is needed to run this application
* Go 1.21 or higher
* Makefile installed

### How to run
* Clone this repository
* Run `make run` command
* Open `http://localhost:3000` in your browser

### How to use
Currently there are 4 endpoints available:
* `GET /api/v1/store` - returns all keys and values
* `GET /api/v1/store/{key}` - returns value for given key
* `POST /api/v1/store` - creates or updates value for given key
* `DELETE /api/v1/store/{key}` - deletes value for given key

### API docs
API docs are available [here](http://localhost:3000/docs/index.html)