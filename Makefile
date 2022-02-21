BINARY=${type}
test: 
	go test -v -cover ./...

build:
	go build -o portto-${BINARY} ${type}/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	docker-compose up --build -d

stop:
	docker-compose down