package models

import (
	"splitbill_api/db"
	"github.com/go-playground/validator"
	"net/http"
)

type DebtorBill struct {
	Id             int    `json:"id"`
	Total          string `json:"total"`
	Tax            string `json:"tax"`
	Service_Charge string `json:"service_charge"`
	Discount       string `json:"discount"`
	Other          string `json:"other"`
	IsConfirmed    string `json:"isConfirmed"`
	Bill_Id        string `json:"bill_id"`
}

func CreateDebtorBill(total string, tax string, service_charge string, discount string, other string, isConfirmed string, bill_id string) (Response, error) {
	var res Response

	v := validator.New()

	debtorBill := DebtorBill{
		Total:          total,
		Tax:            tax,
		Service_Charge: service_charge,
		Discount:       discount,
		Other:          other,
		IsConfirmed:    isConfirmed,
		Bill_Id:        bill_id,
	}

	err := v.Struct(debtorBill)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO debtorbill(total, tax, service_charge, discount, other, isConfirmed , bill_id) VALUES (?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(total, tax, service_charge, discount, other, isConfirmed , bill_id)

	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedID,
	}
	return res, nil

}

func DeleteDebtorBill(id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "DELETE FROM debtorbill WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil

}

func EditDebtorBill(id string, total string, tax string, service_charge string, discount string, other string, isConfirmed string, bill_id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE bill SET total=?, tax=?, service_charge=?, discount=?, other=?, isConfirmed=?, bill_id=? WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(total, tax, service_charge, discount, other, isConfirmed, bill_id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil
}

func ReadDebtorBill(id string)(Response, error){
	var obj DebtorBill
	var arrObj []DebtorBill
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM debtorbill where bill_id="+ id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Total, &obj.Tax, &obj.Service_Charge, &obj.Discount, &obj.Other,&obj. IsConfirmed, &obj.Bill_Id)

		if err != nil{
			return res,err
		}
		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message="Success"
	res.Data = arrObj

	return res, nil
}