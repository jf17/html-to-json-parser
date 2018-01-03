package htmltojsonparser

// author : https://github.com/jf17
import (
	"net/http"
)
import "io/ioutil"
import "strings"


func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func Parse(url string) string {

	var result string = "html-to-json-parser"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)


	result = keepLines(string(body), 40)

	return result

}



func WriteJSON()  {

	//TODO code

}