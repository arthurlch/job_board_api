-- Insert
INSERT INTO "User" (name, email, phone, role) VALUES ($1, $2, $3, $4) RETURNING id;

-- Select
SELECT * FROM "User";

-- Select a specific user by ID
SELECT * FROM "User" WHERE id = $1;

-- Update
UPDATE "User" SET name = $1, email = $2, phone = $3, role = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5;

-- Delete
DELETE FROM "User" WHERE id = $1;
