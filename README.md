# Kasir API

A simple RESTful API for managing products and categories, built with Go and PostgreSQL (Supabase compatible).

## Features
- CRUD for Products
- CRUD for Categories
- JSON API responses
- PostgreSQL/Supabase backend

## Getting Started

### Prerequisites
- Go 1.18+
- PostgreSQL or Supabase database

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/dwirez99/kasir-go.git
   cd kasir-api
   ```
2. Set up your environment variables in a `.env` file:
   ```env
   PORT=8080
   DB_CONN=postgres://user:password@host:port/dbname?sslmode=disable
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

## API Endpoints

### Products
- `GET    /api/product`         - List all products
- `POST   /api/product`         - Create a new product
- `GET    /api/product/{id}`    - Get a product by ID
- `PUT    /api/product/{id}`    - Update a product by ID
- `DELETE /api/product/{id}`    - Delete a product by ID

### Categories
- `GET    /api/category`        - List all categories
- `POST   /api/category`        - Create a new category
- `GET    /api/category/{id}`   - Get a category by ID
- `PUT    /api/category/{id}`   - Update a category by ID
- `DELETE /api/category/{id}`   - Delete a category by ID

## API Documentation

This project uses [Swagger (OpenAPI)](https://swagger.io/) for API documentation.

- To generate docs, run:
  ```sh
  swag init
  ```
- Swagger docs will be available in the `docs/` folder.
- You can serve the docs using [swaggo/http-swagger](https://github.com/swaggo/http-swagger).

## Database Schema

### Product
| Column     | Type        |
|------------|------------|
| id         | BIGINT      |
| name       | VARCHAR     |
| price      | BIGINT      |
| created_at | TIMESTAMPTZ |
| stock      | SMALLINT    |

### Category
| Column     | Type        |
|------------|------------|
| id         | BIGINT      |
| name       | VARCHAR     |
| description| VARCHAR     |
| created_at | TIMESTAMPTZ |

## License
MIT
