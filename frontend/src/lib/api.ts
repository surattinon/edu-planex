import axios from 'axios';

const API_BASE_URL = 'http://localhost:8888/api/v1'

export const axiosInstance = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  }
});

// SWR fetcher
export const fetcher = (url: string) =>
  axiosInstance.get(url).then(res => res.data)

// Course Progress respond
export interface CourseProgress {
  key: string
  required: number
  earned: number
}
export interface ProgressResponse {
  user_id: number
  courses: CourseProgress[]
}

// Enrollment History list respond
export interface Course {
  course_code: string
  course_name: string
  credits: number
}

export interface Enrollment {
  enrollment_id: number
  user_id: number
  semester: {
    year: number
    number: number
  }
  courses: Course[]
}

export interface Category {
  cat_id: number,
  name: string,
  credits: number
}

export interface Prerequisite {
  course_code: string;
  pre_course_code: string;
}

export interface CurriculumCourse {
  course_code: string;
  course_name: string;
  credits: number;
  year_offered: number;
  course_type: string;
  description: string;
  categories: Category[];
  prerequisites: Prerequisite[];
}


export interface Courses {
  code: string;
  name: string;
  desc: string;
  credits: number;
  prerequisites: string[];
}

export interface CourseTypes {
  id: number;
  name: string;
  courses: Courses[];
}

export interface Categories {
  id: number;
  name: string;
  credit_required: number;
  course_types: CourseTypes[];
}

export interface Curriculum {
  id: number;
  name: string;
  categories: Categories[];
}
