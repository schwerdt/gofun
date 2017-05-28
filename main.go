package main

import(
  "fmt"
  "time"

  "github.com/schwerdt/gofun/loan"
)

//Build amortizer that takes in principal, annual interest rate, number of payments, frequency of payments (dates or should we generate this?)

func main() {
  start_date := time.Date(2017, time.May, 23, 0, 0, 0, 0,time.UTC)
  disbursement_date := time.Date(2017, time.May, 18, 0, 0, 0, 0, time.UTC)
//test_date := time.Date(2017, time.May, 35, 0,0, 0, 0, time.UTC)
//fmt.Println("what is test_date", test_date)
//year, month, day := start_date.Date()
//fmt.Println("start date: ", year, month, day)
//fmt.Println("time now is: ", time.Now())
  loan := loan.Loan{ Yearly_interest_rate: 0.5, Principal: 2000.00, Num_installments: 12, Payment_frequency: "monthly", Day_of_month: 30, Start_date: start_date, Disbursement_date: disbursement_date, Draw_fee_percent: 0.01 }
//loan.buildSchedule()
//loan.computeIntervalPayment(7)
//fmt.Println("what is the schedule in main:", loan.schedule.due_dates)
//overpayment := loan.Solve()
//fmt.Println("what is the overpayment", overpayment)
//loan.schedule.payments[loan.num_installments - 1] = round(loan.schedule.payments[loan.num_installments - 1] - overpayment, 2)
  loan.CalculatePaymentSchedule()
  for i:=0; i < loan.Num_installments; i++ {
    fmt.Println("Payment amounts", loan.Schedule.Payments[i])
  }
}
