package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Din-27/Go_job/internal/models"
)

func FetchGetProvinsi() (value []models.Provinsi, err error) {
	url := fmt.Sprintf("%sprovinces.json", os.Getenv("API_STATE")) // Replace with the URL you want to fetch data from

	// You can prepare your request data if needed
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var dataArray []models.Provinsi

	// Ubah respons JSON menjadi array menggunakan json.Unmarshal
	if err := json.NewDecoder(response.Body).Decode(&dataArray); err != nil {
		return nil, err
	}

	return dataArray, nil
}

func FetchGetKabupaten(id_provinsi string) (data []models.Kabupaten, err error) {
	url := fmt.Sprintf("%sregencies/%s.json", os.Getenv("API_STATE"), id_provinsi) // Replace with the URL you want to fetch data from

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var dataArray []models.Kabupaten

	// Ubah respons JSON menjadi array menggunakan json.Unmarshal
	if err := json.NewDecoder(response.Body).Decode(&dataArray); err != nil {
		return nil, err
	}

	return dataArray, nil
}

func FetchGetKecamatan(id_kabupaten string) (data []models.Kecamatan, err error) {
	url := fmt.Sprintf("%sdistricts/%s.json", os.Getenv("API_STATE"), id_kabupaten) // Replace with the URL you want to fetch data from

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var dataArray []models.Kecamatan

	// Ubah respons JSON menjadi array menggunakan json.Unmarshal
	if err := json.NewDecoder(response.Body).Decode(&dataArray); err != nil {
		return nil, err
	}

	return dataArray, nil
}

func FetchGetKelurahan(id_kecamatan string) (value []models.Kelurahan, err error) {
	url := fmt.Sprintf("%svillages/%s.json", os.Getenv("API_STATE"), id_kecamatan) // Replace with the URL you want to fetch data from

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var dataArray []models.Kelurahan

	// Ubah respons JSON menjadi array menggunakan json.Unmarshal
	if err := json.NewDecoder(response.Body).Decode(&dataArray); err != nil {
		return nil, err
	}

	return dataArray, nil
}
