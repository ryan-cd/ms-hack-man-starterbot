# Ms Hack Man Starterbot

This project is the official golang "starter bot" for the [Ms. Hack Man](https://booking.riddles.io/competitions/ms.-hack-man) competition, hosted by booking.com. It can be downloaded from the [how to play](https://booking.riddles.io/competitions/ms.-hack-man/how-to-play) page.

The starter bot is able to make valid (but random) moves around the board. It is used as a baseline point for other developers to copy the code, and implement a search algorithm on top of it. 

# Game rules
The rules are available [here](https://docs.riddles.io/ms-hack-man/rules). In short, the player must navigate a pacman like maze at the same time as another player. Both players must collect snippets, plant mines, and avoid enemies. 

# How it works
The game engine requests information and sends information to the player using stdin. The bot must reply to the engine with stdout, within the timeout period. [API reference](https://docs.riddles.io/ms-hack-man/api).

# Compiling
### Just the bot locally:
```bash
go build -o starterbot
./starterbot
```

### To run the bot with the engine locally (to run a full game):
See the instructions from the [engine's GitHub page](https://github.com/riddlesio/hack-man-2-engine).