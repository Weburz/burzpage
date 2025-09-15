// @ts-check
import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";
import { rehypeMermaid } from "@beoe/rehype-mermaid";
import { getCache } from "@beoe/cache";

const cache = getCache();

export default defineConfig({
  site: "https://burzpage.weburz.com",
  integrations: [
    starlight({
      title: "BurzPage",
      description:
        "A fast & lightweight CMS for your blogging needs at any scale!",
      editLink: {
        baseUrl: "https://github.com/Weburz/burzpage/edit/main/docs",
      },
      social: {
        github: "https://github.com/Weburz/burzpage",
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
          label: "Reference",
          autogenerate: { directory: "reference" },
        },
      ],
      credits: true,
      components: {
        PageFrame: "./src/components/PageFrame.astro",
      },
    }),
  ],
  markdown: {
    rehypePlugins: [
      [
        rehypeMermaid,
        {
          strategy: "file",
          fsPath: "public/beoe",
          webPath: "/beoe",
          darkSchema: "class",
          cache: cache,
        },
      ],
    ],
  },
});
