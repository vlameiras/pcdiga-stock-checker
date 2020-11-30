package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	urls := []string{
		"https://www.pcdiga.com/placa-grafica-msi-geforce-rtx-3070-ventus-3x-8g-oc",
		"https://www.pcdiga.com/placa-grafica-msi-geforce-rtx-3070-gaming-x-trio-8g",
		"https://www.pcdiga.com/placa-grafica-msi-geforce-rtx-3070-ventus-2x-8g-oc",
		"https://www.pcdiga.com/placa-grafica-gigabyte-geforce-rtx-3070-gaming-8gb-gddr6-oc",
		"https://www.pcdiga.com/placa-grafica-gigabyte-geforce-rtx-3070-vision-oc-8g-gddr6",
		"https://www.pcdiga.com/placa-grafica-pny-geforce-rtx-3070-8gb-gddr6-dual-fan",
		"https://www.pcdiga.com/placa-grafica-asus-dual-geforce-rtx-3070-8gb-gddr6",
		"https://www.pcdiga.com/placa-grafica-asus-rog-strix-geforce-rtx-3070-8gb-gddr6",
		"https://www.pcdiga.com/placa-grafica-gigabyte-geforce-rtx-3070-aorus-master-8g-gddr6",
		"https://www.pcdiga.com/placa-grafica-msi-geforce-rtx-3070-suprim-x-8g",
		"https://www.pcdiga.com/placa-grafica-zotac-gaming-geforce-rtx-3070-8gb-gddr6-twin-edge-oc",
		"https://www.pcdiga.com/placa-grafica-zotac-gaming-geforce-rtx-3070-8gb-gddr6-twin-edge",
		"https://www.pcdiga.com/placa-grafica-asus-tuf-gaming-geforce-rtx-3070-8gb-gddr6-oc-edition",
		"https://www.pcdiga.com/placa-grafica-pny-geforce-rtx-3070-8gb-gddr6-xlr8-gaming-epic-x-rgb-triple-fan",
		"https://www.pcdiga.com/placa-grafica-asus-dual-geforce-rtx-3070-8gb-gddr6-oc-editon",
		"https://www.pcdiga.com/placa-grafica-asus-rog-strix-geforce-rtx-3070-8gb-gddr6-oc-editon",
		"https://www.pcdiga.com/placa-grafica-gigabyte-geforce-rtx-3070-eagle-8gb-gddr6-oc",
	}

	for {
		var productsStock []string
		for _, url := range urls {

			resp, err := http.Get(url)
			if err != nil {
				log.Info().Msgf("Got an error %s", err)
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				log.Fatal().Msgf("Error parsing HTML for %s", url)
			}

			var productName string
			doc.Find(".page-title").Each(func(i int, s *goquery.Selection) {
				productName = strings.TrimSpace(strings.TrimLeft(s.Text(), "\n"))
			})

			doc.Find(".page-wrapper").Each(func(i int, s *goquery.Selection) {
				hasStock := strings.Contains(s.Text(), "'is_in_stock': 1,")
				if hasStock {
					log.Info().Msgf("Stock for %s !!!", productName)
					productsStock = append(productsStock, productName)
				}

			})
		}
		if len(productsStock) == 0 {
			log.Info().Msg("No stock found...Trying again in 5 minutes")
		}
		time.Sleep(5 * time.Minute)
	}
}
