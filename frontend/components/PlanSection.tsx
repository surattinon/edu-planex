'use client'

import { useState } from 'react'
import { ColumnDef } from '@tanstack/react-table'
import { CourseTable } from '@/components/CourseTable'
import { deletePlan, applyPlan, Plan } from '@/lib/api'


// shadcn components
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from '@/components/ui/select'

interface PlanSectionProps {
  plan: Plan
  columns: ColumnDef<Plan['courses'][number], any>[]
  applied: boolean
  onApplied?: () => void
}

export function PlanSection({
  plan,
  columns,
  applied,
  onApplied,
}: PlanSectionProps) {
  const [open, setOpen] = useState(false)
  const [openDel, setOpenDel] = useState(false)
  const [selYear, setSelYear] = useState(new Date().getFullYear())
  const [selSemester, setSelSemester] = useState(1)

  const years = Array.from({ length: 2026 - 2021 + 1 }, (_, i) => 2021 + i)

  return (
    <>
      <section className={`p-4 rounded-md space-y-4 ${applied ? 'border-none text-card-foreground/40 bg-card' : 'border-1 border-white/30 text-card-foreground bg-gradient-to-b from-primary/2 to-card shadow-[0_0px_30px_rgba(255,_255,_255,_0.1)]'}`}>
        <h3 className="text-2xl font-bold">Plan Name : {plan.plan_name}</h3>
        <CourseTable columns={columns} data={plan.courses} />

        <div className='flex gap-5'>
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button disabled={applied} className='bg-primary/10 text-card-foreground border hover:bg-primary/20'>
                {applied ? 'Applied' : 'Apply Plan'}
              </Button>
            </DialogTrigger>

            <DialogContent>
              <DialogHeader>
                <DialogTitle>Apply “{plan.plan_name}”</DialogTitle>
              </DialogHeader>

              <div className="grid gap-4 py-4">
                <div>
                  <label className="block text-sm font-medium">Year</label>
                  <Select
                    value={String(selYear)}
                    onValueChange={(v) => setSelYear(Number(v))}
                  >
                    <SelectTrigger className="w-full">
                      <SelectValue placeholder="Select year" />
                    </SelectTrigger>
                    <SelectContent>
                      {years.map((y) => (
                        <SelectItem key={y} value={String(y)}>
                          {y}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>

                <div>
                  <label className="block text-sm font-medium">Semester</label>
                  <Select
                    value={String(selSemester)}
                    onValueChange={(v) => setSelSemester(Number(v))}
                  >
                    <SelectTrigger className="w-full">
                      <SelectValue placeholder="Select semester" />
                    </SelectTrigger>
                    <SelectContent>
                      {[1, 2, 3].map((s) => (
                        <SelectItem key={s} value={String(s)}>
                          {s}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
              </div>

              <DialogFooter className="flex justify-end gap-2">
                <Button variant="outline" onClick={() => setOpen(false)}>
                  Cancel
                </Button>
                <Button
                  onClick={async () => {
                    await applyPlan(plan.plan_id, selYear, selSemester)
                    setOpen(false)
                    onApplied?.()
                  }}
                  disabled={applied}
                >
                  Confirm Apply
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
          {!applied && (
            <Dialog open={openDel} onOpenChange={setOpenDel}>
              <DialogTrigger asChild>
                <Button className='bg-red-900 text-card-foreground border hover:bg-red-800'>
                  Delete
                </Button>
              </DialogTrigger>

              <DialogContent>
                <DialogHeader>
                  <DialogTitle>Delete “{plan.plan_name}”</DialogTitle>
                </DialogHeader>
                <DialogFooter className="flex justify-end gap-2">
                  <Button variant="outline" onClick={() => setOpen(false)}>
                    Cancel
                  </Button>
                  <Button
                    onClick={async () => {
                      await deletePlan(plan.plan_id)
                      setOpenDel(false)
                    }}
                  >
                    Confirm Delete
                  </Button>
                </DialogFooter>
              </DialogContent>
            </Dialog>
          )}
        </div>
      </section>
    </>
  )
}
