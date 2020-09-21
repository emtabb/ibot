package ibot

// https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
type Work interface {
	Push(Work)
	Machine(IBot)
	GetMachine() IBot
	Queue() []Work
	Stack() []Work
	IsStack(bool)
	// true == queue, false == stack
	Setup()
	Run()
}