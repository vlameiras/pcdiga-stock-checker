# WARNING
Currently not working due to PCDiga website changes. Will try to fix it ASAP.

# PC Diga Stock Checker

Periodically checks (every 5 minutes) for PCDiga products stock availability.

# Usage
Add your PCDiga parts URL

```go
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
```

Run the application
```bash
go run main.go
```

# Notes
Created for educational purposes, use it responsibly by not faking your `User-Agent` and using a reasonable timeout per check.

Working as of `2020-11-30`
