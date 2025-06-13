"use client"

import useSWR from "swr"
import { fetcher } from "@/lib/api"
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import {
  Table, TableHeader, TableBody,
  TableRow, TableHead, TableCell
} from "@/components/ui/table"
import React, { useMemo } from "react"

import { Enrollment } from "@/lib/api"

export function AllEnrollmentTables() {
  // 1) Hooks at the top level – always executed on every render
  const { data, error } = useSWR<Enrollment[]>(
    "/enrollhist",
    fetcher
  );

  const sorted = useMemo(() => {
    if (!data) return []; // handle missing data
    return [...data].sort((a, b) => {
      if (b.semester.year !== a.semester.year) {
        return b.semester.year - a.semester.year;
      }
      return b.semester.number - a.semester.number;
    });
  }, [data]);

  const groupedByYear = useMemo(() => {
    const groups: Record<number, Enrollment[]> = {};
    for (const enrollment of sorted) {
      const yr = enrollment.semester.year;
      if (!groups[yr]) groups[yr] = [];
      groups[yr].push(enrollment);
    }
    return groups;
  }, [sorted]);

  // 2) Conditional rendering based on state
  if (error) {
    return (
      <Card className="p-5">
        <CardHeader>
          <CardTitle>Error</CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-red-600">Failed to load enrollment history.</p>
        </CardContent>
      </Card>
    );
  }

  if (!data) {
    return (
      <Card className="p-5">
        <CardHeader>
          <CardTitle>Loading…</CardTitle>
        </CardHeader>
      </Card>
    );
  }

  // 3) Render all semester tables grouped by year
  return (
    <div className="w-full flex flex-col gap-7">
      {Object.entries(groupedByYear).map(([year, enrollments]) => (
        <section key={year} className="bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md border mb-7 p-5">
          <h2 className="text-2xl font-semibold mb-4">Year {year}</h2>
          <div className="flex flex-col md:grid-cols-3 gap-6">
            {enrollments.map((enrollment) => (
              <Card key={enrollment.enrollment_id} className="h-full bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md">
                <CardHeader>
                  <CardTitle className="text-lg">
                    Semester {enrollment.semester.number}/{year}
                  </CardTitle>
                </CardHeader>
                <CardContent className="overflow-x-auto">
                  <Table className="w-full">
                    <TableHeader>
                      <TableRow>
                        <TableHead>Course Code</TableHead>
                        <TableHead>Course Name</TableHead>
                        <TableHead className="text-right">
                          Credits
                        </TableHead>
                      </TableRow>
                    </TableHeader>
                    <TableBody>
                      {enrollment.courses.map((course) => (
                        <TableRow key={course.course_code}>
                          <TableCell>{course.course_code}</TableCell>
                          <TableCell>{course.course_name}</TableCell>
                          <TableCell className="text-right">
                            {course.credits}
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </CardContent>
              </Card>
            ))}
          </div>
        </section>
      ))}
    </div>
  );
}
