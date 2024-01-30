USE `fantasy_products`;

-- Add data to table customers
INSERT INTO `customers` (`first_name`,`last_name`,`condition`) VALUES
('John','Doe',1),
('Jane','Doe',1),
('John','Smith',1),
('Jane','Smith',1),
('John','Doe',0),
('Jane','Doe',0),
('John','Smith',0),
('Jane','Smith',0);

-- Add data to table invoices
INSERT INTO `invoices` (`datetime`,`customer_id`,`total`) VALUES
('2019-01-01',1,100.00),
('2019-01-01',2,200.00),
('2019-01-01',3,300.00),
('2019-01-01',4,400.00),
('2019-01-01',5,500.00),
('2019-01-01',6,600.00),
('2019-01-01',7,700.00),
('2019-01-01',8,800.00);

-- Add data to table products
INSERT INTO `products` (`description`,`price`) VALUES
('Product 1',100.00),
('Product 2',200.00),
('Product 3',300.00),
('Product 4',400.00),
('Product 5',500.00),
('Product 6',600.00),
('Product 7',700.00),
('Product 8',800.00);

-- Add data to table sales
INSERT INTO `sales` (`quantity`,`invoice_id`,`product_id`) VALUES
(1,1,1),
(2,2,2),
(3,3,3),
(4,4,4),
(5,5,5),
(6,6,6),
(7,7,7),
(8,8,8);