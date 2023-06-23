# go-updater

## How it works

It uses [update-golang](https://github.com/udhos/update-golang) to do the updating.
It zips the whole repository, embeds it inside of the executable.
When running, it unzips everything to /tmp/ and runs the update-golang.sh script.
The arguments provided to this program are passed to the script itself.

## Note

DON'T USE IT ON WINDOWS.
Only tested on my machine, on an Ubuntu 22.04 installation.
I expect that it works on any linux distro.

## How to build

```bash
make build
```

Or, alternatively

```bash
make install
```

This copies the executable to /usr/bin/
