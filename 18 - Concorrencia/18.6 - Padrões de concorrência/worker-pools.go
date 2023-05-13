package main

func fibo(pos int) int {
	if pos <= 1 {
		return pos
	}

	return fibo(pos-2) + fibo(pos-1)
}

func worker(tarefas chan int, resultados chan int) {

}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
}
