[![CI](https://github.com/samdouble/biblio-api/actions/workflows/checks.yml/badge.svg)](https://github.com/samdouble/biblio-api/actions/workflows/checks.yml)

[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](https://go.dev/)
[![Google Cloud](https://img.shields.io/badge/Google%20Cloud-%234285F4.svg?logo=google-cloud&logoColor=white)](https://cloud.google.com/)
[![DigitalOcean](https://img.shields.io/badge/DigitalOcean-%230167ff.svg?logo=digitalOcean&logoColor=white)](https://www.digitalocean.com/)

# biblio-api

## Development

Create a `.env` file with the following variables:

```
GOOGLE_BOOKS_API_TOKEN=
MONGO_DBNAME=
MONGO_URL=
```

## Docker

Instantiate the MongoDB replica set:

```sh
docker compose up -d
```

Build the Docker image:

```sh
docker build -t biblio-api .
```

Run the Docker container with the book's ISBN as a command line argument:

```sh
docker run --env-file .env --network biblio-api_default -e "MONGO_URL=mongodb://biblio-api-mongo0:27017,biblio-api-mongo1:27017,biblio-api-mongo2:27017/?replicaSet=rs0" biblio-api <isbn>
```
