-- Courses Catalog
CREATE TABLE IF NOT EXISTS courses (
  course_id   VARCHAR(10) PRIMARY KEY,
  course_name VARCHAR(100) NOT NULL,
  credits     INT         NOT NULL
);

-- Curricula
CREATE TABLE IF NOT EXISTS curricula (
  curriculum_id SERIAL PRIMARY KEY,
  name           VARCHAR(100) NOT NULL,
  version        VARCHAR(10)  NOT NULL,
  effective_date DATE         NOT NULL,
  retired_date   DATE
);

-- Course Category
CREATE TABLE IF NOT EXISTS course_categories (
  category_id   SERIAL PRIMARY KEY,
  curriculum_id INT    NOT NULL REFERENCES curricula(curriculum_id),
  name          VARCHAR(50) NOT NULL
);

-- Curriculum-Course link
CREATE TABLE IF NOT EXISTS curriculum_courses (
  id             SERIAL    PRIMARY KEY,
  category_id    INT       NOT NULL REFERENCES course_categories(category_id),
  course_id      VARCHAR(10) NOT NULL REFERENCES courses(course_id),
  year_offered   INT       NOT NULL,
  prerequisite   VARCHAR(10)
);

-- Users --
CREATE TABLE IF NOT EXISTS advisors (
  advisor_id VARCHAR(20) PRIMARY KEY,
  fname      VARCHAR(50) NOT NULL,
  lname      VARCHAR(50) NOT NULL,
  pass_hash  VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS students (
  student_id VARCHAR(20) PRIMARY KEY,
  fname      VARCHAR(50) NOT NULL,
  lname      VARCHAR(50) NOT NULL,
  pass_hash  VARCHAR(255) NOT NULL,
  advisor_id VARCHAR(20)  NOT NULL REFERENCES advisors(advisor_id)
);

-- Terms
CREATE TABLE IF NOT EXISTS terms (
  term_id    SERIAL      PRIMARY KEY,
  term_name  VARCHAR(50) NOT NULL,
  start_date DATE        NOT NULL,
  end_date   DATE        NOT NULL
);

-- Study Plan
CREATE TABLE IF NOT EXISTS study_plans (
  plan_id    SERIAL      PRIMARY KEY,
  student_id VARCHAR(20) NOT NULL REFERENCES students(student_id),
  term_id    INT         NOT NULL REFERENCES terms(term_id),
  status     VARCHAR(20) NOT NULL,
  created_at TIMESTAMP   DEFAULT NOW(),
  updated_at TIMESTAMP   DEFAULT NOW()
);

-- Sections
CREATE TABLE IF NOT EXISTS sections (
  section_id  SERIAL      PRIMARY KEY,
  course_id   VARCHAR(10) NOT NULL REFERENCES courses(course_id),
  term_id     INT         NOT NULL REFERENCES terms(term_id),
  section_no  INT         NOT NULL,
  schedule    VARCHAR(100),
  capacity    INT,
  instructor  VARCHAR(100)
);

-- Enrollments
CREATE TABLE IF NOT EXISTS enrollments (
  enroll_id   SERIAL      PRIMARY KEY,
  plan_id     INT         NOT NULL REFERENCES study_plans(plan_id),
  section_id  INT         NOT NULL REFERENCES sections(section_id),
  status      VARCHAR(20),
  grade       VARCHAR(5)
);

