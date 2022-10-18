CREATE TABLE `template` (
  `id` bigint unsigned NOT NULL,
  `id_kafka` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_input` varchar(50) NOT NULL DEFAULT 'SYSTEM',
  `tgl_input` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_update` varchar(50) NOT NULL DEFAULT 'SYSTEM',
  `tgl_update` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `tgl_kafka` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `uuid` varchar(36) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id_kafka` (`id_kafka`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;