CREATE TABLE "User" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255),
  phone VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE JobSeeker (
  id SERIAL PRIMARY KEY,
  user_id INT UNIQUE,
  resume TEXT,
  skills TEXT[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES "User"(id)
);

CREATE TABLE Company (
  id SERIAL PRIMARY KEY,
  user_id INT UNIQUE,
  name VARCHAR(255),
  email VARCHAR(255),
  phone VARCHAR(255),
  website VARCHAR(255),
  logo VARCHAR(255),
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES "User"(id)
);

CREATE TABLE Education (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  institution VARCHAR(255),
  degree VARCHAR(255),
  field_of_study VARCHAR(255),
  start_date DATE,
  end_date DATE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id)
);

CREATE TABLE Experience (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  title VARCHAR(255),
  company VARCHAR(255),
  location VARCHAR(255),
  start_date DATE,
  end_date DATE,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id)
);

CREATE TABLE Job (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  description TEXT,
  requirements TEXT,
  location VARCHAR(255),
  salary INTEGER,
  company_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES Company(id)
);

CREATE TABLE Application (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  job_id INT,
  cover_letter TEXT,
  resume TEXT,
  status VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id),
  FOREIGN KEY (job_id) REFERENCES Job(id)
);
