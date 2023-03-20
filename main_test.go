package main

import (
	"testing"

	"github.com/keploy/go-sdk/keploy"
)

// func TestKeploy(t *testing.T) {
// 	// change port so that test server can run concurrently
// 	os.Setenv("PORT", "8090")

// 	keploy.SetTestMode()
// 	go main()
// 	keploy.AssertTests(t)
// }

// package main

// import (
// 	"testing"

// 	"github.com/keploy/go-sdk/keploy"
// )

func TestMain(t *testing.T) {
	// port := "1323"
	// k := keploy.New(keploy.Config{
	// 	App: keploy.AppConfig{
	// 		Name: "echo-gorm",
	// 		Port: port,
	// 	},
	// 	Server: keploy.ServerConfig{
	// 		URL: "http://localhost:6789/api",
	// 	},
	// })

	keploy.SetTestMode()
	// go main()
	keploy.AssertTests(t)
}
