-- Insert products
INSERT INTO products(name,price,vat) VALUES('Desktop Computer', 1000, 18);
INSERT INTO products(name,price,vat) VALUES('Macbook Air', 1200, 18);
INSERT INTO products(name,price,vat) VALUES('iPhone 13 Max 512G', 1299, 18);
INSERT INTO products(name,price,vat) VALUES('T-shirt', 30.99, 8);
INSERT INTO products(name,price,vat) VALUES('Hand Cream', 8, 8);
INSERT INTO products(name,price,vat) VALUES('Shoes', 99, 18);
INSERT INTO products(name,price,vat) VALUES('Honey 1KG', 15, 1);
INSERT INTO products(name,price,vat) VALUES('Tomato 1KG', 10, 1);
INSERT INTO products(name,price,vat) VALUES('Food 1', 15, 1);
INSERT INTO products(name,price,vat) VALUES('Food 2', 8, 1);
INSERT INTO products(name,price,vat) VALUES('Food 2', 25, 1)

-- Insert Users
INSERT INTO users(name) VALUES('User1');
INSERT INTO users(name) VALUES('User2');
INSERT INTO users(name) VALUES('User3');
INSERT INTO users(name) VALUES('User4');
INSERT INTO users(name) VALUES('User5');
INSERT INTO users(name) VALUES('User6');
INSERT INTO users(name) VALUES('User7');
INSERT INTO users(name) VALUES('User8');
INSERT INTO users(name) VALUES('User9');
INSERT INTO users(name) VALUES('User10');

-- Create carts for each user
INSERT INTO carts(user_id) VALUES(1);
INSERT INTO carts(user_id) VALUES(2);
INSERT INTO carts(user_id) VALUES(3);
INSERT INTO carts(user_id) VALUES(4);
INSERT INTO carts(user_id) VALUES(5);
INSERT INTO carts(user_id) VALUES(6);
INSERT INTO carts(user_id) VALUES(7);
INSERT INTO carts(user_id) VALUES(8);
INSERT INTO carts(user_id) VALUES(9);
INSERT INTO carts(user_id) VALUES(10);

-- Add products to some users carts
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(1, 1, 1);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(1, 2, 1);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(1, 3, 1);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(2, 4, 4);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(2, 6, 2);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(3, 7, 4);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(3, 8, 5);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(3, 9, 6);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(4, 10, 5);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(4, 11, 12);
INSERT INTO cart_products(cart_id, product_id, quantity) VALUES(4, 9, 2);


-- Add some order history for some Users
-- User 1 has 3 orders. Next order is 4th order
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(1, '2022-08-01', 1200);
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(1, '2022-08-01', 1500);
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(1, '2022-08-01', 1500);

-- User 2 has big orders last month but small orders this month.
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(2, '2022-07-01', 1200);
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(2, '2022-07-01', 2500);
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(2, '2022-08-01', 50);
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(2, '2022-08-01', 100);
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(2, '2022-08-01', 50);
INSERT INTO orders(user_id,ordered_at,total_paid) VALUES(2, '2022-08-01', 100);

