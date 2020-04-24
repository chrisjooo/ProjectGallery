CREATE TABLE Account (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,          
username    VARCHAR(255) NOT NULL,
password    VARCHAR(255) NOT NULL,
full_name   VARCHAR(255) NOT NULL,
email       VARCHAR(255) NOT NULL,
description TEXT,
created_at  DATETIME
);

CREATE TABLE Project (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
name        VARCHAR(255) NOT NULL,
rate      INT(1) DEFAULT '0',
author    VARCHAR(255) NOT NULL,
description TEXT,
created_at  DATETIME,
updated_at  DATETIME
);

CREATE TABLE Rating (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
author    VARCHAR(255) NOT NULL,
project_id  INT UNSIGNED NOT NULL,
rate      INT NOT NULL,
created_at  DATETIME,
updated_at  DATETIME,
FOREIGN KEY (project_id) REFERENCES Project(id)
);