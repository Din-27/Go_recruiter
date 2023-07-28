package models

type Provinsi struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Kabupaten struct {
	Id         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}

type Kecamatan struct {
	Id        string `json:"id"`
	RegencyId string `json:"regency_id"`
	Name      string `json:"name"`
}

type Kelurahan struct {
	Id         string `json:"id"`
	DistrictId string `json:"district_id"`
	Name       string `json:"name"`
}
