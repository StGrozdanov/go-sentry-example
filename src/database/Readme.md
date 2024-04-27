### How to start the project

1. Get your sentry DSN and add it in the config.env
2. Run the following commands in your terminal: ``docker-compose up -d go-sentry`` then `docker logs --follow go-sentry`
3. The server should start on localhost:8080 where you will have two endpoints - ``GET /course`` and `GET /frequent-questions` to work with

### Where to look for

the attempt to create a query span is located in the `crud.go` file. You can find different versions of my approach there ..

i see that out of the box my transactions get tracked from gin cdk and reflected correctly on the performance tab. Problems are when i try to go custom mode with them.