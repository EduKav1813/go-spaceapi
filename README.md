# go-spaceapi

A simple webserver for [SpaceAPI](https://spaceapi.io/#schema), written in Go.

## How to use

1. Modify `spaceinfo.json` with information about your hackerspace.
2. Compile the program:
```sh
$ cd spaceapi
$ go build
```
3. Run the `main` (or `main.exe` on Windows) binary. 
It will host the server under 8080, and sending the GET request to `/get` 
endpoint will return the info from the file
