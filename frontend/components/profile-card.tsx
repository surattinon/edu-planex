import { Badge } from "@/components/ui/badge"
import Image from "next/image"
import {
  Card,
  CardAction,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

export function ProfileCard() {
  return (
    <div className="w-full *:data-[slot=card]:from-primary/5 *:data-[slot=card]:to-card dark:*:data-[slot=card]:bg-card grid grid-cols-1 gap-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:shadow-xs ">
      <Card className="@container/card p-10 flex flex-row justify-center space-x-4">
        <Image className="rounded-md border" src={`/imgs/profile-pic/profile-studio.jpg`} alt="profile" width={256} height={256} />
        <div className="flex-1 content-between h-full">
          <CardHeader className="mb-10">
            <CardDescription className="text-xl mb-5">Profile</CardDescription>
            <CardTitle className="text-3xl font-semibold tabular-nums @[250px]/card:text-5xl">
              Surattinon
            </CardTitle>
            <CardAction>
              <Badge variant="outline">
                Year 4
              </Badge>
            </CardAction>
          </CardHeader>
          <CardFooter className="flex-col items-start gap-1.5 text-sm">
            <div className="line-clamp-1 flex gap-2 font-medium">
              ID: 2105250007
            </div>
            <div className="text-muted-foreground">
              Email: 2105250007@students.stamford.edu
            </div>
          </CardFooter>
        </div>
      </Card>
    </div>
  )
}
