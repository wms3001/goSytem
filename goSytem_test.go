package goSytem

import (
	"fmt"
	"testing"
)

func TestGoSytem_Info(t *testing.T) {
	goSytem := GoSytem{}
	resp := goSytem.Info()
	fmt.Println(resp)
}
