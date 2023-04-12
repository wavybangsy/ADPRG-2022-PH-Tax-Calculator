/********************
Name:
	Cabungcal, Mary Joselle
	Culala, Mary Erika
	Tapia, John Lorenzo
Language: Go
Paradigm: Procedural
********************/

package main

import (
	"fmt"
	"math"
)

func philHealthCalc(salary float32) float32 {
	// if salary is 10k or below
	if salary > 0 && salary <= 10000 {
		return 400 / 2
		// if salary is over 10k and below 90k
	} else if salary > 10000.00 && salary <= 89999.99 {
		return ((salary * 0.040) / 2)
		// if salary 90k and above
	} else {
		return 3200.00 / 2
	}
}

/**
* TODO: Fill out the function
**/
func SSScalc(salary float32) float32 {
	var sss float32 = 135.00
	var multiplier int = 1
	if salary >= 24750 {
		sss = 1125.0
	} else if salary >= 3250 {
		salary -= 3250.0
		multiplier += int(math.Floor(float64(salary) / float64(500)))
		sss += float32(multiplier) * 22.50
	}
	return sss
}

func PAGIBIGcalc(salary float32) float32 {
	// maximum limit that PAGIBIG includes in the computation from employee's SALARY
	var maxLimit float32 = 5000

	// if salary is less than max
	if salary > 0 && salary < 5000 {

		// if salary is < max (5000) and less than 1500 pesos
		if salary > 0 && salary < 1500 {
			return salary * 0.01
			// if salary is < max (5000) and over 1500 pesos
		} else if salary >= 1500 {
			return salary * 0.02
		}
	}
	// otherwise, return 100 pesos default.
	return maxLimit * 0.02
}

/**
* TODO: Fill out the function
**/
func wHDTax(taxableIncome float32) float32 {
	switch {
	case taxableIncome >= 20833 && taxableIncome <= 33332:
		return (taxableIncome - 20833) * 0.20
	case taxableIncome >= 33333 && taxableIncome <= 66666:
		return 2500.00 + (taxableIncome-33333)*0.25
	case taxableIncome >= 66667 && taxableIncome <= 166666:
		return 10833.33 + (taxableIncome-66667)*0.30
	case taxableIncome >= 166667 && taxableIncome <= 666666:
		return 40833.33 + (taxableIncome-166667)*0.32
	case taxableIncome >= 666667:
		return 200833.33 + (taxableIncome-666667)*0.35
	default:
		return 0.0
	}
}

func printRes(phl, sss, pgb, tCont, incTax, nPTax, tDeduct, nPDeduct float32) {
	println("")
	println("")
	println("")
	println("\n-------------------------------------------------")
	println("|	MONTHLY CONTRIBUTION			|")
	fmt.Printf("|PhilHealth Contribution: %.2f\t\t|\n", phl)
	fmt.Printf("|SSS Contribution	: %.2f\t\t|\n", sss)
	fmt.Printf("|PAGIBIG Contribution	: %.2f\t\t|\n", pgb)
	fmt.Printf("|Total Contribution	: %.2f\t\t|\n", tCont)
	println("|						|")
	println("|						|")

	println("|	   TAX COMPUTATION			|")
	fmt.Printf("|Income Tax 		: %.2f		|\n", incTax)
	fmt.Printf("|Net Pay after tax	: %.2f		|\n", nPTax)
	println("|						|")
	println("|						|")
	fmt.Printf("|Total Deduction	: %.2f		|\n", tDeduct)
	fmt.Printf("|Net Pay after deduction: %.2f		|", nPDeduct)
	println("\n-------------------------------------------------")
}

func main() {
	var salary float32

	// monthly contributions
	// all are initialized to 0
	var (
		phlContri float32
		pgbContri float32
		sssContri float32
	)

	// extra variables
	var (
		taxableIncome  float32
		withholdingTax float32
	)

	// fields shown in Calculator
	var (
		incomeTax         float32
		totalContribution float32
		totalDeduction    float32
		netPayTaxed       float32
		netPayDeducted    float32
	)

	// placeholder only. To not generate errors for unused variables. DELETE LATER
	_ = incomeTax
	_ = taxableIncome
	_ = withholdingTax
	_ = totalContribution
	_ = totalDeduction
	_ = netPayTaxed
	_ = netPayDeducted

	println("Tax Calculator")
	fmt.Printf("Salary: ")
	fmt.Scan(&salary)

	// Compute for taxable income:

	// call philHealth Calculator function
	phlContri = philHealthCalc(salary)

	// call SSS Calculator function
	sssContri = SSScalc(salary)

	// call PAGIBIG Calculator function
	pgbContri = PAGIBIGcalc(salary)

	totalContribution = phlContri + sssContri + pgbContri
	taxableIncome = salary - totalContribution

	// call wHDTax function to compute for Tax
	incomeTax = wHDTax(taxableIncome)

	// Compute for Net pay after Tax
	netPayTaxed = salary - incomeTax

	// Compute for Total Deductions
	totalDeduction = totalContribution + incomeTax

	// Compute for Net Pay after Deductions
	netPayDeducted = salary - totalDeduction

	printRes(phlContri, sssContri, pgbContri, totalContribution, incomeTax, netPayTaxed, totalDeduction, netPayDeducted)
}
