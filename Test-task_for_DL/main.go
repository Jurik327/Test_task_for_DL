package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)


type Variation struct {
	LastStationArr[] float64
	Arr []float64
	arr2 [] time.Time

}

type VariationInterface interface {

	MinPrice() float64
	MinTime() time.Time
}



func (v Variation) MinPrice() float64 {
	res := v.Arr[0]
	for _ , c := range v.Arr {
	if c < res {
		res = c
		}
}

	return res


}

func (v Variation) MinTime() []time.Time {

	return v.arr2
}

func readLines(path string) (lines []string, err error) {
	var (
		file *os.File
		part []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func writeLines(lines []string, path string) (err error) {
	var (
		file *os.File
	)

	if file, err = os.Create(path); err != nil {
		return
	}
	defer file.Close()

	for _,item := range lines {
		_, err := file.WriteString(strings.TrimSpace(item) + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return
}



func main() {

	//var i interface{}
	NewArrData := make([][]string, 250)
	NewPart1ArrData := make([][]string, 250)
	NewPart2ArrData := make([][]string, 250)
	MinPriceArr := make([]float64,250)

	lines, err := readLines("test_task_data.csv")
	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
		return
	}

	for i, line := range lines {
		NewArrData[i] = append(NewArrData[i], line)
		NewArrData[i] = strings.Split(line, ";")

	}
	//fmt.Print(NewArrData)

	for i, line1 := range NewArrData {
		NewPart1ArrData[i] = append(NewPart1ArrData[i], line1[2], line1[3])
		if line1[2] == line1[2] {
			for range line1 {

				p, _ := strconv.ParseFloat(line1[3], 64)

				Variation1 := Variation{Arr: []float64{p}}

				 c := Variation.MinPrice(Variation1)
				MinPriceArr = append(MinPriceArr, c)

				fmt.Print("Мінімальна ціна - ", line1[2], " - ",c, "\n")

break

			}

		}

	}

	for i, line2 := range NewArrData {
		NewPart2ArrData[i] = append(NewPart2ArrData[i], line2[2], line2[4], line2[5])
		if line2[2] == line2[2] {
			p, _ := time.Parse("15:04:5", line2[4])
			f, _ := time.Parse("15:04:5", line2[5])
			//c := p.Format("15:04:50")
			//h := f.Format("15:04:50")

			Variation2 := Variation{arr2: []time.Time{p,f}}
			c := Variation.MinTime(Variation2)

fmt.Print("Час - ",line2[2]," - ",c,"\n")
		}
	}

}