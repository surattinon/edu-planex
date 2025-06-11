'use client'
import { ProfileCard } from "@/components/profile-card"
import { Progress } from "@/components/ui/progress"

import { useEffect, useState } from "react"
import type { ApiCourse, Course } from "@/lib/types/course"
import { EnrolledCourses } from "@/components/SemesterCoursesTable"
import { mapApiCourseToCourse } from "@/lib/mapper"


export default function Home() {
  const [data, setData] = useState<Course[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetch("http://localhost:8888/courses")
      .then((res) => res.json() as Promise<ApiCourse[]>)
      .then((apiCourses) => {
        // 1) Map each ApiCourse to our Course
        const courses = apiCourses.map(mapApiCourseToCourse)
        setData(courses)
      })
      .catch((err) => {
        console.error("Failed to load courses:", err)
      })
      .finally(() => setLoading(false))
  }, [])

  if (loading) return <div className="w-screen h-screen flex flex-col align-middle">
    <h1 className="self-center">Loading courses…</h1>
  </div>
  if (!data.length) return <div>No courses found.</div>
  return (
    <main className="flex h-fit grow flex-col overflow-auto px-20 pl-30 py-10">
      <div className="mb-10 mt-10">
        <ProfileCard />
      </div>
      <div className="flex">
        <div className="w-[750px] h-[500px] bg-gradient-to-b from-primary/2 to-card text-card-foreground rounded-md border p-5 mr-10">
          <div className="h-full flex flex-col justify-start w-full max-w-md">
            <h1 className="mb-6 text-3xl">Progress</h1>
            <div className="w-full flex flex-col gap-4 justify-center mx-auto">
              <div className="w-full">
                <h3 className="text-md mb-3">General Education</h3>
                <Progress value={33} className="mb-3" />
                <h3 className="w-full text-sm text-end">20 / 100%</h3>
              </div>
              <div className="w-full">
                <h3 className="text-md mb-3">General Education</h3>
                <Progress value={33} className="mb-3" />
                <h3 className="w-full text-sm text-end">20 / 100%</h3>
              </div>
              <div className="w-full">
                <h3 className="text-md mb-3">General Education</h3>
                <Progress value={33} className="mb-3" />
                <h3 className="w-full text-sm text-end">20 / 100%</h3>
              </div>
              <div className="w-full">
                <h3 className="text-md mb-3">General Education</h3>
                <Progress value={33} className="mb-3" />
                <h3 className="w-full text-sm text-end">20 / 100%</h3>
              </div>
            </div>
          </div>
        </div>
        <div className="overflow-auto w-full h-[500px] bg-gradient-to-b from-primary/2 to-card text-card-foreground rounded-md border p-5">
          <EnrolledCourses />
        </div>
      </div>

    </main>
  );
}
