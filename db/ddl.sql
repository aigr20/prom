DROP TABLE IF EXISTS issue_tags;
DROP TABLE IF EXISTS project_tags;
DROP TABLE IF EXISTS issues;
DROP TABLE IF EXISTS issue_statuses;
DROP TABLE IF EXISTS projects;
DROP TABLE tags;
DROP TABLE IF EXISTS users;

CREATE TABLE projects (
  project_id INT NOT NULL AUTO_INCREMENT,
  project_name VARCHAR(35) NOT NULL,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_changed TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (project_id)
);

CREATE TABLE issue_statuses (
  status_id INT NOT NULL AUTO_INCREMENT,
  status_text VARCHAR(30) NOT NULL,

  PRIMARY KEY (status_id)
);

CREATE TABLE issues (
  issue_id INT NOT NULL AUTO_INCREMENT,
  issue_title VARCHAR(50) NOT NULL,
  issue_description VARCHAR(300) NOT NULL DEFAULT "",
  estimate INT NOT NULL DEFAULT 0,
  creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_changed TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  project INT NOT NULL,
  issue_status INT NULL DEFAULT 1,

  PRIMARY KEY (issue_id),
  FOREIGN KEY project_fk (project) REFERENCES projects (project_id) ON DELETE CASCADE,
  FOREIGN KEY status_fk (issue_status) REFERENCES issue_statuses (status_id)
);

CREATE TABLE tags (
  tag_id INT NOT NULL AUTO_INCREMENT,
  tag_text VARCHAR(15) NOT NULL,
  tag_color CHAR(7) NOT NULL,

  PRIMARY KEY (tag_id),
  UNIQUE INDEX (tag_text, tag_color)
);

CREATE TABLE issue_tags (
  issue_id INT NOT NULL,
  tag_id INT NOT NULL,

  PRIMARY KEY (issue_id, tag_id),
  FOREIGN KEY issue_fk (issue_id) REFERENCES issues (issue_id) ON DELETE CASCADE,
  FOREIGN KEY tag_fk (tag_id) REFERENCES tags (tag_id) ON DELETE CASCADE
);

CREATE TABLE project_tags (
  project_id INT NOT NULL,
  tag_id INT NOT NULL,

  PRIMARY KEY (project_id, tag_id),
  FOREIGN KEY p2t_project_fk (project_id) REFERENCES projects (project_id) ON DELETE CASCADE,
  FOREIGN KEY p2t_tag_fk (tag_id) REFERENCES tags (tag_id) ON DELETE CASCADE
);

CREATE TABLE users (
  user_id INT NOT NULL AUTO_INCREMENT,
  forename VARCHAR(15) NOT NULL DEFAULT "",
  surname VARCHAR(20) NOT NULL DEFAULT "",
  user_email VARCHAR(30) NOT NULL,
  user_password VARCHAR(20) NOT NULL,
  user_role TINYINT DEFAULT 1 NOT NULL,

  PRIMARY KEY (user_id),
  UNIQUE INDEX (user_email)
);

DROP VIEW IF EXISTS project_tag_counts;

CREATE VIEW project_tag_counts AS
  SELECT
    ptags.project_id,
    t.tag_text,
    SUM(IF(i.project IS NULL, 0, 1)) AS tag_count
  FROM tags AS t
  JOIN project_tags AS ptags ON ptags.tag_id = t.tag_id
  LEFT JOIN issue_tags AS itags ON itags.tag_id = t.tag_id
  LEFT JOIN issues AS i ON i.issue_id = itags.issue_id AND i.project = ptags.project_id
  GROUP BY ptags.project_id, t.tag_text, t.tag_color
  ORDER BY tag_count DESC
;
