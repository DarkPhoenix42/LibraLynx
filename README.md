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
3. Start the docker containers using `docker compose up` and connect at `localhost:8080`.

## Note
The first user is made admin by default.