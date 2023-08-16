-- name: InsertJobSeekerSkill :exec
INSERT INTO JobSeekerSkill (job_seeker_id, technical_skill, passive_skill)
VALUES ($1, $2, $3);

-- name: SelectJobSeekerSkillsByJobSeekerID :many
SELECT job_seeker_id, technical_skill, passive_skill
FROM JobSeekerSkill
WHERE job_seeker_id = $1;

-- name: DeleteJobSeekerSkill :exec
DELETE FROM JobSeekerSkill WHERE job_seeker_id = $1 AND technical_skill = $2 AND passive_skill = $3;
