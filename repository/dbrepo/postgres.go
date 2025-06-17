package dbrepo

import (
	"jamiuafolabi/web-service-gin/models"
	"log"
)

func (m *postgresDBRepo) GetAllRows() ([]models.ProductInformation, error) {
	rows, err := m.DB.Query("select id, name from productinformation")

	if err != nil {
		log.Println("error in connecting to the productinformation table")
		return nil, err
	} else {
		log.Println("connecting to the productinformation table")
	}
	defer rows.Close()

	var products []models.ProductInformation

	for rows.Next() {
		var p models.ProductInformation
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			log.Println("error scanning row:", err)
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		log.Println("rows error:", err)
		return nil, err
	}

	return products, nil
}
