DELETE FROM issues;
DELETE FROM projects;
DELETE FROM users;

INSERT INTO projects (project_id, project_name) VALUES
  (1, "Glada Schemat"),
  (2, "Max Power"),
  (3, "Mina Recept")
;

INSERT INTO issues (issue_title, issue_description, project) VALUES
  ("Make accessible", "Appen måste gå att använda av alla!", 1),
  ("Make more fun", "Lägg in mycket färger, appen skall vara rolig!", 1),
  ("Update damage system", "Damage system must be more complicated", 2),
  ("Receptmodell", NULL, 3)
;

INSERT INTO users (forename, surname, user_email, user_password) VALUES
  ("Test", "Testson", "testis@test.com", "gargamel")
;
