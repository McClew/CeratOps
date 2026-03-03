# CeratOps

A web application built with the GOTH stack (Go, Templ, HTMX, Tailwind).

## 🚀 Getting Started

### Prerequisites
- [Go](https://go.dev/dl/)
- [Templ](https://templ.guide/quick-start/installation)
- [Tailwind CSS CLI](https://tailwindcss.com/docs/installation)

### Running the Application

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Generate Templ components:
   ```bash
   templ generate
   ```

3. Build Tailwind CSS:
   ```bash
   tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
   ```

4. Run the server:
   ```bash
   go run main.go
   ```
