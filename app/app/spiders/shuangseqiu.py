import scrapy
import re

class UrlSpider(scrapy.Spider):
    name = "url"
    start_urls = [
        "http://kaijiang.500.com/shtml/ssq/18022.shtml"
    ]

    def parse(self, response):
        r = response.css("body").re('<a href="http://kaijiang.500.com/shtml/ssq/[0-9]{5}.shtml">[0-9]{5}</a>')
        r = map(lambda i: re.search('<a href="http://kaijiang.500.com/shtml/ssq/([0-9]{5}).shtml">([0-9]{5})</a>', i).group(1), r)
        print("*"*100)
        print(r)
        print("*"*100)

        with open("urls.py", 'w') as f:
            f.write("urls = " + str(r))

