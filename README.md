# Clean-architecture-go-redis
*Implement album micro service.*
*This album service database implements redis which use [redigo library](https://github.com/gomodule/redigo)*

***

## Build
	make

***

## Run
	./bin/api

***

## Find album
```
curl "http://localhost:8080/v1/album/{id}"
```

***

## Create album
```
curl -X "POST" "http://localhost:8080/v1/album" -d "`{"title": "Billi jean", "artist": "Micheal jackson", "id": 1}`"
```

