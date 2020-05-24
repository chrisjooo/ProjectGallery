CREATE TABLE account (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,          
username    VARCHAR(255) NOT NULL UNIQUE,
password    VARCHAR(255) NOT NULL,
full_name   VARCHAR(255) NOT NULL,
email       VARCHAR(255) NOT NULL,
profile_pic VARCHAR(255),
description TEXT,
created_at  DATETIME
);

CREATE TABLE project (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
name        VARCHAR(255) NOT NULL,
author      VARCHAR(255) NOT NULL,
project_pic VARCHAR(255),
description TEXT,
created_at  DATETIME,
FOREIGN KEY (author) REFERENCES account(username) ON DELETE CASCADE
);

CREATE TABLE vote (
id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
author      VARCHAR(255) NOT NULL,
project_id  INT UNSIGNED NOT NULL,
vote        BOOLEAN NOT NULL DEFAULT 0,
FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE,
FOREIGN KEY (author) REFERENCES account(username) ON DELETE CASCADE
);