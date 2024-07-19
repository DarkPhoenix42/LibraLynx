# LibraLynx

## Description

This project is a library management system twritten in go using the MVC architecture.

## Features

- Separate admin and client portals
  
### Admin can
- Manage the book catalog (list, update, add/remove books).
- Approve/deny borrow and return requests from clients.
- Approve/deny requests from users seeking admin privileges.

### Users can
- View the list of available books.
- Request borrowing and returning of books from the admin.
- View their transaction history.

## Setup
1. Clone the repository.
2. Run `mv .envsample .env` and modify the .env file as you wish.

### Using Docker
1. Set the DB_HOST variable in .env file to `mysqldb`.
1. Start the docker containers using `docker compose up` and connect at `localhost:8080`.

### Without Using Docker
1. Install golang, and golang migrate and setup mysql server on your device.
2. run `./run.sh`
3. Connect to the website at `libralynx.org`.

## Note
If you change DB_NAME in .env file then change init.sql accordingly.
The first user is made admin by default.