"use client";

import React, { useMemo } from "react";
import useSWR from "swr";
import { fetcher } from "@/lib/api";
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from "@/components/ui/table";
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
} from "@/components/ui/card";

import { Curriculum } from "@/lib/api";

export function CurriculumTable2() {
  const { data, error } = useSWR<Curriculum[]>("/curriculumtable", fetcher);

  if (error) {
    return <p className="text-red-600">Failed to load curriculum.</p>;
  }
  if (!data) {
    return <p>Loading curriculum…</p>;
  }

  const curriculum = data[0];
  const { categories } = curriculum;

  return (
    <div className="w-full flex flex-col gap-5 mb-5 justify-center">
      {categories.map((cat) => (
        <section key={cat.id} className="">
          <Card className="w-full h-full bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl">
            <CardHeader className="">
              <CardTitle className="text-2xl text-center font-light">{cat.name} ( Required {cat.credit_required} Credits )</CardTitle>
            </CardHeader>
            <CardContent className="space-y-8">
              {cat.course_types.map((ct) => (
                <div key={ct.id}>
                  <h3 className="text-xl font-light mb-3">
                    {ct.name}
                  </h3>
                  <div className="overflow-x-auto border rounded-md">
                    <Table className="min-w-full">
                      <TableHeader className="bg-gradient-to-b from-zinc-950 to-card">
                        <TableRow className="">
                          <TableHead>Code</TableHead>
                          <TableHead>Name</TableHead>
                          <TableHead>Prerequisites</TableHead>
                          <TableHead>Credits</TableHead>
                        </TableRow>
                      </TableHeader>
                      <TableBody>
                        {ct.courses.map((course) => (
                          <TableRow key={course.code}>
                            <TableCell>{course.code}</TableCell>
                            <TableCell>{course.name}</TableCell>
                            <TableCell>
                              {course.prerequisites.length > 0
                                ? course.prerequisites.join(", ")
                                : ""}
                            </TableCell>
                            <TableCell>{course.credits}</TableCell>
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
}
