import { createColumnHelper } from '@tanstack/react-table';
import { Course } from '@/types/course';

const columnHelper = createColumnHelper<Course>();

export const courseColumns = [
  columnHelper.accessor('course_code', { header: 'Code' }),
  columnHelper.accessor('course_name', { header: 'Name' }),
  columnHelper.accessor('credits', { header: 'Credits' }),
  columnHelper.accessor('year_offered', { header: 'Year' }),
  columnHelper.accessor('course_type', { header: 'Type' }),
];
