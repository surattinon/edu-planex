'use client'

import useSWR from 'swr'

import { fetcher, ProgressResponse } from '@/lib/api'
import { Progress } from '@/components/ui/progress'

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

export const StatsCard = () => {
  const { data, error } = useSWR<ProgressResponse>('/progress', fetcher)

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


  const progressClasses: Record<string, string> = {
    general_education: "bg-gradient-to-r from-blue-200 to-blue-300",
    professional: "bg-gradient-to-r from-green-200 to-green-300",
    free_elective: "bg-gradient-to-r from-yellow-200 to-yellow-300",
    internship: "bg-gradient-to-r from-red-200 to-red-300",
  }


  return (
    <Card className="p-5 h-full w-full flex flex-col bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md border">
      <CardHeader className="mt-2">
        <CardTitle className="text-3xl font-light">Category Credit Stats</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        {data.courses.map(course => {
          const percent = Math.round((course.earned / course.required) * 100)  // calculate completion percentage
          const colorClass = progressClasses[course.key] ?? "[&>*]:bg-gray-400"

          return (
            <div key={course.key} className="flex flex-col gap-1">
              <div className='flex justify-between'>
              <h1 className='capitalize'>{course.key.replace('_', ' ')}</h1>
              <h1 className="text-md text-zinc-300">{percent}%</h1>
              </div>
              <Progress value={percent} />
              <h1 className="self-end text-md text-white">{course.earned}<span className='text-sm text-zinc-400'> / {course.required} credits</span> </h1>
            </div>
          )
        })}
      </CardContent>
    </Card>
  )
}

