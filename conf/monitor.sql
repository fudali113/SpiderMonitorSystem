CREATE TABLE `monitor`.`all_data` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `step` INT NULL,
  `bid` VARCHAR(45) NULL,
  `sid` VARCHAR(128) NULL,
  `stc` INT DEFAULT 0 NULL,
  `all` VARCHAR(10000) NULL,
  `exception` VARCHAR(5000) NULL,
  `time` DATETIME NOT NULL,
  PRIMARY KEY (`id`));
ALTER TABLE all_data add UNIQUE KEY (sid,step);





CREATE TABLE exception
(
    id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    adid INT(11),
    exception VARCHAR(5000),
    time DATETIME
);





CREATE TABLE `monitor`.`heartbeat` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `deadtime` DATETIME NOT NULL,
  PRIMARY KEY (`id`));





CREATE TABLE `monitor`.`finish` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL ,
  `bid` VARCHAR(45) NULL;
  `sid` VARCHAR(128) NOT NULL,
  `step` VARCHAR(45) NULL,
  `time` DATETIME NOT NULL,
  PRIMARY KEY (`id`, `sid`));

ALTER TABLE `monitor`.`finish` 
ADD UNIQUE INDEX `sid_UNIQUE` (`sid` ASC);





CREATE TABLE monitor.traffic
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    count INT,
    time DATETIME
);
CREATE UNIQUE INDEX traffic_id_uindex ON monitor.traffic (id);



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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;