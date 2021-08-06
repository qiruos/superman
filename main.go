package main

// import "github.com/qiruos/study"

// var studyMaster study.Master

func main() {
	// go study()
	go studyPassive()

	// start io system

	// exit
}

func init() {
	// init env

	// check

}

func studyPassive() {
	for {
		select {
		case <-bus:
			// studyMaster.Process(m)
		}
	}
}
