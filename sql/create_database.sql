CREATE SCHEMA IF NOT EXISTS `stream_media` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin;

USE mysql;
DELETE FROM user WHERE User='stream_media' AND Host='%';
FLUSH PRIVILEGES;
GRANT ALL PRIVILEGES ON stream_media.* TO 'stream_media'@'%' IDENTIFIED BY 'stream_media';
GRANT ALL PRIVILEGES ON stream_media.* TO 'stream_media'@'localhost' IDENTIFIED BY 'stream_media';
FLUSH PRIVILEGES;