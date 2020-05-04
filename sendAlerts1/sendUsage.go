package main

import (
	"database/sql"
	"fmt"
	_ "github.com/sudesh35139/FourGtest/go-sql-driver/mysql"
	"log"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)
type UsageAlert struct {
	custerName string
	siteName string
	currentUsage float32
	mac string
	usageLimit int8

}

var db *sql.DB

func init() {
	var err error
	//db,err = sql.Open("mysql","root:root@tcp(localhost:3306)/usage?charset=utf8")
	db, err = sql.Open("mysql", "root:rFl5Ye<XW0Sa@tcp(localhost:3306)/usage?charset=utf8")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(200)

	fmt.Println("You connected to your database")
}
func alerts() {
	//row,err  := db.Query("SELECT  max(CustomerName) as 'CustomerName',Site_Name as 'SiteName',sum(monthfgusageval) as 'CurrentUsage'FROM `usage`.selectdata where CustomerName = 'Adline Media' group by Site_Name having sum(monthfgusageval) > 10/100*50")
	row,err  := db.Query("SELECT  max(CustomerName) as 'CustomerName',max(Site_Name) as 'SiteName', mac, sum(monthfgusageval) as 'CurrentUsage' FROM `usage`.selectdata  where CustomerName = 'Adline Media' and  MONTH(alert_50) <> MONTH(CURRENT_DATE()) and MONTH(alert_50) <> MONTH(CURRENT_DATE()) group by mac having  sum(monthfgusageval) > 10/100*50 and sum(monthfgusageval) <= 10/100*80 order by max(site_name)")
	if err != nil{
		fmt.Println("Empty row",err)
	}
	defer row.Close()
	usageRouter := make([]UsageAlert,0)
	for row.Next(){
		usageR := UsageAlert{}
		err := row.Scan(&usageR.custerName,&usageR.siteName,&usageR.mac,&usageR.currentUsage)
		if err != nil{
			fmt.Println("empty result",err)
		}
		usageRouter = append(usageRouter,usageR)

	}
	if err = row.Err();err != nil{
		panic(err)
	}
	var sb strings.Builder


	for _,usageR := range usageRouter{
		currentusagembf:= usageR.currentUsage*1024

		currentusagemb:=fmt.Sprintf("%.0f",currentusagembf)
		usageString  := fmt.Sprintf("%.2f", usageR.currentUsage)
		sb.WriteString(usageR.custerName)
		sb.WriteString("\t")
		sb.WriteString(usageR.siteName)
		sb.WriteString("\t")
		sb.WriteString(usageR.mac)
		sb.WriteString("\t")
		sb.WriteString(usageString)
		sb.WriteString("\n")
		if usageRouter != nil{
			sendEmailWithOutlook50(sb.String(),usageR.mac,usageR.custerName,usageR.siteName,usageString,currentusagemb)
			sb.Reset()
		}
	}
}
func sendEmailWithOutlook50(em,mac,customerName,sitename,usageGB,usageMb string){
	c, err := smtp.Dial("192.168.1.28:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("192.168.1.28"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("sl_team@onewifi.com.au"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "Subject:"+customerName+"- Data Usage Warning: 50%% Reached at "+sitename+"\r\n" +
		"\r\n" +
		"Hi there,\r\n\n 50%% of the monthly data quota of 10GB has been used. There is less than 5GB remaining for the month.\r\n\nCurrent Usage: "+usageMb+"MB | "+usageGB+"GB\r\n\nPlease contact support@onewifi.com.au if you have any concern.\r\n\nThanks\r\nYour OneWIFI Team\r\n\n|OneWiFi Data Usage Monitoring System|")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(em)
	addFlagToDb50(mac)
}



/*
func sendingEmail(em,mac,customerName,sitename string)   {
	//auth := smtp.PlainAuth("", "sunru.au@gmail.com", "mStrack3", "smtp.gmail.com")
	auth := smtp.PlainAuth("", "sudeshn@onewifi.com.au", "Chamila1972", "smtp-mail.outlook.com")
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"sudesh@sunru.com.au"}
	msg := []byte("To:sudesh@sunru.com.au\r\n" +
		"Subject:"+customerName+"- Data Usage Warning: 50% reached at \t"+sitename+"\r\n" +
		"\r\n" +
		"Hi there,\r\n\n 50 % of the monthly data quota of 10GB has been used. There is less than 5GB remaining for the month.\r\n\nPlease contact support@onewifi.com.au if you have any concern.\r\n\n"+em+"\r\n\n\nThanks\r\n\nYour OneWIFI Team")
	err := smtp.SendMail("smtp-mail.outlook.com:25", auth, "sudeshn@onewifi.com.au", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	addFlagToDb50(mac)
	fmt.Println("send mail to Mahesh")
}*/



func alerts80() {
	//row,err  := db.Query("SELECT  max(CustomerName) as 'CustomerName',Site_Name as 'SiteName',sum(monthfgusageval) as 'CurrentUsage'FROM `usage`.selectdata where CustomerName = 'Adline Media' group by Site_Name having sum(monthfgusageval) > 10/100*50")
	row,err  := db.Query("SELECT  max(CustomerName) as 'CustomerName',max(Site_Name) as 'SiteName', mac, sum(monthfgusageval) as 'CurrentUsage' FROM `usage`.selectdata  where CustomerName = 'Adline Media' and  MONTH(alert_80) <> MONTH(CURRENT_DATE()) and MONTH(alert_80) <> MONTH(CURRENT_DATE()) group by mac having  sum(monthfgusageval) > 10/100*80 and sum(monthfgusageval) <= 10/100*100 order by max(site_name)")
	if err != nil{
		fmt.Println("Empty row",err)
	}
	defer row.Close()
	usageRouter := make([]UsageAlert,0)
	for row.Next(){
		usageR := UsageAlert{}
		err := row.Scan(&usageR.custerName,&usageR.siteName,&usageR.mac,&usageR.currentUsage)
		if err != nil{
			fmt.Println("empty result",err)
		}
		usageRouter = append(usageRouter,usageR)

	}
	if err = row.Err();err != nil{
		panic(err)
	}
	var sb strings.Builder
	for _,usageR := range usageRouter{
		currentusagembf:= usageR.currentUsage*1024
		currentusagemb:=fmt.Sprintf("%.0f",currentusagembf)

		usageString  := fmt.Sprintf("%.2f", usageR.currentUsage)
		sb.WriteString(usageR.custerName)
		sb.WriteString("\t")
		sb.WriteString(usageR.siteName)
		sb.WriteString("\t")
		sb.WriteString(usageR.mac)
		sb.WriteString("\t")
		sb.WriteString(usageString)
		sb.WriteString("\n")
		if usageRouter != nil{
			sendEmailWithOutlook80(sb.String(),usageR.mac,usageR.custerName,usageR.siteName,usageString,currentusagemb)

			sb.Reset()
		}
	}
}
func sendEmailWithOutlook80(em,mac,customerName,sitename,usageGB,usageMb string){
	c, err := smtp.Dial("192.168.1.28:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("192.168.1.28"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("sl_team@onewifi.com.au"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "Subject:"+customerName+"- Data Usage Warning: 80%% Reached at "+sitename+"\r\n" +
		"\r\n" +
		"Hi there,\r\n\n 80%% of the monthly data quota of 10GB has been used. There is less than 2GB remaining for the month.\r\n\nCurrent Usage: "+usageMb+"MB | "+usageGB+"GB\r\n\nPlease contact support@onewifi.com.au if you have any concern.\r\n\nThanks\r\nYour OneWIFI Team\r\n\n|OneWiFi Data Usage Monitoring System|")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(em)
	addFlagToDb80(mac)
}





func alerts100() {
	//row,err  := db.Query("SELECT  max(CustomerName) as 'CustomerName',Site_Name as 'SiteName',sum(monthfgusageval) as 'CurrentUsage'FROM `usage`.selectdata where CustomerName = 'Adline Media' group by Site_Name having sum(monthfgusageval) > 10/100*50")
	//row,err  := db.Query("   ")
	//orgin query
	row,err  := db.Query("SELECT  max(CustomerName) as 'CustomerName',max(Site_Name) as 'SiteName', mac, sum(monthfgusageval) as 'CurrentUsage' , max(UsageLimit) as 'UsageLimit' FROM `usage`.selectdata  where MONTH(alert_100) <> MONTH(CURRENT_DATE())and MONTH(alert_100) <> MONTH(CURRENT_DATE()) group by mac having  sum(monthfgusageval) > UsageLimit/100*100 order by max(site_name)")
	//checking testing query
	//row,err := db.Query("SELECT  max(CustomerName) as 'CustomerName',max(Site_Name) as 'SiteName', mac, sum(monthfgusageval) as 'CurrentUsage', max(UsageLimit) as 'UsageLimit' FROM `usage`.selectdata  where MONTH(alert_100) <> MONTH(CURRENT_DATE()) <> MONTH(CURRENT_DATE())group by mac having  sum(monthfgusageval) > UsageLimit/100*100 order by max(site_name)")
	if err != nil{
		fmt.Println("Empty row",err)
	}
	defer row.Close()
	usageRouter := make([]UsageAlert,0)
	for row.Next(){
		usageR := UsageAlert{}
		err := row.Scan(&usageR.custerName,&usageR.siteName,&usageR.mac,&usageR.currentUsage,&usageR.usageLimit)
		if err != nil{
			fmt.Println("empty result",err)
		}
		usageRouter = append(usageRouter,usageR)

	}
	if err = row.Err();err != nil{
		panic(err)
	}
	var sb strings.Builder

	for _,usageR := range usageRouter{
		currentusagembf:= usageR.currentUsage*1024
		currentusagemb:=fmt.Sprintf("%.0f",currentusagembf)

		//convert usage limit int to string
		fmt.Println("Usage Value ****",usageR.usageLimit)
		usageLimitMaxSites := strconv.Itoa(int(usageR.usageLimit))
		fmt.Println("usageLimitMaxSites ***",usageLimitMaxSites)

		usageString  := fmt.Sprintf("%.2f", usageR.currentUsage)
		sb.WriteString(usageR.custerName)
		sb.WriteString("\t")
		sb.WriteString(usageR.siteName)
		sb.WriteString("\t")
		sb.WriteString(usageR.mac)
		sb.WriteString("\t")
		sb.WriteString(usageString)

		sb.WriteString("\n")
		if usageRouter != nil{
			sendEmailWithOutlook100(sb.String(),usageR.mac,usageR.custerName,usageR.siteName,usageString,currentusagemb,usageLimitMaxSites)
			fmt.Println("String Buffer Result 100% usage",sb.String())
			sb.Reset()
		}
	}
}
func sendEmailWithOutlook100(em,mac,customerName,sitename,usageGB,usageMb,usageLimitMaxSites string){
	fmt.Println("maximum Usage limit ***%",usageLimitMaxSites)
	c, err := smtp.Dial("192.168.1.28:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("192.168.1.28"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("sl_team@onewifi.com.au"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "Subject:"+customerName+"- Data Usage Warning: 100%% Reached at "+sitename+"\r\n" +
		"\r\n" +
		"Hi there,\r\n\n 100%% of the monthly data quota of "+usageLimitMaxSites+" has been used.\r\n\nCurrent Usage: "+usageMb+"MB | "+usageGB+"GB\r\n\nPlease contact support@onewifi.com.au if you have any concern.\r\n\nThanks\r\nYour OneWIFI Team\r\n\n|OneWiFi Data Usage Monitoring System|")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(em)
	addFlagToDb100(mac)
	fmt.Println("send mail to sl-team 100%"+customerName+"usage"+usageGB+".")
}


/*
//send mail with out smtp authenticate
func mailWithOutAut100(em,mac,customerName,sitename string)  {
	formEmail :="support@onewifi.com.au"
	toEmail :="sudeshn@onewifi.com.au"

	msg := "Subject:"+customerName+"- Data Usage Warning:100% reached at \t"+sitename+"\r\n" +
		"\r\n" +
		"Hi there,\r\n\n 100% of the monthly data quota of 10GB has been used. There is less than 0GB remaining for the month.\r\n\nPlease contact support@onewifi.com.au if you have any concern.\r\n\nThanks\r\nYour OneWIFI Team"

	sendmail := exec.Command("/usr/sbin/sendmail", "-f", formEmail, toEmail)
	stdin, err := sendmail.StdinPipe()
	if err != nil {
		panic(err)
	}

	stdout, err := sendmail.StdoutPipe()
	if err != nil {
		panic(err)
	}

	sendmail.Start()
	stdin.Write([]byte(msg))
	stdin.Close()
	sentBytes, _ := ioutil.ReadAll(stdout)
	sendmail.Wait()

	fmt.Println("Send Command Output\n")
	fmt.Println(string(sentBytes))
	fmt.Println(em)
	addFlagToDb100(mac)

}*/


func addFlagToDb50(macAd string){
	dt := time.Now()
	var fdt = dt.Format("2006-01-02 15:04:05")
	//fmt.Println(fdt)
	_,err :=db.Exec("update forgtb_livedata set alert_50 = ? where mac = ?",fdt,macAd)
	if err!=nil{
		fmt.Println("no Records insert")
	}


}

func addFlagToDb80(macAd string){
	dt := time.Now()
	var fdt = dt.Format("2006-01-02 15:04:05")
	//fmt.Println(fdt)
	_,err :=db.Exec("update forgtb_livedata set alert_80 = ? where mac = ?",fdt,macAd)
	if err!=nil{
		fmt.Println("no Records insert")
	}


}
func addFlagToDb100(macAd string){
	dt := time.Now()
	var fdt = dt.Format("2006-01-02 15:04:05")
	//fmt.Println(fdt)
	_,err :=db.Exec("update forgtb_livedata set alert_100 = ? where mac = ?",fdt,macAd)
	if err!=nil{
		fmt.Println("no Records insert")
	}


}

func main()  {
	alerts()
	alerts80()
	alerts100()


}
