import { CourseType } from './course';

export interface Course {
  course_code: string;
  course_name: string;
  credits: number;
  year_offered: number;
  course_type: string;
  description: string;
  categories: any | null;
  prerequisites: any | null;
}

export interface Semester {
  semester_id: number;
  year: number;
  semester_number: number;
}

export interface Enrollment {
  enrollment_id: number;
  user_id: number;
  semester_id: number;
  created_at: string;
  semester: Semester;
  courses: Course[];
}

//

export interface EnrollmentType {
  enrollment_id: number;
  user_id: number;
  semester_id: number;
  created_at: string;
  semester: { semester_id: number; year: number; semester_number: number };
  courses: CourseType[];
}
