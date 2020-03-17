package main

import (
	"errors"
	"fmt"
)

type employee struct {
	name   string
	salary int
}

// func try(v int) (string, error) {
// 	if v > 10 {
// 		return "", errors.New("too much")
// 	} else {
// 		return "good", nil
// 	}
// }

func catch(e error) {
	fmt.Println(e)
}

func finaly() {
	fmt.Println("Recover")
}

func main() {
	defer finaly()
	msg, err := try(11)
	if err != nil {
		catch(err)
	} else {
		fmt.Println(msg)
	}
}

try {
	.....
} catch(e) {
	.....
} finaly {
	.....
}


defer finaly()
msg, err := try()
if err == nil {
	catch(err)
}