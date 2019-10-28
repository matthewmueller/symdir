# Symdir

Recursively symlink from one directory to another directory. I made this to symlink my dot files from Dropbox.

## Features

- It doesn't override any existing files
- Recursively creates the folders for you

## Install

```sh
go get -u github.com/matthewmueller/symdir
```

## Usage

```sh
symdir ~/Dropbox/Settings/Home ~/
```

## License

MIT
