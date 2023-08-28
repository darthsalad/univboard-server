<div align="center">

# Univboard

<!-- badges -->

`Uniboard` is a cross-platform universal clipboard and file-sharing platform. Share text, URLs, files and more across all your devices with ease.

</div>


## Built with
- [Golang](https://golang.org/) - Backend server
  - [net/http](https://golang.org/pkg/net/http/) - HTTP server
  - [gorilla/mux](https://pkg.go.dev/github.com/gorilla/mux) - HTTP router
  - [gorilla/websocket](https://pkg.go.dev/github.com/gorilla/websocket) - Websocket server
  - [database/sql](https://pkg.go.dev/database/sql) - SQL database driver
- [Planetscale](https://planetscale.com/) - Remote MySQL database

## TODO
### General
- [x] Planetscale schema for fields:
	- created_at
	- modified_at
- [x] NewUser method integrate instead of decoding directly to struct
- [x] Errors thrown in api route functions should be thrown through JSON responses and shouldn't the server
- [x] Remove unnecessary log.Fatal calls causing server to crash - only log.Default calls and Fatal calls in [necessary places](https://stackoverflow.com/questions/33885235/should-a-go-package-ever-use-log-fatal-and-when#:~:text=73-,It,-might%20be%20just)
- [x] custom log file location of source file instead of logger.go
- [ ] Protected Routes middleware wrapper



### API
- [X] Auth
  - [x] Register
  - [x] Login
  - [x] Profile 
    - [x] GET - Get profile
    - [ ] PUT - Update profile *(Later)* [Profile Picture, Name, Email]
  - [x] Logout
  - [x] Delete
  - [ ] Change password *(Later)* 
  - [ ] Forgot password *(Later)*

- [ ] Device
  - [ ] GET - Get all devices
  - [ ] POST - Add new device
    - [ ] Device details like:
      - OS 
      - Version
      - Device Name
      - Device ID
  - [ ] DELETE - Delete device
	- Notes: New Device should be added on `Login` 

- [ ] Pushes (WebSocket server)
	- [ ] Send a message
  	- [ ] Specify TYPE
    	- [ ] Text/URL
    	- [ ] File/Blob
  	- [ ] Specify Device
    	- [ ] Default - Send to all devices(pool)
  	- [ ] Send to all devices
  	- [ ] Send to specific device
	- [ ] Notification
	- [ ] DB operations