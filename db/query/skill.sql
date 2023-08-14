-- name: InsertSkill :exec
INSERT INTO Skill (name) VALUES ($1) RETURNING id;

-- name: SelectAllSkills :many
SELECT * FROM Skill;

-- name: SelectSkillByID :one
SELECT * FROM Skill WHERE id = $1;

-- name: UpdateSkill :exec
UPDATE Skill SET name = $1 WHERE id = $2;

-- name: DeleteSkill :exec
DELETE FROM Skill WHERE id = $1;
