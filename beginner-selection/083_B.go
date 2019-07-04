package main
import "fmt"
import "math"

func main() {
	var N, a, b int
	fmt.Scan(&N, &a, &b)
	res := 0
	for i := 1; i <= N; i++ {
		tmp := i%10
		for j := 1; j <= 5; j++ {
			tmp += (i/int(math.Pow(10, float64(j))))%10
			
		}
		
		if tmp >= a && tmp <= b {
			res += i
		}
	}
	fmt.Println(res)
}