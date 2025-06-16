"use client";

import React, { useState } from "react";
import useSWR, { mutate } from "swr";
import { axiosInstance, fetcher } from "@/lib/api";
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
} from "@/components/ui/card";
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

interface Course {
  course_code: string;
  course_name: string;
  credits: number;
}

interface Plan {
  plan_id: number;
  name: string;
  user_id: number;
  is_apply: boolean;
  courses: Course[];
}

export function PlanList() {
  const { data, error } = useSWR<Plan[]>("/planlist", fetcher);
  const [planId, setPlanId] = useState<number | null>(null);
  const [year, setYear] = useState<number>(new Date().getFullYear());
  const [semesterNo, setSemesterNo] = useState<number>(1);
  const [isApplyPop, setIsApplyPop] = useState(false)
  const [isDeletePop, setIsDeletePop] = useState(false)
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  const onApplyClick = (planId: number) => {
    setPlanId(planId);
    setIsApplyPop(true)
    setYear(new Date().getFullYear());
    setSemesterNo(1);
  };


  const onDeleteClick = (planId: number) => {
    setPlanId(planId);
    setIsDeletePop(true)
  };

  if (error) return <p className="text-red-600">Failed to load plans.</p>;
  if (!data) return <p>Loading plans…</p>;

  const drafts = data.filter((p) => !p.is_apply);
  const applied = data.filter((p) => p.is_apply);

  const submitApply = async () => {
    if (planId === null) return;
    setIsSubmitting(true);
    try {
      await axiosInstance.post(`/plan/${planId}/apply`, {
        year,
        semester_no: semesterNo,
      });
      // Refresh list
      mutate("/planlist");
      closeApplyPopup();
    } catch (err) {
      console.error(err);
    } finally {
      setIsSubmitting(false);
    }
  };

  const deletePlan = async () => {
    if (planId === null) return;
    setIsDeleting(true);
    try {
      await axiosInstance.delete(`/plan/${planId}`);
      // Refresh list
      mutate("/planlist");
      closeDeletePopup();
    } catch (err) {
      console.error(err);
    } finally {
      setIsDeleting(false);
    }
  };

  const closeApplyPopup = () => {
    setIsApplyPop(false)
    setPlanId(null);
  };


  const closeDeletePopup = () => {
    setPlanId(null);
    setIsDeletePop(false)
  };

  return (
    <div className="space-y-12">
      {/* Draft Plans */}
      <section>
        <h2 className="text-2xl font-semibold mb-4">Draft Plans</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {drafts.map((plan) => (
            <Card key={plan.plan_id} className="border">
              <CardHeader>
                <CardTitle>{plan.name}</CardTitle>
              </CardHeader>
              <CardContent className="h-full flex flex-col justify-between">
                {/* Courses Table */}
                <Table className="w-full mb-10">
                  <TableHeader>
                    <TableRow>
                      <TableHead>Code</TableHead>
                      <TableHead>Name</TableHead>
                      <TableHead className="text-right">Credits</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {plan.courses.map((c) => (
                      <TableRow key={c.course_code}>
                        <TableCell>{c.course_code}</TableCell>
                        <TableCell>{c.course_name}</TableCell>
                        <TableCell className="text-right">
                          {c.credits}
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>

                {/* Actions */}
                <div className="flex space-x-2">
                  <Button onClick={() => onApplyClick(plan.plan_id)}>
                    Apply
                  </Button>
                  <Button
                    variant="destructive"
                    onClick={() => onDeleteClick(plan.plan_id)}
                  >
                    Delete
                  </Button>
                </div>
              </CardContent>
            </Card>
          ))}
          {drafts.length === 0 && (
            <p className="col-span-full text-center text-muted-foreground">
              No draft plans.
            </p>
          )}
        </div>
      </section >

      {/* Applied Plans */}
      < section >
        <h2 className="text-2xl font-semibold mb-4">Applied Plans</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {applied.map((plan) => (
            <Card
              key={plan.plan_id}
              className="border opacity-60 cursor-not-allowed"
            >
              <CardHeader>
                <CardTitle>{plan.name}</CardTitle>
              </CardHeader>
              <CardContent>
                <Table className="w-full">
                  <TableHeader>
                    <TableRow>
                      <TableHead>Code</TableHead>
                      <TableHead>Name</TableHead>
                      <TableHead className="text-right">Credits</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {plan.courses.map((c) => (
                      <TableRow key={c.course_code}>
                        <TableCell>{c.course_code}</TableCell>
                        <TableCell>{c.course_name}</TableCell>
                        <TableCell className="text-right">
                          {c.credits}
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </CardContent>
            </Card>
          ))}
          {applied.length === 0 && (
            <p className="col-span-full text-center text-muted-foreground">
              No applied plans.
            </p>
          )}
        </div>
      </section >
      {
        (planId !== null && isApplyPop) && (
          <>
            <div
              className="fixed inset-0 bg-black/50 z-40"
              onClick={closeApplyPopup}
            />
            <Card className="fixed h-fit inset-0 m-auto w-80 p-4 z-50">
              <CardHeader>
                <CardTitle>Apply Plan</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div>
                  <label className="block mb-1 text-sm">Year</label>
                  <Input
                    type="number"
                    value={year}
                    onChange={e => setYear(+e.target.value)}
                  />
                </div>
                <div>
                  <label className="block mb-1 text-sm">Semester No</label>
                  <Input
                    type="number"
                    value={semesterNo}
                    onChange={e => setSemesterNo(+e.target.value)}
                  />
                </div>
                <div className="flex space-x-2 justify-end">
                  <Button variant="ghost" onClick={closeApplyPopup}>
                    Cancel
                  </Button>
                  <Button
                    onClick={submitApply}
                    disabled={isSubmitting}
                  >
                    {isSubmitting ? "Applying…" : "Apply"}
                  </Button>
                </div>
              </CardContent>
            </Card>
          </>
        )
      }
      {
        ( planId !== null && isDeletePop ) && (
          <>
            <div
              className="fixed inset-0 bg-black/50 z-40"
              onClick={closeApplyPopup}
            />
            <Card className="fixed h-fit inset-0 m-auto w-80 p-4 z-50">
              <CardHeader>
                <CardTitle>Delete Plan</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="flex space-x-2 justify-end">
                  <Button variant="ghost" onClick={closeDeletePopup}>
                    Cancel
                  </Button>
                  <Button
                    onClick={deletePlan}
                    disabled={isDeleting}
                  >
                    {isDeleting ? "Deleting..." : "Delete"}
                  </Button>
                </div>
              </CardContent>
            </Card>
          </>
        )
      }
    </div >
  );
}

