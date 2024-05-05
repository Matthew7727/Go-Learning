package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)

type TaxInclPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxInclPrices map[string]float64
}

func NewTaxInclPriceJob(taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{10 ,20 ,30},
		TaxRate: taxRate,
	}
}

func (job *TaxInclPriceJob) Process() {

	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInclPrice := price * (1+job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInclPrice)
	}

	fmt.Println(result)
}

func (job *TaxInclPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Cold not open File")
		fmt.Println(err)
		return
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for index, line := range lines{
		prices[index], err = strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("An Error occured during type conversion")
			fmt.Println(err)
			file.Close()
			return
		}
	}

	job.InputPrices = prices
}
