select
    courses.course_code as "Course Code",
    courses.course_name as "Course Name",
    curriculum_categories.name as "Category"
from curriculum_categories
left join
    curriculum_courses
    on curriculum_courses.curriculum_cat_id = curriculum_categories.cat_id
left join courses on courses.course_code = curriculum_courses.course_code
order by curriculum_categories.cat_id, courses.course_code
;

