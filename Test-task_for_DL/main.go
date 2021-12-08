package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	//"strconv"
	"strings"
	"time"
)


type Variation struct {
	LastStationArr[] int
	Arr []int64
	NewDurationArr[] time.Duration

}

type VariationInterface interface {

	MinPrice() int
	MinTime() time.Duration
	Counter() int

}

func MinTime(arr[] time.Duration) time.Duration {

	var Res time.Duration
	for _, l := range arr {
		if l <= Res {
			Res = l

		}
	}
	return Res
}

func (v Variation) MinTime() time.Duration {

	var Res = v.NewDurationArr[0]
	for _, l := range v.NewDurationArr {
		if l > Res {
			l = Res
		}
		continue

	}
	return Res
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

	var Stage1 string
	var Stage2 string
	fmt.Print("Введіть станцію 1", "\n")
	fmt.Fscan(os.Stdin, &Stage1)
	fmt.Print("Введіть станцію 2", "\n")
	fmt.Fscan(os.Stdin, &Stage2)
	//fmt.Print(Stage1," - ",Stage2,"\n")

	//	k := Variation{stage1: Stage1, stage2: Stage2}
	//	l, _ := Variation.MinPrice(k)
	//fmt.Print(l)
	//fmt.Print(Variation.Counter(c))

	NewArrData := make([][]string, 250)
	NewPart1ArrData := make([][]string, 250)
	//NewPart2ArrData := make([][]string, 250)
	//NewStationArr := make([][] int64, 250)
	NewDurationArr := make([]time.Duration,250)
	//res := NewDurationArr[0]
	//var N time.Duration
	var Z time.Duration


	lines, err := readLines("test_task_data.csv")
	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
		return
	}

	for i, line := range lines {
		NewArrData[i] = append(NewArrData[i], line)
		NewArrData[i] = strings.Split(line, ";")

	}

	for i, line1 := range NewArrData {
		NewPart1ArrData[i] = append(NewPart1ArrData[i], line1[1], line1[2], line1[4], line1[5])

		for range line1 {


			if line1[1] == Stage1 {
				if line1[2] == Stage2 {

					p, _ := time.Parse("15:04:5", line1[4])
					f, _ := time.Parse("15:04:5", line1[5])

					C := p.Sub(f)
					NewDurationArr = append(NewDurationArr, C)
					//Variation1 := Variation{NewDurationArr: []time.Duration{C}}

					 //:= Variation.MinTime(Variation1)
					Z = MinTime(NewDurationArr)
					//Variation1 := Variation{NewDurationArr: []time.Duration{C}}
					//n := Variation.MinTime(Variation1)
					//fmt.Print("Співпало і 1 і 2", " - ", line1[1], ":", line1[2], " - ", Stage1, ":", Stage2, " - Час - ",z, "\n")

					break
				}
			}
		}
	}

	fmt.Print("\n Станція відправлення - ", Stage1 ," : Станція прибуття - ",Stage2," - мінімальний час ", Z, "\n")



	//fmt.Print("Співпало і 1 і 2", " - ",z, "\n")

}