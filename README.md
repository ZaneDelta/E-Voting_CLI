# Aplikasi Sistem Pemungutan Suara Digital (E-Voting) 🗳️

**Projek Tugas Besar (Tubes) - CLI Based**

Aplikasi ini adalah program berbasis *Command Line Interface* (CLI) yang ditulis menggunakan bahasa pemrograman Golang. Aplikasi ini menyimulasikan sistem e-voting yang mencakup manajemen kandidat, manajemen pemilih, proses voting, perhitungan statistik, serta implementasi algoritma *Sorting* dan *Searching*.

## 🗂️ Struktur Proyek Asli

Pada awalnya, proyek ini dipisah menjadi beberapa modul/file agar lebih rapi dan dihubungkan menggunakan `go.mod`. Berikut adalah tampilan struktur aslinya seperti yang terlihat pada gambar di bawah:

![Struktur Folder](image_e0ab70.png)

Daftar file pada struktur asli:
- `main.go`: Entry point program.
- `model.go`: Mendefinisikan tipe data (struct) dan variabel global.
- `menu.go`: Menangani logika menu utama dan navigasi.
- `crud_kandidat.go`: Modul untuk Tambah, Tampil, Edit, dan Hapus (CRUD) data Kandidat.
- `crud_pemilih.go`: Modul untuk Tambah, Tampil, Edit, dan Hapus (CRUD) data Pemilih.
- `voting.go`: Menangani proses inti pemungutan suara.
- `statistik.go`: Modul kalkulasi jumlah suara, persentase, dan pencarian suara terbanyak.
- `sorting_searching.go`: Implementasi algoritma *Selection Sort*, *Insertion Sort*, *Sequential Search*, dan *Binary Search*.
- `helper.go`: Berisi fungsi-fungsi utilitas pendukung (input, konfirmasi, clear screen, dll).
- `dummy.go`: Data awal (dummy data) untuk mempercepat proses testing.
- `go.mod`: File deklarasi module Go.

*(Catatan: File yang ada saat ini sudah digabung menjadi satu untuk mempermudah eksekusi).*

## ✨ Fitur Utama

1. **Sistem Banner Dinamis**: Terdapat pilihan 3 ASCII Art Banner menarik saat aplikasi pertama kali dijalankan.
2. **Kelola Kandidat (CRUD)**: 
   - Tambah kandidat baru dengan nomor urut, nama, dan visi misi.
   - Tampilkan seluruh daftar kandidat beserta jumlah suaranya.
   - Edit dan Hapus data kandidat.
3. **Kelola Pemilih (CRUD)**:
   - Tambah pemilih berdasarkan ID khusus.
   - Tampilkan daftar dan status pemilih (Sudah Memilih / Belum Memilih).
   - Edit dan Hapus data pemilih.
4. **Proses Voting**: Pemilih yang terdaftar dan belum menggunakan suaranya dapat memilih kandidat sesuai nomor urut. Dilengkapi validasi dan *timestamp* waktu voting.
5. **Statistik Hasil Voting**:
   - Tampilan kalkulasi total suara secara akumulatif.
   - Persentase perolehan suara masing-masing kandidat.
   - Penentuan kandidat dengan suara tertinggi/terbanyak.
6. **Sorting (Pengurutan)**:
   - *Selection Sort*: Mengurutkan kandidat berdasarkan jumlah perolehan suara (Bisa Ascending / Descending).
   - *Insertion Sort*: Mengurutkan kandidat berdasarkan nomor urut (Bisa Ascending / Descending).
7. **Searching (Pencarian)**:
   - *Sequential Search*: Mencari kandidat berdasarkan nomor urut tanpa perlu diurutkan sebelumnya.
   - *Binary Search*: Mencari kandidat berdasarkan nomor urut (sistem otomatis mengurutkan data terlebih dahulu sebelum mencari).

## 🚀 Cara Menjalankan

1. Pastikan Anda sudah menginstal [Golang](https://go.dev/dl/) di komputer Anda.
2. Clone atau unduh repositori proyek ini ke dalam komputer Anda.
3. Buka terminal atau *Command Prompt* dan arahkan ke dalam folder proyek.
4. Karena kode sudah digabung, jalankan perintah berikut:
   ```bash
   go run main.go
   ```
   Atau jika menggunakan format modul aslinya, jalankan:
   ```bash
   go run .
   ```

## 👨‍💻 Data Dummy (Untuk Testing)

Untuk mempermudah pengujian saat presentasi, aplikasi sudah terisi dengan data *dummy*:
- **Kandidat**: 
  - Prabowo Subianto (No. 1)
  - Rafi Maulana (No. 2)
  - Ilman Baruna (No. 3)
  - Nicole Reeyn (No. 4)
- **Pemilih**: Terdapat 6 Pemilih terdaftar dengan ID 101 - 106, di mana beberapa pemilih sudah di-*set* belum memberikan suara sehingga dapat langsung diuji pada menu Voting.
