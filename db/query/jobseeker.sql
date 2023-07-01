INSERT INTO JobSeeker (
  user_id,
  resume,
  skills,
  created_at,
  updated_at
) VALUES (
  $1, -- user_id
  $2, -- resume
  $3, -- skills
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;
