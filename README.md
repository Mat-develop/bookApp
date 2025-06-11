# Book Register App

This is a simple CRUD application for managing books using Go and MySQL. The application allows users to register, retrieve, update, and delete book records.

## Project Structure

```
book-register-app
├── src
│   ├── main.go                # Entry point of the application
│   ├── controllers            # Contains the book controller for handling requests
│   │   └── book_controller.go
│   ├── models                 # Contains the book model
│   │   └── book.go
│   ├── routes                 # Contains the route definitions
│   │   └── routes.go
│   └── database               # Contains database connection logic
│       └── connection.go
├── go.mod                     # Go module definition
└── README.md                  # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd book-register-app
   ```

2. **Install dependencies:**
   Make sure you have Go installed. Run the following command to install the necessary dependencies:
   ```
   go mod tidy
   ```

3. **Set up the MySQL database:**
   Create a MySQL database and update the connection details in `src/database/connection.go`.

4. **Run the application:**
   ```
   go run src/main.go
   ```

5. **API Usage:**
   You can use Postman or any other API client to interact with the application. Below are the available endpoints:

   - **Create a Book**
     - **POST** `/books`
     - **Request Body:**
       ```json
       {
         "title": "Book Title",
         "author": "Author Name",
         "published_year": 2023,
         "image_url": "example.png"
       }
       ```

   - **Get All Books**
     - **GET** `/books`

   - **Get a Book by ID**
     - **GET** `/books/{id}`

   - **Update a Book**
     - **PUT** `/books/{id}`
     - **Request Body:**
       ```json
       {
         "title": "Updated Title",
         "author": "Updated Author",
         "published_year": 2024,
         "image_url": "example.png"
       }
       ```

   - **Delete a Book**
     - **DELETE** `/books/{id}`

## License

This project is licensed under the MIT License.
