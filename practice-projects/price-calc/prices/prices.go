package prices

import "fmt"

type TaxInclPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxInclPrices map[string]float64
}

func NewTaxInclPriceJob( taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{10 ,20 ,30},
		TaxRate: taxRate,
	}
}

func (job TaxInclPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1+job.TaxRate)
	}

	fmt.Println(result)
}
