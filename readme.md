# User Service 🚀

A user management service built in Go to provide authentication and authorization functionality for applications. 🔐


## ✨ Features

- ✅ Create users (`/users`)
- ✅ Validate email uniqueness
- ✅ Integrate with authentication service
- ✅ Manage user status
- ✅ Get all admins (`/admin-users`)
- ✅ Docker and Makefile for easy setup

---

## 🛍 Available Routes

### `POST /users`
Create a new user.

**Example Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role": "user"
}
```

**Response:**
``` json
{
  "id": "12345678-1234-1234-1234-123456789012",
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "status": true,
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6...",
  "created_at": "2025-04-08T13:33:19-03:00",
  "updated_at": "2025-04-08T13:33:19-03:00"
}
```

### 🛠️ Installation

**Requirements**
- Go 1.20 or higher
- Docker
- Make (optional, for ease of use)

### 🐳 Running MariaDB on Docker

```bash
docker-compose up --build -d
```

### 📝 Makefile Commands
Useful commands for development:

```bash
make deps      # Install dependencies
make run       # Run the app locally
make build     # Build the project
make clean     # Clean build files
```

### ⚙️ Environment Variables
Create a .env file with the following content:

``` env
PORT=8081
DB_HOST=localhost
DB_PORT=3306
DB_NAME=users_db
DB_USER=root
DB_PASSWORD=root
AUTH_SERVICE_URL=http://localhost:8080
```

## 📌 Future Roadmap
- [ ] Get normal users
- [ ] Update user
- [ ] Soft delete user

## Developed by
Michael Araujo Rodrigues — @michaelrodriguess