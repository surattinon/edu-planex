SELECT
    c.course_code,
    c.course_name,
    c.credits,
    c.year_offered,
    c.course_type,
    c.description
FROM
    courses AS c;

SELECT
    cc.cat_id,
    cc.name AS category_name,
    cc.credit_req
FROM
    curriculum_categories AS cc;

SELECT
    e.enrollment_id,
    s.year,
    s.semester_number,
    ec.course_code,
    c.course_name,
    c.credits,
    e.created_at
FROM
    enrollments AS e
    JOIN semesters AS s ON e.semester_id = s.semester_id
    JOIN enrollment_courses AS ec ON e.enrollment_id = ec.enrollment_id
    JOIN courses AS c ON ec.course_code = c.course_code
WHERE
    e.user_id = 1
ORDER BY
    s.year,
    s.semester_number;

SELECT
    p.plan_id,
    p.name AS plan_name,
    pc.course_code,
    c.course_name,
    c.credits
FROM
    plans AS p
    JOIN plan_courses AS pc ON p.plan_id = pc.plan_id
    JOIN courses AS c ON pc.course_code = c.course_code
WHERE
    p.user_id = 1
ORDER BY
    p.created_at DESC,
    p.plan_id;

SELECT
    pr.pre_course_code AS prerequisite_code,
    c.course_name AS prerequisite_name
FROM
    prerequisites AS pr
    JOIN courses AS c ON pr.pre_course_code = c.course_code
WHERE
    pr.course_code = 'ITE102';

UPDATE
    user_profile
SET
    display_name = 'Bas'
WHERE
    display_name = 'Surattinon';


DELETE FROM
    plan_courses
WHERE
    plan_id = 23;
