SELECT
  courses.course_code AS "Course Code",
  courses.course_name AS "Course Name",
  courses.credits AS "Credits",
  prerequisites.pre_course_code AS "Prerequisite"
FROM courses
LEFT JOIN prerequisites
  ON courses.course_code = prerequisites.course_code
ORDER BY courses.course_code, prerequisites.pre_course_code;
