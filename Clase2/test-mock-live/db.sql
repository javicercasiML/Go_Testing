DROP DATABASE IF EXISTS item_db;
CREATE DATABASE IF NOT EXISTS item_db;

USE item_db;

-- DDL
CREATE TABLE `items` (
    `id`           varchar(50) NOT NULL,
    `name`         varchar(50) NOT NULL,
    `weight`       decimal NOT NULL,
    `price`        decimal NOT NULL,
    `release_date` date    NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_items_name` (`name`)
);

-- DATASET
INSERT INTO `items` (`id`,`name`,`weight`,`price`,`release_date`) VALUES ("AAA", "Pepsi", 1.5, 180, "1950-11-03");
INSERT INTO `items` (`id`,`name`,`weight`,`price`,`release_date`) VALUES ("BBB", "Marroc", 0.2, 60, "1990-12-24");
INSERT INTO `items` (`id`,`name`,`weight`,`price`,`release_date`) VALUES ("CCC", "Coca", 1.5, 240, "1930-03-15");