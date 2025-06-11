import { ApiCourse, Course } from "@/lib/types/course"

export function mapApiCourseToCourse(api: ApiCourse): Course {
  return {
    courseCode: api.course_code,
    courseName: api.course_name,
    credits: api.credits,
    yearOffered: api.year_offered,
    courseType: api.course_type,
    description: api.description,
    categories: api.categories.map(c => ({
      catId: c.cat_id,
      name: c.name,
      creditRequired: c.credit_required,
    })),
    prerequisites: api.prerequisites,
  }
}
