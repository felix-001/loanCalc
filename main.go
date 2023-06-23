package main

import (
	"fmt"
	"math"
)

func main() {
	// 输入贷款本金、年利率、还款期限
	principal := 500000.0       // 贷款本金
	interestRate := 0.05 / 12   // 年利率，月利率 = 年利率 / 12
	paymentPeriod := 30 * 12    // 还款期限，以月为单位
	paymentPerMonth := 0.0     // 每月还款额
	totalInterest := 0.0       // 总利息
	totalPayment := 0.0        // 还款总额

	// 计算每月还款额
	paymentPerMonth = (principal * interestRate * math.Pow(1+interestRate, float64(paymentPeriod))) / (math.Pow(1+interestRate, float64(paymentPeriod)) - 1)

	// 计算总利息和还款总额
	totalInterest = paymentPerMonth*float64(paymentPeriod) - principal
	totalPayment = paymentPerMonth * float64(paymentPeriod)

	// 输出计算结果
	fmt.Printf("每月还款额为：%.2f 元
", paymentPerMonth)
	fmt.Printf("总利息为：%.2f 元
", totalInterest)
	fmt.Printf("还款总额为：%.2f 元
", totalPayment)
}
