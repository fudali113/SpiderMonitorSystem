CREATE TABLE `all_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pcid` varchar(45) DEFAULT NULL,
  `ip` varchar(45) DEFAULT NULL,
  `step` int(11) DEFAULT NULL,
  `bid` varchar(45) DEFAULT NULL,
  `sid` varchar(128) DEFAULT NULL,
  `all` varchar(10000) DEFAULT NULL,
  `exception` varchar(5000) DEFAULT NULL,
  `time` datetime NOT NULL,
  `stc` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sid` (`sid`,`step`)
) ENGINE=InnoDB AUTO_INCREMENT=35530 DEFAULT CHARSET=utf8;



CREATE TABLE `comp_status` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pcid` varchar(128) DEFAULT NULL,
  `cpu` int(11) DEFAULT NULL,
  `mem` int(11) DEFAULT NULL,
  `io` int(11) DEFAULT NULL,
  `net` int(11) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  `data` varchar(3000) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36486 DEFAULT CHARSET=utf8;



CREATE TABLE `exception` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `adid` int(11) DEFAULT NULL,
  `exception` varchar(5000) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=536 DEFAULT CHARSET=utf8;



CREATE TABLE `finish` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` varchar(128) NOT NULL,
  `step` varchar(45) DEFAULT NULL,
  `pcid` varchar(45) DEFAULT NULL,
  `bid` varchar(45) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`,`sid`),
  UNIQUE KEY `finish_sid_uindex` (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=42963 DEFAULT CHARSET=utf8;


CREATE TABLE `heartbeat` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pcid` varchar(45) DEFAULT NULL,
  `ip` varchar(45) DEFAULT NULL,
  `deadtime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=288 DEFAULT CHARSET=utf8;


CREATE TABLE `traffic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `count` int(11) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `traffic_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2963 DEFAULT CHARSET=utf8;