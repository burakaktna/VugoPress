# VugoPress

VugoPress is a robust project developed using the Go programming language combined with the GoFiber framework, serving as a practical example of a Golang Fiber JWT implementation. This project encompasses a comprehensive GoFiber sample REST API, seamlessly integrated with PostgreSQL databases, demonstrating effective Golang PostgreSQL utilization. Through its well-structured GoFiber project structure, VugoPress exemplifies best practices in Golang project examples, ensuring maintainability and scalability.

## Features

- **Authentication:** Implements JWT-based authentication, showcasing a Golang Fiber JWT example for secure user management.
- **REST API:** Provides a Golang Fiber sample REST API example, enabling efficient handling of HTTP requests and responses.
- **Database Integration:** Utilizes PostgreSQL with GoFiber, highlighting Golang PostgreSQL integration for reliable data storage and retrieval.
- **Validation:** Incorporates GoFiber validation examples to ensure data integrity and adherence to defined schemas.
- **Project Structure:** Demonstrates an organized GoFiber project structure, facilitating easy navigation and scalability of the codebase.

## Usage

The VugoPress API facilitates various CRUD operations. Below are the available endpoints and their descriptions:

### Users

- `POST /register`: Register a new user account.
- `POST /login`: Authenticate a user and provide a JWT token.

### Articles

- `GET /articles`: Retrieve a list of all articles.
- `POST /articles`: Create a new article.
- `GET /articles/:id`: Retrieve an article by its ID.
- `PUT /articles/:id`: Update an existing article by its ID.
- `DELETE /articles/:id`: Remove an article by its ID.

### Tags

- `GET /tags`: Retrieve all tags.
- `POST /tags`: Create a new tag.
- `GET /tags/:id`: Retrieve a tag by its ID.
- `PUT /tags/:id`: Update an existing tag by its ID.
- `DELETE /tags/:id`: Remove a tag by its ID.

### Useful Links

- `GET /useful_links`: Retrieve all useful links.
- `POST /useful_links`: Create a new useful link.
- `GET /useful_links/:id`: Retrieve a useful link by its ID.
- `PUT /useful_links/:id`: Update an existing useful link by its ID.
- `DELETE /useful_links/:id`: Remove a useful link by its ID.

### Social Media Links

- `GET /social_media_links`: Retrieve all social media links.
- `POST /social_media_links`: Create a new social media link.
- `GET /social_media_links/:id`: Retrieve a social media link by its ID.
- `PUT /social_media_links/:id`: Update an existing social media link by its ID.
- `DELETE /social_media_links/:id`: Remove a social media link by its ID.

### Contacts

- `GET /contacts`: Retrieve all contact entries.
- `POST /contacts`: Create a new contact entry.
- `GET /contacts/:id`: Retrieve a contact entry by its ID.
- `PUT /contacts/:id`: Update an existing contact entry by its ID.
- `DELETE /contacts/:id`: Remove a contact entry by its ID.

The API accepts and returns JSON data, ensuring seamless integration with various client applications. VugoPress demonstrates practical implementations of RESTful API development with GoFiber, incorporating JWT authentication, PostgreSQL database management, input validation, and a scalable project structure, making it an exemplary Golang project example.
