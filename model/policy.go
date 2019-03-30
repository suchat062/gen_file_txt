package model

import (
	"database/sql"
	"time"
)

type Policy struct {
	Policy_id sql.NullString
	Status sql.NullString
	Date_start time.Time
	Date_end time.Time
	Contract_no sql.NullString
	Contract_no_individual sql.NullString
	Policy_no sql.NullString
	Policy_no_individual sql.NullString
	Beneficiary_id sql.NullString
	Beneficiary_address_id sql.NullString
	Emergency_contact_id sql.NullString
	Emergency_contact_address_id sql.NullString
	Insured_id sql.NullString
	Insured_address_id sql.NullString
	Owner_id sql.NullString
	Owner_address_id sql.NullString
	Plan_id sql.NullString
	Premium_gross sql.NullString
	Premium_net sql.NullString
	Log_create time.Time
	Log_create_by sql.NullString
	Log_update time.Time
	Log_update_by sql.NullString
	Member_no sql.NullString
	Member_exclusion sql.NullString
	External_plan_code sql.NullString
	Mailing_address sql.NullString
	Flag_card_created sql.NullString
	Emp_no sql.NullString
	Emp_date time.Time
	Old_member_no sql.NullString
	Med_1 sql.NullString
	Med_2 sql.NullString
	Policy_year sql.NullString
	Certificate_no sql.NullString
	Transaction_type sql.NullString
	Sequence_no sql.NullString
	Risk_no sql.NullString
	Customer_risk_no sql.NullString
	Remark sql.NullString
	Import_date time.Time
	Deductible sql.NullString
	Deductible_ipd sql.NullString
	Deductible_opd sql.NullString
	Partner_contract_info_id sql.NullString
	Policy_type sql.NullString
	Subsidiary_id sql.NullString
}

func (Policy) TableName() string {
	return "policies"
}
