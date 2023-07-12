# Go Fiber JWT Authentication 

## Installation

1. Clone the repository:
  ```shell
  git clone https://github.com/yusuftalhaklc/go-jwt-auth.git
  ```

2. Navigate to the project directory:
  ```shell
  cd go-fiber-authentication
  ```
3. Install the dependencies:
  ```shell
  go mod tidy
  ```
4. Start the API server:
  ```shell
  go run main.go
  ```

**It is currently running on localhost port 8080.** 

## API Endpoints
### Login 

- **Endpoint:** `/api/v1/login`
- **Method:** `POST`
- **Description:** Login.

### Protected Area 

- **Endpoint:** `/api/v1/protected`
- **Method:** `GET`
- **Description:** Access the protected area.


## Request Body and Response Examples

### Login
- Request Body
```json
{
    "username": "admin",
    "password": "1234"
}
```
- Response
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODkyMjYzNDMsInVzZXJuYW1lIjoiYWRtaW4ifQ.ltO-j9OAeJcDJetYGOkEJKmvYDF8FGmWXsJ7s4RkCck",
    "token_type": "bearer"
}
```

### Protected
- Request <br>
 Authorization: Bearer <access_token>
```http
GET /api/v1/protected
```

- Response
```json
{
    "message": "Welcome to the protected area",
    "status": 200
}
```
