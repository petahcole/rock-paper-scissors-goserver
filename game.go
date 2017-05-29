package main

import (
      "log"
      "net/http"

      "golang.org/x/net/websocket"
)

type Player1 struct {
  user          string
  choice        string
  startBalance       int
  newBalance   int
  bet           int
}

type Player2 struct {
  user          string
  choice        string
  startBalance       int
  newBalance   int
  bet           int
}

func compareForGame(p1 *Player1, p2 *Player2) string  {
  if p1.choice == p2.choice {
    return "draw"
  } if  p1.choice == "rock" && p2.choice == "paper" {
    return p2.user
  } if p1.choice == "rock" && p2.choice == "scissors" {
    return p1.user
  } if p1.choice == "paper" && p2.choice == "scissors"  {
    return p2.user
  }
}

func calculateBalances(p1 *Player1, p2 *Player2) string {
  if compareForGame() == p1.user  {
    p1.newBalance += p1.bet + p2.bet && p2.newBalance -= p2.bet
  } if compareForGame() == p2.user  {
    p2.newBalance += p1.bet + p2.bet && p1.newBalance -= p1.bet
  } if p1.bet > p1.newBalance | p2.bet > p2.newBalance {
    return "invalid bet"
  }
}
