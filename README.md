# Testing menggunakan Mockery

## Install Mockery
Baca dokumentasi Mockery di : https://github.com/vektra/mockery

### Install mockery via github release
- download Mockery di https://github.com/vektra/mockery/releases
- sesuaikan dengan OS yang anda gunakan. (macOS --> darwin. Windows --> Windows, Linux --> Linux)

### Install mockery di Windows
- extract file yang sudah di download.
- Masukkan ke dalam 1 folder, dan beri nama "mockery" 
- letakkan folder yang sudah terextract ke tempat yang mudah diakses.
- copy path directory tempat anda meletakkan folder mockery, dan tambahkan ke path environment variable di windows. 

### Install mockery di MacOS
- extract file yang sudah di download.
- Masukkan ke dalam 1 folder, dan beri nama "mockery" 
- letakkan folder yang sudah terextract ke tempat yang mudah diakses. (contoh /users/namauser/mockery)
- Buat alias untuk menjalankan mockery, dengan cara edit zshrc atau bashrc (sesuai kan dengan terminal yang digunakan)
```bash
Contoh: 

nano ~/.zshrc

Tambahkan tulisan berikut dibagian bawah: 
alias mockery=”~/mockery/mockery”

Close dengan cara ctrl + x → dan pilih “y” untuk menyimpan.

```
- Close terminal dan buka lagi
- ketik `mockery –help`
- Jika ada propt failed dikarenakan security, silakan masuk ke folder tempat mockery berada, dan coba jalankan `chmod +x mockery`
- Lalu buka setting → security and privacy → klik tab general → klik “Allow”

### Install mockery di Linux
- buka terminal, ketik `wget https://github.com/vektra/mockery/releases/download/v2.14.0/mockery_2.14.0_Linux_x86_64.tar.gz`
- setelah berhasil download, ketik `tar -C /usr/local -xzf mockery_2.14.0_Linux_x86_64.tar.gz`
- check apakah sudah berhasil extract, ketik `cd ~/usr/local` lalu `ls -a`
- jika diperlukan, bisa lakukan rename folder mockery agar lebih mudah saat menjalankan mockery.
- Tambahkan ke path untuk menjalankan mockery, dengan cara edit zshrc atau bashrc (sesuai kan dengan terminal yang digunakan)
```bash
Contoh: 

nano ~/.zshrc

Tambahkan tulisan berikut dibagian bawah: 
export PATH=$PATH:/usr/local/mockery

Tambahkan alias untuk mempermudah pemanggilan:
alias mockery=”/usr/local/mockery”

Close dengan 
cara ctrl + x → dan pilih “y” untuk menyimpan.

```
- Close terminal dan buka lagi
- ketik `mockery –help`
- Jika ada propt failed dikarenakan security, silakan masuk ke folder tempat mockery berada, dan coba jalankan `chmod +x mockery`


## Generate Mock 
- untuk melihat perintah yang bisa kita lakukan menggunakan mockery, silakan ketik `mockery --help`
- masuk ke folder project
- kita dapat men generate mock seluruh interface/kontrak yang ada dengan menjalankan perintah `mockery --all`
- atau kita juga bisa men generate mock satu persatu. jalankan perintah `mockery --dir=features/users --name=Data --filename=UserData.go --structname=UserData`
- note: sesuaikan isi dari `dir`, `name`, `filename`, dan `structname` dengan project yg dibuat.

[*] perintah diatas akan men generate mock interface `Data` yang ada di folder `features/users`. dan akan membuat file baru dengan nama `UserData.go`, dengan nama struct mock `UserData`

[*] `--dir` : directory to search for interfaces (default ".")

[*] `--name` : name or matching regular expression of interface to generate mock for

[*] `--filename` : name of generated file (only works with -name and no regex)

[*] `--structname` : name of generated struct (only works with -name and no regex)
