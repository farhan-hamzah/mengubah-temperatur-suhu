package main
import "fmt"

func temperatur() {
	var celciusAwal, celciusAkhir, step, fahrenheit, kelvin, reamur float64
	fmt.Scan(&celciusAwal, &celciusAkhir, &step)
	for celciusAwal <= celciusAkhir {
		reamur = (4.0 / 5.0) * celciusAwal
		fahrenheit = ((9.0/5.0)*celciusAwal + 32)
		kelvin = (celciusAwal + 273)
		fmt.Println(celciusAwal, reamur, fahrenheit, kelvin)
		celciusAwal = celciusAwal + step
	}
}

func main() {
	temperatur()
}
