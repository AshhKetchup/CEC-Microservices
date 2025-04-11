-- Create databases for each microservice
CREATE DATABASE IF NOT EXISTS auth_db;
CREATE DATABASE IF NOT EXISTS product_db;
CREATE DATABASE IF NOT EXISTS order_db;

-- Create users
CREATE USER IF NOT EXISTS 'auth_user'@'%' IDENTIFIED BY 'authpass';
CREATE USER IF NOT EXISTS 'product_user'@'%' IDENTIFIED BY 'productpass';
CREATE USER IF NOT EXISTS 'order_user'@'%' IDENTIFIED BY 'orderpass';

-- Grant permissions to respective users
GRANT ALL PRIVILEGES ON auth_db.* TO 'auth_user'@'%';
GRANT ALL PRIVILEGES ON product_db.* TO 'product_user'@'%';
GRANT ALL PRIVILEGES ON order_db.* TO 'order_user'@'%';

-- Apply changes
FLUSH PRIVILEGES;
