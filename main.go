package main

import (
	"flag"
	"fmt"
	"math"
)

// 提前还款之后，如果是月供不变，获取剩余的还款期限
func getPaymentPerid(leftPrincipal, interestRate, paymentPerMonth float64) float64 {
	v := math.Log(paymentPerMonth/(paymentPerMonth-interestRate*leftPrincipal)) / math.Log(1+interestRate)
	//fmt.Println("v:", v, "leftPrincipal:", leftPrincipal, "interestRate:", interestRate, "paymentPerMonth:", paymentPerMonth)
	return math.Floor(v)
	//return math.Ceil(v)
}

func main() {
	var principal float64 // 贷款本金
	var years int
	var months int
	var rate float64
	var start int
	var prepayment float64
	var duration int
	var mode string

	flag.Float64Var(&principal, "p", 1140000.0, "贷款本金")
	flag.IntVar(&years, "y", 20, "贷款年限")
	flag.Float64Var(&rate, "r", 0.0475, "贷款年利率")
	flag.IntVar(&months, "m", 0, "贷款月数")
	flag.IntVar(&start, "s", 1, "已还月数")
	flag.Float64Var(&prepayment, "e", 300000.0, "提前还款金额")
	flag.IntVar(&duration, "d", 12, "每隔多少个月提前还款")
	flag.StringVar(&mode, "o", "fixed", "还款之后，fixed:月供不变 other:期数不变 ")
	flag.Parse()

	interestRate := rate / 12   // 年利率，月利率 = 年利率 / 12
	paymentPeriod := years * 12 // 还款期限，以月为单位
	if months != 0 {
		paymentPeriod = months
	}

	paymentPerMonth := 0.0  // 每月还款额
	oriTotalInterest := 0.0 // 总利息
	totalPayment := 0.0     // 还款总额

	// 每月还款额:贷款本金×[月利率×(1+月利率)^还款月数]÷[(1+月利率)^还款月数-1]
	paymentPerMonth = (principal * interestRate * math.Pow(1+interestRate,
		float64(paymentPeriod))) / (math.Pow(1+interestRate, float64(paymentPeriod)) - 1)

	// 计算总利息和还款总额
	oriTotalInterest = paymentPerMonth*float64(paymentPeriod) - principal
	totalPayment = paymentPerMonth * float64(paymentPeriod)

	// 输出计算结果
	fmt.Println("每月还款额(元), 总利息(元), 还款总额(元)")
	fmt.Printf("%.2f, %.2f, %.2f\n", paymentPerMonth, oriTotalInterest, totalPayment)

	totalInterest := 0.0
	cnt := 1
	left := principal
	fmt.Println("期, 月供(元), 本金(元), 利息(元), 剩余贷款(元)")
	totalPaymentReal := 0.0
	// 打印每个月的月供和利息
	for i := start; i <= paymentPeriod; i++ {
		extra := 0
		// 每月应还本金=贷款本金×月利率×(1+月利率)^(还款月序号-1)÷〔(1+月利率)^还款月数-1〕
		principalPerMon := (principal * interestRate * math.Pow(1+interestRate, float64(i-1))) /
			(math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
		//每月应还利息：贷款本金×月利率×〔(1+月利率)^还款月数-(1+月利率)^(还款月序号-1)〕÷〔
		// 			(1+月利率)^还款月数-1〕
		monthInterest := principal * interestRate * (math.Pow(1+interestRate,
			float64(paymentPeriod)) - math.Pow(1+interestRate, float64(i-1))) /
			(math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
		totalInterest += monthInterest
		// 剩余贷款
		left -= principalPerMon
		if i == duration {
			if left < prepayment {
				fmt.Println("还款结束")
				paymentPeriod = 0
				paymentPerMonth = left + principalPerMon
				principalPerMon = 0.0
				monthInterest = 0.0
				left = 0.0
				extra = 0.0
			} else {
				extra = int(prepayment)
				left -= prepayment
				principal = left
				if mode == "fixed" {
					paymentPeriod = int(getPaymentPerid(left, interestRate, paymentPerMonth))
					//fmt.Println("选择月供不变，缩短期限，新的还款期限:", paymentPeriod)
				} else {
					paymentPeriod -= i
				}
				paymentPerMonth = (principal * interestRate * math.Pow(1+interestRate,
					float64(paymentPeriod))) / (math.Pow(1+interestRate, float64(paymentPeriod)) - 1)
				i = 0
			}
		}

		totalPaymentReal += paymentPerMonth + float64(extra)
		fmt.Printf("%03d, %.2f, %.2f, %.2f, %.2f\n", cnt, paymentPerMonth+float64(extra), principalPerMon, monthInterest, left)
		cnt++
	}
	fmt.Printf("实际总支出: %.2f 实际总利息:%.2f\n", totalPaymentReal, totalInterest)
}
