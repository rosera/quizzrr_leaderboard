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

The API uses the route signature `/games/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl http://localhost:8080/games/test
```

## ADD Leaderboard

The API uses the route signature `/points/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl -X POST http://localhost:8080/points/test \
   -H "Content-Type: application/json" \
   -d '{ "game": "quizzrr", "name": "Rich", "score": 1000 }'
[{"game":"quizzrr","name":"Rich","score":1000}]
```

## DELETE Leaderboard

The API uses the route signature `/cancels/{id}` where id is the gameID.

__EXAMPLE:__
```bash
curl -X DELETE http://localhost:8080/cancels/test
```
