// app/enrollment-history/EnrolledCourses.tsx
'use client'

import { useState, useEffect } from 'react'
import { ColumnDef } from '@tanstack/react-table'
import { CourseTable } from '@/components/CourseTable'
import { Enrollment, Course } from '@/types/enrollment'

const YEAR_INIT = 2021
const MAX_SEMESTERS = 3

export function EnrolledCourses() {
  const [year, setYear] = useState(() => new Date().getFullYear() + 1)
  const [semester, setSemester] = useState(1)

  const [courses, setCourses] = useState<Course[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    setLoading(true)
    fetch(
      `http://localhost:8888/enrollhistory`
      + `?year=${year}&semester=${semester}`
    )
      .then(res => {
        if (!res.ok) throw new Error(res.statusText)
        return res.json() as Promise<Enrollment[]>
      })
      .then(enrollments => {
        // flatten all courses into one array
        setCourses(
          enrollments.flatMap(e => e.courses)
        )
      })
      .catch(err => setError(err.message))
      .finally(() => setLoading(false))
  }, [year, semester])

  // --- 3) define the three columns ---
  const columns = [
    {
      accessorKey: 'course_code',
      header: 'Code',
    },
    {
      accessorKey: 'course_name',
      header: 'Name',
    },
    {
      accessorKey: 'credits',
      header: 'Credits',
      cell: info => info.getValue<number>(),
    },
  ] as ColumnDef<Course, any>[]

  return (
    <div className="p-4 space-y-6">
      <h2 className="text-2xl mb-20">
        Enrollment History — Year {year} Sem {semester}
      </h2>

      {/* ─── Prev / Next controls ─────────────────────────── */}
      <div className="flex items-center gap-2">
        <button
          onClick={() => {
            if (semester > 1) {
              setSemester(s => s - 1)
            } else {
              setSemester(MAX_SEMESTERS)
              setYear(y => y - 1)
            }
          }}
          disabled={year === YEAR_INIT && semester === 1}
          className="px-3 py-1 border rounded disabled:opacity-50"
        >
          ◀ Previous
        </button>

        <button
          onClick={() => {
            if (semester < MAX_SEMESTERS) {
              setSemester(s => s + 1)
            } else {
              setSemester(1)
              setYear(y => y + 1)
            }
          }}
          className="px-3 py-1 border rounded"
        >
          Next ▶
        </button>
      </div>

      {/* ─── error/loading states ─────────────────────────── */}
      {loading && <p>Loading courses…</p>}
      {error && <p className="text-red-600">Error: {error}</p>}

      {/* ─── the actual table ─────────────────────────────── */}
      {!loading && !error && (
        <CourseTable columns={columns} data={courses} />
      )}
    </div>
  )
}
