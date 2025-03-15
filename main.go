package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

var dir string

func lotus(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if dir == "" {
			if _, err := os.Stat("lotus.zsh"); err == nil {
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error getting working directory:", err)
					continue
				}

				dir = currentDir

				cmd := exec.Command("zsh", "./lotus.zsh")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Println("Error running lotus.zsh:", err)
				}
			} else if !os.IsNotExist(err) {
				fmt.Println("Error checking lotus.zsh:", err)
			}
		} else {
			return
		}

		time.Sleep(5 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go lotus(&wg)
	wg.Wait()
}
