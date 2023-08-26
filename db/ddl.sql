DROP TABLE IF EXISTS issue_tags;
DROP TABLE IF EXISTS project_tags;
DROP TABLE IF EXISTS sprint_issues;
DROP TABLE IF EXISTS sprints;
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

CREATE TABLE sprints (
  sprint_id INT NOT NULL AUTO_INCREMENT,
  sprint_name VARCHAR(30) NOT NULL,
  project INT NOT NULL,
  sprint_start DATE NOT NULL,
  sprint_end DATE NOT NULL,
  finished BOOLEAN NOT NULL DEFAULT FALSE,
  current BOOLEAN NOT NULL DEFAULT TRUE,

  PRIMARY KEY (sprint_id),
  FOREIGN KEY sprint_project_fk (project) REFERENCES projects (project_id) ON DELETE CASCADE
);

CREATE TABLE sprint_issues (
  issue_id INT NOT NULL,
  sprint_id INT NOT NULL,
  issue_priority INT NOT NULL DEFAULT 0,

  PRIMARY KEY (issue_id, sprint_id),
  FOREIGN KEY sprint_issue_fk (issue_id) REFERENCES issues (issue_id),
  FOREIGN KEY sprint_fk (sprint_id) REFERENCES sprints (sprint_id)
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
DROP VIEW IF EXISTS sprint_issues_v;
DROP VIEW IF EXISTS issues_no_sprint;

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

CREATE VIEW sprint_issues_v AS
  SELECT
    sprints.sprint_id,
    sprints.sprint_name,
    sprints.sprint_start,
    sprints.sprint_end,
    sprints.finished,
    sprints.current,
    sprints.project AS sprint_project,
    issues.issue_id,
    issues.issue_title,
    issues.issue_description,
    issues.estimate,
    issues.creation_date,
    issues.last_changed,
    issues.project,
    issue_statuses.status_text,
    COALESCE(tags.tag_id, -1) AS tag_id,
    COALESCE(tags.tag_text, "") AS tag_text,
    COALESCE(tags.tag_color, "") AS tag_color,
    sprint_issues.issue_priority
  FROM sprints
  JOIN sprint_issues ON sprint_issues.sprint_id = sprints.sprint_id
  JOIN issues ON issues.issue_id = sprint_issues.issue_id
  JOIN issue_statuses ON issue_statuses.status_id = issues.issue_status
  LEFT JOIN issue_tags ON issue_tags.issue_id = issues.issue_id
  LEFT JOIN tags ON tags.tag_id = issue_tags.tag_id
  ORDER BY sprints.project ASC, sprint_issues.issue_priority DESC
;

CREATE VIEW issues_no_sprint AS
  SELECT
    sq.issue_id,
    sq.issue_title,
    sq.issue_description,
    sq.estimate,
    sq.creation_date,
    sq.last_changed,
    sq.project,
    sq.status_text,
    sq.tag_id,
    sq.tag_text,
    sq.tag_color
  FROM (
    SELECT
      sprints.sprint_id,
      issues.issue_id,
      issues.issue_title,
      issues.issue_description,
      issues.estimate,
      issues.creation_date,
      issues.last_changed,
      issues.project,
      issue_statuses.status_text,
      COALESCE(tags.tag_id, -1) AS tag_id,
      COALESCE(tags.tag_text, "") AS tag_text,
      COALESCE(tags.tag_color, "") AS tag_color
    FROM sprints
    JOIN sprint_issues ON sprint_issues.sprint_id = sprints.sprint_id
    RIGHT JOIN issues ON issues.issue_id = sprint_issues.issue_id
    RIGHT JOIN issue_statuses ON issue_statuses.status_id = issues.issue_status
    LEFT JOIN issue_tags ON issue_tags.issue_id = issues.issue_id
    LEFT JOIN tags ON tags.tag_id = issue_tags.tag_id
    WHERE sprints.sprint_id IS NULL
  ) AS sq
  ORDER BY project ASC
;

