package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/everdev/mack"
)

func main() {
	navn := []string{"Bjørn", "Nils", "Amalie", "Elin"}
	rand.Seed(time.Now().UTC().UnixNano())
	mack.Say("Participants " + strings.Join(navn, ", "))
	time.Sleep(time.Second * 5)
	mack.Say("And the winner is " + navn[rand.Intn(len(navn))])
}
