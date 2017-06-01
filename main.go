package main

import(
  "fmt"
  "time"
  "io/ioutil"
  "strings"
  "strconv"

  "github.com/schwerdt/gofun/loan"
)

//Build amortizer that takes in principal, annual interest rate, number of payments, frequency of payments (dates or should we generate this?)
func buildDate(date string) time.Time {
  new_date, _ := time.Parse("2006-01-02", date)
  fmt.Println("what is the date", new_date)
  return new_date
}


func buildLoanFromInputs(filename string) loan.Loan {
  var day_of_month int
  file_data, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("File was not found", err) }

  input_data := strings.Split(string(file_data), "\n")
  input_map := make(map[string]string)

  for i :=0; i < len(input_data) - 1; i++ {
    pair := strings.Split(input_data[i], ":")
    key := strings.Replace(pair[0], " ", "", -1)
    value := strings.Replace(pair[1], " ", "", -1)
    input_map[key] = value
  }

  interest_rate, _ := strconv.ParseFloat(input_map["interest_rate"], 64)
  principal, _ := strconv.ParseFloat(input_map["principal"], 64)
  num_installments, _ := strconv.Atoi(input_map["num_installments"])
  disbursement_date := buildDate(input_map["disbursement_date"])
  draw_fee_percent, _ := strconv.ParseFloat(input_map["draw_fee_percent"], 64)
  start_date := buildDate(input_map["start_date"])
  if day_of_month_string, ok := input_map["day_of_month"]; ok {
    day_of_month, _ = strconv.Atoi(day_of_month_string)
   }
  fmt.Println("what is day of month", day_of_month)
  return loan.Loan{ Yearly_interest_rate: interest_rate,
                    Principal: principal,
                    Num_installments: num_installments,
                    Payment_frequency: input_map["payment_frequency"],
                    Disbursement_date: disbursement_date,
                    Draw_fee_percent: draw_fee_percent,
                    Start_date: start_date,
                    Day_of_month: day_of_month }


}



func main() {
  loan := buildLoanFromInputs("loan_data.txt")
  loan.CalculatePaymentSchedule()
  for i:=0; i < loan.Num_installments; i++ {
    fmt.Println("Payment amounts", loan.Schedule.Payments[i])
  }
}
