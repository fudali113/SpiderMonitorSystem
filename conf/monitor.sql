CREATE TABLE `monitor`.`all_data` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pcid` VARCHAR(45) NULL,
  `ip` VARCHAR(45) NULL,
  `step` INT NULL,
  `bid` VARCHAR(45) NULL,
  `sid` VARCHAR(128) NULL,
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
  `data` VARCHAR(45) NULL,
  PRIMARY KEY (`id`, `sid`));

ALTER TABLE `monitor`.`finish` 
ADD UNIQUE INDEX `sid_UNIQUE` (`sid` ASC);