import { defineConfig } from "vitepress";

export default defineConfig({
  title: "PixiePress",
  description: "A fast, tiny and secure CMS built for the minimalists!",
  themeConfig: {
    logo: "/favicon.svg",
    socialLinks: [
      {
        icon: "github",
        link: "https://github.com/Weburz/pixiepress",
      },
    ],
    footer: {
      message: "Developed with Open-Source <3",
      copyright: "Copyright &copy; Weburz LLC",
    },
    editLink: {
      pattern: "https://github.com/Weburz/pixiepress/edit/main/docs/:path",
      text: "Edit this page on GitHub",
    },
    lastUpdated: {
      text: "Updated at",
      formatOptions: {
        dateStyle: "full",
        timeStyle: "medium",
      },
    },
    externalLinkIcon: true,
    nav: [
      {
        text: "Home",
        link: "/",
      },
      {
        text: "User Guide",
        link: "/user-guide",
      },
      {
        text: "Dev Guide",
        link: "/dev-guide",
      },
      {
        text: "Specifications",
        link: "/dev-guide/spec-sheet",
      },
    ],
  },
  cleanUrls: true,
  srcDir: "./src",
  lastUpdated: true,
});
