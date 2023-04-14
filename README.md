# Webflow
Go SDK for the Webflow Data API

Usage
-------------

> go get -u github.com/seriallink/webflow

Demo :
```go
package main

import (
    "fmt"
    "github.com/seriallink/webflow"
)

func main() {

	client := webflow.NewClient(
		"webflow_api_token_goes_here",
	)

	info, err := client.AuthorizedInfo()
	if err != nil {
		panic(err)
	}
	
    fmt.Print(info.Status)

}
```
