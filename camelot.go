package main

import (
	"fmt"

	"github.com/yaa110/go-persian-calendar/ptime"
)

type Taqvim struct {
}

func main() {
	mah := [6]string{"Odin", "Thor", "Loki", "hel", "Menglad", "surtr"}
	week := [9]string{"Asgard", "Alfheim", "Vanaheim", "Midgard", "Jotunheim", "Svartalfaheim", "Muspelheim", "Niflheim", "Helheim"}
	gmo := GetMonth()
	weekr := ""
	if gmo[0] < 3 {
		wok := 9
		if gmo[1]-gmo[1]/9 > 0 {
			wok = gmo[1] - 9*(gmo[1]/9)
		}
		weekr += week[wok-1]
	} else {
		// fmt.Println(gmo[1] / 6)
	}
	fmt.Println(mah[gmo[0]-1], gmo[1], weekr)
}
func GetMonth() [2]int {

	pt := ptime.Now(ptime.Iran())

	if pt.YearDay() < 144 {
		if pt.YearDay() <= 72 {
			return [2]int{1, pt.YearDay()}
		} else {
			return [2]int{2, pt.YearDay() - 72}

		}
	} else {
		mah := 4/(pt.RYearDay()/54) + 1
		return [2]int{mah + 2, pt.RYearDay() - mah*54 - 5}
	}
}
