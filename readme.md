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
- randomly generated or hard coded?
	
movement
- wasd - move
- smooth movement
- client sends movement
- server sends new position (player requests to move, server has final say)
	
## Project Layout
The root directory contains Go source files that will make up the server.  
The `site` folder contains the webpage and associated assets (html, css, js, etc).
