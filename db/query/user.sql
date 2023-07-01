INSERT INTO "User" (
  name,
  email,
  phone,
  created_at,
  updated_at
) VALUES (
  $1, -- name
  $2, -- email
  $3, -- phone
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;
