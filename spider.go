package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"net/http"
)

func main()  {
	BTC()
	LTC()
	VTC()
	// XAU()
	// OuYuanToYingBang()
	// YingBangToAoYuan()
	// YingBangToMeiYuan()
	// Gold()
}

// BTC data
func BTC() {	
	url := "https://coinmarketcap.com/zh/currencies/bitcoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`<span data-currency-value>(.*?)</span>`)
	// fmt.Println(  r.Find(bodyBytes)  )
	res := r.FindAllString(body, -1)
	// fmt.Println(len(r.FindAllString(body, -1)))
	fmt.Println("****************BTC")
	fmt.Println("PRICE", res[2])
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])
	fmt.Println("****************")


}

// LTC data
func LTC() {	
	url := "https://www.feixiaohao.com/currencies/litecoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`</span> <span class="info_val"><span class="convert">(.*?)</span></span></div> <div class="info_cell"><span class="info_tit">`)
	res := r.FindAllString(body, -1)	
	fmt.Println("****************LTC")
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])

	// r2, _ := regexp.Compile(`<span class="convert textRed">(.*?)</span>`)
	r2, _ := regexp.Compile(`<div class="nav">1 LTC ≈ (.*?) CNY</div>`)	
	res2 := r2.FindAllString(body, -1)
	fmt.Println("PRICE", res2)
	// fmt.Println("PRICE", res2[0])
	fmt.Println("****************")
}

// VTC data
func VTC() {	
	url := "https://www.feixiaohao.com/currencies/vtcoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`</span> <span class="info_val"><span class="convert">(.*?)</span></span></div> <div class="info_cell"><span class="info_tit">`)
	res := r.FindAllString(body, -1)	
	fmt.Println("****************VTC")
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])

	r2, _ := regexp.Compile(`<span class="convert textRed">(.*?)</span>`)
	res2 := r2.FindAllString(body, -1)
	fmt.Println("PRICE", res2[0])
	fmt.Println("****************")
}

// XAU data
func XAU() {	
	url := "https://coinmarketcap.com/zh/currencies/bitcoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`<span data-currency-value>(.*?)</span>`)
	res := r.FindAllString(body, -1)
	// fmt.Println(len(r.FindAllString(body, -1)))
	fmt.Println("****************BTC")
	fmt.Println("PRICE", res[2])
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])
	fmt.Println("****************")
}

// OuYuanToYingBang data
func OuYuanToYingBang() {	
	url := "https://coinmarketcap.com/zh/currencies/bitcoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`<span data-currency-value>(.*?)</span>`)
	res := r.FindAllString(body, -1)
	// fmt.Println(len(r.FindAllString(body, -1)))
	fmt.Println("****************BTC")
	fmt.Println("PRICE", res[2])
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])
	fmt.Println("****************")
}

// YingBangToAoYuan data
func YingBangToAoYuan() {	
	url := "https://coinmarketcap.com/zh/currencies/bitcoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`<span data-currency-value>(.*?)</span>`)
	res := r.FindAllString(body, -1)
	// fmt.Println(len(r.FindAllString(body, -1)))
	fmt.Println("****************BTC")
	fmt.Println("PRICE", res[2])
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])
	fmt.Println("****************")
}

// YingBangToMeiYuan data
func YingBangToMeiYuan() {	
	url := "https://coinmarketcap.com/zh/currencies/bitcoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`<span data-currency-value>(.*?)</span>`)
	res := r.FindAllString(body, -1)
	// fmt.Println(len(r.FindAllString(body, -1)))
	fmt.Println("****************BTC")
	fmt.Println("PRICE", res[2])
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])
	fmt.Println("****************")
}

// Gold data
func Gold() {	
	url := "https://coinmarketcap.com/zh/currencies/bitcoin/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	// fmt.Println(resp.StatusCode)	
	// fmt.Println(body)
	r, _ := regexp.Compile(`<span data-currency-value>(.*?)</span>`)
	res := r.FindAllString(body, -1)
	// fmt.Println(len(r.FindAllString(body, -1)))
	fmt.Println("****************BTC")
	fmt.Println("PRICE", res[2])
	fmt.Println("MAX", res[5])
	fmt.Println("MIN", res[6])
	fmt.Println("****************")
}


