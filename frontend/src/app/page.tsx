import { ProfileCard } from "@/components/profile-card";
import { StatsCard } from "@/components/stats-card";
import { CredProgCard } from "@/components/credits-prog-card";
import { EnrollHistSlide } from "@/components/enrollhist-table-slides";

export default function Dashboard() {
  return (
    <>
      <div className="h-full flex flex-col gap-3">
        <h1 className="text-4xl mb-3 font-light">Dashboard</h1>
        <div className="w-full h-full flex flex-col md:flex-row gap-3 ">
          <div className="h-full w-full md:w-2/3 ">
            <ProfileCard />
          </div>
          <div className="h-full w-full md:w-1/3 ">
            <CredProgCard />
          </div>
        </div>
        <div className="w-full h-full flex flex-col md:flex-row gap-3 ">
          <div className="h-full w-full md:w-1/2 ">
            <EnrollHistSlide />
          </div>
          <div className="h-full w-full md:w-1/2 ">
            <StatsCard />
          </div>
        </div>
      </div>
    </>
  );
}
