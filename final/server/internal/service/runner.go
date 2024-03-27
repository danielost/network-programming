package service

import (
	"pvms-final/internal/service/handlers"
	"pvms-final/internal/service/requests"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
)

// Game state
const (
	inProgress = iota
	oWin
	xWin
	draw
)

const (
	xTick = "X"
	oTick = "0"

	boardSize = 100
)

var (
	currMove = xTick
)

func TicTacToeRunner(log *logan.Entry, queue chan requests.MoveRequest) {
	board := [100][100]string{}
	xQueue := make([]requests.MoveRequest, 0)
	oQueue := make([]requests.MoveRequest, 0)

	go func() {
		for {
			var request requests.MoveRequest

			if currMove == xTick && len(xQueue) > 0 {
				request = xQueue[0]
				xQueue = xQueue[1:]
				log.Infof("Received a request from the internal X queue {%s: %d %d}", request.Tick, request.X, request.Y)
			} else if currMove == oTick && len(oQueue) > 0 {
				request = oQueue[0]
				oQueue = oQueue[1:]
				log.Infof("Received a request from the internal 0 queue {%s: %d %d}", request.Tick, request.X, request.Y)
			} else {
				request = <-handlers.HttpQueue
				log.Infof("Received a request from HTTP queue {%s: %d %d}", request.Tick, request.X, request.Y)
			}

			time.Sleep(time.Millisecond * 1000)

			if request.Tick != currMove {
				if request.Tick == xTick {
					log.Warn("wrong turn, putting the move to X internal queue")
					xQueue = append(xQueue, request)
				} else {
					log.Warn("wrong turn, putting the move 0 to internal queue")
					oQueue = append(oQueue, request)
				}
				continue
			}

			if currMove == xTick {
				currMove = oTick
			} else {
				currMove = xTick
			}

			x, y := request.X-1, request.Y-1

			if board[x][y] == xTick || board[x][y] == oTick {
				log.Error("position is taken")
				return
			}

			board[x][y] = request.Tick
			state := checkTable(board)

			switch state {
			case xWin:
				log.Info("X win!")
				board = [boardSize][boardSize]string{}
			case oWin:
				log.Info("0 win!")
				board = [boardSize][boardSize]string{}
			case draw:
				log.Info("Draw!")
				board = [boardSize][boardSize]string{}
			case inProgress:
				log.Info("The game goes on!")
			}
		}

	}()
}

func checkTable(board [boardSize][boardSize]string) int {
	fieldsTaken := 0

	// Перевіряємо горизонтальні рядки
	for i := 0; i < boardSize; i++ {
		for j := 0; j < 96; j++ {
			if len(board[i][j]) == 1 && board[i][j] == board[i][j+1] && board[i][j] == board[i][j+2] && board[i][j] == board[i][j+3] && board[i][j] == board[i][j+4] {
				fieldsTaken++
				if board[i][j] == xTick {
					return xWin
				}
				return oWin
			}
		}
	}

	// Перевіряємо вертикальні рядки
	for i := 0; i < 96; i++ {
		for j := 0; j < boardSize; j++ {
			if len(board[i][j]) == 1 && board[i][j] == board[i+1][j] && board[i][j] == board[i+2][j] && board[i][j] == board[i+3][j] && board[i][j] == board[i+4][j] {
				fieldsTaken++
				if board[i][j] == xTick {
					return xWin
				}
				return oWin
			}
		}
	}

	// Перевіряємо діагоналі зліва направо та зверху вниз
	for i := 0; i < 96; i++ {
		for j := 0; j < 96; j++ {
			if len(board[i][j]) == 1 && board[i][j] == board[i+1][j+1] && board[i][j] == board[i+2][j+2] && board[i][j] == board[i+3][j+3] && board[i][j] == board[i+4][j+4] {
				fieldsTaken++
				if board[i][j] == xTick {
					return xWin
				}
				return oWin
			}
		}
	}

	// Перевіряємо діагоналі зліва направо та знизу вгору
	for i := 4; i < boardSize; i++ {
		for j := 0; j < 96; j++ {
			if len(board[i][j]) == 1 && board[i][j] == board[i-1][j+1] && board[i][j] == board[i-2][j+2] && board[i][j] == board[i-3][j+3] && board[i][j] == board[i-4][j+4] {
				fieldsTaken++
				if board[i][j] == xTick {
					return xWin
				}
				return oWin
			}
		}
	}

	if fieldsTaken == boardSize*boardSize {
		return draw
	}
	return inProgress
}
