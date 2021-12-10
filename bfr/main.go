package main

import "fmt"

var (
    name string = "hank"
    age = 44
    sex = 1
    height float64 = 1.78
    weight float64 = 81.2
)

var bmi = weight / (height * height)

var bfr = 1.2 * bmi + 0.23 * float64(age) - 5.4 - 10.8 * float64(sex)


func main() {
if sex == 1 {
    if bfr >= 0 && bfr <= 13 {
        fmt.Println("您的体脂偏瘦")
    } else if bfr > 13 && bfr <= 16 {
        fmt.Println("您的体脂标准")
    } else if bfr > 16 && bfr <= 21 {
        fmt.Println("您的体脂偏重")
    } else if bfr > 21 && bfr <= 27 {
        fmt.Println("您的体脂肥胖")
    } else {
        fmt.Println("您的体脂严重肥胖")
    }
}
}
