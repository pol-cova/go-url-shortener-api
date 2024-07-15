```markdown
# Go URL Shortener API Documentation

## Overview
The Go URL Shortener API allows users to create shortened URLs for easy sharing and management. It includes endpoints for user authentication, URL shortening, and user profile management, using JWT for authentication.

## Base URL
`http://localhost:8080`

## Authentication
Include the JWT token in the `Authorization` header for protected endpoints.
`Authorization: <your-token>`

## Endpoints

### Auth Routes

#### Signup

- **Endpoint:** `/auth/signup`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```

#### Login

- **Endpoint:** `/auth/login`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```
- **Response:**
  ```json
  {
    "token": "your-jwt-token"
  }
  ```

#### Logout

- **Endpoint:** `/auth/logout`
- **Method:** `GET`

### URL Routes

#### Shorten URL

- **Endpoint:** `/short`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "url": "string"
  }
  ```
- **Response:**
  ```json
  {
    "shortened_url": "http://localhost:8080/:key"
  }
  ```

#### Redirect URL

- **Endpoint:** `/:key`
- **Method:** `GET`

### User Routes

> These routes require authentication.

#### Shorten URL (Authenticated User)

- **Endpoint:** `/user/short`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "url": "string"
  }
  ```

#### Home

- **Endpoint:** `/user/home`
- **Method:** `GET`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "url": "https://example.com",
      "key": "abc123",
      "created_at": "2023-01-01T00:00:00Z",
      "clicks": 100,
      "user_id": 1
    },
    {
      "id": 2,
      "url": "https://anotherexample.com",
      "key": "xyz789",
      "created_at": "2023-01-02T00:00:00Z",
      "clicks": 50,
      "user_id": 1
    }
  ]

#### Profile

- **Endpoint:** `/user/me`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "email": "string"
  }
  ```

#### Delete Account

- **Endpoint:** `/user/delete`
- **Method:** `DELETE`

## Running the Project

1. Clone the repository:
   ```sh
   git clone https://github.com/pol-cova/go-url-shortener-api.git
   ```
2. Navigate to the project directory:
   ```sh
   cd <go-url-shortener-api>
   ```
3. Run the project:
   ```sh
   go run .
   ```

The server will start on port 8080.
```