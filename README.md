# RESTful API for Object Management

This project provides a RESTful API for managing game submissions. It supports basic CRUD (Create, Read, Update, Delete) operations for JSON objects, with in-memory storage.

## Features

- Retrieve all submitted games.
- Retrieve a specific game submission by ID.
- Submit a new game.
- Update an existing game submission by ID.
- Delete a game submission by ID.

## Prerequisites

1. Go (Golang) should be installed on your system.

## Installation & Usage

1. Clone or download the program's source code.
2. Open a terminal and navigate to the directory containing the program.
3. Compile the program using the following command: `go build main.go`
4. Run the compiled program: `./main`
5. The API will start on port 8080.

## API Endpoints

### Retrieve All games

**Request:** 
`GET /api/games`

**Response:** 
JSON array of all games.

### Retrieve Specific game

**Request:** 
`GET /api/game?id={ID}`

**Response:** 
JSON object of the game with the specified ID.

### Submit New game

**Request:** 
`POST /api/game/create`

**Body:** 
```json
{
 "name": "The Legend of Zelda: Skyward Sword",
 "console": "Nintendo Switch",
 "genre": "Adventure",
 "datePub": "2023-05-12"
}

```
**Response:**
JSON object of the created game, including its ID.

### Update Existing game

**Request:**
`PUT /api/game/update?id={ID}`

**Body:**
```
{
    "name": "Stardew Valley",
    "console": "PC",
    "genre": "Indie",
    "datePub": "2016-02-26"
}

```
**Response:**
JSON object of the updated game.

### Delete game

**Request**:
`DELETE /api/game/delete?id={ID}`

**Response**:
204 No Content if successful.

## Limitations

The current implementation uses in-memory storage, meaning data won't persist across server restarts. For a more robust solution, consider integrating a database.

