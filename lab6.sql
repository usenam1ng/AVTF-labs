CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price NUMERIC(10, 2),
    description TEXT,
    category_id INT REFERENCES categories(id)
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_date DATE,
    status VARCHAR(50),
    customer_id INT REFERENCES customers(id)
);

CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    product_id INT REFERENCES products(id),
    quantity INT,
    unit_price NUMERIC(10, 2)
);

-- Заполним таблицу категорий
INSERT INTO categories (name) VALUES ('Smartphones'), ('Laptops'), ('Headphones'), ('Kettle');

-- Заполним таблицу продуктов
INSERT INTO products (name, price, description, category_id) VALUES
    ('iPhone 13', 999.99, 'phone', 1),
    ('MacBook Pro', 1499.99, 'i7/16gb/512gb/radeon pro 555x', 2),
    ('Sony WH-1000XM4', 349.99, 'High-quality noise-canceling headphones', 3);

-- Заполним таблицу клиентов
INSERT INTO customers (first_name, last_name, email) VALUES
    ('John', 'Doe', 'john.doe@email.com'),
    ('Alice', 'Smith', 'alice.smith@email.com');

-- Заполним таблицу заказов
INSERT INTO orders (order_date, status, customer_id) VALUES
    ('2023-12-20', 'Shipped', 1),
    ('2023-12-21', 'Processing', 2);


-- Заполним таблицу деталей заказов
INSERT INTO order_details (order_id, product_id, quantity, unit_price) VALUES
    (1, 1, 1, 999.99),
    (2, 3, 2, 349.99);

SELECT * FROM products;

SELECT p.id, p.name AS product_name, c.name AS category_name
FROM products p
INNER JOIN categories c ON p.category_id = c.id;

UPDATE orders SET status = 'Delivered' WHERE id = 1;

DELETE FROM customers WHERE email = 'alice.smith@email.com';

SELECT * FROM products ORDER BY price DESC;

SELECT * FROM orders WHERE customer_id = 1;

SELECT customer_id, COUNT(*) AS order_count
FROM orders
GROUP BY customer_id;

SELECT * FROM products WHERE name LIKE '%Book%';

SELECT o.id, o.status, c.first_name, c.last_name
FROM orders o
INNER JOIN customers c ON o.customer_id = c.id;

SELECT order_id, SUM(quantity * unit_price) AS total_cost
FROM order_details
GROUP BY order_id;


drop table if exists customers cascade;
drop table if exists products cascade;
drop table if exists categories cascade;
drop table if exists order_details cascade;
drop table if exists orders cascade;
