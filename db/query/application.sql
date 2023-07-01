-- application.sql

-- Insert statement for the "Application" table
INSERT INTO Application (
  job_seeker_id,
  job_id,
  cover_letter,
  resume,
  status,
  created_at,
  updated_at
) VALUES (
  $1, -- job_seeker_id
  $2, -- job_id
  $3, -- cover_letter
  $4, -- resume
  $5, -- status
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;


-- Select statement for retrieving data from the "Application" table
SELECT * FROM Application;


-- Update statement for the "Application" table
UPDATE Application
SET status = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2; -- Specify the condition for updating, such as the "id" column


-- Delete statement for the "Application" table
DELETE FROM Application
WHERE id = $1; -- Specify the condition for deletion, such as the "id" column
