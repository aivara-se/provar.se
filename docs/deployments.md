# Deployments

## Google Cloud Platform

We leverage Google Cloud Services as our chosen cloud platform.

## Github Actions (CI/CD)

Our CI/CD pipeline is managed through Github Actions. It serves the purpose of deploying services and validating pull requests. Sensitive information is securely stored as Github Actions Secrets.

## Web Apps and APIs

Our web applications and APIs are deployed as Google Cloud Run services. You can find service-specific deployment instructions in the respective readme files for each service, such as `apps/webapp/readme.md` and `apps/webapi/readme.md`.

## Public npm modules

NPM modules are manually deployed by authorized developers using the command line. This process involves specific commands tailored to each module's deployment requirements. Modules like `libs/provar-js` and `libs/mockgen` follow this deployment method.

## Dependency Updates

We rely on Dependabot to automatically generate pull requests for updating dependencies. This automated process helps us maintain the health and security of our projects by staying up-to-date with the latest dependencies.

For detailed instructions on deploying specific services and modules, refer to the respective readme files within each service or module directory.
