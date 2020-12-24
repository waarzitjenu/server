# Running the front-end with Docker


- Clone the front-end repository [waarzitjenu/map][].
  
> Note: If you have cloned it as a submodule of `waarzitjenu/server`, go into `./web/`, that directory will contain the Vue 3 app.

- Put your Mapbox access token in `./.env.local`. It looks like this:

  ```
  VUE_APP_MAPBOX_ACCESS_TOKEN="..."
  ```

- Let's build the front-end.

  - Let's build the image. All dependencies will be installed, `yarn` will run for the first time.

  ```sh
  docker build -t waarzitjenu-map .
  ```

  - Now, we need to run the created Docker image in a container.

  ```sh
  docker run --rm -p 8080:8080 waarzitjenu-map
  ```

  Change `-p 8080:8080` to `-p 4000:8080` to make it run on port 4000.



[waarzitjenu/map]: https://github.com/waarzitjenu/map