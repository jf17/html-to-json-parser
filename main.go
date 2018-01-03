package htmltojsonparser

// author : https://github.com/jf17
import (
	"net/http"
)
import "io/ioutil"
import "strings"

func keepLines(s string, n int) string {
	var strslice []string

	result := strings.Join(strings.Split(s, "\n")[:n], "\n")

	str := strings.Replace(result, "\r", "", -1)

	strArray := strings.Split(str, "\n")

	for _, num := range strArray {
		if strings.Contains(num, "<span class=\"hour\">") {
			strslice = append(strslice, num)
		}

	}

	return strings.Join(strslice, "")
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
