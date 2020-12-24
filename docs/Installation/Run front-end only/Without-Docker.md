# Running the front-end without Docker

- Make sure you've a working installation of npm, all required permissions, etc.

- Clone the front-end repository [waarzitjenu/map][].

  > Note: If you have cloned it as a submodule of `waarzitjenu/server`, go into `./web/`, that directory will contain the Vue 3 app.

- Go into the directory contains the Vue 3 app.

- Let's fetch the dependencies: `yarn install` will do it.

- Put your Mapbox access token in `./.env.local`. It looks like this:

  ```
  VUE_APP_MAPBOX_ACCESS_TOKEN="..."
  ```

- Run a development server with `yarn run dev`, or use `yarn run build` to build the application.

[waarzitjenu/map]: https://github.com/waarzitjenu/map