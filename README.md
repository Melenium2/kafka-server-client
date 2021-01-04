# Kafka server client example
## Installation
### Docker
You need pre install kafka and zookeeper in docker.
Enter at deploy folder and run
```sh
$ docker-compose up -d
```
This command automatically installs kafka and zookeeper and starts them.
Kafka default expose port `<you docker ip>:9092`
### Go a client and server
For install all dependencies run
```sh
$ make get
```
Then build your binary
```sh
$ make build
```

## Run
For starting consumer start your binary
```sh
$ ./cmd/main --act consumer
```
For Windows:
```sh
$ ./cmd/main.exe --act consumer
```
For starting producer
```sh
$ ./cmd/main --act producer
```
For Windows:
```sh
$ ./cmd/main.exe --act producer
```
## Usage
For publish message at broker in producer cli simply write next command to cli
```sh
$ send###<your message here>
```
`send` this is event. So far, there is one event in the app.

Then in consumer cli you will see that message delivered.

