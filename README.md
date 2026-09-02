# backend-5053241015

Repo tugas mata kuliah **Pengembangan Backend Dasar**, dibuat dari template [`webdev-if-its/backend-template`](https://github.com/webdev-if-its/backend-template). Ganti judul di atas jadi nama repo kalian sendiri (`backend-nrp`, contoh: `backend-5025201012`).

## Aturan Umum

- Tugas tiap pertemuan disimpan di folder `pertemuan-XX/` pada repo ini.
- Commit message wajib menyebut level yang dicapai: `pertemuan-XX: level N selesai`.
- Deadline push: sebelum pertemuan berikutnya dimulai.
- Semua level dicek otomatis lewat `go test` — baca `pertemuan-XX/SOAL.md` tiap minggu untuk detail levelnya.

## Mengambil Pertemuan Baru Tiap Minggu

Repo ini **tidak otomatis sinkron** dengan template dosen. Begitu ada pertemuan baru, jalankan (ganti `pertemuan-02` sesuai minggu berjalan):

```bash
git fetch https://github.com/webdev-if-its/backend-template.git main
git checkout FETCH_HEAD -- pertemuan-02
```

Perintah ini **aman dijalankan kapan pun** — tidak akan menimpa folder pertemuan lain yang sudah kalian kerjakan, karena hanya mengambil folder yang disebutkan. Setelah itu, commit folder barunya seperti biasa.

Kalau dosen memperbaiki sesuatu di pertemuan yang sudah dirilis (mis. ada bug di test), biasanya cukup ambil ulang file yang diperbaiki saja, bukan seluruh folder — akan diumumkan file mana yang berubah.

---

Bagian di bawah ini **isi bertahap** sesuai level yang sedang kalian kerjakan (lihat `pertemuan-01/SOAL.md`) — heading-nya dicek otomatis, jangan diganti namanya.

## Identitas
- Nama: Kagendra Amadeo Reynara Pratista
- NRP: 5053241015
- Kelas: M

## Commit vs Push
commit itu melakukan sebuah checkpoint untuk file yang nanti akan di simpan file .git local. Push untuk memberitahu remote repository,untuk mengupdate sesuai dengan yang ada di .git local project 

## Reproducibility
jika anggota menjalankan program ini dengan versi yang berbeda, maka anggota yang mempunyai versi go yang lebih tua dibanding dengan versi projectnya akan mendapatkan compilation error

## Catatan Merge Conflict
barik yang mentok berada di file main.go line 32. hal ini menyebabkan merge conflict karena adanya perbedaaan commit dari line di file dan line yang sama. Untuk menyelesaikannya kita tinggal menghapus penanda merge conflictnya seperti head dan nama branch yang di merge. Lalu hapus baris yang tidak tidak di inginkan. 

## Kenapa .gitignore Penting
git ignore penting untuk memisahkan file mana yang bisa di masukan ke git dan file mana yang harusnya hanya ada di local. Contohnya file package yang dimana mempunyai ukuran yang besar tidak perlu di masukan ke git karena bisa di install via runtime masing masing. Juga file seperti .env yang mengandung beberapa rahasia  

## Refleksi
untuk yang paling membingung kan si karena pertama kali pakai golang. Contohnya di level 3, karena terbiasa pakai javascript, ngecek null nya cuman dengan if(variable), eh ternyata error. dan yang saya pakai sekarang masih ngecek nullnya dengan cara apakah parameter args punya panjang lebih dari 1, yang paadahal ada cara yang lebih proper setelah saya cari cari lagi di internet.
