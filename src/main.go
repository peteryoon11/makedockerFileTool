package main

import "os"

func main() {
	os.Executable("docker", "version")
}
