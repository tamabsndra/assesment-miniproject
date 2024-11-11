export const siteConfig = {
  name: "Miniproject",
  description: "Simple Post Management",
  apiUrl: process.env.NEXT_PUBLIC_API_URL,
  defaultMetadata: {
    title: "Miniproject",
    description: "Simple Post Management",
    keywords: ["nextjs", "react", "typescript"],
  },
  nav: {
    dashboard: [
      { title: "Dashboard", href: "/dashboard" },
      { title: "Posts", href: "/posts" },
    ],
  },
}

export type SiteConfig = typeof siteConfig
