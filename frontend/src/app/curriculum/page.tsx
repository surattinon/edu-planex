import { CurriculumTable } from "@/components/curriculum-table"
import { CurriculumTable2 } from "@/components/curriculum-table-2"

const Curriculum = () => {
  return (
    <>
      <div className="w-full flex flex-col items-center">
        <div className="space-y-10 max-w-6xl w-full">
          <h1 className="w-full text-4xl font-light text-start">Curriculum Courses</h1>
          {/* <CurriculumTable /> */}
          <CurriculumTable2 />
        </div>
      </div>
    </>
  )
}

export default Curriculum
