# Item Order System

This repository contains the source code for an Item Order System application.

## Requirements

- Go >= 1.2
- PostgreSQL >= 9.4
- Redis >= 3.0

## Getting Started

1. Clone the repository

```sh
git clone https://github.com/wellisrite/sr-skilltest.git
```

2. Install the dependencies

```sh
go mod tidy
go mod vendor
```

3. Set up the environment variables in a `.env` file in the root directory of the project

```
APP_PORT=
SQL_HOST= host.docker.internal // for local installation
SQL_PORT=
SQL_DATABASE=
SQL_USER=
SQL_PASSWORD=

REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=
```

4.Start the server

```sh
go run . -local
```

Docker 
```sh
docker-compose up
```

The server will be available at `http://localhost:8080`

## Contributing

We welcome contributions from the community. To contribute, please follow the following steps:

1. Fork the repository
2. Create a new branch with a descriptive name (e.g. `add-new-feature`)
3. Make your changes and commit them to your branch
4. Push your changes to your fork
5. Create a pull request from your fork to the `master` branch of this repository

## License

This repository is licensed under the MIT License. See the `LICENSE` file for more information.
