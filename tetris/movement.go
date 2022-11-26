package tetris

import "time"

func (gm *Game) move(gameOverChanel chan<- bool) {

	for i := 0; i < 3; i++ {
		gm.block.moveDown()
		time.Sleep(1 * time.Second)
	}
	gameOverChanel <- true
	// нет блока - создаём
	// есть блок - двигаем
	// некуда двигать - значит упал - проверяем

}

func (gm *Game) needRepaintAllBlocks() bool {
	return false
}
