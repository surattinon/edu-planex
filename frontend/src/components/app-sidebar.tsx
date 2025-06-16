import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuItem,
  SidebarMenuButton,
  SidebarHeader,
  SidebarFooter
} from "@/components/ui/sidebar"
import { LayoutDashboard, NotebookPen, Table, History } from "lucide-react"

import { ThemeToggle } from "@/components/theme-toggle"

const sidebarItems = [
  {
    title: "Dashboard",
    url: "/",
    icon: LayoutDashboard
  },
  {
    title: "Course Planner",
    url: "/planner",
    icon: NotebookPen
  },
  {
    title: "Course Curriculum",
    url: "/curriculum",
    icon: Table
  },
  {
    title: "Enrollment History",
    url: "/history",
    icon: History
  },
]

export function AppSidebar() {
  return (
    <Sidebar collapsible="icon" variant="sidebar" className="border-none">
      <SidebarHeader className="my-3">
        <SidebarMenu className="w-full flex flex-col justify-center items-start">
          <div className="w-6 h-2 bg-black dark:bg-white dark:shadow-[0_0_10px_1px_rgba(255,_255,_255,_0.5)] ml-1 mt-3 rounded-full" />
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>EDU Planex</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {sidebarItems.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <a href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
          <SidebarGroup />
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem className="mb-2">
            <ThemeToggle />
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  )
}
