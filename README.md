# Go User Management API

A simple REST API built with Go for managing user data. This project demonstrates basic Go web development concepts including database connections, HTTP handlers, and API response formatting.

## Installation & Setup

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
   - Create a MySQL database for the project
   - Create the `users` table:
     ```sql
     CREATE TABLE users (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       phone VARCHAR(20),
       address TEXT
     );
     ```

4. **Configure database connection**:
   - Update the database connection string in `db/connection.go` with your MySQL credentials

## API Endpoints

### Get All Users
- **Endpoint**: `GET /users`
- **Description**: Retrieve all users from the database
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "phone": "08123456789",
      "address": "Jakarta, Indonesia"
    }
  ]
  ```

### Create User
- **Endpoint**: `POST /users`
- **Description**: Create a new user
- **Request Body**:
  ```json
  {
    "name": "Jane Doe",
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