'use client'

import useSWR from 'swr'

import { fetcher, ProgressResponse } from '@/lib/api'

import { CircularProgress } from "@heroui/progress";

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

export const CredProgCard = () => {
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

  const totalRequired = data.courses.reduce((sum, course) => sum + course.required, 0)
  const totalEarned = data.courses.reduce((sum, course) => sum + course.earned, 0)

  // Optional: overall completion percentage
  const overallPercent = Math.round((totalEarned / totalRequired) * 100)

  return (
    <Card className="p-5 h-full w-full flex flex-col bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md border">
      <CardHeader className="mt-5">
        <CardTitle className="text-3xl font-light text-center">Overall Credits</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-1">
        <div className="flex flex-col gap-2">
          <CircularProgress
            classNames={{
              svg: "w-58 h-58 drop-shadow-md",
              indicator: "stroke-green-400 stroke-2",
              track: "stroke-white/10 stroke-2",
              value: "text-4xl font-light text-white",
            }}
            showValueLabel={true}
            strokeWidth={4}
            value={overallPercent}
          />
        </div>
        <h1 className='w-full text-4xl font-light my-3 text-center'>{totalEarned} <span className='text-xl text-zinc-400'> / {totalRequired} credits</span></h1>
      </CardContent>
    </Card>
  )
}

