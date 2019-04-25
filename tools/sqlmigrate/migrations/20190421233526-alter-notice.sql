
-- +migrate Up
ALTER TABLE `notice` ADD COLUMN `detail` varchar(256) NOT NULL;

-- +migrate Down
ALTER TABLE `notice` DROP COLUMN `detail`;
