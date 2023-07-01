INSERT INTO Education (
  job_seeker_id,
  institution,
  degree,
  field_of_study,
  start_date,
  end_date,
  created_at,
  updated_at
) VALUES (
  $1, -- job_seeker_id
  $2, -- institution
  $3, -- degree
  $4, -- field_of_study
  $5, -- start_date
  $6, -- end_date
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;
