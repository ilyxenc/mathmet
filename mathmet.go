package mathmet

import (
	"fmt"
	"math"
)

type Matrix struct {
	nRows    int
	nColumns int
	//n_rows    int
	//n_columns int
	vector map[int][]float64
}

func CreateMatr(i, j int) *Matrix {
	data := make(map[int][]float64)
	fmt.Println("Задайте значения для матрицы ", i, "*", j)
	for y := 0; y < i; y++ {
		for x := 0; x < j; x++ {
			var value float64
			fmt.Scan(&value)
			data[y] = append(data[y], value)
		}
	}
	return &Matrix{
		nRows:    i,
		nColumns: j,
		vector:   data,
	}
}

type funcUsed func(x float64) float64
type funcUsed2 func(x, y float64) float64

//Метод касательных (Ньютона)
func MethNewton(f funcUsed, x0 float64, eps float64) float64 {
	var df float64 = (f(x0+eps) - f(x0)) / eps
	var x1 float64 = x0 - f(x0)/df
	for math.Abs(x1-x0) > eps {
		x0 = x1
		x1 = x1 - f(x1)/((f(x1+eps)-f(x1))/eps)
	}
	return x1
}

//Метод хорд ищет корни на отрезке
func MethChord(f funcUsed, a, b float64, eps float64) float64 {
	for math.Abs(b-a) > eps {
		a = b - (b-a)*f(b)/(f(b)-f(a))
		b = a + (a-b)*f(a)/(f(a)-f(b))
	}
	return b
}

//Метод простых итераций (уточнения корня)
func MethIteration(f funcUsed, x0 float64, eps float64) float64 {
	var lmbd float64 = eps / (f(x0+eps) - f(x0))
	var x1 float64 = x0
	for true {
		x1 = x0
		x0 = x1 + lmbd*f(x0)
		if math.Abs(x0-x1) < eps {
			break
		}
	}
	return x0
}

//Метод интерполяции Лагранжа
func MethInterpolateLagrangePolynomial(f funcUsed, x float64, size int) float64 {
	var polynom float64 = 0
	var basics_polynom float64
	f1_values := make([]float64, size)
	f2_values := make([]float64, size)

	for i := 0; i < size; i++ {
		f1_values[i] = float64(i)
		f2_values[i] = f(float64(i))
	}

	for i := 0; i < size; i++ {
		basics_polynom = 1
		for j := 0; j < size; j++ {
			if j != i {
				basics_polynom *= (x - f1_values[j]) / (f1_values[i] - f1_values[j])
			}
			polynom += basics_polynom * f2_values[i]
		}
	}
	return polynom
}

//Метод наименьших квадратов, интерполирующий к виду ax^2+bx+c
func MethLeastSquare(x_values, y_values []float64) {
	lenVal := len(x_values)
	calc_table := []float64{0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 7; i++ {
		for j := 0; j < lenVal; j++ {
			switch i {
			case 0:
				calc_table[i] += x_values[j]
			case 1:
				calc_table[i] += y_values[j]
			case 2:
				calc_table[i] += x_values[j] * x_values[j]
			case 3:
				calc_table[i] += x_values[j] * x_values[j] * x_values[j]
			case 4:
				calc_table[i] += x_values[j] * x_values[j] * x_values[j] * x_values[j]
			case 5:
				calc_table[i] += x_values[j] * y_values[j]
			case 6:
				calc_table[i] += x_values[j] * x_values[j] * y_values[j]
			}
		}
	}
	data := make(map[int][]float64)
	data[0] = []float64{calc_table[4], calc_table[3], calc_table[2], calc_table[6]}
	data[1] = []float64{calc_table[3], calc_table[2], calc_table[0], calc_table[5]}
	data[2] = []float64{calc_table[2], calc_table[0], float64(lenVal), calc_table[1]}

	//Начало метода Гаусса с выбором главного элемента
	x := make(map[int][]float64)
	n := 3
	for i := 0; i < n; i++ {
		max := math.Abs(data[i][i])
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(data[i][i]) > max {
				max = math.Abs(data[i][i])
				maxRow = k
			}
		}
		for k := i; k < n+1; k++ {
			c := data[maxRow][k]
			data[maxRow][k] = data[i][k]
			data[i][k] = c
		}
		for k := i + 1; k < n; k++ {
			c := -data[k][i] / data[i][i]
			for j := i; j < n+1; j++ {
				if i == j {
					data[k][j] = 0
				} else {
					data[k][j] += c * data[i][j]
				}
			}
		}
		for k := i + 1; k < n; k++ {
			c := -data[k][i] / data[i][i]
			for j := i; j < n+1; j++ {
				if i == j {
					data[k][j] = 0
				} else {
					data[k][j] += c * data[i][j]
				}
			}
		}
	}
	for i := n - 1; i > -1; i-- {
		x[i] = append(x[i], data[i][n]/data[i][i])
		for k := i - 1; k > -1; k-- {
			data[k][n] -= data[k][i] * x[i][0]
		}
	} //Конец метода Гаусса с выбором главного элемента

	fmt.Println("Итоговое уравнение")
	fmt.Println("y =", x[0][0], "* x^2 +", x[1][0], "* x +", x[2][0])
}

//Метод трапеций
func sumTrap(f funcUsed, a, b float64, n int) float64 {
	var h float64 = (b - a) / float64(n)
	var sum float64 = (f(a) + f(b)) / 2
	for x := (a + h); x < (b - h); x += h {
		sum += f(x)
	}
	return sum * h
}

func MethTrap(f funcUsed, a, b, eps float64) float64 {
	var n int = 2
	var i float64 = sumTrap(f, a, b, n)
	var i_half float64
	var r float64
	for true {
		n *= 2
		i_half = sumTrap(f, a, b, n)
		r = (i_half - i) / 3
		if math.Abs(r) < eps {
			break
		}
		i = i_half
	}
	return math.Abs(i_half + r)
}

//Метод Симпсона
func sumSimpson(f funcUsed, a, b, n float64) float64 {
	var h float64 = (b - a) / n
	var h2 = h * 2
	var res1 float64 = f(a) + f(b)
	var res2 float64 = 0
	for x := a + h; x < b; x += h2 {
		res2 += f(x)
	}
	var res3 float64 = 0
	for x := a + h2; x < b; x += h2 {
		res3 += f(x)
	}
	return h * (res1 + 4*res2 + 2*res3)
}

func MethSimpson(f funcUsed, a, b, eps float64) float64 {
	var n float64 = 2
	var i float64 = sumSimpson(f, a, b, n)
	var i_half float64
	var r float64
	for true {
		n *= 2
		i_half = sumSimpson(f, a, b, n)
		r = (i_half - i) / 15
		if math.Abs(r) < eps {
			break
		}
		i = i_half
	}
	return math.Abs(i_half + r)
}

//Метод Эйлера
func MethEuler(f funcUsed2, n int, step float64) []float64 {
	var x float64 = 0
	var y float64 = 0 //Стартовые точки
	for i := 0; i < n; i++ {
		y += step * f(x, y)
		x += step
	}
	return []float64{x, y}
}

//Метод Эйлера-Коши
func MethEulerKoshi(f funcUsed2, a, b, step, x0, y0 float64) {
	var n int = int((b - a) / step)
	fmt.Println(n)
	x := make([]float64, n)
	y := make([]float64, n)
	y1 := make([]float64, n)
	x[0] = x0
	y[0] = y0
	for i := 1; i < n; i++ {
		x[i] = a + float64(i)*step
		y1[i] = y[i-1] + step*f(x[i-1], y[i-1])
		y[i] = y[i-1] + step*(f(x[i-1], y[i-1])+f(x[i], y1[i]))/2
	}
	fmt.Println("X координаты: ", x)
	fmt.Println("Y координаты: ", y)
}
