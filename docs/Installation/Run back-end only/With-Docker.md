# Running the back-end with Docker

These instructions use Docker. If you are interested in running the back-end **without** Docker, [follow this guide instead](./Without-Docker.md).

## Instructions

- Open a terminal in the project's root directory.

- Let's build the back-end.

  - The back-end saves location data to a database. We're going to set up a volume to conserve this database between restarts. Awesome.

    ```sh
    docker volume create waarzitjenu-server-db
    ```

  - Let's build the image. Go will run, fetch all dependencies, do some checks, etc.

    ```sh
    docker build -t waarzitjenu/server .
    ```

  - Now, we need to run the created Docker image in a container. We to attach the database volume we just created.

    ```sh
    docker run --rm \
      --mount 'type=volume,src=go-osmand-tracker-apidb,dst=/go/src/waarzitjenu/server/database/' \
      -p 8080:8080 \
      waarzitjenu/server
    ```

  - Done! The API should be available at http://localhost:8080/. To change the port to 3000, use `-p 3000:8080` instead of `-p 8080:8080`.
