package main

/*
#include <math.h>
double getSqrt(const double test) {
    return sqrt(test);
}
*/
import (
	"C"
)
import "fmt"

func main() {

	for i := 10; i > -10; i-- {
		res, err := C.getSqrt(C.double(float64(i)))

		fmt.Printf("SQRT of %d is %f\n", i, res)

		if err != nil {
			fmt.Printf("Error for finding the sqrt of %d: %v\n", i, err)
		}
	}

}
