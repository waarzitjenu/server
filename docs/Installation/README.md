# Getting started

Waarzitje.nu has a modular application layout which separates the back-end from the front-end. You can run the back-end, front-end or both at the same time. The back-end could be used to serve the front-end map interface as well. :sunglasses:

Because of this flexible and modular set-up, the whole application can be fit to your needs. Perhaps you're planning on running the API on a public server, while providing a map interface to your team on a local server. Or you just want to play with it, and run both locally. All possible. :happy:

:rocket: Here's how to get started...

### Fetch everything.

It always begins with a `git clone --recursive https://github.com/waarzitjenu/server`. This `--recursive` flag will fetch this back-end repository `waarzitjenu/server` as well as the front-end repository `waarzitjenu/map`.

### Dockerise your life.

Docker aims to make life easy by providing a default development / building / running environment, so you don't have to lose precious time on the well-known "... but it works on my computer"-discussion. Docker is recommended when you're planning on running the front-end, as it uses `yarn` / `npm` and they're known for issues with permissions and conflicting package versions. If you don't have time for trouble, just use Docker. :thumbsup:

<details><summary>Tell me more</summary>

Docker will containerise the application(s) and restrict access to the local filesystem. It's more secure. :lock: It also makes mapping ports easier, because you don't have to change the source code of the app itself. Basically, the apps provides the API on port 8080 by default, but Docker can map the ports and makes it run on port 3000. :muscle:

And another cool feature of Docker is that you can run multiple instances simultanously. One _image_ can be run in two _containers_. So when setting up two data volumes and mapping two different ports, you could provide a public endpoint and a testing endpoint for your API. :smirk:

It is up to you. But using Docker is highly recommended. It makes it secure, flexible, reproducable and you can get rid of everything (containers and/or images) once you're done. :tada:

</details>

The Docker setup was tested on Docker version 19.03.8 and docker-compose version 1.26.2.

---

## 1. Prerequisites for the front-end

Follow these instructions if you want to use the front-end before doing anything else. If you're planning on running the back-end (API interface) only, you don't need to follow this.

### 1.1 Fetch front-end sub module

The front-end has been moved into its own repository [waarzitjenu/map][] a while ago, so in order to use it together with the back-end, it is available as a Git submodule. Git will automatically add the map interface into the `web/` directory, while keeping the map interface in its own repository.

In theory, Git should already have added the repository. Check if the `web` directory exists, or fetch it manually:

```sh
git submodule add https://github.com/waarzitjenu/map ./web/
```

### 1.2 Obtain a Mapbox access token

For the front-end, it is important to sign up for an account at Mapbox and geting an access token. The access token should be placed in a file called `/web/.env.local` in the following format:

```plaintext
VUE_APP_MAPBOX_ACCESS_TOKEN="<Your Mapbox Access Token>"
```

---

## 2. Choose a guide to follow

- [Running the back-end only](./Run back-end only/README.md)
- [Running both the front-end and the back-end](./Run both/README.md)
- [Running the front-end only](./Run front-end only/README.md)

