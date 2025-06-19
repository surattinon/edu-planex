INSERT INTO
    curriculum_categories (cat_id, name, credit_req)
VALUES
    (1, 'General Education Courses', 40),
    (2, 'Professional Courses', 100),
    (3, 'Free Electives', 8),
    (4, 'Internship', 12);

INSERT INTO
    courses (
        course_code,
        course_name,
        credits,
        year_offered,
        course_type,
        description
    )
VALUES
    (
        'ART101',
        'Art Appreciation',
        4,
        1,
        'General Education',
        'Art Appreciation'
    ),
    (
        'ART102',
        'Film Appreciation',
        4,
        1,
        'General Education',
        'Film Appreciation'
    ),
    (
        'ATH101',
        'Introduction to Cultural Anthropology',
        4,
        1,
        'General Education',
        'Introduction to Cultural Anthropology'
    ),
    (
        'STA101',
        'Introduction to Statistics (Pre: MAT101)',
        4,
        1,
        'General Education',
        'Introduction to Statistics (Pre: MAT101)'
    ),
    (
        'ITE101',
        'Information Technology Fundamentals',
        4,
        1,
        'Basic Core',
        'Information Technology Fundamentals'
    ),
    (
        'ITE102',
        'Discrete Mathematics Structure (Pre: MAT101)',
        4,
        1,
        'Basic Core',
        'Discrete Mathematics Structure (Pre: MAT101)'
    ),
    (
        'ITE103',
        'Introduction to Data Structure and Algorithms Analysis',
        4,
        3,
        'Basic Core',
        'Introduction to Data Structure and Algorithms Analysis'
    ),
    (
        'ITE104',
        'Computer Organization',
        4,
        3,
        'Basic Core',
        'Computer Organization'
    ),
    (
        'ITE254',
        'Human Computer Interaction',
        4,
        2,
        'General Education',
        'Introduces students to analysis, design, and evaluation of the interaction between user and technologies'
    ),
    (
        'FREE001',
        'Free Elective Placeholder 1',
        4,
        0,
        'Free Electives',
        'Free Elective Placeholder 1'
    ),
    (
        'FREE002',
        'Free Elective Placeholder 2',
        4,
        0,
        'Free Electives',
        'Free Elective Placeholder 2'
    ),
    (
        'INTRNSHP',
        'Internship',
        12,
        0,
        'Internship',
        'Internship'
    );

INSERT INTO
    curriculum_courses (curriculum_cat_id, course_code)
VALUES
    (1, 'ART101'),
    (1, 'ART102'),
    (1, 'ATH101'),
    (1, 'STA101'),
    (2, 'ITE101'),
    (2, 'ITE102'),
    (2, 'ITE103'),
    (2, 'ITE104'),
    (2, 'ITE254'),
    (3, 'FREE001'),
    (3, 'FREE002'),
    (4, 'INTRNSHP');

INSERT INTO
    prerequisites (course_code, pre_course_code)
VALUES
    ('STA101', 'MAT101'),
    ('ITE102', 'MAT101');

INSERT INTO
    semesters (year, semester_number)
VALUES
    (2021, 1),
    (2021, 2),
    (2021, 3),
    (2022, 1),
    (2022, 2),
    (2022, 3);

INSERT INTO
    user_profile (display_name, email)
VALUES
    (1, "USERNAME", "xxxxxx@students.stamford.edu");

INSERT INTO
    plans (plan_id, user_id, name)
VALUES
    (1, 1, "Test Plan 1"),
    (2, 1, "Test Plan 2");

INSERT INTO
    plan_courses (plan_id, course_code)
VALUES
    (1, "ART101"),
    (1, "ART102"),
    (1, "ITE101"),
    (2, "ITE102"),
    (2, "ITE103"),
    (2, "ITE104");

INSERT INTO
    enrollments (user_id, semesters_id)
VALUES
    (1, 1);

INSERT INTO
    enrollment_courses (enrollment_id, course_code)
VALUES
    (1, "ART101"),
    (1, "ART102"),
    (1, "ITE101");
