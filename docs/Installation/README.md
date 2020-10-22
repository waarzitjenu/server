# Getting started

`go-osmand-tracker` has a modular application layout which separates the back-end from the front-end. This makes it possible to run either of one, or both. Here's how to get started.

## 0. Prerequisites

### Obtaining a Mapbox access token

For the front-end, it is important to sign up for an account at Mapbox and geting an access token. The access token should be placed in a file called `/web/.env.local` in the following format:

```plaintext
VUE_APP_MAPBOX_ACCESS_TOKEN="<Your Mapbox Access Token>"
```

If you're planning on running the back-end only, you obviously don't need to do this.



### Back-end, front-end, or both, with or without Docker?

This application has a flexible set-up. Back-end and front-end are separated and the internal code is also separated into chunks. Because of that flexible (or modular) set-up, the whole application can be fit to your needs. 

Perhaps you're planning on running the API on a public server, while providing a map interface to your team on a local server. Or you just want to play with it, and run both locally. All possible. :happy:

Then, there's also Docker. Docker aims to make life easy by providing a default development / building / running environment, so you don't have to ~~spend~~ lose precious time on the well-known "... but it works on my computer"-discussion. Docker is recommended when you're planning on running the front-end, as it uses `yarn` / `npm` and they're known for issues with permissions and conflicting package versions. If you don't have time for that, just use Docker. :thumbsup:

If you're planning on running the back-end only, it's a different story. The back-end is written in Go, and that is stable enough to be used without Docker. But still, Docker is recommended, because it will containerise the application(s) and restrict access to the local filesystem. It's more secure. :lock:

Docker also makes mapping ports easier, because you don't have to change the source code of the app itself. Basically, the apps provides the API on port 8080 by default, but Docker can map the ports and makes it run on port 3000. :muscle:

Another cool feature of Docker is that you can run multiple instances simultanously. One _image_ can be run in two _containers_. So when setting up two data volumes and mapping two different ports, you could provide a public endpoint and a testing endpoint for your API. :smirk:

It is up to you. But using Docker is highly recommended. It makes it secure, flexible, reproducable and you can get rid of everything (containers and/or images) once you're done. :tada:



### Docker and docker-compose versions

Make sure you have a working installation of Docker, if you're planning on using Docker. The setup was tested on Docker version 19.03.8 and docker-compose version 1.26.2. If you're planning on running this app without Docker, you can skip this. :wink: ~~But just set up and use Docker, it saves time.~~

---

## 1. Running both the back-end and the front-end.

To run the back-end and the front-end simultaneously, you could either use Docker with docker-compose, two separate Docker containers, or run them without Docker.

Your options for running both the back-end and front-end comes down to this:

- Run both as one single package: [Use docker-compose](#11-run-both-using-docker-compose)
- Run separately: [Using Docker](#12-run-both-in-separate-docker-containers) or [run it without Docker](#13-run-both-locally) (the latter needs `go` 1.14+, `npm`, `yarn`, a proper set-up etc.)



### 1.1 Run both using docker-compose

The best way to run both the back-end and front-end simultanously is to use Docker in combination with `docker-compose`. Docker-compose builds the docker images for you, sets up data volumes for persistence across restarts, link the data volumes between both containers and it will also run the freshly built images in the right order (i.e. the front-end gets built first, so the back-end could serve the front-end at `/` :wink:).

>  TLDR: no bs, no headaches. docker-compose will take care of everything. It's magic! :rainbow:

So, what are the steps?

- Basically, open a terminal in the project's root directory.
- Take a look at [docker-compose.yml](../../docker-compose.yml).
  - Volumes are automatically created for the build of the front-end and one for the back-end database.
  - If you want to run it on a different port, like **3000**, just change **8080**:8080 to **3000**:8080.
- Let's start the application!
  - The first time, you might have to use `docker-compose up --build`. This tells docker-compose to build the images from scratch, do all the necessary work, and then run them. Stop the containers with Control-C (Cmd-C on macOS).
  - The second time, you would just use `docker-compose up`. Again, stop the containers with C-C.
  - If you want to run the containers in the background (**d**aemonize them), use `docker-compose up -d`. Stop the containers with `docker-compose down`.



### 1.2 Run both in separate Docker containers

Perhaps you're looking for more flexibility or you couldn't get `docker-compose` to work. It is also possible to build two separate Docker images and run them by yourself.

**Building the front-end**

The front-end needs to be built first, so the back-end could serve the front-end at `/`.

- Open a terminal in the project's `./web/`.

- Be sure to have obtained a [Mapbox access token](#obtaining-a-mapbox-access-token) and configured it.

- Let's build the front-end.

  - The first thing you need is a Docker volume to save the files that `yarn` creates during the build process.
  ```sh
  docker volume create go-osmand-tracker-map
  ```

  - Let's build the image. All dependencies will be installed, `yarn` will run for the first time.
  ```sh
  docker build -t ricardobalk/go-osmand-tracker-map .
  ```

  - Now, we need to run the created Docker image in a container, while attaching the volume, so `yarn` saves the files (from the `dist` directory) on that data volume.
  ```sh
  docker run --rm --mount 'type=volume,src=go-osmand-tracker-map,dst=/home/node/go-osmand-tracker/dist' ricardobalk/go-osmand-tracker-map
  ```
  
  - Done! Now we can continue building the back-end, attaching the volume with the front-end to the back-end.

**Building the back-end**

- Open a terminal in the project's root directory.

- Let's build the back-end.

  - The back-end saves location data to a database. We're going to set up a volume to conserve this database between restarts. Awesome.

    ```sh
    docker volume create go-osmand-tracker-apidb
    ```

  - Let's build the image. Go will run, fetch all dependencies, do some checks, etc.

    ```sh
    docker build -t ricardobalk/go-osmand-tracker-api .
    ```

  - Now, we need to run the created Docker image in a container. We need to attach a read-only volume of the front-end map interface and we need to attach the database volume we just created.

    ```sh
    docker run --rm \
      --mount 'type=volume,src=go-osmand-tracker-map,dst=/go/src/go-osmand-tracker/web/dist/,ro' \
      --mount 'type=volume,src=go-osmand-tracker-apidb,dst=/go/src/go-osmand-tracker/database/' \
      -p 8080:8080 \
      ricardobalk/go-osmand-tracker-api
    ```

  - Done! Both the API and map interface should be available at http://localhost:8080/. To change the port to 3000, use `-p 3000:8080` instead of `-p 8080:8080`.



### 1.3 Run both locally

While it is not recommended, you can also run both locally. Refer to [2.2 Running the back-end without Docker](#22-running-the-back-end-without-docker) and [3.2 Running the front-end without Docker](#32-running-the-front-end-without-docker).

---

## 2. Running the back-end only

### 2.1 Running the back-end with Docker

If you only need to run the back-end, you could use Docker.

- Open a terminal in the project's root directory.

- Let's build the back-end.

  - The back-end saves location data to a database. We're going to set up a volume to conserve this database between restarts. Awesome.

    ```sh
    docker volume create go-osmand-tracker-apidb
    ```

  - Let's build the image. Go will run, fetch all dependencies, do some checks, etc.

    ```sh
    docker build -t ricardobalk/go-osmand-tracker-api .
    ```

  - Now, we need to run the created Docker image in a container. We to attach the database volume we just created.

    ```sh
    docker run --rm \
      --mount 'type=volume,src=go-osmand-tracker-apidb,dst=/go/src/go-osmand-tracker/database/' \
      -p 8080:8080 \
      ricardobalk/go-osmand-tracker-api
    ```

  - Done! The API should be available at http://localhost:8080/. To change the port to 3000, use `-p 3000:8080` instead of `-p 8080:8080`.



### 2.2 Running the back-end without Docker

Assuming you've installed a working installation of Go, have set all required environmental variables and have set your personal preferences... Here's how to get started:

**The easy way** &mdash; If you just want to download it and use it, just fetch it with Go.

```sh
go get github.com/ricardobalk/go-osmand-tracker
```

Then, execute the application.

```sh
go-osmand-tracker
```

Alternatively, you could clone this repository, for example, when you want to build the source code yourself and/or help with the development of this server application.

Clone this repository to a place of choice, and run the app straight from the source code with `go run main.go`.

```sh
go run main.go
```

---

## 3. Running the front-end only

### 3.1 Running the front-end with Docker


- Clone this repository

- Go into `./web/`, that directory contains the Vue 3 app.

- Put your Mapbox access token in `./web/.env.local`. It looks like this:
  ```
  VUE_APP_MAPBOX_ACCESS_TOKEN="..."
  ```

- Let's build the front-end.

  - Let's build the image. All dependencies will be installed, `yarn` will run for the first time.
  ```sh
  docker build -t ricardobalk/go-osmand-tracker-map .
  ```

  - Now, we need to run the created Docker image in a container.
  ```sh
  docker run --rm -p 8080:8080 ricardobalk/go-osmand-tracker-map
  ```
  Change `-p 8080:8080` to `-p 4000:8080` to make it run on port 4000.



### 3.2 Running the front-end without Docker

- Make sure you've a working installation of npm, all required permissions, etc.

- Clone this repository

- Go into `./web/`, that directory contains the Vue 3 app.

- Let's fetch the dependencies: `yarn install` will do it.

- Put your [Mapbox access token](#obtaining-a-mapbox-access-token) in `./web/.env.local`. It looks like this:
  ```
  VUE_APP_MAPBOX_ACCESS_TOKEN="..."
  ```

- Run a development server with `yarn run dev`, or use `yarn run build` to build the application.
