import { AllSemestersList } from "@/components/AllSemesterList"


export default function Page() {
  return (
    <div className="ml-15 p-6 pt-10 flex flex-col">
      <h1 className="text-4xl p-6 mb-4 w-6xl self-center">Enrollment History</h1>
      <AllSemestersList />
    </div>
  )
}
