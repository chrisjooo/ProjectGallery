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
rating      INT(1) DEFAULT '0',
authorID    INT UNSIGNED NOT NULL,
description TEXT,
created_at  DATETIME,
updated_at  DATETIME,
FOREIGN KEY (authorID) REFERENCES Account(id)
);

CREATE TABLE Rating (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
authorID    INT UNSIGNED NOT NULL,
projectID   INT UNSIGNED NOT NULL,
rating      INT NOT NULL,
created_at  DATETIME,
updated_at  DATETIME,
FOREIGN KEY (authorID) REFERENCES Account(id),
FOREIGN KEY (projectID) REFERENCES Project(id)
);