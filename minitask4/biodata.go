package minitask4

import "fmt"

type Pendidikan struct {
	Nama    string
	Jurusan string
}

type DataDiri struct {
	Nama              string
	Foto              string
	Email             string
	Umur              int
	NomorTelepon      string
	StatusPernikahan  string
	RiwayatPendidikan []Pendidikan
}

func TampilkanBiodata() {
	data := DataDiri{
		Nama:             "Bian Albaihaqi",
		Foto:             "bian.jpg",
		Email:            "bianalbaihaqi190@gmail.com",
		Umur:             20,
		NomorTelepon:     "087727726316",
		StatusPernikahan: "Belum",
		RiwayatPendidikan: []Pendidikan{
			{"SDN Tikukur 03", "-"},
			{"SMPN 1 Leles", "-"},
			{"SMKS Alfarisi", "Teknik Komputer Jaringan"},
		},
	}

	fmt.Println("=== BIODATA ===")
	fmt.Println("Nama:", data.Nama)
	fmt.Println("Email:", data.Foto)
	fmt.Println("Umur:", data.Email)
	fmt.Println("Umur:", data.Umur)
	fmt.Println("Umur:", data.NomorTelepon)
	fmt.Println("Umur:", data.StatusPernikahan)

	fmt.Println("Riwayat Pendidikan:")
	for _, p := range data.RiwayatPendidikan {
		fmt.Printf("- %s (%s)\n", p.Nama, p.Jurusan)
	}
}
