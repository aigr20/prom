DELETE FROM issues;
DELETE FROM projects;
DELETE FROM users;

INSERT INTO projects (project_id, project_name, creation_date, last_changed) VALUES
  (1, "Glada Schemat", "2023-06-24 10:00:00", "2023-06-24 10:15:00"),
  (2, "Max Power", "2023-06-24 09:10:00", "2023-06-24 16:13:12"),
  (3, "Mina Recept", "2023-06-25 15:22:35", "2023-06-26 12:00:14")
;

INSERT INTO issues (issue_id, issue_title, issue_description, creation_date, last_changed, project) VALUES
  (1, "Make accessible", "Appen måste gå att använda av alla!", "2023-06-28 14:00:00", "2023-06-28 15:15:00", 1),
  (2, "Make more fun", "Lägg in mycket färger, appen skall vara rolig!", "2023-06-28 14:03:00", "2023-06-28 15:20:00", 1),
  (3, "Update damage system", "Damage system must be more complicated", "2023-06-29 13:12:03", "2023-06-29 16:12:00", 2),
  (4, "Receptmodell", "", "2023-06-30 10:00:00", "2023-06-30 15:00:34", 3)
;

INSERT INTO users (forename, surname, user_email, user_password) VALUES
  ("Test", "Testson", "testis@test.com", "gargamel")
;
