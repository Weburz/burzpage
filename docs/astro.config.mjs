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
      social: [
        {
          icon: "github",
          label: "GitHub",
          href: "https://github.com/Weburz/burzpage",
        },
        {
          icon: "discord",
          label: "Discord",
          href: "https://discord.gg/QeYqwyxBhR",
        },
        { icon: "email", label: "Email", href: "mailto:contact@weburz.com" },
        {
          icon: "facebook",
          label: "Facebook",
          href: "https://www.facebook.com/Weburz",
        },
        {
          icon: "instagram",
          label: "Instagram",
          href: "https://www.instagram.com/weburzit",
        },
        {
          icon: "linkedin",
          label: "LinkedIn",
          href: "https://www.linkedin.com/company/weburz",
        },
        {
          icon: "youtube",
          label: "YouTube",
          href: "https://www.youtube.com/@Weburz",
        },
        { icon: "twitter", label: "Twitter", href: "https://x.com/weburz" },
      ],
      lastUpdated: true,

      head: [
        {
          tag: "script",
          attrs: {
            async: true,
            src: "https://analytics.weburz.com/script.js",
            "data-website-id": "74274170-19a7-46b4-ab7e-e3547032855b",
          },
        },
      ],
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
