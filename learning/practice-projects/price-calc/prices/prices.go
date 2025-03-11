package prices

import (
	"fmt"

	conversion "exapmle.com/practice-cacl/conversion"
	iomanager "exapmle.com/practice-cacl/iomanager"
)

type TaxInclPriceJob struct {
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxInclPrices map[string]string `json:"tax_incl_prices"`
	IOManager iomanager.IOManager `json:"-"`
}

func NewTaxInclPriceJob(iom iomanager.IOManager, taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{10 ,20 ,30},
		TaxRate: taxRate,
		IOManager: iom,
	}
}

func (job *TaxInclPriceJob) Process()  error{

	err  := job.LoadData()

	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInclPrice := price * (1+job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInclPrice)
	}

	job.TaxInclPrices = result
	return job.IOManager.WriteJson(job)

}

func (job *TaxInclPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println("An Error occured during type conversion")
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices
	return nil
}
