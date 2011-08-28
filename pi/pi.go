package pi
import ("big"; "math"; "time"; "fmt")

// pi = 4*sum 0 -> inf (((-1)^k)/(2k+1))

func leibnizTerm(ch chan *big.Rat, k int) {
	ch <- big.NewRat(int64(4*math.Pow(-1,float64(k))),int64(((2*k)+1)))
}

func LeibnizPi(iterations int) *big.Rat {
	p := big.NewRat(0,0)
	v := big.NewRat(0,0)
	ch := make(chan *big.Rat, iterations)
	
	//Starts the term calculations and adds to p value if data is available
	for i:=0; i < iterations; i++ {
		go leibnizTerm(ch, i)
	}
	
	var q = false
	var w = 0
	for q != true {
		select{
		case v = <-ch:
			p.Add(p,v)
			w++
		case <-time.After(1*1e9):
			fmt.Printf("timeout: %d\n", w)
			q = true
		}
	}
	return p
}

func sqrt(n *big.Rat) *big.Rat {
	fmt.Println("Need to implement a \"func sqrt(big.Rat) big.Rat\" method")
	//sqrt2 = big.NewRat(99,70)
	return big.NewRat(0,0)
}

func sqrtChain(i int) *big.Rat {
	var two = big.NewRat(2,1)
	if i > 0 {
		i--
		s2chain := sqrtChain(i)
		two.Add(two,s2chain)
	}
	return sqrt(two)
}

func VietesPi(iterations int) *big.Rat {

	chain := sqrtChain(iterations)

	two:= big.NewRat(2,1)
	two.Sub(two,chain)
	two = sqrt(two)

	pi:= big.NewRat(int64(math.Pow(2, float64(iterations+1))),1)
	pi.Mul(two,pi)

	return pi
}
