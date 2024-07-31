CREATE DATABASE gobrax;

USE gobrax;

CREATE TABLE `driver` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `cnh` varchar(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `cnh_UNIQUE` (`cnh`)
);

CREATE TABLE `truck` (
  `id` varchar(36) NOT NULL,
  `brand` varchar(100) NOT NULL,
  `plate` varchar(7) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `plate_UNIQUE` (`plate`)
);

CREATE TABLE `driver_truck` (
  `id_driver` varchar(36) NOT NULL,
  `id_truck` varchar(36) NOT NULL,
  PRIMARY KEY (`id_driver`,`id_truck`),
  KEY `truck_idx` (`id_truck`),
  CONSTRAINT `driver` FOREIGN KEY (`id_driver`) REFERENCES `driver` (`id`),
  CONSTRAINT `truck` FOREIGN KEY (`id_truck`) REFERENCES `truck` (`id`)
);