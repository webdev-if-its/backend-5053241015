package main

import (
	"errors"
	"fmt"
)

// TODO(Level 1): lihat SOAL.md untuk kontrak lengkap tiap fungsi di bawah.
// Ganti setiap "panic" dengan implementasi yang benar.

func HitungSubtotal(qty int, hargaSatuan float64) float64 {
	return float64(qty * int(hargaSatuan))
}

func HitungTotalPesanan(qty []int, hargaSatuan []float64) float64 {
	if len(qty) != len(hargaSatuan) {
		return 0
	}
	var total float64
	for i := 0; i < len(qty); i++ {
		total += float64(qty[i] * int(hargaSatuan[i]))
	}

	return total
}

func TerapkanPajak(total float64, tarifPajak float64) float64 {
	return total + total*tarifPajak
}

func HitungDiskon(total float64) float64 {
	if total >= 1000000 {
		return total * 0.10
	} else if total >= 500000 {
		return total * 0.05
	} else {
		return 0
	}
}

func TotalSetelahDiskon(qty []int, hargaSatuan []float64, tarifPajak float64) float64 {
	total := HitungTotalPesanan(qty, hargaSatuan)
	disc := HitungDiskon(total)
	totalDisc := total - disc

	return TerapkanPajak(totalDisc, tarifPajak)
}

func ValidasiPesanan(qty []int, hargaSatuan []float64) (bool, string) {
	if len(qty) != len(hargaSatuan) {
		return false, "jumlah qty dan harga satuan tidak sesuai"
	}

	for i := 1; i < len(qty); i++ {
		if qty[i] < 0 {
			return false, "jumlah qty ada yang negatif"
		}
	}

	for i := 1; i < len(hargaSatuan); i++ {
		if hargaSatuan[i] < 0 {
			return false, "jumlah qty ada yang negatif"
		}
	}
	return true, ""
}

func TentukanStatus(total float64) string {

	if total > 1000000 {
		return "Prioritas"
	} else if total > 100000 {
		return "Reguler"
	} else {
		return "Hemat"
	}
}

func RingkasanPesanan(qty []int, hargaSatuan []float64, tarifPajak float64) string {
	var subtotal = HitungTotalPesanan(qty, hargaSatuan)
	var diskon = HitungDiskon(subtotal)
	var totalAkhir = TotalSetelahDiskon(qty, hargaSatuan, tarifPajak)
	var status = TentukanStatus(totalAkhir)

	return fmt.Sprintf("%f", subtotal) + fmt.Sprintf("%f", diskon) + fmt.Sprintf("%f", totalAkhir) + status
}

// TODO(Level 9): signature ini SUDAH benar (cari tahu sendiri kenapa
// bentuknya begini - lihat SOAL.md) - tinggal implementasikan isinya.
func Total(harga ...float64) float64 {
	var sum float64
	for _, h := range harga {
		sum += h
	}
	return sum
}

// TODO(Level 10, bonus): signature ini SUDAH benar (cari tahu sendiri
// kenapa ada dua nilai balik - lihat SOAL.md) - tinggal implementasikan isinya.
func HitungOngkosKirim(beratKg float64, jarakKm float64) (float64, error) {

	if beratKg <= 0 || jarakKm <= 0 {
		return 0, errors.New("berakt dan jarak tidak boleh sama atau kurang dari 0 :(")
	}

	return beratKg*2000 + jarakKm*3000, nil
}

func main() {
	fmt.Println("Sales Order Processor - pertemuan 2")

	qty := []int{10}
	harga := []float64{10}

	fmt.Println(HitungTotalPesanan(qty, harga))

	RingkasanPesanan([]int{10, 10}, []float64{60000, 60000}, 0.11)
}
