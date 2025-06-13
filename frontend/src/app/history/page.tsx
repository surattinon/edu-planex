import { AllEnrollmentTables } from "@/components/enrollhist-tables"

const History = () => {
  return (
    <>
      <div className="w-full flex flex-col items-center">
        <div className="space-y-10">
          <h1 className="w-full text-4xl font-light text-start">Enrollment History</h1>
          <AllEnrollmentTables />
        </div>
      </div>
    </>
  )
}

export default History
