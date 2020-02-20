package main

import (
	"flag"
	"fmt"

	greetingpb "github.com/gonzojive/example-go-bazel/proto/greeting"
)

var (
	word   = flag.String("name", "hello", "The word to use for hello.")
	suffix = flag.String("suffix", "!", "The suffix to add to the greeting.")
)

func main() {
	flag.Parse()
	var greeting *greetingpb.Greeting = &greetingpb.Greeting{
		Word:   *word,
		Suffix: *suffix,
	}
	fmt.Printf("%s, world%s\n", greeting.GetWord(), greeting.GetSuffix())
}
