-- productinformation with PRIMARY KEY on id
CREATE TABLE IF NOT EXISTS productinformation (
    id VARCHAR NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    price DOUBLE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- salestransaction with proper FK (no ON DELETE/UPDATE)
CREATE TABLE IF NOT EXISTS salestransaction (
    id VARCHAR NOT NULL PRIMARY KEY,
    product_id VARCHAR NOT NULL,
    quantity INTEGER NOT NULL,
    date TIMESTAMP NOT NULL,
    customer_info VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (product_id) REFERENCES productinformation(id)
);
