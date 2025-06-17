## Create Migration Using Soda

1. Product Information Table: 
    `soda generate fizz CreateProductInformationTable`
2. Sales transactions Table:
    `soda generate fizz CreateSalesTransactionTable`
3. Generate Foreign key for Transactions Table:
    `soda generate fizz CreateFKForSalesTransactionTable`

## Create Init for the Project
    `go mod init example/web-service-gin`

## Get Gin from Repo
    `go get -u github.com/gin-gonic/gin`

## Get Postgres from Repo
    `go get -u github.com/lib/pq`

## Get DuckDb from Repo
    `go get github.com/marcboeker/go-duckdb`
