// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	//waitForResult()
	//fanOut()

	//waitForTask()
	//pooling()

	// Advanced patterns
	//fanOutSem()
	// fanOutBounded()
	//drop()
	cancellation()
}

// waitForResult: You are a manager and you hire a new employee. Your new
// employee knows immediately what they are expected to do and starts their
// work. You sit waiting for the result of the employee's work. The amount
// of time you wait on the employee is unknown because you need a
// guarantee that the result sent by the employee is received by you.
func waitForResult() {
	ch := make(chan string) // built in function
	// make(chan type)
	//      ~~~~
	//            ^
	//          any type

	go func() { // goroutine want to talk
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper" // goroutine send message to ch
		fmt.Println("employee : sent signal")
	}()

	p := <-ch // p receive data from ch
	fmt.Println("manager : recv'd signal :", p)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
//manager : recv'd signal : paper
//employee : sent signal
//-------------------------------------------------------------

// fanOut: You are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. You sit waiting for all the results
// of the employees work. The amount of time you wait on the employees is
// unknown because you need a guarantee that all the results sent by employees
// are received by you. No given employee needs an immediate guarantee that you
// received their result.
func fanOut() {
	students := 100
	//ch := make(chan string)// unbuffered channel => throuput problem
	ch := make(chan string, students) // buffered channel
	// buffered channel
	// improve performance
	// remains data in channel

	// students send quiz to teacher.
	for e := 0; e < students; e++ {
		go func(emp int) {
			fmt.Print(".")
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) //stop
			ch <- "quiz" // send quiz
			// if do you want to send bulk => use loop + send one
			fmt.Println("student : sent signal :", emp)
		}(e)
	}

	// teacher receive quiz from student.
	for students > 0 { //until student is 0
	    // do with teacher's speed.
		p := <-ch // receive quiz from student
		students--
		fmt.Println(p)
		fmt.Println("teacher : recv'd signal :", students)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
// blocked:unbuffered channel
//  => throughput problem
//  SIZE; none. (like 0)
// like: TCP =>guaranteed

// buffered channel
//  => block when capacity is full
//  => if not non block
// you don't know other goroutine
//  SIZE: 1.....
// like: udp =>best effort
//  サーバーが落ちるとchannel内のデータはlost

// tcp/udp joke
//https://www.reddit.com/r/ProgrammerHumor/comments/6p8hmy/hello_would_you_like_to_hear_a_tcp_joke/
//https://www.reddit.com/r/ProgrammerHumor/comments/14wv9p/i_was_gonna_tell_you_guys_a_joke_about_udp/


// which is good? => it depends.

//student : sent signal : 34
//quiz
//teacher : recv'd signal : 99
//student : sent signal : 74
//quiz
//teacher : recv'd signal : 98
//student : sent signal : 64
//quiz
// :
// :
//quiz
//teacher : recv'd signal : 1
//student : sent signal : 70
//quiz
//teacher : recv'd signal : 0


// waitForTask: You are a manager and you hire a new employee. Your new
// employee doesn't know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them. The
// amount of time they wait is unknown because you need a guarantee that the
// work your sending is received by the employee.
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch // wait for receiving
		fmt.Println("employee : recv'd signal :", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager : sent signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// pooling: You are a manager and you hire a team of employees. None of the new
// employees know what they are expected to do and wait for you to provide work.
// When work is provided to the group, any given employee can take it and you
// don't care who it is. The amount of time you wait for any given employee to
// take your work is unknown because you need a guarantee that the work your
// sending is received by an employee.
//
// Channel
//https://github.com/gophercon/2017-talks/blob/master/KavyaJoshi-UnderstandingChannels/Kavya%20Joshi%20-%20Understanding%20Channels.pdf
//https://www.youtube.com/watch?v=KBZlN0izeiY
func pooling() {
	ch := make(chan string)

	// fixed size of pool of goroutines

	g := runtime.NumCPU()
	for e := 0; e < g; e++ { // 8core=> 8 goroutines
		go func(emp int) {
			//
			//            ch <=== [] <= 'paper', 'paper'....
			// emp1 ------^
			// emp2 ------^
			// emp3 ------^
			// emp4 ------^
			// emp5 ------^
			// emp6 ------^

			for p := range ch { // range over channel til channel is closed
			// who close the channel?
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "paper"
		fmt.Println("manager : sent signal :", w)
	}

	close(ch) // close the channelel: no more data(signal) coming to channel
	// ^ This closes the channel!!
	// notice channel 'no more data'
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
//with close => shutdown on goroutines
//manager : sent signal : 97
//manager : sent signal : 98
//employee 2 : recv'd signal : paper
//manager : sent signal : 99
//manager : sent shutdown signal
// :
//employee 3 : recv'd shutdown signal
//employee 1 : recv'd shutdown signal
//employee 7 : recv'd signal : paper
//employee 4 : recv'd signal : paper
//employee 4 : recv'd shutdown signal
//employee 5 : recv'd shutdown signal
//employee 7 : recv'd shutdown signal
//-------------------------------------------------------------

//without close => no shutdown on goroutines
//manager : sent shutdown signal
//employee 3 : recv'd signal : paper
//employee 6 : recv'd signal : paper
//employee 0 : recv'd signal : paper
//employee 7 : recv'd signal : paper

// fanOutSem: You are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. However, you don't want all the
// employees working at once. You want to limit how many of them are working at
// any given time. You sit waiting for all the results of the employees work.
// The amount of time you wait on the employees is unknown because you need a
// guarantee that all the results sent by employees are received by you. No
// given employee needs an immediate guarantee that you received their result.
func fanOutSem() {
	emps := 2000
	ch := make(chan string, emps)

	g := runtime.NumCPU()
	sem := make(chan bool, g)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee : sent signal :", emp)
			}
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", emps)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOutBounded: You are a manager and you hire a team of employees. None of
// the new employees know what they are expected to do and wait for you to
// provide work. The amount of work that needs to get done is fixed and staged
// ahead of time. Any given employee can take work and you don't care who it is
// or what they take. The amount of time you wait on the employees to finish
// all the work is unknown because you need a guarantee that all the work is
// finished.
func fanOutBounded() {
	work := []string{"paper", "paper", "paper", "paper", "paper", 2000: "paper"}

	g := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for e := 0; e < g; e++ {
		go func(emp int) {
			defer wg.Done()
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// drop: You are a manager and you hire a new employee. Your new employee
// doesn't know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them. The
// amount of time they wait is unknown because you need a guarantee that the
// work your sending is received by the employee. You won't wait for the
// employee to take the work if they are not ready to receive it. In that case
// you drop the work on the floor and try again with the next piece of work.
func drop() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		// receive message as possible as it can
		for p := range ch {
			fmt.Println("employee : recv'd signal :", p)
		}
	}()

	// select
	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper": // happy
			fmt.Println("manager : sent signal :", w)
		default: // default case (instead of...)
			fmt.Println("manager : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
//manager : sent signal : 0
//manager : sent signal : 1
//manager : sent signal : 2
//manager : sent signal : 3
//manager : sent signal : 4
//manager : sent signal : 5
//manager : sent signal : 6
//manager : sent signal : 7
//manager : sent signal : 8
// :
//manager : sent signal : 1344
//manager : sent signal : 1345
//manager : sent signal : 1346
//manager : sent signal : 1347
//manager : dropped data : 1348
//manager : dropped data : 1349
// :
//employee : recv'd signal : paper
//employee : recv'd signal : paper
//employee : recv'd signal : paper
//employee : recv'd signal : paper
//employee : recv'd signal : paper
//employee : recv'd signal : paper
//employee : recv'd signal : paper
//-------------------------------------------------------------





// cancellation: You are a manager and you hire a new employee. Your new
// employee knows immediately what they are expected to do and starts their
// work. You sit waiting for the result of the employee's work. The amount
// of time you wait on the employee is unknown because you need a
// guarantee that the result sent by the employee is received by you. Except
// you are not willing to wait forever for the employee to finish their work.
// They have a specified amount of time and if they are not done, you don't
// wait and walk away.
func cancellation() {
	duration := 150 * time.Millisecond // DEADLINE 150ms
	// context package
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	// context : deadline 150ms
	// cancel: function
	defer cancel() // cleanup the context with `cancel function`

	ch := make(chan string, 1) // buffered channel (cap 1) USE BUFFERED CHANNEL
	//ch := make(chan string) // unbuffered => weak

	go func() {
		//time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // work 0-200ms
		//time.Sleep(10 * time.Millisecond) // work 10ms okay
		time.Sleep(300 * time.Millisecond) // work 300ms NG
		ch <- "paper"
	}()

	select {
	case d := <-ch: //receive
		fmt.Println("work complete", d)

	case <-ctx.Done(): // context canceled: => exceed DEADLINE
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
//EXEC 1
//work complete paper
//-------------------------------------------------------------
//EXEC 3
//work cancelled
//-------------------------------------------------------------
//EXEC 3
//work complete paper
//-------------------------------------------------------------