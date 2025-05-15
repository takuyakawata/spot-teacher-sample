-- Modify "grades" table
ALTER TABLE `grades` ADD COLUMN `code_number` bigint NOT NULL, ADD UNIQUE INDEX `code_number` (`code_number`);
