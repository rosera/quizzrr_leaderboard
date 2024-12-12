const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

const app = express();
const port = 8080;

const Leaderboard = function(game) {
  this.game = game;
  this.data = [];
}

// Middleware
app.use(bodyParser.json());
app.use(cors());

// In-memory leaderboard storage
let leaderboards = {};

function getLeaderboard(game) {
  if (!leaderboards[game]) {
    leaderboards[game] = new Leaderboard(game);
  }
  return leaderboards[game];
}

app.get('/leaderboard', (req, res) => {
  const game = req.query.game;
  if (!game) {
    return res.status(400).json({ error: 'Game parameter is required' });
  }
  const leaderboard = getLeaderboard(game);
  res.json(leaderboard.data);
});

app.post('/leaderboard', (req, res) => {
  const newPlayer = req.body;
  const game = newPlayer.game;

  // Validation logic
  if (!game || !newPlayer.name || !newPlayer.score) {
    return res.status(400).json({ error: 'Game, Name and score are required' });
  }

  const leaderboard = getLeaderboard(game);
  leaderboard.data.push(newPlayer);
  leaderboard.data.sort((a, b) => b.score - a.score);
  res.json(leaderboard.data);
});

app.delete('/leaderboard', (req, res) => {
  const game = req.query.game;

  if (!game) {
    return res.status(400).json({ error: 'Game parameter is required' });
  }
  const leaderboard = getLeaderboard(game);
  leaderboard.data = [];

  res.json({ message: `Leaderboard for ${game} reset successfully` });
});

app.listen(port, () => {
    console.log(`Leaderboard API listening on port ${port}`);
});
