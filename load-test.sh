#!/bin/bash

if [[ -z "$1" || -z "$2" ]]; then
  echo "Usage: `basename "$0"` <GAME> <ENDPOINT>"
  echo "  Example: `basename "$0"` 'game1' 'https://api.example.com/game1'"
  exit 1
fi

GAME="$1"
ENDPOINT="$2"

# Set the number of times to run the loop
NUM_ITERATIONS=1000

for ((i=10; i<=NUM_ITERATIONS; i++)); do
  curl -X POST "$ENDPOINT"/points/"$GAME" -H "Content-Type: application/json" -d "{ \"game\": \"$GAME\", \"name\": \"tester-$i\", \"score\": 1000 }"

  sleep 5
done
