package database

import (
	"aigr20/prom/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type IssueRepository struct {
	db          *sql.DB
	columnNames []string
}

func NewIssueRepository(db *sql.DB) *IssueRepository {
	return &IssueRepository{
		db:          db,
		columnNames: []string{"issue_id", "issue_title", "issue_description", "estimate", "creation_date", "last_changed", "project", "issue_status"},
	}
}

func (rep *IssueRepository) Close() error {
	return rep.db.Close()
}

func (rep *IssueRepository) GetIssuesFromProject(projectId int) ([]models.Issue, error) {
	const query = `
	SELECT
		i.issue_id,
		i.issue_title,
		i.issue_description,
		i.estimate,
		i.creation_date,
		i.last_changed,
		i.project,
		s.status_text,
		COALESCE(t.tag_id, -1),
		COALESCE(t.tag_text, ''),
		COALESCE(t.tag_color, '')
	FROM issues AS i
	JOIN issue_statuses AS s
	ON i.issue_status = s.status_id
	LEFT JOIN issue_tags AS itags
	ON itags.issue_id = i.issue_id
	LEFT JOIN tags AS t
	ON t.tag_id = itags.tag_id
	WHERE project = ?`
	rows, err := rep.db.Query(query, projectId)
	if err != nil {
		log.Println(err)
		return nil, ErrIssuesNotFound
	}
	defer rows.Close()

	issues, err := models.ScanIssues(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return issues, nil
}

func (rep *IssueRepository) GetOne(issueId int) (models.Issue, error) {
	const query = `
	SELECT
		i.issue_id,
		i.issue_title,
		i.issue_description,
		i.estimate,
		i.creation_date,
		i.last_changed,
		i.project,
		s.status_text,
		COALESCE(t.tag_id, -1),
		COALESCE(t.tag_text, ''),
		COALESCE(t.tag_color, '')
	FROM issues AS i
	JOIN issue_statuses AS s
	ON i.issue_status = s.status_id
	LEFT JOIN issue_tags AS itags
	ON itags.issue_id = i.issue_id
	LEFT JOIN tags AS t
	ON t.tag_id = itags.tag_id
	WHERE i.issue_id = ?`
	rows, err := rep.db.Query(query, issueId)
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueNotFound
	}

	issues, err := models.ScanIssues(rows)
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueNotFound
	} else if len(issues) > 1 {
		log.Println("Received multiple issues")
		return models.Issue{}, ErrIssueNotFound
	} else if len(issues) < 1 {
		log.Println("Found no issues")
		return models.Issue{}, ErrIssueNotFound
	}

	return issues[0], nil
}

func (rep *IssueRepository) CreateIssue(body models.IssueCreateForm) (models.Issue, error) {
	if body.Title == "" || body.ProjectID == 0 {
		return models.Issue{}, ErrIssueCreate
	}

	result, err := rep.db.Exec("INSERT INTO issues (issue_title, issue_description, estimate, project) VALUES (?, ?, ?, ?)", body.Title, body.Description, body.Estimate, body.ProjectID)
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueCreate
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueCreate
	}
	createdIssue, err := rep.GetOne(int(lastInsertId))
	if err != nil {
		return models.Issue{}, err
	}

	return createdIssue, nil
}

func (rep *IssueRepository) isLegalColumnName(toTest string) bool {
	for _, column := range rep.columnNames {
		if toTest == column {
			return true
		}
	}
	return false
}

func (rep *IssueRepository) UpdateIssue(target int, fields []string, values []any) (models.Issue, error) {
	if len(fields) != len(values) {
		return models.Issue{}, ErrUpdateFieldCount
	} else if len(fields) == 0 || len(values) == 0 {
		return models.Issue{}, ErrNoFields
	}

	queryBuilder := strings.Builder{}
	args := make([]any, 0)
	for i, fieldName := range fields {
		if !rep.isLegalColumnName(fieldName) {
			return models.Issue{}, ErrIllegalFieldName
		}

		queryBuilder.WriteString(fieldName)
		queryBuilder.WriteString("=?")
		if i != len(fields)-1 {
			queryBuilder.WriteString(",")
		}
	}
	args = append(args, values...)
	args = append(args, target)

	query := fmt.Sprintf("UPDATE issues SET %s WHERE issue_id = ?", queryBuilder.String())
	_, err := rep.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrUpdateFailed
	}

	updatedIssue, err := rep.GetOne(target)
	if err != nil {
		return models.Issue{}, ErrIssueNotFound
	}
	return updatedIssue, nil
}

func (rep *IssueRepository) AddTags(target int, tags []int) ([]models.IssueTag, error) {
	issueArgs := make([]interface{}, 0, len(tags)*2)
	builder := strings.Builder{}
	for i := range tags {
		builder.WriteString("(?, ?)")
		if i < len(tags)-1 {
			builder.WriteString(",")
		}
		issueArgs = append(issueArgs, target, tags[i])
	}

	query := fmt.Sprintf("INSERT IGNORE INTO issue_tags (issue_id, tag_id) VALUES %s", builder.String())
	_, err := rep.db.Exec(query, issueArgs...)
	if err != nil {
		log.Println(err)
		return []models.IssueTag{}, err
	}
	_, err = rep.db.Exec("UPDATE issues SET last_changed = CURRENT_TIMESTAMP() WHERE issue_id = ?", target)
	if err != nil {
		log.Println(err)
		return []models.IssueTag{}, err
	}

	newTags, err := rep.getTags(target)
	if err != nil {
		return []models.IssueTag{}, err
	}

	return newTags, nil
}

func (rep *IssueRepository) RemoveTags(target int, tags []int) ([]models.IssueTag, error) {
	builder := strings.Builder{}
	args := make([]interface{}, len(tags))
	for i := range tags {
		builder.WriteString("tag_id = ?")
		if i < len(tags)-1 {
			builder.WriteString(" OR ")
		}
		args[i] = tags[i]
	}
	args = append([]interface{}{target}, args...)

	query := fmt.Sprintf("DELETE FROM issue_tags WHERE issue_id = ? AND (%s)", builder.String())
	_, err := rep.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return []models.IssueTag{}, err
	}
	_, err = rep.db.Exec("UPDATE issues SET last_changed = CURRENT_TIMESTAMP() WHERE issue_id = ?", target)
	if err != nil {
		log.Println(err)
		return []models.IssueTag{}, err
	}

	newTags, err := rep.getTags(target)
	if err != nil {
		return []models.IssueTag{}, err
	}

	return newTags, nil
}

func (rep *IssueRepository) getTags(issueID int) ([]models.IssueTag, error) {
	const query = `
	SELECT t.tag_id, t.tag_text, t.tag_color FROM tags AS t
	JOIN issue_tags AS itag ON itag.tag_id = t.tag_id
	WHERE itag.issue_id = ?`
	rows, err := rep.db.Query(query, issueID)
	if err != nil {
		log.Println(err)
		return []models.IssueTag{}, ErrTagNotFound
	}

	tags, err := models.ScanTags(rows)
	if err != nil {
		return []models.IssueTag{}, ErrTagNotFound
	}

	return tags, nil
}

// Should only be used in tests
func (rep *IssueRepository) CustomQuery(query string, args ...any) (sql.Result, error) {
	return rep.db.Exec(query, args...)
}
