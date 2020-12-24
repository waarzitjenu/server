# Running the back-end without Docker

These instructions use a local installation of Go. If you are interested in running the back-end **with** Docker, [follow this guide instead](./With-Docker.md).

Assuming you've installed a working installation of Go, have set all required environmental variables and have set your personal preferences... Here's how to get started:

- Clone this repository to a place of choice, with a Git GUI, Git CLI over HTTPS/SSH or the [GitHub CLI][].

   ```sh
  git clone https://github.com/waarzitjenu/server
  # or: git clone git@github.com:waarzitjenu/server.git
  # or: gh repo clone waarzitjenu/server
   ```

- Move into the directory where you cloned the repository, obviously

- Verify the code's modules and source code

  ```sh
  go mod download
  go mod verify
  go vet
  ```

- Run the server from source with `go run .`, or build a binary using `go build .`

  The build command can also be used to build a binary for a different platform and processor, like a Raspberry Pi, but that's outside of the scope of this document.

- That's it. The server should have created a settings.json which can be altered to change the port. You can also set up TLS (some still call it "SSL") via the settings file.

[GitHub CLI]: https://cli.github.com/	"GitHub's Command Line Interface"