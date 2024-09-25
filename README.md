This project is a simple RESTful API built using Go and the chi router framework. It serves a hardcoded JSON response to illustrate how to create an HTTP server in Go, handle GET requests, and return JSON responses.

Features
Uses Go's built-in net/http package for HTTP requests handling.
Utilizes the chi lightweight router for route management.
Returns a simple hardcoded JSON response as an example.
Encodes response data into JSON format.
Endpoints
Method	Endpoint	Description
GET	/json	Returns a hardcoded JSON response
Response Structure
The API returns a JSON response in the following structure:

json
Copy code
{
  "message": "Hello, this is a hardcoded JSON response!"
}
Technologies Used
Go - A statically typed, compiled programming language designed for simplicity and efficiency.
Chi Router - A lightweight, idiomatic, and composable router for building HTTP services in Go.
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/go-json-api.git
cd go-json-api
Install dependencies (Chi package):

bash
Copy code
go get github.com/go-chi/chi/v5
Run the application:

bash
Copy code
go run main.go
The server will start on http://localhost:8000. You can access the endpoint by visiting http://localhost:8000/json in your browser or using curl.

Example Request
GET /json

bash
Copy code
curl -X GET http://localhost:8000/json
Response:

json
Copy code
{
  "message": "Hello, this is a hardcoded JSON response!"
}
