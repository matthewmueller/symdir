package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matthewmueller/commander"
)

func main() {
	cli := commander.New("symdir", "recursively symlink a source directory to a target directory.")
	verbose := cli.Flag("verbose", "Verbose output.").Short('v').Default("false").Bool()
	from := cli.Arg("from", "Source directory.").Required().String()
	to := cli.Arg("to", "Destination directory.").Required().String()
	cli.Run(func() error { return run(*from, *to, *verbose) })
	cli.MustParse(os.Args[1:])
}

func run(from, to string, verbose bool) error {
	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relpath, err := filepath.Rel(from, path)
		if err != nil {
			return err
		}
		newpath := filepath.Join(to, relpath)
		// create the directory
		if info.IsDir() {
			return os.MkdirAll(newpath, 0755)
		}
		// if stat returns an error and it's not a isNotExist error
		// then return the error. Otherwise if stats works, then the
		// file exists and we ignore
		if _, err := os.Stat(newpath); err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil {
			return nil
		}
		// symlink the file to the destination directory
		if verbose {
			fmt.Fprintf(os.Stdout, "%s -> %s\n", path, newpath)
		}
		return os.Symlink(path, newpath)
	}
	if err := filepath.Walk(from, filepath.WalkFunc(walk)); err != nil {
		return err
	}
	return nil
}
