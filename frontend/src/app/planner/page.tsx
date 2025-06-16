import { PlanList } from "@/components/plan-list"
import { PlanCreator } from "@/components/plan-creator"

const Planner = () => {
  return (
    <>
      <div className="w-full flex flex-col items-center gap-10">
        <div className="flex w-full max-w-6xl justify-between">
          <h1 className="w-full text-4xl font-light">Planner</h1>
          <PlanCreator />
        </div>
        <div className="">
          <PlanList />
        </div>
      </div>
    </>
  )
}

export default Planner
