package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gocloud.dev/blob/fileblob"

	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {
	// Connect to a bucket when your program starts up.
	// This example uses the file-based implementation in fileblob, and creates
	// a temporary directory to use as the root directory.
	dir, cleanup := newTempDir()
	defer cleanup()
	bucket, err := fileblob.OpenBucket(dir, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer bucket.Close()

	// We now have a *blob.Bucket! We can write our application using the
	// *blob.Bucket type, and have the freedom to change the initialization code
	// above to choose a different service-specific driver later.

	// In this example, we'll write a blob and then read it.
	ctx := context.Background()
	if err := bucket.WriteAll(ctx, "foo.txt", []byte("Go Cloud Development Kit"), nil); err != nil {
		log.Fatal(err)
	}
	b, err := bucket.ReadAll(ctx, "foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}

func newTempDir() (string, func()) {
	dir, err := ioutil.TempDir("", "go-cloud-blob-example")
	if err != nil {
		panic(err)
	}
	return dir, func() { os.RemoveAll(dir) }
}
