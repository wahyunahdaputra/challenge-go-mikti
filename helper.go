package main

import (
	"fmt"
	"math"
)

// fungsi helper untuk membulatkan dua angka di belakang koma
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision)) // Menghitung pembulatan ke dua angka di belakang koma
	return math.Round(val*ratio) / ratio      // Melakukan pembulatan sesuai presisi yang diberikan
}

// fungsi untuk memilih pemenang
func PickAWinner(team1 Team, team2 Team) {
	// membuat variabel untuk menampung data yang diperlukan
	var winner Team                   // variabel untuk menampung tim pemenang
	var isWinnerScoreMoreThan100 bool // variabel untuk menampung informasi apakah tim pemenang memiliki skor lebih dari 100
	var isDraw bool                   // variabel untuk menampung informasi apakah hasilnya seri

	// membandingkan rata-rata skor dari masing-masing tim
	if team1.Avg > team2.Avg {
		winner = team1 // Menentukan pemenang jika tim 1 memiliki rata-rata skor lebih tinggi
	} else if team1.Avg < team2.Avg {
		winner = team2 // Menentukan pemenang jika tim 2 memiliki rata-rata skor lebih tinggi
	} else {
		isDraw = true // Menandakan hasil seri jika kedua tim memiliki rata-rata skor yang sama
	}

	// jika hasilnya seri
	if isDraw {
		// membuat variabel untuk menampung informasi apakah kedua tim memiliki skor lebih dari 100
		var isTeam1ScoreMoreThan100 bool
		var isTeam2ScoreMoreThan100 bool

		// memeriksa apakah kedua tim memiliki skor lebih dari 100
		for _, score := range team1.Scores {
			if score > 100 {
				isTeam1ScoreMoreThan100 = true
				break
			}
		}
		for _, score := range team2.Scores {
			if score > 100 {
				isTeam2ScoreMoreThan100 = true
				break
			}
		}

		// jika kedua tim memiliki skor lebih dari 100
		if isTeam1ScoreMoreThan100 && isTeam2ScoreMoreThan100 {
			fmt.Println("Hasilnya seri, kedua tim memiliki skor lebih dari 100.") // Menampilkan pesan jika kedua tim memiliki skor lebih dari 100
			fmt.Println("Kedua tim mendapat trofi.")                              // Menampilkan pesan bahwa kedua tim mendapat trofi
		} else {
			fmt.Println("Hasilnya seri, kedua tim memiliki skor kurang dari 100.") // Menampilkan pesan jika kedua tim memiliki skor kurang dari 100
			fmt.Println("Tidak ada tim yang mendapat trofi.")                      // Menampilkan pesan bahwa tidak ada tim yang mendapat trofi karena hasil seri
		}
	} else { // jika hasilnya tidak seri, atau ada pemenang di antara kedua tim
		fmt.Printf("Tim %s adalah pemenangnya karena memiliki skor yang lebih tinggi dari tim lawan.\n", winner.Name) // Menampilkan pesan bahwa tim pemenang karena memiliki rata-rata skor lebih tinggi dari lawan

		// memeriksa apakah pemenang memiliki skor lebih dari 100
		for _, score := range winner.Scores {
			if score > 100 {
				isWinnerScoreMoreThan100 = true
				break
			}
		}

		// jika pemenang memiliki skor lebih dari 100
		if isWinnerScoreMoreThan100 {
			fmt.Println("dan memiliki skore lebih dari 100") // Menampilkan pesan jika tim pemenang memiliki skor lebih dari 100
		}
	}
}

// fungsi untuk menghitung rata-rata skor
func (team Team) CalculateScoreAverage() float32 {
	// membuat variabel untuk menampung total skor dan rata-rata skor
	var averageScore float32
	var totalScores float32

	// menghitung total skor
	for _, score := range team.Scores {
		totalScores += float32(score)
	}

	// menghitung rata-rata skor
	averageScore = totalScores / float32(len(team.Scores))

	// membulatkan rata-rata skor menjadi dua angka di belakang koma
	parsedFloat := roundFloat(float64(averageScore), 2) // Memanggil fungsi pembulatan untuk rata-rata skor
	return float32(parsedFloat)                         // Mengembalikan rata-rata skor yang sudah dibulatkan
}

// fungsi untuk menghitung BMI
func (person Person) CalculateBMIScore() float32 {
	bmi := person.Mass / (person.Height * person.Height) // Menghitung BMI menggunakan rumus berat badan dibagi tinggi badan kuadrat

	parsedFloat := roundFloat(float64(bmi), 2) // Memanggil fungsi pembulatan untuk BMI
	return float32(parsedFloat)                // Mengembalikan nilai BMI yang sudah dibulatkan
}
