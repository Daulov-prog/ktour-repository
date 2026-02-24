CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE tours (
    id VARCHAR(36) PRIMARY KEY,
    country VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    max_slots INT NOT NULL,
    days INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    type VARCHAR(255) NOT NULL,
    in_stock BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE bookings (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    tour_id VARCHAR(36) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL
);  