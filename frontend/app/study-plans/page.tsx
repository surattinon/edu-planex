// app/plan/page.tsx
import { Planner } from '@/components/Planner';
import { PlanList } from './planlist';

export default function Page() {
  return (
    <div className="h-fit min-h-screen pt-10 ml-15 mr-10">
      <div className='py-6 flex justify-around'>
        <h1 className="text-4xl mb-6">Course Planner</h1>
        <div className=''>
          <Planner />
        </div>
      </div>
      <PlanList />
    </div>
  );
}
