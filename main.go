package htmltojsonparser

// author : https://github.com/jf17
import (
	"net/http"
)
import "io/ioutil"
import "strings"
import "strconv"

type bus_schedule struct {
	Hour, Minute int
}

func replaceHour(s string) string {

	str := strings.Replace(s, "<span class=\"hour\">", "hour=", -1)

	result := strings.Replace(str, "</span></td>", "", -1)

	return result
}

func buildStruct(hourIN string, minuteIN string) bus_schedule {

	min, _ := strconv.Atoi(minuteIN)

	hour, _ := strconv.Atoi(hourIN)

	return bus_schedule{hour, min}

}

func createSlice(strIN string) []bus_schedule {

	var temphour string
	var tempminute string

	var busResult []bus_schedule = []bus_schedule{{0, 0}}

	strArray := strings.Split(strIN, "\n")

	for _, num := range strArray {
		if strings.Contains(num, "minute=\">") {
			break
		} else if strings.Contains(num, "hour") {
			temphour = num[len(num)-2:]
		} else if strings.Contains(num, "minute") {
			tempminute = num[len(num)-2:]
			busResult = append(busResult, buildStruct(temphour, tempminute))
		}

	}

	return busResult

}

func replaceMinute(s string) string {

	str := strings.Replace(s, "<span class=\"minute\">", "minute=", -1)

	result := strings.Replace(str, "</span><br>", "", -1)

	return result
}
func replaceLongMinute(s string) string {

	result := strings.Replace(s, "</span><br>", "", -1)

	slice := result[len(result)-2:]

	result1 := "minute=" + slice

	return result1
}

func keepLines(s string, n int) string {
	var strslice []string

	result := strings.Join(strings.Split(s, "\n")[:n], "\n")

	str := strings.Replace(result, "\r", "", -1)

	strArray := strings.Split(str, "\n")

	for _, num := range strArray {
		if strings.Contains(num, "<span class=\"hour\">") {
			strslice = append(strslice, replaceHour(num))
		} else if strings.Contains(num, "<td align=\"left\" valign=\"middle\" width=") {
			strslice = append(strslice, replaceLongMinute(num))
		} else if strings.Contains(num, "<span class=\"minute\">") {
			strslice = append(strslice, replaceMinute(num))
		}

	}

	return strings.Join(strslice, "\n")
}

func Parse(url string) []bus_schedule {

	var result string = "html-to-json-parser"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	result = keepLines(string(body), 162)

	bus := createSlice(result)
	return bus

}

func WriteJSON() {

	//TODO code

}
