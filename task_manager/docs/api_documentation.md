# Task Management API Documentation

## Endpoints

### GET /tasks
- **Description:** Get all tasks.
- **Response:** `200 OK`
```json
[
  {
    "id": 1,
    "title": "Sample Task",
    "description": "Description",
    "due_date": "2025-07-21",
    "status": "pending"
  }
]
```

### GET /tasks/:id
- **Description:** Get a specific task.
- **Response:** `200 OK` or `404 Not Found`

### POST /tasks
- **Description:** Create a new task.
- **Request:**
```json
{
  "title": "New Task",
  "description": "Details",
  "due_date": "2025-07-21",
  "status": "pending"
}
```
- **Response:** `201 Created`

### PUT /tasks/:id
- **Description:** Update a task.
- **Request:** Same as POST.
- **Response:** `200 OK` or `404 Not Found`

### DELETE /tasks/:id
- **Description:** Delete a task.
- **Response:** `204 No Content` or `404 Not Found`