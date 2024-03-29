CREATE TABLE `user` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `username` VARCHAR(255) UNIQUE NOT NULL,
  `fullname` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE `event` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `id_user` BIGINT UNSIGNED NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `event_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `event` ADD FOREIGN KEY (`id_user`) REFERENCES `user` (`id`);
