import { CourseType } from './course';

export interface Plan {
  id: number
  name: string
  created_at: string
  status: 'draft' | 'applied'
  applied_year?: number
  applied_semester?: number
}

//

export interface PlanType {
  plan_id: number;
  plan_name: string;
  user_id: number;
  courses: CourseType[];
}
