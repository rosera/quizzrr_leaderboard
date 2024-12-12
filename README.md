# Leaderboard

A simple (in-memory) leaderboard implemented in:

- [ ] Node.js
- [ ] Go

The Leaderboard APIs are functionally equivalent.
JSON format includes :

- [ ] game: game name
- [ ] name: user name
- [ ] score: user score

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

## LIST Leaderboard

```bash
curl https://localhost:8080/leaderboard?game=test
```

## ADD Leaderboard

```bash
curl -X POST http://localhost:8080/leaderboard \ 
   -H "Content-Type: application/json" \ 
   -d '{ "game": "test", "name": "Rich", "score": 1000 }'
```

## DELETE Leaderboard

```bash
curl -X DELETE http://localhost:8080/leaderboard/?game=test
```
