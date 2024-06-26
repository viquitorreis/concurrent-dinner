package main

// import (
// 	"sync"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	var wg sync.WaitGroup
// 	wg.Add(3) // Adding 3 tasks to the wait for

// 	go func() {
// 		defer wg.Done() // Marking the task as done
// 		chopVegetables()
// 	}()

// 	go func() {
// 		defer wg.Done() // Marking the task as done
// 		boilWater()
// 	}()

// 	go func() {
// 		defer wg.Done() // Marking the task as done
// 		mixDessert()
// 	}()

// 	// Waiting for all tasks to be done
// 	wg.Wait()

// 	println("All tasks are done in", time.Since(now).Milliseconds(), "ms")
// 	prepareDinner()
// 	println("Dinner is served in", time.Since(now).Milliseconds(), "ms")
// }

// func chopVegetables() {
// 	now := time.Now()
// 	println("Chopping vegetables...")
// 	time.Sleep(400 * time.Millisecond)
// 	println("Chopped vegetables in", time.Since(now).Milliseconds(), "ms")
// }

// func boilWater() {
// 	now := time.Now()
// 	println("Boiling water...")
// 	time.Sleep(200 * time.Millisecond)
// 	println("Boiled water in", time.Since(now).Milliseconds(), "ms")
// }

// func mixDessert() {
// 	now := time.Now()
// 	println("Mixing dessert...")
// 	time.Sleep(300 * time.Millisecond)
// 	println("Mixed dessert in", time.Since(now).Milliseconds(), "ms")
// }

// func prepareDinner() {
// 	now := time.Now()
// 	println("Preparing dinner...")
// 	time.Sleep(500 * time.Millisecond)
// 	println("Dinner is ready", time.Since(now).Milliseconds(), "ms")
// }
