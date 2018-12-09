package main

import "fmt"

var input_chan = make(chan byte)
var output_chan = make(chan byte)
func input(input string){
	var data []byte;
	data =	[]byte(input)
	for _, v := range data {
		input_chan <- v
	}
	close(input_chan)
}



func contain(match string) {
	var datas =	[]byte(match)
	for  input:= range input_chan {

		for  _, data := range datas {
			if data == input {
				output_chan <- input
			}
		}
	}
	close(output_chan)

}

func output(){
	var output []byte
	for  input:= range output_chan {
		output = append(output,  input)
	}
	fmt.Println(string(output))
}



func main() {
	go input("abcdefg")
	go contain("begij")
	output()

	//c := make(chan int, 3)
	//
	//go func() {
	//	defer fmt.Println("子协程结束")
	//
	//	fmt.Println("子协程正在运行……")
	//
	//	time.Sleep(1000 * time.Millisecond)
	//	c <- 333
	//}()
	//
	//num := <-c //从c中接收数据，并赋值给num
	//
	//fmt.Println("num = ", num)
	//fmt.Println("main协程结束")

}
