package main

import (
	"fmt"
	"github.com/svynnytska/go_crash_course/03_packages/stutil"
	"math"
)

func main(){
	fmt.Println(math.Floor(2.97))
	fmt.Println(math.Ceil(2.97))
	fmt.Printf(stutil.Reverse("qwert"))
}
