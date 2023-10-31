package rps

import (
	"strconv"

	"math/rand"
)

const (
	ROCK     = 0 // Piedra. vence a las tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 // Papel. vence a la piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 // Tijeras. vence al papel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// mensajes para cuando gana
var winMessages = []string{
	"!Bien echo!",
	"!Buen trabajo!",
	"Deberias comprar un boleto de loteria",
}

// mensajes para cuando pierde
var loseMessages = []string{
	"!Qué lástima!",
	"!Intentalo de nuevo!",
	"Hoy simplemente no es tu dia",
}

// mensajes de empate
var drawMessages = []string{
	"Las grandes mentes piensan igual",
	"Oh oh. Inténtalo de nuevo",
	"Nadie gana, pero puedes inetntarlo de nuevo",
}

// vriables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	// mensaje dependiendo de loq que eligio la computadora

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = ROCK
		computerChoice = "aLa computadora eligio TIJERA"
	}

	// generar un numero aleatorio de 0-2 que usamos para elegir el mensaje aleatorio
	messageInt := rand.Intn(3)

	var message string
	if playerValue == computerValue {
		roundResult = "Es un empate"
		// seleccionar mensaje de drawMensajes
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "!El jugador gana!"

		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "!La computadora gana!"

		message = loseMessages[messageInt]
	}
	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
