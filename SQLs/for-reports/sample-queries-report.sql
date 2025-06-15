select
    c.course_code,
    c.course_name,
    c.credits,
    c.year_offered,
    c.course_type,
    c.description
from courses as c
;


select cc.cat_id, cc.name as category_name, cc.credit_req
from curriculum_categories as cc
;


select
    e.enrollment_id,
    s.year,
    s.semester_number,
    ec.course_code,
    c.course_name,
    c.credits,
    e.created_at
from enrollments as e
join semesters as s on e.semester_id = s.semester_id
join enrollment_courses as ec on e.enrollment_id = ec.enrollment_id
join courses as c on ec.course_code = c.course_code
where e.user_id = 1
order by s.year, s.semester_number
;


select p.plan_id, p.name as plan_name, pc.course_code, c.course_name, c.credits
from plans as p
join plan_courses as pc on p.plan_id = pc.plan_id
join courses as c on pc.course_code = c.course_code
where p.user_id = 1
order by p.created_at desc, p.plan_id
;


select pr.pre_course_code as prerequisite_code, c.course_name as prerequisite_name
from prerequisites as pr
join courses as c on pr.pre_course_code = c.course_code
where pr.course_code = 'ITE102'
;

