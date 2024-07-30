CREATE DATABASE gobrax;

USE gobrax;

CREATE TABLE driver(`id`, `name`, `cpf`);
CREATE TABLE trucks(`id`, `modelo`,  `placa`);
CREATE TABLE driver_truck(`driver_id`, `truck_id`, `created_at`, `updated_at`);