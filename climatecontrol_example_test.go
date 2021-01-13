package climatecontrol_test

import (
	"fmt"
	"os"

	cc "github.com/jimeh/climatecontrol"
)

func ExampleWithEnv() {
	// existing environment variables
	os.Setenv("MYAPP_HOSTNAME", "myapp.com")
	os.Setenv("MYAPP_PORT", "80")

	fmt.Println("Before:")
	fmt.Printf(" - MYAPP_HOSTNAME=%s\n", os.Getenv("MYAPP_HOSTNAME"))
	fmt.Printf(" - MYAPP_PORT=%s\n", os.Getenv("MYAPP_PORT"))
	fmt.Printf(" - MYAPP_THEME=%s\n", os.Getenv("MYAPP_THEME"))
	fmt.Printf(" - MYAPP_TESTING=%s\n", os.Getenv("MYAPP_TESTING"))

	// temporary environment variables
	env := map[string]string{
		"MYAPP_HOSTNAME": "testing-myapp.test",
		"MYAPP_TESTING":  "unit",
	}
	cc.WithEnv(env, func() {
		os.Setenv("MYAPP_THEME", "dark")

		fmt.Println("Inside func:")
		fmt.Printf(" - MYAPP_HOSTNAME=%s\n", os.Getenv("MYAPP_HOSTNAME"))
		fmt.Printf(" - MYAPP_PORT=%s\n", os.Getenv("MYAPP_PORT"))
		fmt.Printf(" - MYAPP_THEME=%s\n", os.Getenv("MYAPP_THEME"))
		fmt.Printf(" - MYAPP_TESTING=%s\n", os.Getenv("MYAPP_TESTING"))
	})

	// original environment variables restored
	fmt.Println("After:")
	fmt.Printf(" - MYAPP_HOSTNAME=%s\n", os.Getenv("MYAPP_HOSTNAME"))
	fmt.Printf(" - MYAPP_PORT=%s\n", os.Getenv("MYAPP_PORT"))
	fmt.Printf(" - MYAPP_THEME=%s\n", os.Getenv("MYAPP_THEME"))
	fmt.Printf(" - MYAPP_TESTING=%s\n", os.Getenv("MYAPP_TESTING"))
	// Output:
	// Before:
	//  - MYAPP_HOSTNAME=myapp.com
	//  - MYAPP_PORT=80
	//  - MYAPP_THEME=
	//  - MYAPP_TESTING=
	// Inside func:
	//  - MYAPP_HOSTNAME=testing-myapp.test
	//  - MYAPP_PORT=80
	//  - MYAPP_THEME=dark
	//  - MYAPP_TESTING=unit
	// After:
	//  - MYAPP_HOSTNAME=myapp.com
	//  - MYAPP_PORT=80
	//  - MYAPP_THEME=
	//  - MYAPP_TESTING=
}

func ExampleWithCleanEnv() {
	// existing environment variables
	os.Setenv("MYAPP_HOSTNAME", "myapp.com")
	os.Setenv("MYAPP_PORT", "80")

	fmt.Println("Before:")
	fmt.Printf(" - MYAPP_HOSTNAME=%s\n", os.Getenv("MYAPP_HOSTNAME"))
	fmt.Printf(" - MYAPP_PORT=%s\n", os.Getenv("MYAPP_PORT"))
	fmt.Printf(" - MYAPP_THEME=%s\n", os.Getenv("MYAPP_THEME"))
	fmt.Printf(" - MYAPP_TESTING=%s\n", os.Getenv("MYAPP_TESTING"))

	// temporary environment variables
	env := map[string]string{
		"MYAPP_HOSTNAME": "testing-myapp.test",
		"MYAPP_TESTING":  "unit",
	}
	cc.WithCleanEnv(env, func() {
		os.Setenv("MYAPP_THEME", "dark")

		fmt.Println("Inside func:")
		fmt.Printf(" - MYAPP_HOSTNAME=%s\n", os.Getenv("MYAPP_HOSTNAME"))
		fmt.Printf(" - MYAPP_PORT=%s\n", os.Getenv("MYAPP_PORT"))
		fmt.Printf(" - MYAPP_THEME=%s\n", os.Getenv("MYAPP_THEME"))
		fmt.Printf(" - MYAPP_TESTING=%s\n", os.Getenv("MYAPP_TESTING"))
	})

	// original environme
	fmt.Println("After:")
	fmt.Printf(" - MYAPP_HOSTNAME=%s\n", os.Getenv("MYAPP_HOSTNAME"))
	fmt.Printf(" - MYAPP_PORT=%s\n", os.Getenv("MYAPP_PORT"))
	fmt.Printf(" - MYAPP_THEME=%s\n", os.Getenv("MYAPP_THEME"))
	fmt.Printf(" - MYAPP_TESTING=%s\n", os.Getenv("MYAPP_TESTING"))
	// Output:
	// Before:
	//  - MYAPP_HOSTNAME=myapp.com
	//  - MYAPP_PORT=80
	//  - MYAPP_THEME=
	//  - MYAPP_TESTING=
	// Inside func:
	//  - MYAPP_HOSTNAME=testing-myapp.test
	//  - MYAPP_PORT=
	//  - MYAPP_THEME=dark
	//  - MYAPP_TESTING=unit
	// After:
	//  - MYAPP_HOSTNAME=myapp.com
	//  - MYAPP_PORT=80
	//  - MYAPP_THEME=
	//  - MYAPP_TESTING=
}
