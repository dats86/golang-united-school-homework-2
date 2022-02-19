package square

import (
	"fmt"
	"math"
)
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type snum int

const SidesTriangle snum = 3
const SidesSquare snum = 4
const SidesCircle snum = 0

func CalcSquare(sideLen float64, sidesNum snum) float64 {

	var square float64

	if sidesNum == SidesTriangle {
		fmt.Println("Square of equilateral Triangle:")
		square = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		fmt.Println("Square of Square:")
		square =  math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		fmt.Println("Square of Circle:")
		square =  math.Pi * math.Pow(sideLen, 2)
	}

	return square
}
