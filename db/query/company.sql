INSERT INTO Company (
  user_id,
  name,
  email,
  phone,
  website,
  logo,
  description,
  created_at,
  updated_at
) VALUES (
  $1, -- user_id
  $2, -- name
  $3, -- email
  $4, -- phone
  $5, -- website
  $6, -- logo
  $7, -- description
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;
