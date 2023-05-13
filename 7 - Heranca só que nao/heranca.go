package main

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa    //"herança" do go
	matricula string
}

func main() {

}
