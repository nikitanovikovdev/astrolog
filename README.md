# Astrolog service

## Getting started

1. Create copy from .env.dist to .env for local development and .env.docker.dist to .env.docker for configuring docker containers
2. Go to https://api.nasa.gov/ , generate your own api_key and set it into APOD_API_KEY environment variable
3. Use `make` commands for building, launching and suspending a docker containers, migrations will automatically start

## API documentation 
    http://localhost:5632/ping - return `pong`, simple healthcheck
    http://localhost:5632/content/by-date?date=2023-11-29 - return content for a specific date
    http://localhost:5632/content/all - return all content

## Potential improvements

1. Add unit and integration tests
2. Add pagination for getting all content. For example by date
3. Upgrade logger. For example github.com/sirupsen/logrus
4. Standardize server responses
5. Add swagger documentation