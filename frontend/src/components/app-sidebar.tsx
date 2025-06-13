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
      <SidebarHeader>
        <SidebarMenu>
          <div className="w-4 h-2 bg-black dark:bg-white ml-2 mt-3 rounded-full" />
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
          <SidebarMenuItem className="mb-1">
            <ThemeToggle />
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  )
}
