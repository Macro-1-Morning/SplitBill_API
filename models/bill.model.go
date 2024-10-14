package models

import (
	"splitbill_api/db"
	"github.com/go-playground/validator"
	"net/http"
)

type Bill struct {
	Id              int    `json:"id"`
	Receipt_Image   string `json:"receipt_image"`
	Restaurant_Name string `json:"restaurant_name"`
	Subtotal        string `json:"subtotal"`
	Total_Discount  string `json:"total_discount"`
	Service_Charge  string `json:"service_charge"`
	Tax             string `json:"tax"`
	Other           string `json:"other"`
	Grand_Total     string `json:"grand_total"`
	IsSettled       string `json:"isSettled"`
	Payment_Id      string `json:"payment_id"`
}

func CreateBill(receipt_image string, restaurant_name string, subtotal string, total_discount string, service_charge string, tax string, other string, grand_total string, isSettled string, payment_id string) (Response, error) {
	var res Response

	v := validator.New()

	bill := Bill{
		Receipt_Image:   receipt_image,
		Restaurant_Name: restaurant_name,
		Subtotal:        subtotal,
		Total_Discount:  total_discount,
		Service_Charge:  service_charge,
		Tax:             tax,
		Other:           other,
		Grand_Total:     grand_total,
		IsSettled:       isSettled,
		Payment_Id:      payment_id,
	}

	err := v.Struct(bill)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO bill(receipt_image, restaurant_name, subtotal, total_discount, service_charge, tax, other, grand_total,isSettled,payment_id) VALUES (?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(receipt_image, restaurant_name, subtotal, total_discount, service_charge, tax, other, grand_total, isSettled, payment_id)

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

func DeleteBill(id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "DELETE FROM bill WHERE id=?"
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

func EditBill(id string, receipt_image string, restaurant_name string, subtotal string, total_discount string, service_charge string, tax string, other string, grand_total string, isSettled string, payment_id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE bill SET receipt_image=?, restaurant_name=?, subtotal=?, total_discount=?, service_charge=?, tax=?, other=?, grand_total=?, isSettled=?, payment_id=? WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(receipt_image, restaurant_name, subtotal, total_discount, service_charge, tax, other, grand_total, isSettled, payment_id, id)

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

func ReadAllBill(id string)(Response, error){
	var obj Bill
	var arrObj []Bill
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM bill where user_id="+ id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Receipt_Image, &obj.Restaurant_Name, &obj.Subtotal, &obj.Total_Discount, &obj.Service_Charge, &obj.Tax, &obj.Other, &obj.Grand_Total, &obj.IsSettled, &obj.Payment_Id)

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

func ReadBillByBillId(id string)(Response, error){
	var obj Bill
	var arrObj []Bill
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from bill where id= "+id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Receipt_Image, &obj.Restaurant_Name, &obj.Subtotal, &obj.Total_Discount, &obj.Service_Charge, &obj.Tax, &obj.Other, &obj.Grand_Total, &obj.IsSettled, &obj.Payment_Id)

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
