# Blog Backend API

## Description

This is a simple backend application for a blog website built with Go using the Gin framework. The API allows for user authentication and CRUD operations on blog posts.

## Technologies Used

- **Go**: Programming language used for the backend.
- **Gin**: Web framework for building the API.
- **GORM**: ORM library for Go, used for database interactions.
- **JWT**: Used for user authentication and securing routes.

## Installation

### Prerequisites

- Go installed on your machine
- A PostgreSQL database (or any other database supported by GORM)

### Steps to Run the Project

1. Clone the repository:
   ```bash
   git clone https://github.com/khachatryanhovhannes/blog_backend.git
   ```
2. Navigate to the project directory:

   ```bash
   cd blog_backend

   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Set up your database and configure environment variables (see .env.example for reference).
5. Run the application:

   ```bash
   go run main.go

   ```

   ## API Endpoints

### Authentication

- `POST /login`: Login endpoint for admin. Requires a JSON body with `username` and `password`.
- `POST /register`: Register a new user. Requires a JSON body with `username` and `password`.

### Blog Posts

- `GET /posts`: Retrieve all blog posts.
- `POST /posts`: Create a new blog post. Requires a JSON body with `title` and `content`.
- `GET /posts/:id`: Retrieve a single blog post by ID.
- `PUT /posts/:id`: Update a blog post by ID. Requires a JSON body with `title` and `content`.
- `DELETE /posts/:id`: Delete a blog post by ID.

## Usage Examples

### Login

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "admin", "password": "password"}'
```

### Get All Posts

```bash
curl -X GET http://localhost:8080/posts
```
