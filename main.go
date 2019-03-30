package main

import (
	"fmt"
	"gen_file_txt/model"
	"github.com/gobuffalo/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/vjeantet/jodaTime"
	"os"
)

var fileNameExport string
var path string

func main() {
	db, err := gorm.Open("mssql", "sqlserver://sa:1234@localhost:1433?database=TPA_FWD")
	if err != nil {
		panic(err.Error())
	}

	var Policy []model.Policy

	db.Debug().Find(&Policy)

	uid, err := uuid.NewV4()
	fileNameExport = uid.String()

	path = "C:\\Users\\NENG-DEV\\go\\src\\gen_file_txt\\"+fileNameExport+".txt"

	createFile()
	writeFile(Policy)

	defer db.Close()
}

func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) { return }
		defer file.Close()
	}
	fmt.Println("==> done creating file", path)
}

func writeFile(Policy []model.Policy) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()


	_, err = file.WriteString("policy_id|")
	_, err = file.WriteString("status|")
	_, err = file.WriteString("date_start|")
	_, err = file.WriteString("date_end|")
	_, err = file.WriteString("contract_no|")
	_, err = file.WriteString("contract_no_individual|")
	_, err = file.WriteString("policy_no|")
	_, err = file.WriteString("policy_no_individual|")
	_, err = file.WriteString("beneficiary_id|")
	_, err = file.WriteString("beneficiary_address_id|")
	_, err = file.WriteString("emergency_contact_id|")
	_, err = file.WriteString("emergency_contact_address_id|")
	_, err = file.WriteString("insured_id|")
	_, err = file.WriteString("insured_address_id|")
	_, err = file.WriteString("owner_id|")
	_, err = file.WriteString("owner_address_id|")
	_, err = file.WriteString("plan_id|")
	_, err = file.WriteString("premium_gross|")
	_, err = file.WriteString("premium_net|")
	_, err = file.WriteString("log_create|")
	_, err = file.WriteString("log_create_by|")
	_, err = file.WriteString("log_update|")
	_, err = file.WriteString("log_update_by|")
	_, err = file.WriteString("member_no|")
	_, err = file.WriteString("member_exclusion|")
	_, err = file.WriteString("external_plan_code|")
	_, err = file.WriteString("mailing_address|")
	_, err = file.WriteString("flag_card_created|")
	_, err = file.WriteString("emp_no|")
	_, err = file.WriteString("emp_date|")
	_, err = file.WriteString("old_member_no|")
	_, err = file.WriteString("med_1|")
	_, err = file.WriteString("med_2|")
	_, err = file.WriteString("policy_year|")
	_, err = file.WriteString("certificate_no|")
	_, err = file.WriteString("transaction_type|")
	_, err = file.WriteString("sequence_no|")
	_, err = file.WriteString("risk_no|")
	_, err = file.WriteString("customer_risk_no|")
	_, err = file.WriteString("remark|")
	_, err = file.WriteString("import_date|")
	_, err = file.WriteString("deductible|")
	_, err = file.WriteString("deductible_ipd|")
	_, err = file.WriteString("deductible_opd|")
	_, err = file.WriteString("partner_contract_info_id|")
	_, err = file.WriteString("policy_type|")
	_, err = file.WriteString("subsidiary_id\r\n")
	if isError(err) { return }


	for _, v := range Policy{
		_, err = file.WriteString(v.Policy_id.String+"|")
		_, err = file.WriteString(v.Status.String+"|")

		if jodaTime.Format("YYYY-MM-dd H:m:s", v.Date_start) != "1-01-01 0:0:0" {
			_, err = file.WriteString(jodaTime.Format("YYYY-MM-dd", v.Date_start)+"|")
		}else {
			_, err = file.WriteString("|")
		}

		if jodaTime.Format("YYYY-MM-dd H:m:s", v.Date_end) != "1-01-01 0:0:0" {
			_, err = file.WriteString(jodaTime.Format("YYYY-MM-dd", v.Date_end)+"|")
		}else {
			_, err = file.WriteString("|")
		}
		_, err = file.WriteString(v.Contract_no.String+"|")
		_, err = file.WriteString(v.Contract_no_individual.String+"|")
		_, err = file.WriteString(v.Policy_no.String+"|")
		_, err = file.WriteString(v.Policy_no_individual.String+"|")
		_, err = file.WriteString(v.Beneficiary_id.String+"|")
		_, err = file.WriteString(v.Beneficiary_address_id.String+"|")
		_, err = file.WriteString(v.Emergency_contact_id.String+"|")
		_, err = file.WriteString(v.Emergency_contact_address_id.String+"|")
		_, err = file.WriteString(v.Insured_id.String+"|")
		_, err = file.WriteString(v.Insured_address_id.String+"|")
		_, err = file.WriteString(v.Owner_id.String+"|")
		_, err = file.WriteString(v.Owner_address_id.String+"|")
		_, err = file.WriteString(v.Plan_id.String+"|")
		_, err = file.WriteString(v.Premium_gross.String+"|")
		_, err = file.WriteString(v.Premium_net.String+"|")

		if jodaTime.Format("YYYY-MM-dd H:m:s", v.Log_create) != "1-01-01 0:0:0" {
			_, err = file.WriteString(jodaTime.Format("YYYY-MM-dd H:m:s", v.Log_create)+"|")
		}else {
			_, err = file.WriteString("|")
		}

		_, err = file.WriteString(v.Log_create_by.String+"|")

		if jodaTime.Format("YYYY-MM-dd H:m:s", v.Log_update) != "1-01-01 0:0:0" {
			_, err = file.WriteString(jodaTime.Format("YYYY-MM-dd H:m:s", v.Log_update)+"|")
		}else {
			_, err = file.WriteString("|")
		}

		_, err = file.WriteString(v.Log_update_by.String+"|")
		_, err = file.WriteString(v.Member_no.String+"|")
		_, err = file.WriteString(v.Member_exclusion.String+"|")
		_, err = file.WriteString(v.External_plan_code.String+"|")
		_, err = file.WriteString(v.Mailing_address.String+"|")
		_, err = file.WriteString(v.Flag_card_created.String+"|")
		_, err = file.WriteString(v.Emp_no.String+"|")

		if jodaTime.Format("YYYY-MM-dd H:m:s", v.Emp_date) != "1-01-01 0:0:0" {
			_, err = file.WriteString(jodaTime.Format("YYYY-MM-dd H:m:s", v.Emp_date)+"|")
		}else {
			_, err = file.WriteString("|")
		}

		_, err = file.WriteString(v.Old_member_no.String+"|")
		_, err = file.WriteString(v.Med_1.String+"|")
		_, err = file.WriteString(v.Med_2.String+"|")
		_, err = file.WriteString(v.Policy_year.String+"|")
		_, err = file.WriteString(v.Certificate_no.String+"|")
		_, err = file.WriteString(v.Transaction_type.String+"|")
		_, err = file.WriteString(v.Sequence_no.String+"|")
		_, err = file.WriteString(v.Risk_no.String+"|")
		_, err = file.WriteString(v.Customer_risk_no.String+"|")
		_, err = file.WriteString(v.Remark.String+"|")

		if jodaTime.Format("YYYY-MM-dd H:m:s", v.Import_date) != "1-01-01 0:0:0" {
			_, err = file.WriteString(jodaTime.Format("YYYY-MM-dd H:m:s", v.Import_date)+"|")
		}else {
			_, err = file.WriteString("|")
		}
		_, err = file.WriteString(v.Deductible.String+"|")
		_, err = file.WriteString(v.Deductible_ipd.String+"|")
		_, err = file.WriteString(v.Deductible_opd.String+"|")
		_, err = file.WriteString(v.Partner_contract_info_id.String+"|")
		_, err = file.WriteString(v.Policy_type.String+"|")
		_, err = file.WriteString(v.Subsidiary_id.String+"\r\n")
		if isError(err) { return }
	}

	err = file.Sync()
	if isError(err) { return }

	fmt.Println("==> done writing to file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}