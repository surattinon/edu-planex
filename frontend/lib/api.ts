export type Course = {
  course_code: string;
  course_name: string;
  credits: number;
};

export type Plan = {
  plan_id: number
  plan_name: string
  user_id: number
  courses: Array<{
    course_code: string
    course_name: string
    credits: number
  }>
}

export type Enrollment = {
  enrollment_id: number;
  user_id: number;
  semester_id: number;
  created_at: string;
  semester: {
    semester_id: number;
    year: number;
    semester_number: number;
  };
  courses: Course[];
};

const API_BASE = 'http://localhost:8888';

export async function fetchAllCourses(): Promise<Course[]> {
  const res = await fetch(`${API_BASE}/courses`);
  if (!res.ok) throw new Error('Failed to fetch courses');
  return res.json();
}

export async function fetchPlans(): Promise<Plan[]> {
  const res = await fetch(`${API_BASE}/plantable`);
  if (!res.ok) throw new Error('Failed to fetch plans');
  return res.json();
}

export async function fetchEnrollments(): Promise<Enrollment[]> {
  const res = await fetch(`${API_BASE}/enrollments`);
  if (!res.ok) throw new Error('Failed to fetch enrollments');
  return res.json();
}

export async function savePlanDraft(draft: { name: string; course_codes: string[] }): Promise<void> {
  const res = await fetch(`${API_BASE}/plans`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(draft),
  });
  if (!res.ok) throw new Error('Failed to save draft');
}

export async function applyPlan(planId: number, year: number, semester_no: number): Promise<void> {
  const res = await fetch(`${API_BASE}/plan/${planId}/apply`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ year, semester_no }),
  });
  if (!res.ok) throw new Error('Failed to apply plan');
}
