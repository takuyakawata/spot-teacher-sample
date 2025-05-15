-- Modify "lesson_confirmations" table
ALTER TABLE `lesson_confirmations` ADD COLUMN `created_at` timestamp NOT NULL, ADD COLUMN `updated_at` timestamp NOT NULL;
-- Modify "lesson_reservations" table
ALTER TABLE `lesson_reservations` ADD COLUMN `created_at` timestamp NOT NULL, ADD COLUMN `updated_at` timestamp NOT NULL;
