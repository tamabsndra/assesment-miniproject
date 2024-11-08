export const siteConfig = {
  name: "Your App Name",
  description: "Your app description",
  apiUrl: process.env.NEXT_PUBLIC_API_URL,
  defaultMetadata: {
    title: "Your App Name",
    description: "Your app description",
    keywords: ["nextjs", "react", "typescript"],
  },
  nav: {
    dashboard: [
      { title: "Dashboard", href: "/dashboard" },
      { title: "Posts", href: "/posts" },
      { title: "Settings", href: "/settings" },
    ],
  },
}

export type SiteConfig = typeof siteConfig
