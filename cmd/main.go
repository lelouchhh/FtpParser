package main

import (
	"../pkg"
	"flag"
	"fmt"
	"github.com/secsy/goftp"
	"os"
)
func main()  {
	if len(os.Args) == 1{
		fmt.Println("Используйте 'cmd.exe -h' для получения справки")
		os.Exit(1)
	}
	name := flag.String("region", "Dagestan_Resp", "Строка с названием региона")
	day := flag.String("day", "2022061700", "Дата для выгрузки\nФормат: ГодМесяцДень00\nнапример: 2022061700")
	output := flag.String("output", "./", "Папка для выгрузки")
	flag.Parse()
	config := goftp.Config{
		User:            "free",
		Password:        "free",
		ActiveTransfers: true,
	}
	path := fmt.Sprintf("/fcs_regions/%s/notifications/currMonth/", *name)
	ftpConn, _:= goftp.DialConfig(config, "ftp.zakupki.gov.ru:21")
	c := pkg.GetFolders(ftpConn, fmt.Sprintf("/fcs_regions/%s/notifications/currMonth/", *name))
	if len(c) == 0{
		os.Exit(1)
	} else{
		nextDay := pkg.GetNextDay(*day)
		list := pkg.MatchingFolder(c, fmt.Sprintf("notification_%s_%s_%s_\n",*name, *day, nextDay))
		pkg.Write(ftpConn, list, path, *output)}
		os.Exit(0)
}
