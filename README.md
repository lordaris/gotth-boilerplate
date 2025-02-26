# GOTTH Stack Boilerplate

A minimal yet powerful web application boilerplate using the GOTTH (golang, templ, tailwind, HTMX) stack:

- **Go**
- **Chi** Lightweight and idiomatic router for Go
- **Templ** - Type-safe templates for Go
- **Tailwind CSS** - CSS framework
- **HTMX** - Modern frontend interactivity without JavaScript frameworks
- **PostgreSQL + SQLx** - Database and SQL handling
- **Air** - Live reload for development

## Requirements

Ensure you have the following installed before starting:

- **Go 1.20+**
- **Node.js and npm**
- **PostgreSQL**

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/lordaris/gotth-boilerplate
   cd gotth-boilerplate
   ```

2. **Install Go dependencies**

   ```bash
   go mod tidy
   ```

3. **Install Node dependencies**

   ```bash
   npm install
   ```

4. **Configure the PostgreSQL database**
   - Set up your database and update the connection details in the environment configuration.
   - Create a .env file in the main folder and add the following content:

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gotth_boilerplate

# Server Configuration
PORT=8080
```

This .env file is essential for storing sensitive configuration values and should not be committed to version control.

5. **Run database migrations**

```bash
make db-migrate
```

6. **Start the development server**

```bash
   make dev
```

7. Visit `http://localhost:8080` in your browser.

## Makefile Commands

This project includes a `Makefile` to simplify common operations, just remove the "copy" text from the file name and modify the project and database variables to your own:

```bash
# Install required tools
make install-tools

# Start development environment (Air + Tailwind)
make dev

# Database operations
make db-create     # Create the database
make db-migrate    # Run migrations
make db-reset      # Reset the database

# Generate Templ files
make generate-templ

# See all available commands
make help
```

## Project Structure & Philosophy

This boilerplate follows a kind of **clean architecture** approach, designed to be flexible and scalable for small to medium projects. The backend structure is consistent with my preferred approach, making it easy to navigate and extend.

The application includes basic templates, handlers, and HTMX examples. While the HTMX implementation is still a work in progress ( could and should be improved), it serves as a starting point for building dynamic interactions without excessive JavaScript.

Just head to the main page and experiment with the form interactions!

## Improvement Areas & Future Plans

- **Automating Templ generation**: Currently, you need to run `make generate-templ` every time you modify a `.templ` file or update a handler function. Ideally, this process should be automated.
- **Better HTMX usage**: As I'm still learning HTMX, the implementation isn't perfect. Expect improvements in the future.
- **General refinements**: UI, error handling, and overall experience will be enhanced over time.

For now, if you modify a `templ` file, run:

```bash
make generate-templ
```

and restart the server using:

```bash
make run  # or make dev
```

---

Feel free to improve on this template and make it your own. Happy coding!
