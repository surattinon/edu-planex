import Image from "next/image"
import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

export const ProfileCard = () => {
  return (
    <>
      <Card className="px-10 h-full w-full flex flex-col justify-between bg-gradient-to-b from-zinc-950/60 to-card/60 backdrop-blur-xl rounded-md border">
        <CardHeader className="mt-5">
          <CardTitle className="text-5xl font-light">Student Profile</CardTitle>
          <CardAction>Edit Profile</CardAction>
        </CardHeader>
        <CardContent className="flex gap-14 h-full items-center">
          <Image src={'/profile-pics/profile-studio.jpg'} alt="profile-pics" width={220} height={220} className="rounded-md border" />
          <div className="flex flex-col gap-12">
            <h1 className="text-5xl">Surattinon Husen</h1>
            <div className="space-y-2">
              <p className="text-sm text-zinc-400">Student ID: <span className="text-white">2105250007</span></p>
              <p className="text-sm text-zinc-400">Email: <span className="text-white">2105250007@students.stamford.edu</span></p>
            </div>
          </div>
        </CardContent>
      </Card>
    </>
  )
}

