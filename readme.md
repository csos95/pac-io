# Pac IO

This is a repository for a possible future project.  
The project is a multiplayer "io" game that builds on pacman.

## Basic Planning

client
- send movements
- receive entity positions

server
- send entity positions
- receive player movements
- calculate item pickups/damage dealt
- move enemies
- store user accounts (gorm/sqlite)
	
connection
- use websockets to send/receive data (gorilla)
	
levels
- Hard coded level(s) for initial testing, procedurally/user generated levels later
	
movement
- wasd - move
- smooth movement
- client sends movement
- server sends new position (player requests to move, server has final say)

gameplay
- drop in/out gameplay
- modes:
  - co-op: All players connect as Pac-man with the goal of clearing the stage against AI ghosts
  - FFA: all players connect as Pac-man with the goal of having the highest score when the level is cleared. Players that run out of lives before level clear can respawn as ghosts.
  - vs: 4 players connect as ghosts for every 1 that connects as Pac-man (ghosts goal?)
	
## Project Layout
The root directory contains Go source files that will make up the server.  
The `site` folder contains the webpage and associated assets (html, css, js, etc).
