-- DDL
DROP DATABASE IF EXISTS `fantasy_products`;

CREATE DATABASE `fantasy_products`;

USE `fantasy_products`;

-- Table structure for table `customers`
CREATE TABLE `customers` (
    `id` int NOT NULL AUTO_INCREMENT,
    `first_name` varchar(45) DEFAULT NULL,
    `last_name` varchar(45) DEFAULT NULL,
    `condition` tinyint(1) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

-- Table structure for table `invoices`
CREATE TABLE `invoices` (
    `id` int NOT NULL AUTO_INCREMENT,
    `datetime` datetime DEFAULT NULL,
    `customer_id` int DEFAULT NULL,
    `total` float DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_invoices_customer_id` (`customer_id`),
    CONSTRAINT `fk_invoices_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Table structure for table `products`
CREATE TABLE `products` (
    `id` int NOT NULL AUTO_INCREMENT,
    `description` varchar(100) DEFAULT NULL,
    `price` float DEFAULT NULL,
    PRIMARY KEY (`id`)
);

-- Table structure for table `sales`
CREATE TABLE `sales` (
    `id` int NOT NULL AUTO_INCREMENT,
    `quantity` int DEFAULT NULL,
    `invoice_id` int DEFAULT NULL,
    `product_id` int DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_sales_invoice_id` (`invoice_id`),
    KEY `idx_sales_product_id` (`product_id`),
    CONSTRAINT `fk_sales_invoice_id` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_sales_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);