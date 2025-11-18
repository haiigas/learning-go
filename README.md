# Go Management API

A simple REST API built with Go for managing user data. This project demonstrates basic Go web development concepts including database connections, HTTP handlers, and API response formatting.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/haiigas/learning-go.git
   cd learning-go
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up the database**:
   
   Create the `users` and `biodatas` tables:
   ```sql
   CREATE TABLE users (
     id INT AUTO_INCREMENT PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     email VARCHAR(255) UNIQUE NOT NULL,
     password VARCHAR(80) NOT NULL,
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     deleted_at DATETIME NULL
   );
   
   CREATE TABLE biodatas (
     id INT AUTO_INCREMENT PRIMARY KEY,
     user_id INT NOT NULL UNIQUE,
     phone VARCHAR(30),
     address TEXT,
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     deleted_at DATETIME NULL,
     FOREIGN KEY (user_id) REFERENCES users(id)
   );
   ```

4. **Configure database connection**:
   
   Update the database connection string in `db/connection.go` with your MySQL credentials

## API Endpoints

### Get All Users
- **Endpoint**: `GET /v1/users`
- **Description**: Retrieve all users from the database
- **Response**:
  ```json
  {
    "status": true,
    "message": "fetch all users",
    "data": [
      {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "phone": "08123456789",
        "address": "Jakarta, Indonesia"
      }
    ]
  }
  ```

### Create User
- **Endpoint**: `POST /v1/users`
- **Description**: Create a new user
- **Request Body**:
  ```json
  {
    "name": "Jane Doe",
    "email": "jane@example.com",
    "password": "secure",
    "phone": "08198765432",
    "address": "Bandung, Indonesia"
  }
  ```
- **Response**: Returns the created user with ID

## Running with Auto-reload

The project includes configuration for `air` (Go live-reload tool). Run:
```bash
air
```

This will automatically rebuild and restart the application when you make changes.

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.