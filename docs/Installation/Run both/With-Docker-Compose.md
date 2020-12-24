### 1.1 Run both using docker-compose

The best way to run both the back-end and front-end simultanously is to use Docker in combination with `docker-compose`. Docker-compose builds the docker images for you, sets up data volumes for persistence across restarts, link the data volumes between both containers and it will also run the freshly built images in the right order (i.e. the front-end gets built first, so the back-end could serve the front-end at `/` :wink:).

>  TLDR: no bs, no headaches. docker-compose will take care of everything. It's magic! :rainbow:

So, what are the steps?

- Basically, open a terminal in the project's root directory.
- Take a look at [docker-compose.yml](../../../docker-compose.yml).
  - Volumes are automatically created for the build of the front-end and one for the back-end database.
  - If you want to run it on a different port, like **3000**, just change **8080**:8080 to **3000**:8080.
- Let's start the application!
  - The first time, you might have to use `docker-compose up --build`. This tells docker-compose to build the images from scratch, do all the necessary work, and then run them. Stop the containers with Control-C (Cmd-C on macOS).
  - The second time, you would just use `docker-compose up`. Again, stop the containers with Control-C / Cmd-C.
  - If you want to run the containers in the background (**d**aemonize them), use `docker-compose up -d`. Stop the containers with `docker-compose down`.