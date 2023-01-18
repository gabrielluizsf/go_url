package arch_analyst

import (
	"fmt"
	"runtime"
);
//this function is a systems architecture analyst
func ARCH_ANALYZER(){
	switch MyARCH(){
	case "amd64":
		fmt.Println(MyARCH());
	case "amd64p32":
		fmt.Println(MyARCH());
	default:
		fmt.Println("ARCH UNDEFINIED")

}
}
//this function returns the architecture of the system
func MyARCH()string{
	return runtime.GOARCH;
}