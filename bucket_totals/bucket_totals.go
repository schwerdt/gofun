package bucket_totals

import(
  "fmt"
  "time"
)

type BucketTotals struct {
  Fee_current float64
  Interest_current float64
  Principal_current float64
  Fee_paid float64
  Interest_paid float64
  Principal_paid float64
}

func (running_totals *BucketTotals) ResetValues(principal float64, draw_fee_percent float64) {
  running_totals.Principal_current = principal
  running_totals.Fee_current = principal * draw_fee_percent
  running_totals.Interest_current = 0.0
  running_totals.Principal_paid = 0.0
  running_totals.Fee_paid = 0.0
  running_totals.Interest_paid = 0.0
}

func (running_totals *BucketTotals) ComputeOverPayment(num_installments int, payments []float64) float64 {
  total_paid := running_totals.Principal_paid + running_totals.Interest_paid + running_totals.Fee_paid
  fmt.Println("what is total_paid", total_paid)
  total_payments := 0.0
  for i := 0; i < num_installments; i++ {
    total_payments += payments[i]
  }
  fmt.Println("what is total_payments", total_payments)
  return (total_payments - total_paid)
}

func (totals *BucketTotals) ApplyAnyPayments(today time.Time, due_dates []time.Time, payments []float64) float64 {
  payment_amount := 0.0
  for i := 0; i< len(due_dates); i ++ {
    if equalDates(today,due_dates[i]) == true {
      payment_amount = payments[i]
      fmt.Println("there is a payment today", today, payment_amount)
      if totals.Fee_current < payment_amount {
        totals.Fee_paid += totals.Fee_current
        payment_amount -= totals.Fee_current
        totals.Fee_current = 0.0
      } else {
        totals.Fee_paid += payment_amount
        totals.Fee_current -= payment_amount
        payment_amount = 0.0
     }
     if totals.Interest_current < payment_amount {
       totals.Interest_paid += totals.Interest_current
       payment_amount -= totals.Interest_current
       totals.Interest_current = 0.0
     } else {
       totals.Interest_paid += payment_amount
       totals.Interest_current -= payment_amount
       payment_amount = 0.0
     }
     if totals.Principal_current < payment_amount {
       totals.Principal_paid += totals.Principal_current
       payment_amount -= totals.Principal_current
       totals.Principal_current = 0.0
     } else {
       totals.Principal_paid += payment_amount
       totals.Principal_current -= payment_amount
       payment_amount = 0.0
    }
  }
  }
   return payment_amount //This is the overpayment
}

func equalDates(date1 time.Time, date2 time.Time) bool {
  year1, month1, day1 := date1.Date()
  year2, month2, day2 := date2.Date()
  if year1 == year2 && month1 == month2 && day1 == day2 {
    return true
  } else {
    return false
  }
}
