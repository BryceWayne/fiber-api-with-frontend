# fiber-api-with-frontend
Go Fiber API with HTML Templates and HTMX

A modular web application built with Go Fiber framework that demonstrates modern web development practices including modular architecture, HTML templating, HTMX integration, and RESTful API endpoints.

## Features

- **Go Fiber Framework**: Fast and lightweight web framework for Go
- **Modular Architecture**: Clean separation of concerns with dedicated packages
- **HTML Template Engine**: Server-side rendering with Go's HTML templates
- **HTMX Integration**: Dynamic content loading without full page reloads
- **Static File Serving**: CSS, JavaScript and other static assets
- **RESTful API Endpoints**: JSON API alongside HTML pages
- **Modern UI**: Clean, professional design with responsive layout

## Project Structure

```
.
├── main.go              # Main application entry point
├── config/              # Configuration package
│   └── config.go        # App configuration and setup
├── handlers/            # HTTP handlers
│   ├── page_handlers.go # Page rendering handlers
│   └── api_handlers.go  # API endpoint handlers
├── routes/              # Route definitions
│   └── routes.go        # All application routes
├── views/               # HTML templates
│   ├── index.html       # Home page template
│   └── about.html       # About page template
├── static/              # Static assets
│   ├── css/
│   │   └── style.css    # Stylesheet
│   └── js/
│       └── htmx.min.js  # HTMX library
├── go.mod               # Go module file
└── go.sum               # Go dependencies
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/BryceWayne/fiber-api-with-frontend.git
cd fiber-api-with-frontend
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

1. Build the application:
```bash
go build -o fiber-app
```

2. Run the application:
```bash
./fiber-app
```

The server will start on `http://localhost:3000`

## Available Routes

### Page Routes
- **GET /** - Home page with HTML template and HTMX demo
- **GET /about** - About page with feature list

### API Routes
- **GET /api/hello** - JSON API endpoint
- **GET /api/hello-html** - HTML fragment endpoint for HTMX

## Architecture

### Modular Design
The application follows a modular architecture pattern:

- **config**: Handles application configuration and Fiber app setup
- **handlers**: Contains HTTP request handlers organized by functionality
- **routes**: Centralized route definitions and setup

This structure makes the codebase more maintainable and scalable.

### HTMX Integration
The frontend uses HTMX for dynamic content updates without full page reloads. The API endpoint returns HTML fragments that are swapped into the page, providing a smooth user experience.

## Screenshots

### Home Page
![Home Page](https://github.com/user-attachments/assets/4c1cf2b9-76e0-40c8-8525-dfb938087496)

### HTMX in Action
![HTMX Demo](https://github.com/user-attachments/assets/ce38770c-6336-4126-a5ca-7d4ae954b49a)

### About Page
![About Page](https://github.com/user-attachments/assets/ad99f26c-913c-4f74-a7db-8a776e03914d)

## Development

To run in development mode with auto-reload, you can use tools like `air`:

```bash
go install github.com/cosmtrek/air@latest
air
```

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- [Fiber Template HTML](https://github.com/gofiber/template) - HTML template engine

## License

MIT
