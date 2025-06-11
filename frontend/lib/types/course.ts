// 1) The raw shape coming from your API:
export interface ApiCourse {
  course_code: string
  course_name: string
  credits: number
  year_offered: number
  course_type: string
  description: string
  categories: { cat_id: number; name: string; credit_required: number }[]
  prerequisites: string[]
}

// 2) The shape your DataTable will use (camelCase):
export interface Course {
  courseCode: string
  courseName: string
  credits: number
  yearOffered: number
  courseType: string
  description: string
  categories: { catId: number; name: string; creditRequired: number }[]
  prerequisites: string[]
}

export interface SemesterCourses {
  courseCode: string
  courseName: string
  credits: number
}
