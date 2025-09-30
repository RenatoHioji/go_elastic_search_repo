CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price NUMERIC NOT NULL
);

INSERT INTO users (name, email) VALUES
   ('Alice Johnson', 'alice.johnson@example.com'),
   ('Bob Smith', 'bob.smith@example.com'),
   ('Charlie Brown', 'charlie.brown@example.com'),
   ('Diana Prince', 'diana.prince@example.com'),
   ('Ethan Hunt', 'ethan.hunt@example.com'),
   ('Fiona Gallagher', 'fiona.gallagher@example.com'),
   ('George Miller', 'george.miller@example.com'),
   ('Hannah Davis', 'hannah.davis@example.com'),
   ('Ian Wright', 'ian.wright@example.com'),
   ('Julia Roberts', 'julia.roberts@example.com');

INSERT INTO products (name, description, price) VALUES
    ('Laptop Pro 15"', 'High-performance laptop with 32GB RAM and 1TB SSD', 2499.99),
    ('Gaming Mouse', 'RGB optical gaming mouse with 6 programmable buttons', 79.95),
    ('Mechanical Keyboard', 'Backlit mechanical keyboard with blue switches', 129.50),
    ('4K Monitor', '27-inch 4K UHD monitor with HDR support', 399.00),
    ('Smartphone X', 'Latest generation smartphone with 256GB storage', 1099.99),
    ('Wireless Earbuds', 'Bluetooth 5.2 wireless earbuds with noise isolation', 149.99),
    ('External SSD', '1TB portable USB-C solid state drive', 189.99),
    ('Smartwatch', 'Fitness smartwatch with heart rate and GPS tracking', 229.00),
    ('Desk Chair', 'Ergonomic office chair with adjustable lumbar support', 349.49),
    ('Backpack', 'Water-resistant backpack with laptop compartment', 89.99);
