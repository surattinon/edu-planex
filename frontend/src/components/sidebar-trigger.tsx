import { useSidebar } from "@/components/ui/sidebar"

export function CustomTrigger() {
  const { toggleSidebar } = useSidebar()

  return <button onFocus={toggleSidebar}>Toggle Sidebar</button>
}
