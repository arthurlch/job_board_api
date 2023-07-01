INSERT INTO Job (
  title,
  description,
  requirements,
  location,
  salary,
  company_id,
  created_at,
  updated_at
) VALUES (
  $1, -- title
  $2, -- description
  $3, -- requirements
  $4, -- location
  $5, -- salary
  $6, -- company_id
  CURRENT_TIMESTAMP, -- created_at
  CURRENT_TIMESTAMP -- updated_at
) RETURNING id;
