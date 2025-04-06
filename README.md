# Leaderboard

A simple (in-memory) leaderboard implemented in:

- [ ] Node.js
- [x] Go

![API Architecture](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-http-api.png "Architecture")

The Leaderboard APIs are functionally equivalent.

An array is defined for each game dataset.

```json
[
  {
    "game": "test",
    "name": "John Doe",
    "score": 100
  }
]
```

JSON format includes :

| Field | Description |
|-------|-------------|
| game  | game identifier |
| name  | user identifier |
| score | points allocation |


## LIST Leaderboard

The API supports listing gameID Leaderboard entries.

![List Leaderboard](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-leaderboard-empty.png "View Leaderboard")

The API uses the route signature `/games/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl -X GET http://localhost:8080/games/test
```

## ADD Leaderboard

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

## DELETE Leaderboard

The API supports deleting gameID Leaderboard.

![Delete Leaderboard](https://github.com/rosera/quizzrr_leaderboard/blob/main/screenshots/glb-leaderboard-delete.png "Delete Leaderboard")

The API uses the route signature `/cancels/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl -X DELETE http://localhost:8080/cancels/test
```

## DELETE User

The API supports deleting a user from the Leaderboard.

```
curl -X DELETE http://localhost:8080/remove/test/Rich
```
