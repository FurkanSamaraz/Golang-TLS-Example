package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*import (



	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xuri/excelize/v2"
)

//Iç içe Modeller.
type TimeModel struct {
	Updated    string
	UpdatedISO string
	Updateduk  string
}
type CruModel struct {
	Code       string
	Symbol     string
	Rate       string
	Rate_Float float64
}
type BpiModel struct {
	USD CruModel
	EUR CruModel
	GBP CruModel
}

type ApiResModel struct {
	Time      TimeModel
	ChartName string
	Bpi       BpiModel
}

func main() {
	resp, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	data, _ := ioutil.ReadAll(resp.Body)
	var api ApiResModel
	json.Unmarshal(data, &api)

	usd := api.Bpi.USD
	eur := api.Bpi.EUR
	gbp := api.Bpi.GBP

	categories := map[string]string{ // Başlıkları tanımlama.
		"A2": "$$$", "B1": "USD", "C1": "EUR", "D1": "GBP"}

	values := map[string]int{ //Json Apiden gelen para birimi dalgalanma oranları çekme.
		"B2": int(usd.Rate_Float),
		"C2": int(eur.Rate_Float), "D2": int(gbp.Rate_Float)}

	f := excelize.NewFile() //excel oluşturma.

	for k, v := range categories { // For ile categories içindekileri yazdırma.
		f.SetCellValue("Sheet1", k, v)
	}

	for k, v := range values { // For ile values içindekilerini yazdırma.
		f.SetCellValue("Sheet1", k, v)
	}

	// 3D Grafik Oluşturma.

	if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$B$1",
			"categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2"
        },
        {
            "name": "Sheet1!$C$1",
			"categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$C$2"
        },
        {
            "name": "Sheet1!$D$1",
			"categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$D$2:$D$2"
        }],
        "title":
        {
            "name": "-Column Chart-"
        }
    }`); err != nil {
		fmt.Println(err)
		return
	}
	// Kayıt etme.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}
*/
