package main

import (
    "fmt"
    "bufio"
    "encoding/csv"
    "os"
    "io"
    "strconv"
)

var LOCATION = "US East (N. Virginia)"


type Price struct {
    Location        string
    InstanceType    string
    LeaseLength     string
    PurchaseOpt     string
    OfferingClass   string
    Unit            string
    PricePerUnit    float64
}


func headers(data []string) []string {
    var list []string
    for field := range data {
        list = append(list, string(data[field]))
    }
    return list
}


func objectize(data, headers []string) map[string]string{
    object := make(map[string]string)
    for field := range data {
        object[headers[field]] = data[field]
    }

    return object
}


func priceObject(obj map[string]string) (price Price) {
    price.Location = obj["Location"]
    price.InstanceType = obj["Instance Type"]
    price.LeaseLength = obj["LeaseContrctLength"]
    price.PurchaseOpt = obj["PurchaseOption"]
    price.OfferingClass = obj["OfferingClass"]
    price.Unit = obj["Unit"]
    price.PricePerUnit, _ =  strconv.ParseFloat(obj["PricePerUnit"], 64)

    return
}


func main() {
    i := 0
    f, _ := os.Open("index.csv")
    r := csv.NewReader(bufio.NewReader(f))
    var priceList []Price
    var headerList []string
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if i == 5 {
            headerList = headers(record)
        } else if i > 5 {
            obj := objectize(record, headerList)
            if ( (obj["Operating System"] == "Linux") && (obj["TermType"] == "Reserved") && (obj["Location"] == LOCATION) &&
                 (obj["Tenancy"] == "Shared") && (obj["PricePerUnit"] != "0") && (obj["PricePerUnit"] != "0.0000000000") ) {
                priceList = append(priceList, priceObject(obj))
            }
        }
        i += 1
    }
    fmt.Println(priceList)
}
