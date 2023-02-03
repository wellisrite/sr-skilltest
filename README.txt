# Item Order System

This repository contains the source code for an Item Order System application.

## Requirements

- Go >= 1.2
- PostgreSQL >= 9.4
- Redis >= 3.0

## Getting Started

1. Clone the repository

git clone https://github.com/wellisrite/sr-skilltest.git


2. Install the dependencies

go mod tidy
go mod vendor


3. Set up the environment variables in a `.env` file in the root directory of the project

APP_PORT=
SQL_HOST=
SQL_PORT=
SQL_DATABASE=
SQL_USER=
SQL_PASSWORD=

REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=


4. Run the migration

CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  expired_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  full_name VARCHAR(255) NOT NULL,
  first_order TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE order_histories (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  order_item_id INTEGER NOT NULL,
  descriptions VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (order_item_id) REFERENCES order_items (id)
);



5. Start the server

go run . -local


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
