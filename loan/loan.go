package loan

import(
  "fmt"
  "time"
  "math"
  "github.com/schwerdt/gofun/bucket_totals"
)

type Loan struct {
  Yearly_interest_rate float64
  Principal float64
  Draw_fee_percent float64
  Disbursement_date time.Time
  Num_installments int
  Payment_frequency string
  Day_of_month int
  Start_date time.Time
  Schedule Schedule
}

type Schedule struct {
  Due_dates []time.Time
  Payments []float64
}

func (loan *Loan) estimateSchedule() {
  // weekly, add 7 days, biweekly add 14 days, monthly choose same date each month (logic for when date does not exist + later when it is a weekend)
  switch loan.Payment_frequency {
  case "weekly":
    loan.buildScheduleDueDatesNDays(7)
    loan.approximatePayment(7)
  case "biweekly":
    loan.buildScheduleDueDatesNDays(14)
    loan.approximatePayment(14)
  case "monthly":
    loan.buildScheduleMonthlyDueDates()
    loan.approximatePayment(30)
  }
}

func (loan *Loan) buildScheduleMonthlyDueDates() {
  due_dates := make([]time.Time, loan.Num_installments)
  year, month, _ := loan.Start_date.Date()
  //Move one month ahead
  month += 1
  for i := 0; i< loan.Num_installments ; i++ {
    due_date := time.Date(year, month, loan.Day_of_month, 0, 0, 0, 0, time.UTC)
    due_dates[i] = due_date
    month += 1
  }
  loan.Schedule = Schedule{ Due_dates: due_dates }
  fmt.Println("what are the due_dates", due_dates)
}

func (loan *Loan) buildScheduleDueDatesNDays(date_interval int) {
  due_dates := make([]time.Time, loan.Num_installments)
  year, month, day := loan.Start_date.Date()

  for i := 0; i < loan.Num_installments; i++ {
    end_day := day + i * date_interval
    due_dates[i] = time.Date(year, month, end_day, 0, 0 ,0 ,0, time.UTC)
  }
  loan.Schedule = Schedule{ Due_dates: due_dates }
  fmt.Println("what are due_dates", due_dates)
}

func (loan *Loan) approximatePayment(period_length int) {
  per_period_interest_rate := float64(period_length) * loan.Yearly_interest_rate / 365.25
  numerator := per_period_interest_rate * math.Pow((1 + per_period_interest_rate), float64(loan.Num_installments))
  denominator := ( math.Pow((1 + per_period_interest_rate), float64(loan.Num_installments)) - 1 )
  payment_amount := loan.Principal * numerator / denominator
  payment_amount = round(payment_amount, 2)

  payment_array := make([]float64, loan.Num_installments)
  for i :=0; i< len(payment_array); i++ {
    payment_array[i] = payment_amount }
  loan.Schedule.Payments = payment_array
  fmt.Println("what is the schedule", loan.Schedule.Payments)
}

func (loan *Loan) solve() float64 {
  running_totals := &bucket_totals.BucketTotals{ Principal_current: loan.Principal, Fee_current: loan.Principal * loan.Draw_fee_percent }
//running_totals := &bucket_totals.BucketTotals{}
//running_totals.ResetValues(loan.principal, loan.draw_fee_percent)
  var overpayment float64
  for i := 0; ; i++ {
    running_totals = loan.simulateLoanLife(running_totals)
    underpayment := (loan.Principal - running_totals.Principal_paid)
    overpayment = running_totals.ComputeOverPayment(loan.Num_installments, loan.Schedule.Payments)
    fmt.Println("what is overpayment: ", overpayment)
    fmt.Println("what is underpayment: ", underpayment)

    if  overpayment > 0.5 || underpayment > 0.001 {
      delta := 0.0
         if underpayment > 0.001 {
           delta = underpayment / float64(loan.Num_installments)
           fmt.Println("what is underpayment", underpayment)
          } else {
            delta = -1*overpayment / float64(loan.Num_installments)
            fmt.Println("what is overpayment, num_installments", overpayment, loan.Num_installments)
          }

      delta = round(delta, 2)
      if delta == 0.0 {
        delta += 0.01 }
      fmt.Println("what is delta", delta)
      for j :=0; j < loan.Num_installments; j++ {
        loan.Schedule.Payments[j] = round((loan.Schedule.Payments[j] + delta), 2)
      }
    running_totals.ResetValues(loan.Principal, loan.Draw_fee_percent)
    } else { break }
  }
  return overpayment

}

func (loan *Loan) simulateLoanLife(running_totals *bucket_totals.BucketTotals) *bucket_totals.BucketTotals {
  start_year, start_month, start_day := loan.Disbursement_date.Date()
  // Start day after disbursement and simulate interest and payments
  for i := 1;; i++ {
    day_offset := start_day + i
    today := time.Date(start_year, start_month, day_offset, 0, 0, 0, 0, time.UTC)

    if greaterThanDate(today,loan.Schedule.Due_dates[loan.Num_installments - 1]) {
      break }
    // Compute interest for today
    interest_today := running_totals.Principal_current * loan.Yearly_interest_rate/365.25
    running_totals.Interest_current += roundDown(interest_today, 2)
    fmt.Println("today is", today, round(interest_today, 2))
    // Apply any payments for today
    running_totals.ApplyAnyPayments(today, loan.Schedule.Due_dates, loan.Schedule.Payments)
  }

  fmt.Println("what is running_totals", running_totals)
  return running_totals

}

func (loan *Loan) CalculatePaymentSchedule() {
  loan.estimateSchedule()
  overpayment := loan.solve()
  loan.Schedule.Payments[loan.Num_installments - 1] = round(loan.Schedule.Payments[loan.Num_installments - 1] - overpayment, 2)
}


func round(val float64, num_decimals int) float64 {

//val_sign := val / math.Abs(val)
//val = val*100
//val = math.Floor(math.Abs(val))
//fmt.Println("what is floored value", val)
//return val_sign * val / 100.0
  if val < 0.0 {
    val -= 0.005
    val = val * 100.0
    val = math.Ceil(val)
    val = val / 100.0
  } else {
    val += 0.005
    val = val * 100.0
    val = math.Floor(val)
    val = val / 100.0
   }
  return val

}

func roundDown(val float64, num_decimals int) float64 {
  val += 0.001  //Just to deal with precision 0.019999999
  val = math.Floor(100*val)
  val = val / 100.0
  return val
}

func greaterThanDate(date1 time.Time, date2 time.Time) bool {
  year1, month1, day1 := date1.Date()
  year2, month2, day2 := date2.Date()
  if year1 > year2 {
    return true
  } else if year1 == year2 && month1 > month2 {
    return true
  } else if year1 == year2 && month1 == month2 && day1 > day2 {
    return true
  } else {
    return false
  }
}
