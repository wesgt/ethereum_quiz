# Ethereum Quiz
Implement two Ethereum blockchain services
- portto-app : API service
- portto-worker : Ethereum block indexer service

# Run the Applications
Here is the steps to run it with `docker-compose`

```bash
# Clone into YOUR $GOPATH/
$ git clone https://github.com/wesgt/ethereum_quiz.git

# move to project
$ cd ethereum_quiz

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:8080/blocks/16683936

# Stop
$ make stop
```