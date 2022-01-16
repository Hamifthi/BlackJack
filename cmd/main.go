package main

import (
	"Blackjack"
	"Blackjack/pkg"
	"fmt"
)

func main() {
	var gs Blackjack.GameState
	gs = pkg.Shuffle(gs)
	gs = pkg.Deal(gs)

	var input string
	for gs.State == Blackjack.StatePlayerTurn {
		fmt.Println("Player:", gs.Player)
		fmt.Println("Dealer:", gs.Dealer)
		fmt.Println("What do you want to do hit or stay?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "hit":
			gs = pkg.Hit(gs)
		case "stay":
			gs = pkg.Stay(gs)
		}
	}
	gs.Dealer.FlipHiddenCard()
	gs.Dealer.CalculateScore()
	for gs.State == Blackjack.StateDealerTurn {
		if gs.Dealer.Score <= 16 || (gs.Dealer.Score == 17 && gs.Dealer.MinScore != 17) {
			gs = pkg.Hit(gs)
		} else {
			gs = pkg.Hit(gs)
		}
	}
	gs = pkg.End(gs)
}
