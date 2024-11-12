package ex2

type PointX struct {
	R int
	C int
}

func FindRectangle(A [][]bool) (res [2]PointX) {
	startPoint := false
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] && !startPoint {
				res[0] = PointX{R: i, C: j}
				startPoint = true
			}
			if A[i][j] {
				res[1] = PointX{R: i, C: j}
			}

		}
	}

	return res

}
