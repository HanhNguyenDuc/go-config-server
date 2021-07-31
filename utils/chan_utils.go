package utils

func PushToChan(c chan int) {
	c <- 1
}
