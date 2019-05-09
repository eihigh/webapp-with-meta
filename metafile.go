package main

import "github.com/eihigh/meta/metafile"

func main() {
	metafile.New(
		metafile.Tools(),
		metafile.Tasks(
			metafile.Task("test", test),
			metafile.Task("build", build),
		),
	)
}

func test([]string) error {
	panic("no test specified.")
}

func build([]string) error {

	// build server program
	metafile.RunVIn("server", "go", "build")

	// execute server program
	metafile.RunVIn("server", "./server")

	// build client program
	metafile.RunVIn("client", "npm", "run", "build")

	return nil
}
