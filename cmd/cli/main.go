//go:build !test

package main

import (
	"docker-color-output/internal/cli"
	"docker-color-output/internal/stdin"
	"docker-color-output/internal/stdout"
)

func main() {
	in, err := stdin.Get()
	if err != nil {
		cli.Exit(err)
	}

	out, err := cli.Execute(in)
	if err != nil {
		cli.Exit(err)
	}

	for _, v := range out {
		stdout.Println(v)
	}
}
