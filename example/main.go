package main

import (
	"github.com/jf17/html-to-json-parser"
)

func main() {

	fmt.Println("Расписание автобусного маршрута 17 от остановки Троицк (мкр. \"В\") (к/ст, выс., пос.) до остановки Торговый центр (выс.)")
	fmt.Println(" сайта Мосгортранс ")
	fmt.Println(" парсится версия \"для печати\" ")
	fmt.Println(htmltojsonparser.Parse("http://www.mosgortrans.org/pass3/shedule.printable.php?type=avto&way=17&date=1111100&direction=AB&waypoint=6"))

}
