# Sales Analytics Web Service

A high-performance REST API built with Go and Gin framework that provides comprehensive sales analytics and business intelligence endpoints. The service supports both PostgreSQL and DuckDB databases.


## Setup Instructions

### Prerequisites

- Go 1.21 or higher
- DuckDb PostgreSQL (optional)
- Git

### Get Gin from Repo
```bash
    go get -u github.com/gin-gonic/gin
```

### Get Postgres from Repo (Optional)
```bash
    go get -u github.com/lib/pq
```

## Get DuckDb from Repo

### Installation

1. **Clone the repository**
   ```bash
   git clone git@github.com:JamiuAfolabi/web-gin.git
   cd web-service-gin
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Database Setup**

   **Option A: DuckDB (Default)**
   - No additional setup required
   - Database file `sales_data.duckdb` will be created automatically

   **Option B: PostgreSQL**
   - Install PostgreSQL
   - Create a database named `postgres`
   - Update connection details in `main.go`:
     ```go
     dbConfig.Host = "your-host"
     dbConfig.Port = 5432
     dbConfig.Username = "your-username"
     dbConfig.Password = "your-password"
     ```

     - Create Migration Using Soda
     ```soda migrate```


4. **Run the application**
   ```bash
   go run *.go
   ```

The server will start on `http://localhost:8080`

## API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

#### Basic Data
- **GET** `/getproducts`
  - Returns all products in the catalog
  - Response: `[{id, name, category_id, price}]`

#### Financial Analytics
- **GET** `/finance/totalsalesbycat`
  - Total sales aggregated by category
  - Response: `[{category, total_sales}]`

- **GET** `/finance/revenuebymonth`
  - Monthly revenue trends over time
  - Response: `[{month, revenue}]`

- **GET** `/finance/performingproducts`
  - Top 5 revenue-generating products
  - Response: `[{product, total_revenue}]`

- **GET** `/finance/salesperformance`
  - Overall sales KPIs and metrics
  - Response: `[{total_orders, total_items_sold, total_revenue, avg_order_value}]`

#### Sales Analytics
- **GET** `/sales/salesdistributionbycat`
  - Market share distribution by category
  - Response: `[{category, category_sales, percentage}]`

- **GET** `/sales/productperformancebycat`
  - Product performance within each category
  - Response: `[{category, product, revenue}]`

- **GET** `/sales/topcustomerpurchase`
  - Top 20 customer purchase patterns
  - Response: `[{name, email, orders, total_spent}]`



## Architecture Overview


```
┌─────────────────┐    ┌─────────────────┐
│   HTTP Handlers │────│   Repository    │
│   (Gin Routes)  │    │   Interface     │
└─────────────────┘    └─────────────────┘
                              │
                       ┌─────────────────┐
                       │  Database Repo  │
                       │ Implementation  │
                       └─────────────────┘
                              │
                       ┌─────────────────┐
                       │  Database Layer │
                       │ (PostgreSQL/    │
                       │  DuckDB)        │
                       └─────────────────┘
```

### Key Components

1. **Models Package**: Data structures and business entities
2. **Repository Pattern**: Abstract database operations with interface
3. **Database Driver**: Connection management and SQL execution
4. **Generic Query Engine**: Type-safe result mapping using Go generics
5. **Handlers**: HTTP request/response processing
6. **Configuration**: Environment and database settings


### Production Build

1. **Build the application**
   ```bash
   go build -o sales-analytics .
   ```

2. **Run the binary**
   ```bash
   ./sales-analytics
   ```
## Database Schema

The application uses four main tables:

- **category**: Product categories (Electronics, Books, etc.)
- **product**: Product catalog with category relationships
- **customer**: Customer information and demographics
- **sales_transaction**: Transaction records linking products and customers

Sample data is automatically inserted on application startup for demonstration purposes.

## Configuration

Database configuration can be modified in `main.go`:

```go
// For DuckDB
dbConfig.DSN = "duckdb"
dbConfig.DBSchemaFile = "migrations/duckdbschema.sql"

// For PostgreSQL  
dbConfig.Host = "localhost"
dbConfig.Port = 5432
dbConfig.Dbname = "your-database"
dbConfig.Username = "your-username"
dbConfig.Password = "your-password"
```

## Improvements
- Inclusion of Unit test
- Inclusion of Page limiting 
- Inclusion of Authentication and Authorization
- Inclusion of API rate limiting