package main

import (
	"github.com/markbates/pkger"
)

func main() {
	pkger.Include("/migrations")

}

// pkger list -include /migrations/
// pkger -include /migrations/
