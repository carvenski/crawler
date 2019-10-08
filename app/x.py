import requests
import re

url = "https://mifengcha.com/coin/bitcoin"

r = requests.get(url, headers={"User-Agent": "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36"})
# print(r.text)

r = re.search(r.text, '<span data-v-0283d255="">.*</span>')
print( r.group(0) )


