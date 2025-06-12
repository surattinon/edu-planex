'use client';

import { useRouter } from 'next/navigation';
import { useState, useEffect } from 'react';
import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { fetchAllCourses, fetchEnrollments, savePlanDraft } from '@/lib/api';
import { CourseType } from '@/types/course';

export function Planner() {
  const [courses, setCourses] = useState<CourseType[]>([]);
  const [loaded, setLoaded] = useState(false);
  const [enrolledCodes, setEnrolledCodes] = useState<Set<string>>(new Set());
  const router = useRouter()

  useEffect(() => {
    Promise.all([fetchAllCourses(), fetchEnrollments()])
      .then(([all, enrolls]) => {
        setCourses(all);
        setEnrolledCodes(new Set(enrolls.flatMap(e => e.courses.map(c => c.course_code))));
      })
      .finally(() => setLoaded(true));
  }, []);

  const [open, setOpen] = useState(false);
  const [selected, setSelected] = useState<Set<string>>(new Set());
  const toggleCourse = (code: string) => {
    setSelected(prev => {
      const next = new Set(prev);
      if (next.has(code)) next.delete(code);
      else if (next.size < 5) next.add(code);
      return next;
    });
  };

  const [planName, setPlanName] = useState("")

  const creditSum = Array.from(selected).reduce((sum, code) => {
    const crs = courses.find(c => c.course_code === code);
    return sum + (crs?.credits || 0);
  }, 0);

  const handleSaveDraft = async () => {
    await savePlanDraft({ name: `${planName}`, course_codes: Array.from(selected) });
    router.push('/study-plans')
    setOpen(false);
  };

  if (!loaded) return <p>Loading…</p>;

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild><Button>New Plan</Button></DialogTrigger>
      <DialogContent className="py-9 px-10 max-h-[1000px] min-w-[700px] lg:min-w-[1400px] bg-gradient-to-b from-primary/10 to-card/30 backdrop-blur-sm border-2">
        <DialogHeader><DialogTitle className='text-3xl text-gray-200'>Create New Course Plan</DialogTitle></DialogHeader>
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 py-4">
          <div className="h-full col-span-2 grid grid-col-2 lg:grid-cols-9 gap-2 overflow-x-hidden p-2">
            {courses.map(c => {
              const isEnrolled = enrolledCodes.has(c.course_code);
              const isSelected = selected.has(c.course_code);
              return (
                <button key={c.course_code} disabled={isEnrolled} onClick={() => !isEnrolled && toggleCourse(c.course_code)}
                  className={`bg-green-900/70 border p-3 rounded text-center hover:scale-105 hover:bg-green-700 transition-all ${isEnrolled ? 'bg-gradient-to-b from-card to-card text-gray-600 cursor-not-allowed hover:scale-100' : ''} ${isSelected ? 'border-green-400 text-white border-2 shadow-[0_0px_10px_rgba(147,_255,_156,_0.3)]' : ''}`}
                >
                  <div className="font-semibold text-gray-300">{c.course_code}</div>
                </button>
              );
            })}
          </div>
          <div className="col-span-1 space-y-4">
            <div>
              <label className="block text-3xl mb-3">Plan Name</label>
              <Input id='planname' value={planName} onChange={e => setPlanName(e.target.value)} />
            </div>
            <div className="text-lg">Selected <strong>{selected.size}/5</strong><br />Total Credits: <strong>{creditSum}</strong></div>
            <div className="flex justify-end space-x-2"><Button variant="outline" onClick={handleSaveDraft} disabled={selected.size === 0}>Save Draft</Button></div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
