package main

import (
	"fmt"
	"os"
)

/*
	File Features:
	x	-		xxx	  -   xxx   -   xxx
	_			___		  ___		___
	d or file   owner	  group		other
	_			___		  ___	    ___
	d or -      rwx	  	  rwx	    rwx

	example: 0644
	x	-		xxx	  -   xxx   -   xxx
	_			___		  ___		___
    -	        rw-		  r--		r--

	example: 0755
	x	-		xxx	  -   xxx   -   xxx
	_			___		  ___		___
	d	        rwx		  r-x		r-x

*/

func main() {
	// open file
	f, err := os.OpenFile("./code/os/write/file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write file
	f.Write([]byte("Hello World\n"))
}