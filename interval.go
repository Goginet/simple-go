package simple_go

import (
  "time"
)

func NowTime() float64 {
  return float64(time.Now().UnixNano()) / 1000000000
}

// Interval - структура для зранения промежутков времени
type Interval struct {
  StartTime, EndTime float64
}

// Start - установить маркер начала на текущее время
func (interval *Interval)Start() {
  interval.StartTime = NowTime()
  return
}

// End - установить маркер конца на текущее время
func (interval *Interval)End() {
  interval.EndTime = NowTime()
  return
}

// GetStart - возвращает время начала отсчета
func (interval Interval)GetStart() float64 {
  return interval.StartTime
}

// GetEnd - возвращает время окончания отсчета
func (interval Interval)GetEnd() float64 {
  return interval.EndTime
}

// Len - возвращает длинну промежутка
func (interval Interval)Len() float64 {
  if interval.EndTime == 0 {
    if interval.StartTime == 0 {
      return 0
    }
    return float64(NowTime() - interval.StartTime)
  } else {
    return float64(interval.EndTime - interval.StartTime)
  }
}
