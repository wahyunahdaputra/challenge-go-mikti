package main

import "fmt"

func main() {
	// memanggil fungsi untuk Soal 1
	fmt.Println("Soal No.1")
	Soal1() // Memanggil fungsi Soal1 untuk menyelesaikan soal pertama

	fmt.Println()

	// memanggil fungsi untuk Soal 2
	fmt.Println("Soal No.2")
	Soal2() // Memanggil fungsi Soal2 untuk menyelesaikan soal kedua
}

// fungsi untuk soal 1
func Soal1() {
	// menambahkan tim
	LumbaLumba := Team{
		Name: "Lumba-Lumba",
	}
	Koala := Team{
		Name: "Koala",
	}

	// memasukkan skor untuk masing-masing tim
	fmt.Println("Masukkan Skor untuk Tim Lumba-Lumba")
	for i := 1; i <= 3; i++ {
		var score int
		fmt.Printf("Skor ke %d untuk Tim Lumba-Lumba: ", i)
		_, err := fmt.Scan(&score)
		if err != nil {
			return
		}
		LumbaLumba.Scores = append(LumbaLumba.Scores, Score(score))
	}
	fmt.Println()
	fmt.Println("Masukkan Skor untuk Tim Koala")
	for i := 1; i <= 3; i++ {
		var score int
		fmt.Printf("Masukkan Skor ke %d untuk Tim Koala: ", i)
		_, err := fmt.Scan(&score)
		if err != nil {
			return
		}
		Koala.Scores = append(Koala.Scores, Score(score))
	}
	fmt.Println("=======================================")
	// menghitung rata-rata skor untuk masing-masing tim
	LumbaLumba.Avg = LumbaLumba.CalculateScoreAverage() // Menghitung rata-rata skor untuk tim LumbaLumba
	Koala.Avg = Koala.CalculateScoreAverage()           // Menghitung rata-rata skor untuk tim Koala

	// memilih pemenang
	PickAWinner(LumbaLumba, Koala) // Memilih pemenang antara tim LumbaLumba dan Koala
}

// fungsi untuk soal 2
func Soal2() {
	// menambahkan data untuk Mark dan John
	Mark := Person{
		Name: "Mark",
	}
	John := Person{
		Name: "John",
	}

	// membuat variabel untuk menampung informasi apakah Mark memiliki BMI lebih tinggi dari John
	var markHigherBMI bool

	// memasukkan berat badan dan tinggi badan untuk Mark
	fmt.Print("Masukkan Berat Badan Mark: ")
	_, err := fmt.Scan(&Mark.Mass)
	if err != nil {
		return
	}
	fmt.Print("Masukkan Tinggi Badan Mark: ")
	_, err = fmt.Scan(&Mark.Height)
	if err != nil {
		return
	}

	// menghitung BMI untuk Mark
	Mark.BMI = Mark.CalculateBMIScore() // Menghitung BMI untuk Mark
	fmt.Println()
	// memasukkan berat badan dan tinggi badan untuk John
	fmt.Print("Masukkan Berat Badan John: ")
	_, err = fmt.Scan(&John.Mass)
	if err != nil {
		return
	}
	fmt.Print("Masukkan Tinggi Badan John: ")
	_, err = fmt.Scan(&John.Height)
	if err != nil {
		return
	}
	fmt.Println("=======================================")
	// menghitung BMI untuk John
	John.BMI = John.CalculateBMIScore() // Menghitung BMI untuk John

	// membandingkan BMI Mark dan John
	if Mark.BMI > John.BMI {
		markHigherBMI = true
	} else {
		markHigherBMI = false
	}

	// menampilkan hasil perbandingan BMI Mark dan John
	if markHigherBMI {
		fmt.Println("Hasilnya, BMI Mark lebih tinggi dari BMI John.") // Menampilkan pesan jika BMI Mark lebih tinggi dari BMI John
	} else {
		fmt.Println("Hasilnya, BMI John lebih tinggi dari BMI Mark.") // Menampilkan pesan jika BMI John lebih tinggi dari BMI Mark
	}
}
