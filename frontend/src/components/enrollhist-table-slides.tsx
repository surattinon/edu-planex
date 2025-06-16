"use client"

import { useState, useEffect, useMemo } from "react"
import useSWR from "swr"
import { fetcher } from "@/lib/api"
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
} from "@/components/ui/card"
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from "@/components/ui/table"
import { Button } from "@/components/ui/button"

import { Enrollment } from "@/lib/api"

export function EnrollHistSlide() {
  // 1) Fetch the raw array
  const { data, error } = useSWR<Enrollment[]>(
    "/enrollhist",
    fetcher
  )

  // 2) Sort descending by year, then semester number
  const sorted = useMemo(() => {
    if (!data) return []
    return [...data].sort((a, b) => {
      if (b.semester.year !== a.semester.year) {
        return b.semester.year - a.semester.year
      }
      return b.semester.number - a.semester.number
    })
  }, [data])

  // 3) Track which index in `sorted` is “current”
  const [idx, setIdx] = useState(0)

  // On data load, default to the most-recent (idx = 0 after sorting)
  useEffect(() => {
    if (sorted.length) {
      setIdx(0)
    }
  }, [sorted])

  if (error) return <>
    <Card className="h-full w-full flex flex-col items-center justify-center bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md border">
      <p className='text-red-500 font-light'>Failed to load contents</p>
    </Card>
  </>


  if (!data) return <>
    <Card className="h-full w-full flex flex-col items-center justify-center bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md border">
      <p className='font-light'>Loading...</p>
    </Card>
  </>

  const current = sorted[idx]
  const { year, number } = current.semester

  return (
    <Card className="p-7 h-full w-full flex flex-col bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md border">
      <CardHeader className="flex items-center justify-between">
        <CardTitle className="text-3xl font-light">
          Enrollment History <span className="ml-3 text-lg text-zinc-300">Semester {number}/{year}</span>
        </CardTitle>

        <div className="space-x-2">
          <Button
            size="sm"
            variant="outline"
            onClick={() => setIdx(i => Math.min(i + 1, sorted.length - 1))}
            disabled={idx >= sorted.length - 1}
          >
            Previous
          </Button>
          <Button
            size="sm"
            variant="outline"
            onClick={() => setIdx(i => Math.max(i - 1, 0))}
            disabled={idx <= 0}
          >
            Next
          </Button>
        </div>
      </CardHeader>

      <CardContent className="mt-14">
        <div className="rounded-md border bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl ">
          <Table className="rounded-md">
            <TableHeader className="">
              <TableRow>
                <TableHead>Course Code</TableHead>
                <TableHead>Course Name</TableHead>
                <TableHead className="text-right">Credits</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {current.courses.map(course => (
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
        </div>
      </CardContent>
    </Card>
  )
}
