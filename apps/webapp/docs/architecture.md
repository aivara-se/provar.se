# Architecture

This is a web application where users can log in and view collected feedback. It is also where users will create and manage users, organizations, projects, and credentials.

This application is built using the following frameworks and libraries:

- **SvelteKit:** A full-stack framework that uses Svelte and NodeJS. [Read more](https://kit.svelte.dev)
- **Auth.js:** Authentication library that supports SvelteKit. [Read more](https://authjs.dev/)
- **Daisy UI:** A UI component library to speed up development. [Read more](https://daisyui.com/)

## Frontend Elements

All routes can be found under `src/routes`. These routes will be rendered on server side using the SvelteKit SSR feature. In addition to them, the web application serves all static content under the `static` directory. This includes the `manifest.json` file which sets the application icon on browsers.

## Backend Elements

Check `+page.server.ts` files and `+layout.server.ts` files for server side code affecting to a specific route or all routes matching a prefix.

In addition to that, SvelteKit hooks are used to authenticate users.

## Authentication

To manage authentication, the application utilizes SvelteKit's `hooks.server.ts` file to enforce user authentication for all routes, except those starting with `/auth`. This setup ensures that only authenticated users can access the application. User login data is stored in MongoDB (see below).

## Persisting Data

The application uses MongoDB to persist data which is hosted by MongoDB Atlas. Check `*.repository.ts` files for code using the database. Database migrations are done by engineers using the MongoDB Shell.

## API Integrations

As of now, it only has an integration with Provar itself using the `provar-js` library.

## Deployment (prod)

For deployment, we leverage GitHub Actions to automate the deployment process. Currently, the deployment is targeted only to the production environment as a staging environment is not yet in place. The deployment workflow involves using Google Cloud Platform services:

- **Google Cloud Artifact Registry:** This is used to store Docker images, ensuring a centralized repository for our application's container images. Repositories are configured to automatically delete older docker images.
- **Google Cloud Run:** Our service is deployed and run on Google Cloud Run, which provides a serverless environment for containerized applications. This allows for scalability and ease of management without worrying about infrastructure.

The deployment pipeline involves continuous integration and deployment using GitHub Actions, enabling seamless updates to the production environment.
