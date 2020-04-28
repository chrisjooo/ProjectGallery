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
author    VARCHAR(255) NOT NULL,
description TEXT,
created_at  DATETIME
);

CREATE TABLE Vote (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
author    VARCHAR(255) NOT NULL,
project_id  INT UNSIGNED NOT NULL,
vote      BOOLEAN NOT NULL DEFAULT 0,
FOREIGN KEY (project_id) REFERENCES Project(id)
);