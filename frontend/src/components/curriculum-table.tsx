"use client";

import React from "react";
import useSWR from "swr";
import { fetcher } from "@/lib/api";
import { Card, CardTitle, CardHeader, CardContent } from "./ui/card";
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from "@/components/ui/table";
import { useMemo } from "react";

import { CurriculumCourse } from "@/lib/api";

export function CurriculumTable() {
  const { data, error } = useSWR<CurriculumCourse[]>(
    "/curriculum",
    fetcher
  );

  // 2) Build category → type → [courses] map via reduce :contentReference[oaicite:6]{index=6}
  const grouped = useMemo(() => {
    if (!data) return {} as Record<string, Record<string, CurriculumCourse[]>>;

    return data.reduce<Record<string, Record<string, CurriculumCourse[]>>>(
      (acc, course) => {
        // Each course may belong to multiple categories; group under each
        course.categories.forEach((cat) => {
          const catName = cat.name; // e.g. "General Education Courses"
          acc[catName] ||= {};
          const typeGroup = acc[catName]!;
          const type = course.course_type; // e.g. "General Education"
          typeGroup[type] ||= [];
          typeGroup[type]!.push(course);
        });
        return acc;
      },
      {}
    );
  }, [data]);

  // 3) Early return for loading & error states
  if (error) {
    return <p className="text-red-600">Failed to load curriculum.</p>;
  }
  if (!data) {
    return <p>Loading curriculum…</p>;
  }

  // 4) Render each Category as a section, each Course Type as a sub-table
  return (
    <div className="space-y-12">
      {Object.entries(grouped).map(([categoryName, types]) => (
        <section key={categoryName}>
          <Card className="p-4">
            <CardHeader>
              <CardTitle className="text-2xl">
                {categoryName} ({typesPrimaryCredit(types)} Credits)
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              {Object.entries(types).map(([courseType, courses]) => (
                <div key={courseType}>
                  <h3 className="text-xl font-semibold mb-2">
                    {courseType}
                  </h3>
                  <div className="overflow-x-auto">
                    <Table className="min-w-full">
                      <caption className="sr-only">
                        {categoryName} – {courseType}
                      </caption>
                      <TableHeader>
                        <TableRow>
                          <TableHead>Course Code</TableHead>
                          <TableHead>Course Name</TableHead>
                          <TableHead className="text-center">Year{" "}
                            <span className="sr-only">(Offered)</span>
                          </TableHead>
                          <TableHead className="text-right">Credits</TableHead>
                          <TableHead>Description</TableHead>
                          <TableHead>Prerequisites</TableHead>
                        </TableRow>
                      </TableHeader>
                      <TableBody>
                        {courses.map((course) => (
                          <TableRow key={course.course_code}>
                            <TableCell>{course.course_code}</TableCell>
                            <TableCell>{course.course_name}</TableCell>
                            <TableCell className="text-center">
                              {course.year_offered}
                            </TableCell>
                            <TableCell className="text-right">
                              {course.credits}
                            </TableCell>
                            <TableCell>{course.description}</TableCell>
                            <TableCell>
                              {course.prerequisites.length > 0
                                ? course.prerequisites
                                    .map((p) => p.pre_course_code)
                                    .join(", ")
                                : "—"}
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </div>
                </div>
              ))}
            </CardContent>
          </Card>
        </section>
      ))}
    </div>
  );

  /** Helper: Sum required credits for the first Category entry */
  function typesPrimaryCredit(
    types: Record<string, CurriculumCourse[]>
  ): number {
    // The API's categories[].credit_required is consistent within each Category
    const firstCourse = Object.values(types)[0]?.[0];
    return firstCourse?.categories[0].credits ?? 0;
  }
}
