# Firestore Storage for OAuth 2.0 Server

## Get it

```bash
go get -u github.com/tslamic/go-oauth2-firestore
```

## Use it

```go
package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/tslamic/go-oauth2-firestore"
	"gopkg.in/oauth2.v3/manage"
	"log"
	"os"
)

func main() {
	cli := client(context.Background())
	defer cli.Close()

	// Create a new Firestore TokenStore with "tokens" as a top-level collection name.
	store := fstore.New(cli, "tokens")

	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(store)
}

// As seen here: https://firebase.google.com/docs/firestore/quickstart#initialize
func client(ctx context.Context) *firestore.Client {
	conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
```

## License

    Copyright (c) 2019 Tadej Slamic
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
