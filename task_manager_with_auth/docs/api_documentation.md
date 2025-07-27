# Task Management API with JWT Authentication & MongoDB

This API provides task management functionality with authentication and authorization using **JWT**. It includes **user roles** (admin and regular user) and **MongoDB** for data storage.

---

## **1. Setup Instructions**

### **1.1 Prerequisites**
- Go 1.20+ installed
- MongoDB running locally on `mongodb://localhost:27017`
- Dependencies installed:
  ```bash
  go mod tidy
1.2 Running the Application
Start MongoDB:

bash
Copy
Edit
sudo systemctl start mongod
Run the API:

bash
Copy
Edit
go run main.go
The server will run at:

arduino
Copy
Edit
http://localhost:8080
2. Authentication Flow
A new user registers with POST /register.

If the database is empty, the first user becomes admin.

The user logs in with POST /login.

A JWT token is returned.

The JWT token must be sent in the Authorization header for all protected endpoints:

makefile
Copy
Edit
Authorization: Bearer <your_token_here>
3. Endpoints
3.1 User Management
POST /register
Description: Create a new user.

Request Body:

json
Copy
Edit
{
  "username": "john",
  "password": "password123"
}
Response:

json
Copy
Edit
{
  "id": 1,
  "username": "john",
  "role": "admin"
}
POST /login
Description: Authenticate and get a JWT token.

Request Body:

json
Copy
Edit
{
  "username": "john",
  "password": "password123"
}
Response:

json
Copy
Edit
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5..."
}
PUT /admin/promote/:username
Description: Promote a user to admin.

Headers:
Authorization: Bearer <admin_token>

Response:

json
Copy
Edit
{
  "message": "User promoted to admin"
}
3.2 Task Management
GET /tasks
Description: Retrieve all tasks (accessible by all users).

Response:

json
Copy
Edit
[
  {
    "id": 1,
    "title": "Fix bug",
    "description": "Fix the login bug"
  }
]
GET /tasks/:id
Description: Retrieve a single task by ID.

Response:

json
Copy
Edit
{
  "id": 1,
  "title": "Fix bug",
  "description": "Fix the login bug"
}
POST /admin/tasks
Description: Create a new task (admin only).

Headers:
Authorization: Bearer <admin_token>

Request Body:

json
Copy
Edit
{
  "title": "New Feature",
  "description": "Implement dark mode"
}
Response:

json
Copy
Edit
{
  "id": 2,
  "title": "New Feature",
  "description": "Implement dark mode"
}
PUT /admin/tasks/:id
Description: Update an existing task (admin only).

Headers:
Authorization: Bearer <admin_token>

Request Body:

json
Copy
Edit
{
  "title": "Updated Feature",
  "description": "Update dark mode"
}
Response:

json
Copy
Edit
{
  "message": "Task updated"
}
DELETE /admin/tasks/:id
Description: Delete a task (admin only).

Headers:
Authorization: Bearer <admin_token>

Response:

json
Copy
Edit
{
  "message": "Task deleted"
}
4. Error Responses
401 Unauthorized:
Returned if JWT token is missing or invalid.

json
Copy
Edit
{ "error": "Invalid or expired token" }
403 Forbidden:
Returned if a non-admin user tries to access an admin route.

json
Copy
Edit
{ "error": "Admin access only" }
400 Bad Request:
Returned for invalid input data.

json
Copy
Edit
{ "error": "Invalid task ID" }
5. MongoDB Collections
The database is task_manager_db and includes:

users collection:

json
Copy
Edit
{
  "id": 1,
  "username": "john",
  "password": "<hashed_password>",
  "role": "admin"
}
tasks collection:

json
Copy
Edit
{
  "id": 1,
  "title": "Fix bug",
  "description": "Fix the login bug"
}
6. Testing With cURL
Register a user:

bash
Copy
Edit
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"1234"}'
Login:

bash
Copy
Edit
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"1234"}'
Get all tasks:

bash
Copy
Edit
curl -X GET http://localhost:8080/tasks
Create a task (admin only):

bash
Copy
Edit
curl -X POST http://localhost:8080/admin/tasks \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Fix Bug","description":"Fix login issue"}'

