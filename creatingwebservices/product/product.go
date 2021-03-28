package product

type Product struct {
	ProductID      int    `json:"productID"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityAtHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}
