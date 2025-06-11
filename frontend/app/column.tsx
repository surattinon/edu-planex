"use client"

import { ColumnDef } from "@tanstack/react-table"
import { Course, SemesterCourses } from "@/lib/types/course"

export const columns: ColumnDef<Course>[] = [
  { accessorKey: "courseCode", header: "Code" },
  { accessorKey: "courseName", header: "Name" },
  { accessorKey: "credits", header: "Credits" },
  { accessorKey: "yearOffered", header: "Year" },
]


export const semesterColumns: ColumnDef<SemesterCourses>[] = [
  { accessorKey: "courseCode", header: "Code" },
  { accessorKey: "courseName", header: "Name" },
  { accessorKey: "credits", header: "Credits" },
]
