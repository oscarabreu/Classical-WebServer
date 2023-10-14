# RESTful API for Form Management

This project provides a RESTful API for managing form submissions. It supports basic CRUD (Create, Read, Update, Delete) operations for form data, with in-memory storage.

## Features

- Retrieve all submitted forms.
- Retrieve a specific form submission by ID.
- Submit a new form.
- Update an existing form submission by ID.
- Delete a form submission by ID.

## Prerequisites

1. Go (Golang) should be installed on your system.

## Installation & Usage

1. Clone or download the program's source code.
2. Open a terminal and navigate to the directory containing the program.
3. Compile the program using the following command: `go build main.go`
4. Run the compiled program: `./main`
5. The API will start on port 8080.

## API Endpoints

### Retrieve All Forms

**Request:** 
`GET /api/forms`

**Response:** 
JSON array of all forms.

### Retrieve Specific Form

**Request:** 
`GET /api/form?id={ID}`

**Response:** 
JSON object of the form with the specified ID.

### Submit New Form

**Request:** 
`POST /api/form/create`

**Body:** 
```json
{
 "name": "John Doe",
 "address": "123 Main St"
}
```
**Response:**
JSON object of the created form, including its ID.

### Update Existing Form

**Request:**
`PUT /api/form/update?id={ID}`

**Body:**
```
{
    "name": "Jane Doe",
    "address": "456 Elm St"
}
```
**Response:**
JSON object of the updated form.

### Delete Form

**Request**:
`DELETE /api/form/delete?id={ID}`

**Response**:
204 No Content if successful.

## Limitations

The current implementation uses in-memory storage, meaning data won't persist across server restarts. For a more robust solution, consider integrating a database.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change
