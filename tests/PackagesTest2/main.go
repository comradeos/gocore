// В tests/packagesTest2/main.go
package main

import (
    "packagesTest2/packages/packageA"
    "packagesTest2/packages/packageB"
)

func main() {
    packageA.SayHelloFromPackageA()
    packageB.SayHelloFromPackageB()
}

/*

Инициализируйте модуль в вашем проекте. 
В корневой директории вашего проекта (tests/packagesTest2), 
выполните следующую команду: go mod init packagesTest2
Это создаст файл go.mod в вашем проекте с именем модуля packagesTest2.


*/