"use client"

import { ColumnDef, createColumnHelper } from "@tanstack/react-table"

export type Plans = {
  id: number
  name: string
  create_at: string
  courses: {course_code: string, course_name: string, credits: number}[]
}

export type Courses = Plans['courses'][number];

const columnHelper = createColumnHelper<Courses>();

export const courseColumns = [
  columnHelper.accessor('course_code',   { header: 'Course Code'  }),
  columnHelper.accessor('course_name',   { header: 'Course Name'  }),
  columnHelper.accessor('credits',       { header: 'Credits'      }),
] as Array<ColumnDef<Courses, unknown>>;
