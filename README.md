# fiber-api-with-frontend
Go Fiber API with HTML Templates

A simple web application built with Go Fiber framework that demonstrates HTML templating, static file serving, and RESTful API endpoints.

## Features

- **Go Fiber Framework**: Fast and lightweight web framework for Go
- **HTML Template Engine**: Server-side rendering with Go's HTML templates
- **Static File Serving**: CSS and other static assets
- **RESTful API Endpoints**: JSON API alongside HTML pages
- **Modern UI**: Responsive design with gradient backgrounds

## Project Structure

```
.
├── main.go              # Main application file
├── views/               # HTML templates
│   ├── index.html       # Home page template
│   └── about.html       # About page template
├── static/              # Static assets
│   └── css/
│       └── style.css    # Stylesheet
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

- **GET /** - Home page with HTML template
- **GET /about** - About page with HTML template
- **GET /api/hello** - JSON API endpoint

## Screenshots

### Home Page
![Home Page](https://github.com/user-attachments/assets/7cfe93dc-7199-4d12-ad39-13831d45676e)

### API Demo
![API Demo](https://github.com/user-attachments/assets/70db98a7-b122-4c18-8030-d8eb68e79d8e)

### About Page
![About Page](https://github.com/user-attachments/assets/679cc612-bf3f-462a-93a8-7860a24834be)

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
