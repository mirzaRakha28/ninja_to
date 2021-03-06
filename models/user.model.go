package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mirzaRakha28/ninja_to/db"
	"github.com/mirzaRakha28/ninja_to/helpers"
)

type User struct {
	Id            int    `json:"id"`
	Email         string `json:"email" validate:"required"`
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	Jenjang       string `json:"jenjang" validate:"required"`
	Score         int    `json:"score"`
	Jml_benar     int    `json:"jml_benar"`
	Jml_pelajaran int    `json:"jml_pelajaran"`
	Jml_ratarata  int    `json:"jml_ratarata"`
	Jml_salah     int    `json:"jml_salah"`
	Jml_paket     int    `json:"jml_paket"`
}

func Login(username string, password string) (Response, error) {
	var res Response
	con := db.CreateCon()
	var obj User
	var arrobj []User

	sqlStatement := "SELECT * FROM user WHERE username = ?"

	stmt, err := con.Query(sqlStatement, username)
	if err != nil {
		return res, err
	}

	for stmt.Next() {
		err = stmt.Scan(&obj.Id, &obj.Email, &obj.Username, &obj.Password, &obj.Jenjang,
			&obj.Score, &obj.Jml_benar, &obj.Jml_pelajaran, &obj.Jml_ratarata,
			&obj.Jml_salah, &obj.Jml_paket)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	if obj.Id == 0 {
		res.Status = http.StatusOK
		res.Message = "Username tidak Cocok"
		res.Data = arrobj
		return res, nil
	}
	if obj.Password == password {
		res.Status = http.StatusOK
		res.Message = "Login Berhasil"
		res.Data = arrobj
	} else {
		res.Status = http.StatusOK
		res.Message = "Password  tidak cocok"
		res.Data = nil

	}
	return res, nil
}
func Register(email string, username string, password string, jenjang string) (Response, error) {
	var res Response

	v := validator.New()

	hash, _ := helpers.HashPassword(password)
	peg := User{
		Email:    email,
		Username: username,
		Password: hash,
		Jenjang:  jenjang,
	}
	err := v.Struct(peg)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT user (email, username, password, jenjang) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(email, username, hash, jenjang)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	var obj User
	var arrobj []User

	sqlStatementData := "SELECT * FROM user WHERE id = ?"
	rows, err := con.Query(sqlStatementData, lastInsertedId)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Email, &obj.Username, &obj.Password, &obj.Jenjang,
			&obj.Score, &obj.Jml_benar, &obj.Jml_pelajaran, &obj.Jml_ratarata,
			&obj.Jml_salah, &obj.Jml_paket)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Registrasi Berhasil"
	res.Data = arrobj

	return res, nil
}
