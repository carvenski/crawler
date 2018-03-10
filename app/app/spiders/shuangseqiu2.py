import scrapy
import re
import urls


# crawl shuangseqiu data from http://kaijiang.500.com
class DataSpider(scrapy.Spider):
    name = "data"
    start_urls = map(
        lambda i: 'http://kaijiang.500.com/shtml/ssq/{0}.shtml'.format(i), 
        urls.urls)


    def parse(self, response):
        # use regex to filter data from body text.
        r = response.css("body").re('<li class="ball_red">[0-9]{2}</li>')
        r = map(lambda i: re.search('<li class="ball_red">([0-9]{2})</li>', i).group(1), r)
        r = '$'.join(r) + "\n"
        print("*"*100)
        print(r)

        with open("data.txt", 'a') as f:
            f.write(r)

