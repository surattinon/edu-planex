"use client";

import React, { useState, useMemo } from "react";
import useSWR, { mutate } from "swr";
import { axiosInstance, fetcher } from "@/lib/api";
import { LazyMotion, domAnimation } from "motion/react"
import * as m from "motion/react-m"

// Shadcn UI components
import { Button } from "@/components/ui/button";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";

interface Course {
  code: string;
  is_enrolled: boolean;
  credits: number;
}

interface CourseType {
  id: number;
  name: string;
  courses: Course[];
}

interface Category {
  id: number;
  name: string;
  course_types: CourseType[];
}

interface Curriculum {
  user_id: number;
  categories: Category[];
}

export function PlanCreator() {
  const { data, error } = useSWR<Curriculum>("personal-cur", fetcher);
  const [open, setOpen] = useState(false);
  const [planName, setPlanName] = useState("");
  const [isSaving, setIsSaving] = useState(false);
  const [selected, setSelected] = useState<Record<string, boolean>>({});

  // Flatten course lookup for credits
  const allCourses = useMemo(() => {
    if (!data) return {};
    const map: Record<string, Course> = {};
    data.categories.forEach(cat =>
      cat.course_types.forEach(ct =>
        ct.courses.forEach(c => (map[c.code] = c))
      )
    );
    return map;
  }, [data]);

  // Compute selected count & credits
  const { count, credits } = useMemo(() => {
    let cnt = 0, sum = 0;
    for (const code in selected) {
      if (selected[code] && allCourses[code]) {
        cnt++;
        sum += allCourses[code].credits;
      }
    }
    return { count: cnt, credits: sum };
  }, [selected, allCourses]);

  const toggle = (code: string) => {
    setSelected(s => ({ ...s, [code]: !s[code] }));
  };


  const savePlan = async () => {
    const courses = Object.entries(selected).filter(([_, v]) => v).map(([code]) => code);
    setIsSaving(true);
    try {
      await axiosInstance.post(`/plans`, {
        name: planName,
        course_codes: courses,
      });
      // Refresh list
      mutate("/planlist");
    } catch (err) {
      console.error(err);
    } finally {
      setIsSaving(false);
      setOpen(false);
      setPlanName("");
      setSelected({});
    }
  };

  if (error) return <p className="text-red-600">Failed to load courses.</p>;
  if (!data) return <p>Loading…</p>;

  return (
    <>
      <Button onClick={() => setOpen(true)}>Create Plan</Button>

      {open && (
        // Overlay background
        <LazyMotion features={domAnimation}>
          <m.div animate={`bg-black`} className="fixed inset-0 bg-black/0 backdrop-blur-xs z-40" onClick={() => setOpen(false)} />
        </LazyMotion>

      )}

      {open && (
        <Card className="fixed inset-y-4 right-4 w-1/2 z-50 flex flex-col">
          <CardHeader className="flex justify-between items-center">
            <CardTitle>Create a New Study Plan</CardTitle>
            <Button variant="ghost" onClick={() => setOpen(false)}>✕</Button>
          </CardHeader>
          <CardContent className="flex-1 overflow-auto flex gap-4">
            {/* Left pane: course selector */}
            <div className="flex-1 overflow-auto">
              {data.categories.map(cat => (
                <section key={cat.id} className="mb-6 mx-5">
                  <h3 className="text-lg font-semibold mb-2">{cat.name}</h3>
                  {cat.course_types.map(ct => (
                    <div key={ct.id} className="mb-4">
                      <h4 className="font-medium">{ct.name}</h4>
                      <div className="grid grid-cols-3 gap-2 mt-2">
                        {ct.courses.map(c => (
                          <label
                            key={c.code}
                            className={`flex items-center p-2 border rounded ${c.is_enrolled ? "opacity-50 cursor-not-allowed" : "cursor-pointer"
                              }`}
                          >
                            <Checkbox
                              checked={!!selected[c.code]}
                              disabled={c.is_enrolled}
                              onCheckedChange={() => toggle(c.code)}
                              className="mr-2"
                            />
                            <span>{c.code}</span>
                          </label>
                        ))}
                      </div>
                    </div>
                  ))}
                </section>
              ))}
            </div>

            {/* Right pane: plan summary and save */}
            <div className="w-80 flex-shrink-0 flex flex-col">
              <div className="mb-4">
                <label className="block mb-1 text-sm">Plan Name</label>
                <Input
                  value={planName}
                  onChange={e => setPlanName(e.target.value)}
                  placeholder="e.g. Spring 2025"
                />
              </div>

              <div className="mb-4">
                <p><strong>{count}</strong> course{count !== 1 && "s"} selected</p>
                <p><strong>{credits}</strong> total credits</p>
              </div>

              <Button
                onClick={savePlan}
                disabled={isSaving || !planName || count === 0}
              >
                {isSaving ? "Saving..." : "Save Draft"}
              </Button>
            </div>
          </CardContent>
        </Card>
      )}
    </>
  );
}
