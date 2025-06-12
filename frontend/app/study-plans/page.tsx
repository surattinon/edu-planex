// app/plan/page.tsx
import { Planner } from '@/components/Planner';
import { PlanList } from './planlist';

export default function Page() {
  return (
    <div className="h-fit min-h-screen pt-10 mx-36 mr-10">
      <div className='px-52 py-10 flex justify-between'>
        <h1 className="text-4xl mb-6">Course Planner</h1>
        <div className=''>
          <Planner />
        </div>
      </div>
      <PlanList />
    </div>
  );
}
