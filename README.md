# Leaderboard

A simple (in-memory) leaderboard implemented in:

- [ ] Node.js
- [x] Go

![API Architecture](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-http-api.png "Architecture")

The leaderboard apis handle the following functionality:

- [ ] [Game Management](https://github.com/rosera/quizzrr_leaderboard/tree/main?tab=readme-ov-file#game-management)
- [ ] [Data Management](https://github.com/rosera/quizzrr_leaderboard/tree/main?tab=readme-ov-file#data-management)
- [ ] [User Management](https://github.com/rosera/quizzrr_leaderboard/tree/main?tab=readme-ov-file#user-management)


## Game Management

Game management performs interaction with a defined leaderboard.

###  LIST Leaderboard

The API supports listing gameID Leaderboard entries.

![List Leaderboard](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-leaderboard-empty.png "View Leaderboard")

The API uses the route signature `/games/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl -X GET http://localhost:8080/games/test
```

### ADD Leaderboard

The API supports adding gameID Leaderboard point entries.

![Add Leaderboard](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-leaderboard-add.png "Add to Leaderboard")

The API uses the route signature `/points/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl -X POST http://localhost:8080/points/test \
   -H "Content-Type: application/json" \
   -d '{ "game": "test", "name": "Rich", "score": 1000 }'
```

__EXPECTED OUTPUT__
```json
[{"game":"test","name":"Rich","score":1000}]
```

### DELETE Leaderboard

The API supports deleting gameID Leaderboard.

![Delete Leaderboard](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-leaderboard-delete.png "Delete Leaderboard")

The API uses the route signature `/cancels/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl -X DELETE http://localhost:8080/cancels/test
```

### TEST Leaderboard

```bash
#!/bin/bash

if [[ -z "$1" || -z "$2" ]]; then
  echo "Usage: `basename "$0"` <GAME> <ENDPOINT>"
  echo "  Example: `basename "$0"` 'game1' 'https://api.example.com/game1'"
  exit 1
fi

GAME="$1"
ENDPOINT="$2"

# Set the number of times to run the loop
NUM_ITERATIONS=5

for ((i=1; i<=NUM_ITERATIONS; i++)); do
  curl -X POST "$ENDPOINT"/points/"$GAME" -H "Content-Type: application/json" -d '{ "game": "$GAME", "name": "tester-$i", "score": 1000 }'

  sleep 5
done
```

## Data Management

Data management performs import/export for a defined leaderboard.

### Import Leaderboard

Enable import of leaderboard from a JSON file.

```
curl -X POST -H "Content-Type: application/json" --data-binary @leaderboards.json http://localhost:8080/import
```


### Export Leaderboard

Enable export of leaderboard to a JSON file.

```
curl http://localhost:8080/export > leaderboards.json
```

__EXPECTED OUTPUT__
```json
{"test":[{"game":"test","name":"Abby","score":1000},{"game":"test","name":"Bobby","score":1000},{"game":"test","name":"Catherine","score":1000},{"game":"test","name":"Derick","score":1000},{"game":"test","name":"Ernest","score":1000},{"game":"test","name":"Fred","score":1000},{"game":"test","name":"Gisele","score":1000},{"game":"test","name":"Harold","score":1000},{"game":"test","name":"Ivonne","score":1000},{"game":"test","name":"Jack","score":1000},{"game":"test","name":"Kelly","score":1000}]
```


## User Management

User management performs deletion of a specified user within a defined leaderboard.

### Delete User

The API supports deleting a user from the Leaderboard.

```
curl -X DELETE http://localhost:8080/remove/test/Rich
```
