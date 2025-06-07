package PackageA

import (
	"fmt"
)

func HelloFromPackageA() {
	fmt.Println("Hello from PackageA!")
}

func thisIsPrivate() {
	fmt.Println("thisIsPrivate")
}