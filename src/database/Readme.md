### How to start the project

1. Get your sentry DSN and add it in the config.env
2. Run the following commands in your terminal: ``docker-compose up -d go-sentry`` then `docker logs --follow go-sentry`
3. The server should start on localhost:8080 where you will have two endpoints - ``GET /course`` and `GET /frequent-questions` to work with