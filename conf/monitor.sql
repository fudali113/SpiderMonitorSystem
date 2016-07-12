CREATE TABLE `monitor`.`all_data` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `step` INT NULL,
  `bid` VARCHAR(45) NULL,
  `sid` VARCHAR(128) NULL,
  `all` VARCHAR(1000) NULL,
  `execption` VARCHAR(1000) NULL,
  `deadtime` DATETIME NOT NULL,
  PRIMARY KEY (`id`));


CREATE TABLE `monitor`.`execption` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `step` VARCHAR(45) NULL,
  `bid` VARCHAR(45) NULL,
  `execption` VARCHAR(1000) NULL,
  `data` VARCHAR(1000) NULL,
  `deadtime` DATETIME NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `monitor`.`heartbeats` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `deadtime` DATETIME NOT NULL,
  PRIMARY KEY (`id`));