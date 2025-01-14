// @ts-check
import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

export default defineConfig({
  site: "https://weburz.github.io/burzpress",
  base: "burzpress",
  integrations: [
    starlight({
      title: "BurzPress",
      description:
        "A fast & lightweight CMS for your blogging needs at any scale!",
      editLink: {
        baseUrl: "https://github.com/Weburz/burzpress/edit/main/docs",
      },
      social: {
        github: "https://github.com/Weburz/burzpress",
        discord: "https://discord.gg/QeYqwyxBhR",
        email: "mailto:contact@weburz.com",
        facebook: "https://www.facebook.com/Weburz",
        instagram: "https://www.instagram.com/weburzit",
        linkedin: "https://www.linkedin.com/company/weburz",
        youtube: "https://www.youtube.com/@Weburz",
        twitter: "https://x.com/weburz",
      },
      lastUpdated: true,
      sidebar: [
        {
          label: "Guides",
          items: [
            // Each item here is one entry in the navigation menu.
            { label: "Example Guide", slug: "guides/example" },
          ],
        },
        {
          label: "Reference",
          autogenerate: { directory: "reference" },
        },
      ],
    }),
  ],
});
