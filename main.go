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
			//when i < 0, err contains the string:
			//"The process cannot access the file because another process has locked a portion of the file."
			fmt.Printf("Error for finding the sqrt of %d: %v\n", i, err)
		}
	}

}
