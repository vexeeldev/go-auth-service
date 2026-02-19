package auth

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Contoh struct {
	ID int `json:"id"`
	Pekerjaan string `json:"pekerjaan"`
	Umur int `json:umur`
	Alamat string `json:"alamat"`
}
