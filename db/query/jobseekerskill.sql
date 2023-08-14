-- name: InsertJobSeekerSkill :exec
INSERT INTO JobSeekerSkill (job_seeker_id, skill_id)
VALUES ($1, $2);

-- name: SelectJobSeekerSkillsByJobSeekerID :many
SELECT * FROM JobSeekerSkill WHERE job_seeker_id = $1;

-- name: DeleteJobSeekerSkill :exec
DELETE FROM JobSeekerSkill WHERE job_seeker_id = $1 AND skill_id = $2;
