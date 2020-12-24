# Run both in separate Docker containers

Perhaps you're looking for more flexibility or you couldn't get `docker-compose` to work. It is also possible to build two separate Docker images and run them by yourself.

**Building the front-end**

The front-end needs to be built first, so the back-end could serve the front-end at `/`.

- Open a terminal in the project's `./web/`.

- Be sure to have obtained a [Mapbox access token](#obtaining-a-mapbox-access-token) and configured it.

- Let's build the front-end.

  - The first thing you need is a Docker volume to save the files that `yarn` creates during the build process.

  ```sh
  docker volume create waarzitjenu-map
  ```

  - Let's build the image. All dependencies will be installed, `yarn` will run for the first time.

  ```sh
  docker build -t waarzitjenu/map .
  ```

  - Now, we need to run the created Docker image in a container, while attaching the volume, so `yarn` saves the files (from the `dist` directory) on that data volume.

  ```sh
  docker run --rm --mount 'type=volume,src=waarzitjenu-map,dst=/home/node/waarzitjenu/map/dist' waarzitjenu/map
  ```

  - Done! Now we can continue building the back-end, attaching the volume with the front-end to the back-end.

**Building the back-end**

- Open a terminal in the project's root directory.

- Let's build the back-end.

  - The back-end saves location data to a database. We're going to set up a volume to conserve this database between restarts. Awesome.

    ```sh
    docker volume create waarzitjenu-server-db
    ```

  - Let's build the image. Go will run, fetch all dependencies, do some checks, etc.

    ```sh
    docker build -t waarzitjenu-server .
    ```

  - Now, we need to run the created Docker image in a container. We need to attach a read-only volume of the front-end map interface and we need to attach the database volume we just created.

    ```sh
    docker run --rm \
      --mount 'type=volume,src=waarzitjenu-map,dst=/go/src/waarzitjenu/server/web/dist,ro' \
      --mount 'type=volume,src=waarzitjenu-server-db,dst=/go/src/waarzitjenu/server/database/' \
      -p 8080:8080 \
      waarzitjenu-server
    ```

  - Done! Both the API and map interface should be available at http://localhost:8080/. To change the port to 3000, use `-p 3000:8080` instead of `-p 8080:8080`.
