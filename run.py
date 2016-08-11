import os
import time

# Instal ImageScraper using `pip install ImageScraper`

with open('ingredients.txt','r') as f:
	for line in f:
		line = '-'.join(line.split())
		if os.path.isdir("ingredients/"+line):
			continue
		os.system('mkdir ingredients/' + line.strip())
		os.system('image-scraper -s ingredients/'+line.strip()+' http://www.bing.com/images/search?q='+ line.strip()+'&FORM=HDRSC2')
		time.sleep(20)

