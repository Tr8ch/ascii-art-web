## Description

The "ASCII-ART-WEB" project aims to provide users with the ability to input a string of text and receive the corresponding ASCII art representation of that string on a web page. ASCII art is a form of art that uses characters from the ASCII character set to create visual representations of text or images. This project will enable users to interact with ASCII art in a user-friendly way through a web-based interface.


## Usage to run
#### to run the web server with defaul PORT=3000:
```bash
go run ./cmd/web
```
#### to run the web server with your port:
```bash
export PORT=<YOUR PORT>
go run ./cmd/web
```
## You can also build and run docker
#### build docker:
```bash
docker build -t <TAG> .
```
#### run docker with defaul PORT=3000:
```bash
docer run <TAG>
```
#### run docker with your PORT="your port":
```bash
docer run -e PORT=<YOUR PORT> <TAG>
```

## Usage to docker run
#### build docker:
```bash
docker build -t ascii .
```
#### run docker with defaul PORT=3000:
```bash
docer run ascii
```
#### run docker with your PORT=7777:
```bash
docer run -e PORT=7777 ascii
```
You can also use -p tag for docker

### Go to the server on the port you've specified or on default port 3000/"your port"
```bash
localhost:3000
```


Once the project is accessed through the web browser, you should see a user interface that allows you to input a string of text. After entering the desired text, the web page should generate and display the corresponding ASCII art representation of the provided string.


## Tech
Go version 1.20


## Authors
[@Tr8ch](https://github.com/Tr8ch)
