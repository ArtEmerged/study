package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var line, column, amoutSum int
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)
	fmt.Fscanln(in, &line, &column, &amoutSum)

	//создаем массив префиксных сумм, он на еденицу длинее массива входного массива
	quad := make([][]int, line+1)
	quad[0] = make([]int, column+1)

	//заполняем двухмерный массив префиксными суммами
	for l := 0; l < line; l++ {
		quad[l+1] = make([]int, column+1)
		for r := 0; r < column; r++ {
			var nums int
			fmt.Fscan(in, &nums)

			//преф. сумма = преф. сумма прошлого элемента по горизонтали (quad[l+1][r]) +
			//преф. сумма прошлого элемента по вертикали (quad[l][r+1]) + наш элемент (nums)
			//но так как горизонтали и вертикаль пересекаются, мы прибавили Левый Верхний уголо 2 раза
			//теперь мы его вычтем quad[l][r]
			quad[l+1][r+1] = quad[l+1][r] + quad[l][r+1] + nums - quad[l][r]
		}
	}

	//Нахождение суммы прямоугольника через преф. сумму
	for i := 0; i < amoutSum; i++ {

		// [x][y]
		// L - Left,  R - Right
		var xL, yL, xR, yR int
		fmt.Fscan(in, &xL, &yL, &xR, &yR)

		//Из правого нижнего угла преф. сумм quad[xR][yR] вычитаем quad[xR][yL-1] и quad[xL-1][yR]
		//quad[xR][yL-1] - префиксная сумма вертикали которая не входит
		//quad[xL-1][yR] - префиксная сумма горизонтали которая не входит
		//но так как горизонтали и вертикаль пересекаются, мы вычли Левый Верхний уголо 2 раза
		//теперь мы его прибавляем quad[xL-1][yL-1]
		fmt.Println(quad[xR][yR] - (quad[xR][yL-1] - quad[xL-1][yR] + quad[xL-1][yL-1]))
	}
}
