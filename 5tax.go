package main

import(
  "fmt"
  "time"
)

type TaxRange struct {
  min float64
  max float64
}

var level1 = TaxRange{0,18200}
var level2 = TaxRange{18201,37000}
var level3 = TaxRange{37001,80000}
var level4 = TaxRange{80001,180000}
var level5 = TaxRange{180001,1000000} //max of level5 is redundant, so just fill with any value

func ValueChecker(i float64) float64{
  if (i>level1.max && i < level2.min) ||
  (i>level2.max && i < level3.min) ||
  (i>level3.max && i < level4.min) ||
  (i>level4.max && i < level5.min) {
    return float64(int(i))
  } else {
    return i
  }
}

func CalculateTax(taxIncome float64) float64 {
  var taxAmt float64

  taxIncome = ValueChecker(taxIncome)

  //Range is based on the public information of AU government
  //It doesn't assume that the income amount has a decimal value in between 0 and 1

  switch {
  case taxIncome <= level1.max:
    taxAmt = 0
  case taxIncome <= level2.max && taxIncome >= level2.min:
    taxAmt = (taxIncome-level1.max)*0.19
  case taxIncome <= level3.max && taxIncome >= level3.min:
    taxAmt = 3572 + (taxIncome-level2.max)*0.325
  case taxIncome <= level4.max && taxIncome >= level4.min:
    taxAmt = 17547 + (taxIncome-level3.max)*0.37
  case taxIncome >= level5.min:
    taxAmt = 54547 + (taxIncome-level4.max)*0.45
  }

  return taxAmt
}

func main(){
  var income, tax float64

  now := time.Now()
  fmt.Println("Today's date is",now)

  income = 82000
  tax = CalculateTax(income)

  fmt.Println("In", now.Year(),"I have to pay tax amounted to AUD", tax)
  fmt.Println("The tax is", tax/income, "% from the total amount")

}
