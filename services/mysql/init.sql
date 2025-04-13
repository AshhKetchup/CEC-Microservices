-- Create databases for each microservice
CREATE DATABASE IF NOT EXISTS auth_db;
CREATE DATABASE IF NOT EXISTS product_db;
CREATE DATABASE IF NOT EXISTS order_db;
CREATE DATABASE IF NOT EXISTS cart_db;
CREATE DATABASE IF NOT EXISTS delivery_db;

-- Create users
CREATE USER IF NOT EXISTS 'auth_user'@'%' IDENTIFIED BY 'authpass';
CREATE USER IF NOT EXISTS 'product_user'@'%' IDENTIFIED BY 'productpass';
CREATE USER IF NOT EXISTS 'order_user'@'%' IDENTIFIED BY 'orderpass';
CREATE USER IF NOT EXISTS 'cart_user'@'%' IDENTIFIED BY 'cartpass';
CREATE USER IF NOT EXISTS 'delivery_user'@'%' IDENTIFIED BY 'deliverypass';

-- Grant permissions to respective users
GRANT ALL PRIVILEGES ON auth_db.* TO 'auth_user'@'%';
GRANT ALL PRIVILEGES ON product_db.* TO 'product_user'@'%';
GRANT ALL PRIVILEGES ON order_db.* TO 'order_user'@'%';
GRANT ALL PRIVILEGES ON cart_db.* TO 'cart_user'@'%';
GRANT ALL PRIVILEGES ON delivery_db.* TO 'delivery_user'@'%';

-- Apply changes
FLUSH PRIVILEGES;
