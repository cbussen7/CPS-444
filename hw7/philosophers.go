package main

import "fmt"
import "sync"
import "time"
import "math/rand"

var eatWgroup sync.WaitGroup

type fork struct{ 
	sync.Mutex 
}

type philosopher struct {
	name	string
	leftFork	*fork 
	rightFork	*fork
}

func (philo philosopher) eat() {
	defer eatWgroup.Done()

	printState(philo.name, "is seated")
	time.Sleep(time.Duration(rand.Int63n(1e9)))

	//each philosopher must eat three times
	for i := 0; i < 3; i++ {
		printState(philo.name, "is hungry")
		time.Sleep(time.Duration(rand.Int63n(1e9)))

		//philosopher grabs forks, become locked to others
		philo.leftFork.Lock()
		philo.rightFork.Lock()

		//can eat once both forks are acquired
		printState(philo.name, "is eating")
		time.Sleep(time.Duration(rand.Int63n(1e9)))

		//done eating, unlock forks for other philosophers to use
		philo.rightFork.Unlock()
		philo.leftFork.Unlock()

		printState(philo.name, "is thinking")
		time.Sleep(time.Duration(rand.Int63n(1e9)))
	}

	//satisfied after eating three times
	printState(philo.name, "is satisfied")
	time.Sleep(time.Duration(rand.Int63n(1e9)))

}

//function for printing out state of each philosopher
func printState(name string, state string) {
	fmt.Println(name, "is", state)
}

func main() {
	//create array of 5 philosophers
	philosopherNames := [5]string{"Plato", "Socrates", "Aristotle", "Kant", "Marx"}

	numPhilosophers := len(philosopherNames)

	//create as many forks as there are philosophers
	forks := make([]*fork, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = new(fork)
	}

	//add philosophers to the table to begin eating process
	philosophers := make([]*philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &philosopher{
			name: philosopherNames[i], leftFork: forks[i], rightFork: forks[(i+1)%numPhilosophers]}
		eatWgroup.Add(1)
		go philosophers[i].eat()

	}
	eatWgroup.Wait()
}