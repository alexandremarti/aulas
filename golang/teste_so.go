package main

import "fmt"
import "os/exec"

func main() {
	fmt.Println("Teste!!")
	out, err := exec.Command("dir","c:\").CombinedOutput;
	//fmt.Println(string(out))
	//if err != nill {
	//	fmt.Println(String(err))
	//}/
}
