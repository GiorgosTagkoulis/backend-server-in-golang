CREATE SCHEMA IF NOT EXISTS `sitoo_test_assignment`
DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sitoo_test_assignment`.`product`
CASCADE;
DROP TABLE IF EXISTS `sitoo_test_assignment`.`product_barcode`
CASCADE;
DROP TABLE IF EXISTS `sitoo_test_assignment`.`product_attribute`
CASCADE;

CREATE TABLE IF NOT EXISTS `sitoo_test_assignment`.`product` (
  `product_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(32) NOT NULL,
  `sku` VARCHAR(32) NOT NULL,
  `description` VARCHAR(1024) NULL,
  `price` DECIMAL(12,2) NOT NULL DEFAULT 0.00,
  `created` DATETIME NOT NULL,
  `last_updated` DATETIME NULL,
  PRIMARY KEY (`product_id`),
  UNIQUE INDEX (`sku` ASC),
  INDEX (`created`),
  INDEX (`last_updated`)
);

CREATE TABLE IF NOT EXISTS `sitoo_test_assignment`.`product_barcode` (
  `product_id` INT UNSIGNED NOT NULL,
  `barcode` VARCHAR(32) NOT NULL,
  PRIMARY KEY (`product_id`, `barcode`),
  UNIQUE INDEX (`barcode`)
 );

CREATE TABLE IF NOT EXISTS `sitoo_test_assignment`.`product_attribute` (
  `product_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(16) NOT NULL,
  `value` VARCHAR(32) NOT NULL,
  PRIMARY KEY (`product_id`, `name`)
);

INSERT INTO `sitoo_test_assignment`.`product` 
  (title, sku, price, created, last_updated)
  VALUES
  ('product_1', 'sku_1', 10.5, '2019-05-01', '2019-06-03'),
  ('product_2', 'sku_2', 10.0, '2019-05-02', '2019-06-05'),
  ('product_3', 'sku_3', 10.0, '2019-05-03', '2019-06-07'),
  ('product_4', 'sku_4', 15.0, '2019-05-04', '2019-06-12'),
  ('product_5', 'sku_5', 20.0, '2019-05-05', '2019-06-03'),
  ('product_6', 'sku_6', 25.0, '2019-05-06', '2019-06-03'),
  ('product_7', 'sku_7', 30.0, '2019-05-07', '2019-06-04'),
  ('product_8', 'sku_8', 35.0, '2019-05-08', '2019-06-08'),
  ('product_9', 'sku_9', 40.0, '2019-05-09', '2019-06-23'),
  ('product_10', 'sku_10', 45.0, '2019-05-10', '2019-07-10');

INSERT INTO `sitoo_test_assignment`.`product` 
  (title, sku, description, price, created, last_updated)
  VALUES
  ('product_11', 'sku_11', 'desc', 10.5, '2019-05-01', '2019-06-03');

INSERT INTO `sitoo_test_assignment`.`product_barcode` 
  (product_id, barcode)
  VALUES
  (1, 'barcode_1'),
  (1, 'barcode_111'),
  (1, 'barcode_1111'),
  (2, 'barcode_2'),
  (3, 'barcode_3'),
  (4, 'barcode_4'),
  (5, 'barcode_5'),
  (6, 'barcode_6'),
  (7, 'barcode_7'),
  (8, 'barcode_8'),
  (9, 'barcode_9'),
  (10, 'barcode_10'),
  (11, 'barcode_11');

INSERT INTO `sitoo_test_assignment`.`product_attribute`
  (product_id, name, value)
  VALUES
  (1, 'name_1', 'value_1'),
  (1, 'name_11', 'value_11'),
  (1, 'name_111', 'value_111'),
  (2, 'name_2', 'value_2'),
  (3, 'name_3', 'value_3'),
  (4, 'name_4', 'value_4'),
  (5, 'name_5', 'value_5'),
  (6, 'name_6', 'value_6'),
  (7, 'name_7', 'value_7'),
  (8, 'name_8', 'value_8'),
  (9, 'name_9', 'value_9'),
  (10, 'name_10', 'value_10'),
  (11, 'name_11', 'value_11');
