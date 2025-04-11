# Leaderboard

A simple (in-memory) leaderboard implemented in:

- [ ] Node.js
- [x] Go

![API Architecture](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-http-api.png "Architecture")

The leaderboard apis handle the following functionality:

- [ ] game management
- [ ] data management
- [ ] user management




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
