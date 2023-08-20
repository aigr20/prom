DELETE FROM sprint_issues;
DELETE FROM issues;
DELETE FROM tags;
DELETE FROM issue_statuses;
DELETE FROM sprints;
DELETE FROM projects;
DELETE FROM users;

INSERT INTO projects (project_id, project_name, creation_date, last_changed) VALUES
  (1, "Glada Schemat", "2023-06-24 10:00:00", "2023-06-24 10:15:00"),
  (2, "Max Power", "2023-06-24 09:10:00", "2023-06-24 16:13:12"),
  (3, "Mina Recept", "2023-06-25 15:22:35", "2023-06-26 12:00:14")
;

INSERT INTO sprints (sprint_id, sprint_name, project, sprint_start, sprint_end, finished, current) VALUES 
  (1, "Sprint 1", 1, "2023-06-24", "2023-06-30", TRUE, FALSE),
  (2, "Sprint 2", 1, "2023-07-01", "2023-07-08", FALSE, TRUE),
  (3, "Sprint 1", 2, "2023-06-24", "2023-06-28", FALSE, TRUE)
;

INSERT INTO issue_statuses (status_id, status_text) VALUES
 (1, "TODO"),
 (2, "In Progress"),
 (3, "Finished")
;

INSERT INTO tags (tag_id, tag_text, tag_color) VALUES
  (1, "important", "#ff0000"),
  (2, "bug", "#0000ff")
;

INSERT INTO issues (issue_id, issue_title, issue_description, estimate, creation_date, last_changed, project, issue_status) VALUES
  (1, "Make accessible", "Appen måste gå att använda av alla!", 5, "2023-06-28 14:00:00", "2023-06-28 15:15:00", 1, 1),
  (2, "Make more fun", "Lägg in mycket färger, appen skall vara rolig!", 3, "2023-06-28 14:03:00", "2023-06-28 15:20:00", 1, 2),
  (3, "Update damage system", "Damage system must be more complicated", 2, "2023-06-29 13:12:03", "2023-06-29 16:12:00", 2, 3),
  (4, "Receptmodell", "", 1, "2023-06-30 10:00:00", "2023-06-30 15:00:34", 3, 1)
;

INSERT INTO sprint_issues (issue_id, sprint_id, issue_priority) VALUES
  (1, 1, 5),
  (2, 2, 3),
  (3, 3, 2)
;

INSERT INTO issue_tags (issue_id, tag_id) VALUES
  (1, 2),
  (1, 1),
  (2, 1),
  (3, 2)
;

INSERT INTO project_tags (project_id, tag_id) VALUES
  (1, 1),
  (1, 2),
  (2, 2)
;

INSERT INTO users (forename, surname, user_email, user_password) VALUES
  ("Test", "Testson", "testis@test.com", "gargamel")
;
