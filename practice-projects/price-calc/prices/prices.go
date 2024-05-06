package prices

import (
	"fmt"
	conversion "exapmle.com/practice-cacl/conversion"
	filemanager "exapmle.com/practice-cacl/file-manager"
)

type TaxInclPriceJob struct {
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxInclPrices map[string]string `json:"tax_incl_prices"`
	IOManager filemanager.FileManger `json:"-"`
}

func NewTaxInclPriceJob(fm filemanager.FileManger, taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{10 ,20 ,30},
		TaxRate: taxRate,
		IOManager: fm,
	}
}

func (job *TaxInclPriceJob) Process() {

	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInclPrice := price * (1+job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInclPrice)
	}

	job.TaxInclPrices = result
	job.IOManager.WriteJson(job)
	
	fmt.Println(result)
}

func (job *TaxInclPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println("An Error occured during type conversion")
		fmt.Println(err)
	}

	job.InputPrices = prices
}
