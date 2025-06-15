CREATE TABLE curriculum_categories (
  cat_id      SERIAL       PRIMARY KEY,
  name        VARCHAR(100) NOT NULL UNIQUE,
  credit_req  INT          NOT NULL
);

CREATE TABLE courses (
  course_code    VARCHAR(20)  PRIMARY KEY,
  course_name    VARCHAR(255) NOT NULL,
  credits        INT          NOT NULL,
  year_offered   INT          NOT NULL,
  course_type    VARCHAR(100) NOT NULL,
  description    TEXT
);

CREATE TABLE curriculum_courses (
  curriculum_cat_id  INT           NOT NULL,
  course_code        VARCHAR(20)   NOT NULL,
  PRIMARY KEY (curriculum_cat_id, course_code),
  FOREIGN KEY (curriculum_cat_id)
    REFERENCES curriculum_categories(cat_id)
    ON DELETE CASCADE,
  FOREIGN KEY (course_code)
    REFERENCES courses(course_code)
    ON DELETE CASCADE
);

CREATE TABLE prerequisites (
  course_code      VARCHAR(20) NOT NULL,
  pre_course_code  VARCHAR(20) NOT NULL,
  PRIMARY KEY (course_code, pre_course_code),
  FOREIGN KEY (course_code)
    REFERENCES courses(course_code)
    ON DELETE CASCADE,
  FOREIGN KEY (pre_course_code)
    REFERENCES courses(course_code)
    ON DELETE CASCADE
);

CREATE TABLE user_profile (
  user_id       SERIAL        PRIMARY KEY,
  display_name  VARCHAR(100)  NOT NULL,
  email         VARCHAR(255),
  avatar_url    TEXT,
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE plans (
  plan_id   SERIAL    PRIMARY KEY,
  user_id   INT       NOT NULL
             REFERENCES user_profile(user_id) 
             ON DELETE CASCADE,
  name      VARCHAR(100) NOT NULL DEFAULT 'My Plan',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE plan_courses (
  plan_id     INT        NOT NULL
               REFERENCES plans(plan_id)
               ON DELETE CASCADE,
  course_code VARCHAR(20) NOT NULL
               REFERENCES courses(course_code)
               ON DELETE CASCADE,
  PRIMARY KEY (plan_id, course_code)
);
