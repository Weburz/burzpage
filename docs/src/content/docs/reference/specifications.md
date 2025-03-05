---
title: Software Requirements Specifications (SRS)
---

This section of the documentations contains the "Software Requirements
Specifications (SRS)" for BurzContent. All feature enhancements and related
development on the project will be based on the criteria listed here on this
document. Any other feature request or behaviour of the tool which can be
considered out-of-scope of this document will not be worked upon. In case, a
functionality or behaviour has been heavily requested by community members, the
document will first have to be updated accordingly before development on the
functionality can start taking shape.

### Purpose

BurzContent is a lightweight and fast Content Management System (CMS) for blogging
at any scale. The tool is intended to be used by anyone involved in blogging
either as an individual or with a team of writers/editors. The CMS should
prioritise speed and efficiency over any other proposed functionalities.

**NOTE**: BurzContent is intended to **ONLY** serve the purpose of a CMS and not
the actual blogging website, that will have to be developed by the user (or
their team) according to the requirements. The CMS will expose API endpoints
which the developers can hook into for setting up a blog.

### Scope

BurzContent takes much of its inspiration from similar tools like
[Contentful](https://www.contentful.com), [PayloadCMS](https://payloadcms.com)
and [StrapiCMS](https://strapi.io) and intends to provide a better experience
for everyone involved (including editors and developers alike). Hence, the CMS
should be developed with the following pointers in mind:

1. Ensure the client-side UI/UX is usable out-of-the-box without **ANY**
   customisation required and making it look/feel like a modern SPA-based web
   application.
2. Provide a developer-friendly interface to hook into for creating blogs
   (separate from the CMS itself).
3. Optimise the server-side to be fast as much as possible (**WITHOUT**
   compromising on other functionalities) and ensure it runs on the tiniest VPS
   available in the market at the moment.

### References

The CMS depends on a
[Client-Server Model](https://en.wikipedia.org/wiki/Client%E2%80%93server_model)
and hence the following technology stack is used for development:

- [Golang](https://go.dev/) and its router for writing API services,
  [Chi](https://go-chi.io/) on the server-side.
- [Nuxt.js 3](https://nuxt.com) (and
  [TypeScript](https://www.typescriptlang.com)) on the client-side.
- [PostgreSQL](https://www.postgresql.org) for data storage.

For additional queries related (or unrelated) to BurzContent, please reach out to
[Somraj Saha](mailto:somraj.saha@weburz.com).

### Summary

BurzContent is a lightweight, fast Content Management System (CMS) designed for
blogging. It focuses on speed and efficiency, offering API endpoints for users
to build their own blogging websites. The CMS will be easy to use
out-of-the-box, with a modern UI/UX, and optimized for performance, even on
minimal server resources. It uses Python with FastAPI for the server-side,
Nuxt.js 3 with TypeScript for the client-side, and PostgreSQL for data storage.
BurzContent is intended to be a developer-friendly tool for blogging but does not
provide website creation itself.

## Functional Requirements

BurzContent is expected to provide the following functionalities to its users (and
developers to develop the websites):

1. Ability to create and curate articles (or blogposts) through a Rich-Text
   editor on the client-side.
2. Ability to add editors/writers and assign relevant Role-based Access Control
   List (RBAC) rules to them.
3. Ability to manage comments and curate them on the client-side of the CMS.
4. Provide a robust authentication system (for both the users and developers) to
   authenticate to the CMS to fetch/control it either programmatically or
   through the client-side.
