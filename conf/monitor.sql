CREATE TABLE `monitor`.`all_data` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `step` INT NULL,
  `bid` VARCHAR(45) NULL,
  `sid` VARCHAR(128) NULL,
  `all` VARCHAR(10000) NULL,
  `exception` VARCHAR(5000) NULL,
  `deadtime` DATETIME NOT NULL,
  PRIMARY KEY (`id`));


CREATE TABLE `monitor`.`exception` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `step` VARCHAR(45) NULL,
  `bid` VARCHAR(45) NULL,
  `exception` VARCHAR(1000) NULL,
  `data` VARCHAR(1000) NULL,
  `deadtime` DATETIME NOT NULL,
  PRIMARY KEY (`id`));

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
  `data` VARCHAR(45) NULL,
  PRIMARY KEY (`id`, `sid`));

ALTER TABLE `monitor`.`finish` 
ADD UNIQUE INDEX `sid_UNIQUE` (`sid` ASC);