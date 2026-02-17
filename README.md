# N-Queens Solver (Tucil 1 STIMA)

## Penjelasan Singkat
Program ini dibuat untuk menyelesaikan persoalan N-Queens dengan variasi warna pada papan catur. Tujuan utamanya adalah menempatkan N ratu pada papan berukuran N x N sedemikian rupa sehingga tidak ada dua ratu yang saling menyerang (secara horizontal, vertikal, diagonal) dan tidak ada dua ratu yang menempati petak dengan warna yang sama.

Program ini mengimplementasikan algoritma Brute Force dengan tiga pilihan strategi:
1. **Pure Brute Force**: Mencoba semua kemungkinan penempatan ratu secara naif(exhaustive search).
2. **One Queen Per Column**: Optimasi dengan membatasi satu ratu per kolom.
3. **Early Pruning**: Optimasi lebih lanjut dengan memangkas cabang pencarian yang tidak valid lebih awal.

## Requirements dan Instalasi
Untuk menjalankan program ini, diperlukan:
- **Go (Golang)**: Pastikan Go sudah terinstal di komputer Anda. Anda dapat mengunduhnya di [go.dev](https://go.dev/).

Tidak ada instalasi khusus yang diperlukan selain menginstal Go. Cukup unduh atau *clone* repositori ini ke komputer Anda.

## Cara Kompilasi
Jika Anda ingin mengompilasi program menjadi *executable* file, ikuti langkah berikut:

1. Buka terminal atau command prompt.
2. Arahkan direktori ke folder `src`.
   ```sh
   cd src
   ```
3. Jalankan perintah *build*:
   ```sh
   go build -o ../bin/nqueens_solver.exe real.go
   ```
   (Sesuaikan nama output `.exe` jika menggunakan Windows, atau tanpa ekstensi jika di Linux/macOS).

## Cara Menjalankan dan Menggunakan Program
Anda dapat menjalankan program langsung menggunakan `go run` atau menggunakan file *executable* hasil kompilasi.

**Menggunakan `go run`:**
1. Buka terminal dan arahkan ke folder `src`.
2. Jalankan perintah:
   ```sh
   go run real.go
   ```

**Menggunakan Hasil Kompilasi:**
1. Jika sudah dikompilasi, jalankan file *executable* dari folder `bin` atau tempat Anda menyimpannya.

**Langkah Penggunaan:**
1. Setelah program berjalan, Anda akan diminta memasukkan **Path Input File**. Masukkan path relatif atau absolut ke file teks yang berisi papan catur (contoh file tersedia di folder `test`, misal: `../test/test_4x4.txt`).
2. Pilih **Strategi Penyelesaian** dengan memasukkan angka 1, 2, atau 3.
3. Masukkan **Frekuensi Visualisasi**. Ini menentukan seberapa sering program menampilkan progres pencarian ke layar (semakin kecil angkanya, semakin sering update terminal, yang mungkin memperlambat eksekusi). Masukkan angka besar (misal 1000000) jika tidak ingin terlalu banyak output visual.
4. Program akan memproses dan menampilkan hasil:
   - Posisi ratu pada papan (jika solusi ditemukan).
   - Waktu eksekusi.
   - Total iterasi yang dilakukan.
5. Anda akan diberi opsi untuk menyimpan solusi ke dalam file teks. Jika memilih 'y', hasil akan disimpan di folder `test` dengan nama `solusi_[nama_file_input].txt`.

## Author / Identitas Pembuat
**Nama**: Al Farabi
**NIM**: 13524086
**Kelas**: K-02
**Mata Kuliah**: IF2211 - Strategi Algoritma 
**Universitas**: Institut Teknologi Bandung