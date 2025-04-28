-- Create "products" table
CREATE TABLE `products` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `price` double NOT NULL,
    `description` varchar(500) NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
