## Things to think about

#### Image size

- Used multistage build for the UI, and Go containers.

#### Container security

- Only the API and UI ports are exposed to localhost.

#### Container inter-dependencies

- Dependency checks and health checks added to compose.
- Removed pg health check from the Dockerfile.

#### Have people unfamiliar with containerization and possibly a terminal able to stand up the stack on their own.

- A navigational menu is built to assist non-tech users.
- Lifecycle functions are included for non-tech users.


### Out of Scope

- In a production environment, Postgres credentials would be provided to the container either as secrets or from the .env files. Didn't implement it here. In a cloud environment we could also inject with aws vault for instance.

- For password, I could have also added checks to check the environment variables or default to the default values. Also didn't implement it here. I did however moved the environment variables from the Dockerfile onto the compose file.

- There are smaller container images which could be used for FROM, but didn't implement it here.

- I used nginx for serving Angular files, a single page express application could be used also. There are other web servers as well.

- Containers are running as ROOT, which probably would not be too ideal in production environments.

- Resource allocations are not capped, in a production environment it is usually a good idea to limit CPU, memory etc.

- Containers could have been made read only. Didn't implement it here.

- If the audience was QA, developers etc. mounting the PG data directory as a volume could be useful.

- Default nginx configuration is not suitable for production, and could be hardened to block accidental serving of possibly sensitive files (.env etc).

- Instead of using yarn global add, the executable could be launched once with npx or nlx. Reducing the image size.

- Some base images are pinned to specific version. This could be suitable to match the production / development environments.

- The included menu is crafted using GUM for amd64. If a wide variety of users are expected to try the project, startup menu would need to be modified. Didn't implement it here. I could have also used generic sh/bash for a simpler menu. FZF would also be another tool. There are many tools, including some for GUI.

- If the non tech end user is the target audience, Prebuilt binaries should be included in the project targeting different architectures. Some GO projects might not cross compile successfully, and extending Docker into QEMU might be required for MacOS users. Not implemented.

- Custom health check scripts could be nice for the other containers. Not implemented.

- .gitignore files should be kept in the build directories (for dev targets). Not implemented here.

- In production environments, tools like docker slim could be used in pipelines. In production environments, alpine images are better suited.

- In a real-life repository, this readme.md file would have instructions on how to get the project up and running. Not implemented. 