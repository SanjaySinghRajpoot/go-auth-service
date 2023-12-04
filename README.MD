# Backend Project for User Management

## Overview 🚀

Welcome to our robust backend service tailored to provide secure JWT authentication, admin login, refresh token functionality, user-friendly endpoints, and a scalable architecture.

## Getting Started 🛠️

1. **Environment Setup:**
   - Create a `.env` file and configure the following variables:
     ```env
     ADMIN_NAME=admin
     ADMIN_EMAIL=admin
     ADMIN_PASSWORD=admin
     ACCESS_TOKEN_SECRET=your_access_token_secret
     REFRESH_TOKEN_SECRET=your_refresh_token_secret
     ```
     Customize these variables based on your preferences. The `.env` file is crucial for setting up the root admin user.

2. **Database Setup:**
   - Utilizes a PostgreSQL Docker image for the database.
   - Execute `docker compose up` to launch the PostgreSQL image.
   - Customize defaults in the Docker Compose file if needed.

3. **Run the Project:**
   - Install required dependencies.
   - Execute `go run main.go` to start the project.

4. **API Testing:**
   - Import the provided Postman collection file for comprehensive API testing.

## Project Features 🌟

- **MVC Architecture:**
  - Well-organized Model-View-Controller architecture for clarity and maintainability.

- **JWT Implementation:**
  - Implements secure JSON Web Token authentication for enhanced security.

- **Gin Framework:**
  - Utilizes the Gin framework for rapid and efficient API development.

- **Configuration with .env:**
  - Safeguards sensitive information using a `.env` file for enhanced security.

- **Dockerized Setup:**
  - Docker file provided for a hassle-free and consistent setup.

- **Code Quality:**
  - Adheres to best coding practices and methods for improved maintainability.

- **Scalable Architecture:**
  - Designed with scalability in mind to accommodate future growth.

## Contributing 🤝

We welcome contributions and suggestions to enhance this project. Please refer to our [contribution guidelines](CONTRIBUTING.md).

## License 📄

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments 🙌

A big thank you to the open-source community for their valuable contributions.

Happy Coding! 🌐✨