// components/ui/icons.tsx
import {
  Loader2,
  ChevronLeft,
  ChevronRight,
  Mail,
  LucideProps,
  User,
  Key,
  AlertTriangle,
  Check,
  ChevronDown,
  Plus,
  X,
} from "lucide-react"

export const Icons = {
  spinner: Loader2,
  chevronLeft: ChevronLeft,
  chevronRight: ChevronRight,
  mail: Mail,
  user: User,
  key: Key,
  warning: AlertTriangle,
  check: Check,
  chevronDown: ChevronDown,
  plus: Plus,
  close: X,
  logo: ({ ...props }: LucideProps) => (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      {...props}
    >
      <circle cx="12" cy="12" r="10" />
      <path d="M12 6v6l4 2" />
    </svg>
  ),
  // Add more icons as needed
}
