# BeerCrawl backend service

The BeerCrawl backend service provides data from the DB and forms the best route along the beer factories given a starting point (the weight is based on distance / beer count)

# How to run

The simple development server can be run with the command ```go run main.go serve```, while the production should use the ```beer-docker``` repository.

# How to test

Run the command: ```go test ./...```
