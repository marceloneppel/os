package main

import (
	"flag"
	"os"
	"os/user"
	"strconv"
)

func main() {
	pathArgument := flag.String("p", "", "pathArgument")
	userArgument := flag.String("u", "", "userArgument")
	flag.Parse()

	if *pathArgument == "" {
		panic("invalid path")
	}

	if *userArgument == "" {
		panic("invalid user")
	}

	providedUser, err := user.Lookup(*userArgument)
	if err != nil {
		panic(err)
	}

	uid, err := strconv.Atoi(providedUser.Uid)
	if err != nil {
		panic(err)
	}

	gid, err := strconv.Atoi(providedUser.Gid)
	if err != nil {
		panic(err)
	}

	err = os.Chown(*pathArgument, uid, gid)
	if err != nil {
		panic(err)
	}
}
