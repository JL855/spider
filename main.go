package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"main.go/spider"
	"strconv"
)

func main() {
	//创建新的xlsx
	file := xlsx.NewFile()
	//创建新的sheet
	sheet, err := file.AddSheet("spider")
	if err != nil {
		fmt.Printf(err.Error())
	}
	//添加一行
	row := sheet.AddRow()
	//添加一格
	cell1 := row.AddCell()
	cell1.Value = "title"
	cell2 := row.AddCell()
	cell2.Value = "date"
	cell3 := row.AddCell()
	cell3.Value = "readNum"
	cell4 := row.AddCell()
	cell4.Value = "picture"
	for i:=1;i<60;i++ {
		fmt.Printf("正在爬取第 %d 页信息\n",i)
		spider.Spider(strconv.Itoa(i), sheet)
	}
	//保存xlsx文件
	err = file.Save("spider.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

//输出到excel的模块还可弄成一个demo

/*func TestExcel(t *testing.T) {
	//创建新的xlsx
	file := xlsx.NewFile()
	//创建新的sheet
	sheet, err := file.AddSheet("spider")
	if err != nil {
		fmt.Printf(err.Error())
	}
	//添加一行
	row := sheet.AddRow()
	//添加一格
	cell1 := row.AddCell()
	cell1.Value = "a"
	cell2 := row.AddCell()
	cell2.Value = "b"
	cell3 := row.AddCell()
	cell3.Value = "c"

	//保存xlsx文件
	err = file.Save("1.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}*/

