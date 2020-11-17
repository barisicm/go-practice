package main

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
}
