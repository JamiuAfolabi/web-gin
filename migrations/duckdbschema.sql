-- Category Table
CREATE TABLE IF NOT EXISTS category (
    id INTEGER PRIMARY KEY,
    name VARCHAR NOT NULL,
);

-- Product Table
CREATE TABLE IF NOT EXISTS product (
    id INTEGER PRIMARY KEY,
    name VARCHAR NOT NULL,
    category_id INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES category(id)
);

-- Customer Table
CREATE TABLE IF NOT EXISTS customer (
    id INTEGER PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE,
    gender VARCHAR,
    location VARCHAR,
    created_at DATE
);

-- Sales Transaction Table
CREATE TABLE IF NOT EXISTS sales_transaction (
    id INTEGER PRIMARY KEY,
    product_id INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    transaction_date DATE NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    total_amount DECIMAL(12, 2) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id),
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);

