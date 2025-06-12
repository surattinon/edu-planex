'use client'

import { useEffect, useState } from 'react'
import { ColumnDef } from '@tanstack/react-table'
import { fetchPlans, fetchEnrollments, Plan, Enrollment } from '@/lib/api'
import { PlanSection } from '@/components/PlanSection'

export function PlanList() {
  const [plans, setPlans] = useState<Plan[]>([])
  const [enrollments, setEnrollments] = useState<Enrollment[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  // load data
  useEffect(() => {
    setLoading(true)
    Promise.all([fetchPlans(), fetchEnrollments()])
      .then(([p, e]) => {
        setPlans(p)
        setEnrollments(e)
      })
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false))
  }, [])


  // table columns
  const columns: ColumnDef<Plan['courses'][number], any>[] = [
    { accessorKey: 'course_code', header: 'Course Code' },
    { accessorKey: 'course_name', header: 'Course Name' },
    {
      accessorKey: 'credits',
      header: 'Credits',
      cell: (info) => info.getValue<number>(),
    },
  ]

  // helper: mark applied plans
  const enrolledCodes = new Set(
    enrollments.flatMap((e) => e.courses.map((c) => c.course_code))
  )
  const isApplied = (plan: Plan) =>
    plan.courses.every((c) => enrolledCodes.has(c.course_code))

  if (loading) return <p>Loading…</p>
  if (error) return <p className="text-red-600">Error: {error}</p>

  const notApplied = plans.filter((p) => !isApplied(p))
  const applied    = plans.filter(isApplied)

  return (
    <div className="space-y-8 p-6 w-6xl mx-auto">
      {notApplied.length > 0 && (
        <>
          <h2 className="text-2xl font-semibold">Plans to Apply</h2>
          {notApplied.map((plan) => (
            <PlanSection
              key={plan.plan_id}
              plan={plan}
              columns={columns}
              applied={false}
              onApplied={() => {
                // move it over to applied
                setPlans((prev) =>
                  prev.map((p) =>
                    p.plan_id === plan.plan_id ? { ...p } : p
                  )
                )
                setEnrollments((prev) => [
                  ...prev,
                  { ...({} as Enrollment), courses: plan.courses },
                ])
              }}
            />
          ))}
        </>
      )}

      {applied.length > 0 && (
        <>
          <h2 className="text-2xl font-semibold">Already Applied</h2>
          {applied.map((plan) => (
            <PlanSection
              key={plan.plan_id}
              plan={plan}
              columns={columns}
              applied={true}
            />
          ))}
        </>
      )}
    </div>
  )
}
