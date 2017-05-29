#Rock Paper Scissors Server

*Server listens for websocket open command, will receive text from inputs
*Server waits until another request is made (maybe checks to see if websocket is already open?)
**Performs game logic - who wins and what the wagers were
*Server sends data back through sockets
*Server logs interaction
**IPs, RPS choice, bet, winner
*Server closes websocket on close command
