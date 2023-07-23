package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Din-27/Go_job/helpers/models"
)

func FetchGetProvinsi() (value []models.Provinsi, err error) {
	url := "https://din-27.github.io/api-wilayah-indonesia/api/provinces.json" // Replace with the URL you want to fetch data from

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
	url := fmt.Sprintf("https://din-27.github.io/api-wilayah-indonesia/api/regencies/%s.json", id_provinsi) // Replace with the URL you want to fetch data from

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
	url := fmt.Sprintf("https://din-27.github.io/api-wilayah-indonesia/api/districts/%s.json", id_kabupaten) // Replace with the URL you want to fetch data from

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
	url := fmt.Sprintf("https://din-27.github.io/api-wilayah-indonesia/api/villages/%s.json", id_kecamatan) // Replace with the URL you want to fetch data from

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
