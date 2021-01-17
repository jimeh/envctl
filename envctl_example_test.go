package envctl_test

import (
	"fmt"
	"os"

	"github.com/jimeh/envctl"
)

func Example_basic() {
	os.Setenv("PORT", "8080")

	envctl.With(map[string]string{"BIND": "0.0.0.0", "PORT": "3000"}, func() {
		fmt.Println(os.Getenv("BIND") + ":" + os.Getenv("PORT"))
	})

	fmt.Println(os.Getenv("BIND") + ":" + os.Getenv("PORT"))
	// Output:
	// 0.0.0.0:3000
	// :8080
}

func ExampleWith() {
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
	envctl.With(env, func() {
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

func ExampleWithClean() {
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
	envctl.WithClean(env, func() {
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
