# Onion Architecture Kata

A kata to practice Onion Architecture by building a simple microservice.

## Domain

In this kata we build a backend system (no frontend atm) to power a “virtual table-top” application.
The virtual table-top application makes it possible for players to remotely play a games together by
providing an infrastructure to manage those things usually put on a real table.
The system is designed around role-playing games (sometimes called narrative games) and is oriented
by something like the Fate Core system. No knowledge of Fate Core or other RPGs is required, though.
The following list describes the use-cases to be implemented:

1. One user can start a new session. He will become the game master of the session.
2. The game master can produce invitations and send them to the other users.
3. Other users can join a session upon invitation; they will be a player in that session.
   1. Once a player joins, every other player in the session receives a notification about the joining
4. For each player as well as for the game master there is an amount of game points (or Fate points)
   1. The game master can freely adjust the game points for himself as well as any other player
      (he is the master :-) ).
   2. Players _may spent_ game points given their balance is greater then zero.
5. Any time a player spends a game point, the game master should receive a notification
6. Any time the game master modifies a player’s game points balance, that player should receive a
   notification
7. If the game master adjusts the game master’s balance, a notification should only be sent to
   himself

For simplicity the system the following constraints can be applied:

- The system uses no authentication. Everyone can create sessions or join them if they know about
  the invitation link. The invitation link should contain some kind of random part that makes it
  hard for others to guess and thus hijack a session.
- The system can be designed to run as a “highlander” instance (only a single replica).
- All state can be kept in memory; no disk or other database-based persistence is required (we
  might add that later, though).
- The system should offer a simple HTTP-based API (RESTful may be a good idea but that is not
  strictly required).

## Optional extensions

The system is designed to be relatively simple from the functional point of view; this kata is to practice
software architecture and not about building VTTs. Feel free to extend the above lists by any means.
The following can give you some inspiration:

- **Dice Rolls**: Players can make a “dice roll”. A dice roll should be described by some kind of input
  language,
  1. to keep things simple, use the dice system from Fate Core, which is rolling Fudge Dice and
     adding a static modified; an input would be the modifier, i.e. +2
  2. to implement something more complex, implement something used in the D20 system,
     which is - specifying the number of die/dice to roll - specifying the number of sides the die/dice should have - specifying an optional modifier to apply. - example: 3d8+4 means: Roll 3 eight-sided dice, add their results and add 4 to get the
     final result.
  3. The system executes the dice roll
  4. The player making the dice roll as well as the game master receive a notification stating
  - what roll was requested (see below)
  - what the result of the dice roll was
  - No other player will receive that notification
- **Table Management**:
  1. Players can leave the session.
     - All other players including the game master receive a notification
  2. The game master can close the session (kicking all other players)
  3. The game master can kick an idividual player

## Exercise

1. Come up with modelling of the domain. Use some kind of diagram language (i.e. D2, PlantUML
   or Mermaid) to create a visual representation of the domain entities, the aggregates and their
   relationships.
2. Create software that implements the requirements applying a domain centric architecture pattern, i.e.:
   - Onion Architecture
   - Clean Architecture
   - Ports & Adapters Architecture

Some additional points:

- Focus on simplicity; the idea is to practice separation of concerns and dependency inversion.
  The idea is not to build the most advanced VTT RESTful API ever built :-)
- Start with the very basic requirements and structure the system. If you feel like things are working,
  extend the requirements by implementing some of the optional features or come up with your
  own
