https://www.youtube.com/watch?v=5MZWiUHdSBA&list=PLP19RjSHH4aE9pB77yT1PbXzftGsXFiGl&index=2
https://www.youtube.com/watch?v=0s3Jz8Y_cq8&list=PLP19RjSHH4aE9pB77yT1PbXzftGsXFiGl&index=3
https://www.youtube.com/watch?v=NmNjOiKHQt4&list=PLP19RjSHH4aE9pB77yT1PbXzftGsXFiGl&index=4

https://www.youtube.com/watch?v=uixi3pH0fgw&list=PLP19RjSHH4aE9pB77yT1PbXzftGsXFiGl&index=5


$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o hello 001_hello.go

$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o hello.exe 001_hello.go