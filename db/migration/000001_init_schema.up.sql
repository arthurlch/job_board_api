CREATE TYPE job_status AS ENUM ('active', 'closed', 'draft', 'paused');
CREATE TYPE application_status AS ENUM ('applied', 'interviewed', 'offered', 'rejected');
CREATE TYPE interview_status AS ENUM ('pending', 'passed', 'failed');

CREATE TABLE "User" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(320) NOT NULL UNIQUE,
  phone VARCHAR(255),
  role VARCHAR(50) CHECK (role IN ('job_seeker', 'company')),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE JobSeeker (
  id SERIAL PRIMARY KEY,
  user_id INT UNIQUE,
  resume VARCHAR(512),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES "User"(id)
);


CREATE TABLE Company (
  id SERIAL PRIMARY KEY,
  user_id INT UNIQUE,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(320) NOT NULL,
  phone VARCHAR(255),
  website VARCHAR(512),
  logo VARCHAR(512),
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES "User"(id)
);

CREATE TABLE Institution (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE CompanyEntity (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE JobCategory (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE ExperienceType (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE CHECK (name IN ('internship', 'full-time', 'part-time', 'freelance', 'contractor'))
);

CREATE TABLE Education (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  institution_id INT,
  degree VARCHAR(255),
  field_of_study VARCHAR(255),
  start_date DATE,
  end_date DATE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id),
  FOREIGN KEY (institution_id) REFERENCES Institution(id)
);

CREATE TABLE Experience (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  title VARCHAR(255),
  company_id INT,
  location VARCHAR(255),
  start_date DATE,
  end_date DATE CHECK (end_date >= start_date),
  type_id INT, 
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id),
  FOREIGN KEY (company_id) REFERENCES CompanyEntity(id),
  FOREIGN KEY (type_id) REFERENCES ExperienceType(id)
);

CREATE TABLE Job (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  requirements TEXT,
  location VARCHAR(255),
  salary INTEGER,
  company_id INT,
  category_id INT,
  status job_status DEFAULT 'active',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES Company(id),
  FOREIGN KEY (category_id) REFERENCES JobCategory(id)
);

CREATE INDEX idx_job_status ON Job(status);
CREATE INDEX idx_job_company ON Job(company_id);

CREATE TABLE Application (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  job_id INT,
  cover_letter TEXT,
  resume VARCHAR(512),
  status application_status DEFAULT 'applied',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id),
  FOREIGN KEY (job_id) REFERENCES Job(id)
);

CREATE TABLE Skill (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE JobSeekerSkill (
  job_seeker_id INT,
  skill_id INT,
  PRIMARY KEY (job_seeker_id, skill_id),
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id),
  FOREIGN KEY (skill_id) REFERENCES Skill(id)
);

CREATE INDEX idx_skill_id ON JobSeekerSkill(skill_id);

CREATE TABLE JobViews (
  job_id INT PRIMARY KEY,
  view_count INT,
  FOREIGN KEY (job_id) REFERENCES Job(id)
);

CREATE TABLE ChatbotInterview (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  job_id INT,
  status interview_status DEFAULT 'pending',
  review TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id),
  FOREIGN KEY (job_id) REFERENCES Job(id)
);

CREATE TABLE ChatbotConversation (
  id SERIAL PRIMARY KEY,
  chatbot_interview_id INT,
  sender_type VARCHAR(50) CHECK (sender_type IN ('job_seeker', 'bot')),
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (chatbot_interview_id) REFERENCES ChatbotInterview(id)
);

CREATE TABLE Messages (
  id SERIAL PRIMARY KEY,
  sender_id INT NOT NULL,
  receiver_id INT NOT NULL,
  content TEXT NOT NULL,
  sender_type VARCHAR(50) CHECK (sender_type IN ('job_seeker', 'company', 'bot')),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (sender_id) REFERENCES "User"(id),
  FOREIGN KEY (receiver_id) REFERENCES "User"(id)
);

CREATE TABLE ScheduledInterview (
  id SERIAL PRIMARY KEY,
  job_seeker_id INT,
  company_id INT,
  scheduled_at TIMESTAMP,
  location VARCHAR(255),
  notes TEXT,
  meeting_link VARCHAR(1000),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (job_seeker_id) REFERENCES JobSeeker(id),
  FOREIGN KEY (company_id) REFERENCES Company(id)
);
