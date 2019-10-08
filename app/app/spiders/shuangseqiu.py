#encoding=utf8
import scrapy
import re

class UrlSpider(scrapy.Spider):
    name = "btc"
    start_urls = [
        "https://coinmarketcap.com/zh/currencies/bitcoin/",
        "https://coinmarketcap.com/zh/currencies/litecoin/",
        "https://www.feixiaohao.com/currencies/vtcoin/",
    ]

    def parse(self, response):
        if response.url == "https://coinmarketcap.com/zh/currencies/bitcoin/":
            name = "BTC"
            print("*"*100+name)
            # print(response.text)
            # (.*?)
            s = '<span data-currency-value>(.*?)</span>'
            r = response.css("body").re(s)        
            # print(len(r))
            print("PRICE", r[2])
            print("MAX", r[5])
            print("MIN", r[6])
            print("*"*100)

        elif response.url == "https://coinmarketcap.com/zh/currencies/litecoin/":
            name = "LTC"
            print("*"*100+name)
            # print(response.text)
            # (.*?)
            s = '<span data-currency-value>(.*?)</span>'
            r = response.css("body").re(s)        
            # print(len(r))
            print("PRICE", r[2])
            print("MAX", r[5])
            print("MIN", r[6])
            print("*"*100)

        elif response.url == "https://www.feixiaohao.com/currencies/vtcoin/":
            name = "VTC"
            print("*"*100+name)
            # print(response.text)
            # (.*?)
            s = '</span> <span class="info_val"><span class="convert">(.*?)</span></span></div> <div class="info_cell"><span class="info_tit">'
            r = response.css("body").re(s)
            print("MAX", r[5])
            print("MIN", r[6])

            s = '<span class="convert textRed">(.*?)</span>'
            r = response.css("body").re(s)        
            print("PRICE", r[0])
            print("*"*100)

















            
















