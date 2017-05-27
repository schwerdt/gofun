package main

import(
  "fmt"
  "time"
//"math"

  "github.com/schwerdt/gofun/loan"
)

//Build amortizer that takes in principal, annual interest rate, number of payments, frequency of payments (dates or should we generate this?)

//type LoanData struct {
//  yearly_interest_rate float64
//  principal float64
//  draw_fee_percent float64
//  disbursement_date time.Time
//  num_installments int
//  payment_frequency string
//  start_date time.Time
//  schedule Schedule
//}

//type Schedule struct {
//  due_dates []time.Time
//  payments []float64
//}

//type BucketTotals struct {
//  fee_current float64
//  interest_current float64
//  principal_current float64
//  fee_paid float64
//  interest_paid float64
//  principal_paid float64
//}

//func (running_totals *BucketTotals) resetValues(loan LoanData) {
//  running_totals.principal_current = loan.principal
//  running_totals.fee_current = loan.principal * loan.draw_fee_percent
//  running_totals .interest_current = 0.0
//  running_totals.principal_paid = 0.0
//  running_totals.fee_paid = 0.0
//  running_totals.interest_paid = 0.0
//}

//func (running_totals *BucketTotals) computeOverPayment(loan LoanData) float64 {
//  total_paid := running_totals.principal_paid + running_totals.interest_paid + running_totals.fee_paid
//  fmt.Println("what is total_paid", total_paid)
//  total_payments := 0.0
//  for i := 0; i < loan.num_installments; i++ {
//    total_payments += loan.schedule.payments[i]
//  }
//  fmt.Println("what is total_payments", total_payments)
//  return (total_payments - total_paid)
//}


//func (loan *LoanData) buildSchedule() {
//  // weekly, add 7 days, biweekly add 14 days, monthly choose same date each month (logic for when date does not exist + later when it is a weekend)
//  switch loan.payment_frequency {
//  case "weekly":
//    loan.buildScheduleDueDatesNDays(loan.start_date, loan.num_installments, 7)
//  case "biweekly":
//    loan.buildScheduleDueDatesNDays(loan.start_date, loan.num_installments, 14)
//  case "monthly":
//    loan.buildScheduleMonthlyDueDates()
//  }
//}

//func (loan *LoanData) buildScheduleMonthlyDueDates() {
//}

//func (loan *LoanData) buildScheduleDueDatesNDays(start_date time.Time, num_installments int, date_interval int) {
//  due_dates := make([]time.Time, num_installments)
//  year, month, day := start_date.Date()

//  for i := 0; i < num_installments; i++ {
//    end_day := day + i * date_interval
//    due_dates[i] = time.Date(year, month, end_day, 0, 0 ,0 ,0, time.UTC)
//  }
//  loan.schedule = Schedule{ due_dates: due_dates }
//  fmt.Println("what are due_dates", due_dates)
//}

//func (loan *LoanData) computeIntervalPayment(period_length int) {
//  per_period_interest_rate := float64(period_length) * loan.yearly_interest_rate / 365.25
//  numerator := per_period_interest_rate * math.Pow((1 + per_period_interest_rate), float64(loan.num_installments))
//  denominator := ( math.Pow((1 + per_period_interest_rate), float64(loan.num_installments)) - 1 )
//  payment_amount := loan.principal * numerator / denominator
//  payment_amount = round(payment_amount, 2)

//  payment_array := make([]float64, loan.num_installments)
//  for i :=0; i< len(payment_array); i++ {
//    payment_array[i] = payment_amount }
//  loan.schedule.payments = payment_array
//  fmt.Println("what is the schedule", loan.schedule.payments)
//}

//func (loan *LoanData) Solve() float64 {
//  running_totals := &BucketTotals{ principal_current: loan.principal, fee_current: loan.principal * loan.draw_fee_percent }
//  var overpayment float64
//  for i := 0; ; i++ {
//    running_totals = loan.simulateLoanLife(running_totals)
//    underpayment := (loan.principal - running_totals.principal_paid)
//    overpayment = running_totals.computeOverPayment(*loan)
//    fmt.Println("what is overpayment: ", overpayment)
//    fmt.Println("what is underpayment: ", underpayment)

//    if  overpayment > 0.5 || underpayment > 0.001 {
//      delta := 0.0
//         if underpayment > 0.001 {
//           delta = underpayment / float64(loan.num_installments)
//           fmt.Println("what is underpayment", underpayment)
//          } else {
//            delta = -1*overpayment / float64(loan.num_installments)
//            fmt.Println("what is overpayment, num_installments", overpayment, loan.num_installments)
//          }

//      delta = round(delta, 2)
//      if delta == 0.0 {
//        delta += 0.01 }
//      fmt.Println("what is delta", delta)
//      for j :=0; j < loan.num_installments; j++ {
//        loan.schedule.payments[j] = round((loan.schedule.payments[j] + delta), 2)
//      }
//    running_totals.resetValues(*loan)
//    } else { break }
//  }
//  return overpayment

//}

//func (loan *LoanData) simulateLoanLife(running_totals *BucketTotals) *BucketTotals {
////running_totals := &BucketTotals{ principal_current: loan.principal, fee_current: loan.principal * loan.draw_fee_percent }
//  start_year, start_month, start_day := loan.disbursement_date.Date()
//  // Start day after disbursement and simulate interest and payments
//  for i := 1;; i++ {
//    day_offset := start_day + i
//    today := time.Date(start_year, start_month, day_offset, 0, 0, 0, 0, time.UTC)

//    if greaterThanDate(today,loan.schedule.due_dates[loan.num_installments - 1]) {
//      break }
//    // Compute interest for today
//    interest_today := running_totals.principal_current * loan.yearly_interest_rate/365.25
//    running_totals.interest_current += roundDown(interest_today, 2)
//    fmt.Println("today is", today, round(interest_today, 2))
//    // Apply any payments for today
//    running_totals.applyAnyPayments(today, loan.schedule)
//  }

//  fmt.Println("what is running_totals", running_totals)
//  return running_totals

//}

//func (totals *BucketTotals) applyAnyPayments(today time.Time, schedule Schedule) float64 {
//  payment_amount := 0.0
//  for i := 0; i< len(schedule.due_dates); i ++ {
//    if equalDates(today,schedule.due_dates[i]) == true {
//      payment_amount = schedule.payments[i]
//      fmt.Println("there is a payment today", today, payment_amount)
//      if totals.fee_current < payment_amount {
//        totals.fee_paid += totals.fee_current
//        payment_amount -= totals.fee_current
//        totals.fee_current = 0.0
//      } else {
//        totals.fee_paid += payment_amount
//        totals.fee_current -= payment_amount
//        payment_amount = 0.0
//     }
//     if totals.interest_current < payment_amount {
//       totals.interest_paid += totals.interest_current
//       payment_amount -= totals.interest_current
//       totals.interest_current = 0.0
//     } else {
//       totals.interest_paid += payment_amount
//       totals.interest_current -= payment_amount
//       payment_amount = 0.0
//     }
//     if totals.principal_current < payment_amount {
//       totals.principal_paid += totals.principal_current
//       payment_amount -= totals.principal_current
//       totals.principal_current = 0.0
//     } else {
//       totals.principal_paid += payment_amount
//       totals.principal_current -= payment_amount
//       payment_amount = 0.0
//    }
//  }
//  }
//   return payment_amount //This is the overpayment
//}

//func equalDates(date1 time.Time, date2 time.Time) bool {
//  year1, month1, day1 := date1.Date()
//  year2, month2, day2 := date2.Date()
//  if year1 == year2 && month1 == month2 && day1 == day2 {
//    return true
//  } else {
//    return false
//  }
//}

//func greaterThanDate(date1 time.Time, date2 time.Time) bool {
//  year1, month1, day1 := date1.Date()
//  year2, month2, day2 := date2.Date()
//  if year1 > year2 {
//    return true
//  } else if year1 == year2 && month1 > month2 {
//    return true
//  } else if year1 == year2 && month1 == month2 && day1 > day2 {
//    return true
//  } else {
//    return false
//  }
//}

//  func round(val float64, num_decimals int) float64 {

////val_sign := val / math.Abs(val)
////val = val*100
////val = math.Floor(math.Abs(val))
////fmt.Println("what is floored value", val)
////return val_sign * val / 100.0
//  if val < 0.0 {
//    val -= 0.005
//    val = val * 100.0
//    val = math.Ceil(val)
//    val = val / 100.0
//  } else {
//    val += 0.005
//    val = val * 100.0
//    val = math.Floor(val)
//    val = val / 100.0
//   }
//  return val

//}

//func roundDown(val float64, num_decimals int) float64 {
//  val += 0.001  //Just to deal with precision 0.019999999
//  val = math.Floor(100*val)
//  val = val / 100.0
//  return val
//}


func main() {
  start_date := time.Date(2017, time.May, 23, 0, 0, 0, 0,time.UTC)
  disbursement_date := time.Date(2017, time.May, 18, 0, 0, 0, 0, time.UTC)
//test_date := time.Date(2017, time.May, 35, 0,0, 0, 0, time.UTC)
//fmt.Println("what is test_date", test_date)
//year, month, day := start_date.Date()
//fmt.Println("start date: ", year, month, day)
//fmt.Println("time now is: ", time.Now())
  loan := loan.Loan{ Yearly_interest_rate: 0.5, Principal: 2000.00, Num_installments: 26, Payment_frequency: "weekly", Start_date: start_date, Disbursement_date: disbursement_date, Draw_fee_percent: 0.01 }
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
