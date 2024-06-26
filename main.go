package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(3) // Adding 3 tasks to the wait for

	// Using a channel to signal when a task is done
	done := make(chan string, 3)
	errch := make(chan error, 1)

	go func() {
		defer wg.Done() // Marking the task as done
		veg, err := chopVegetables()
		if err != nil {
			errch <- err
			return
		}
		done <- veg // Signaling that the task is done. We are sending a message to the channel
	}()

	go func() {
		defer wg.Done() // Marking the task as done
		boil, err := boilWater()
		if err != nil {
			errch <- err
			return
		}
		done <- boil // Signaling that the task is done. We are sending a message to the channel
	}()

	go func() {
		defer wg.Done() // Marking the task as done
		dessert, err := mixDessert()
		if err != nil {
			errch <- err
			return
		}
		done <- dessert // Signaling that the task is done. We are sending a message to the channel
	}()

	// Using a separate goroutine to wait for all tasks to be done and to close our communication channel
	go func() {
		wg.Wait()
		close(done)
	}()

	// Process messages from done and errch channels using select
	for completedTasks := 0; completedTasks < 3; {
		select {
		case msg := <-done:
			println(msg)
			completedTasks++
		case err := <-errch:
			if err != nil {
				println("Error while preparing dinner:", err.Error())
				println("Error took", time.Since(start).Milliseconds(), "ms to occur")
				return // here we exit the program, in a real world scenario you would handle / return the error properly
			}
		}
	}

	println("All tasks are done in", time.Since(start).Milliseconds(), "ms")
	prepareDinner()
	println("Dinner is served in", time.Since(start).Milliseconds(), "ms")
}

// the following examples are just simulating real world tasks, in a real world scenario you would have to handle errors properly
// especially in go, a language that you are almost forced to handle errors all the time
func chopVegetables() (string, error) {
	now := time.Now()
	println("Chopping vegetables...")
	time.Sleep(400 * time.Millisecond)
	vegetablePart := fmt.Sprintf("Chopped vegetables in %d ms", time.Since(now).Milliseconds())
	return vegetablePart, nil
}

func boilWater() (string, error) {
	now := time.Now()
	println("Boiling water...")
	time.Sleep(200 * time.Millisecond)
	boilWaterPart := fmt.Sprintf("Boiled water in %d ms", time.Since(now).Milliseconds())
	return boilWaterPart, nil
}

func mixDessert() (string, error) {
	now := time.Now()
	println("Mixing dessert...")
	time.Sleep(300 * time.Millisecond)
	dessertPart := fmt.Sprintf("Mixed dessert in %d ms", time.Since(now).Milliseconds())
	return dessertPart, nil
}

func prepareDinner() {
	now := time.Now()
	println("Preparing dinner...")
	time.Sleep(500 * time.Millisecond)
	println("Dinner is ready", time.Since(now).Milliseconds(), "ms")
}
