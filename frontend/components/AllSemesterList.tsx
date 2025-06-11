'use client'

import { useState, useEffect } from 'react'
import { ColumnDef } from '@tanstack/react-table'
import { CourseTable } from '@/components/CourseTable'
import { Enrollment, Course } from '@/types/enrollment'

export function AllSemestersList() {
  const [enrollments, setEnrollments] = useState<Enrollment[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  // 1) Fetch ALL enrollments once
  useEffect(() => {
    setLoading(true)
    fetch('http://localhost:8888/enrollments')
      .then(res => {
        if (!res.ok) throw new Error(res.statusText)
        return res.json() as Promise<Enrollment[]>
      })
      .then(json => setEnrollments(json))
      .catch(err => setError(err.message))
      .finally(() => setLoading(false))
  }, [])

  // 2) Group courses by [year][semester]
  const byYear: Record<number, Record<number, Course[]>> = {}
  enrollments.forEach((enr) => {
    const y = enr.semester.year
    const sem = enr.semester.semester_number

    if (!byYear[y]) byYear[y] = {}
    if (!byYear[y][sem]) byYear[y][sem] = []

    // push all courses in that enrollment into the bucket
    byYear[y][sem].push(...enr.courses)
  })

  // 3) Prepare your columns once
  const columns: ColumnDef<Course, any>[] = [
    { accessorKey: 'course_code', header: 'Code' },
    { accessorKey: 'course_name', header: 'Name' },
    {
      accessorKey: 'credits',
      header: 'Credits',
      cell: info => info.getValue<number>(),
    },
  ]

  if (loading) return <p>Loading enrollment data…</p>
  if (error) return <p className="text-red-600">Error: {error}</p>

  // 4) Sort years and semesters
  const years = Object.keys(byYear)
    .map((y) => parseInt(y, 10))
    .sort((a, b) => a - b)

  return (
    <div className="space-y-12 p-6 w-6xl mx-auto pt-5">
      {years.map((yr) => {
        // get semesters for this year, sorted
        const sems = Object.keys(byYear[yr])
          .map(s => parseInt(s, 10))
          .sort((a, b) => a - b)

        return (
          <section key={yr} className="bg-gradient-to-b from-primary/2 to-card text-card-foreground rounded-md border space-y-8 p-10">
            <h1 className="text-2xl font-semibold">Year {yr}</h1>

            {sems.map((sem) => {
              const courses = byYear[yr][sem]
              return (
                <div key={sem}>
                  <h2 className="text-xl font-medium mb-2">
                    Semester {sem}
                  </h2>
                  {courses.length > 0 ? (
                    <CourseTable columns={columns} data={courses} />
                  ) : (
                    <p className="italic">
                      No courses enrolled in semester {sem}.
                    </p>
                  )}
                </div>
              )
            })}
          </section>
        )
      })}
    </div>
  )
}
