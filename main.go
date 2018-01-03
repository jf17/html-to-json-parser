package htmltojsonparser

// author : https://github.com/jf17
import (
	"net/http"
)
import "io/ioutil"
import "strings"

func replaceHour(s string) string {

	str := strings.Replace(s, "<span class=\"hour\">", "hour=", -1)

	result := strings.Replace(str, "</span></td>", "", -1)

	return result
}
func replaceMinute(s string) string {

	str := strings.Replace(s, "<span class=\"minute\">", "minute=", -1)

	result := strings.Replace(str, "</span><br>", "", -1)

	return result
}
func replaceLongMinute(s string) string {

	str4 := strings.Replace(s, "<td align=\"left\" ", "", -1)
	str3 := strings.Replace(str4, "align=\"middle\" ", "", -1)
	str2 := strings.Replace(str3, "width=\"25\" ", "", -1)
	str1 := strings.Replace(str2, "lass=\"bottomborder\">", "", -1)

	str := strings.Replace(str1, "<span class=\"minute\">", "minute=", -1)

	result := strings.Replace(str, "</span><br>", "", -1)

	return result
}

func keepLines(s string, n int) string {
	var strslice []string

	result := strings.Join(strings.Split(s, "\n")[:n], "\n")

	str := strings.Replace(result, "\r", "", -1)

	strArray := strings.Split(str, "\n")

	for _, num := range strArray {
		if strings.Contains(num, "<span class=\"hour\">") {
			strslice = append(strslice, replaceHour(num))
		} else if strings.Contains(num, "<span class=\"minute\">") {
			strslice = append(strslice, replaceMinute(num))
		} else if strings.Contains(num, "<td align=\"left\" valign=\"middle\" width=") {
			strslice = append(strslice, replaceLongMinute(num))

		}

	}

	return strings.Join(strslice, "\n")
}

func Parse(url string) string {

	var result string = "html-to-json-parser"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	result = keepLines(string(body), 162)

	return result

}

func WriteJSON() {

	//TODO code

}
