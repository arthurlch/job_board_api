INSERT INTO Experience (
  job_seeker_id,
  title,
  company,
  location,
  start_date,
  end_date,
  description,
  created_at,
  updated_at
) VALUES (
  $1, -- job_seeker_id
  $2, -- title
  $3, -- company
  $4, -- location
  $5, -- start_date
  $6, -- end_date
  $7, -- description
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;
