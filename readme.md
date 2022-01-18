# Readme, please, biar kamu paham.

## Instalasi golang

Ikuti instruksi ini untuk melakukan instalasi golang: [Tutorial instalasi golang](https://go.dev/doc/install)

Versi golang yang dipakai pada boilerplate ini: **go version go1.16.4 darwin/amd64**

Disarankan menggunakan versi golang terbaru. Cara update golang adalah seperti ini:

```bash
git clone https://github.com/udhos/update-golang
cd update-golang
sudo ./update-golang.sh
```

Ketika cloning dan error, jalankan perintah:

```bash
go mod tidy
```

## Menjalankan service secara lokal

Jika kamu ingin menjalankan service ini secara lokal dengan database lokal, kamu tinggal copy **run_example.sh** ke **run_local.sh**.

```bash
cp run_example.sh run_local.sh
```

Ubah variabel dalam **run_local.sh**.

```bash
export DB_USERNAME="USERNAME_MYSQL_KAMU"
export DB_PASSWORD="PASSWORD_MYSQL_KAMU"
export DB_NAME="NAMA_DATABASE_MYSQL_KAMU"
```

Jalankan shell script untuk menjalankan service

```bash
sh ./run_local.sh
```

Jangan lupa, kamu perlu melakukan instalasi **nodemon** terlebih dahulu. Lihat [Tutorial instalasi nodemon](https://www.npmjs.com/package/nodemon) ini bila kamu belum mengetahui cara instalasi nodemon.

```bash
npm install -g nodemon
```

## Mengubah boilerplate menjadi service

1. Buat repositori baru pada gitlab: [Meciptakan service baru](https://gitlab-cloud.uii.ac.id/projects/new) (Pastikan VPN BSI kamu sudah aktif ya...)

2. Clone repositori boilerplate

```bash
git clone git@gitlab-cloud.uii.ac.id:finance/backend/svc-boilerplate-golang.git
```

3. Masuk ke direktori boilerplate

```bash
cd svc-boilerplate-golang
```

4. Buang .git bawaan

```bash
rm -Rf .git
```

5. Inisialisasi git baru

```bash
git init
```

6. Ubah alamat repositori

Jangan lupa, ubah variabel **TIM_KAMU** dan **NAMA_SERVICE_KAMU** di bawah ini. Atau, copy alamat url dari repositori yang baru kamu buat barusan.

```bash
git remote add origin git@gitlab-cloud.uii.ac.id:TIM_KAMU/backend/NAMA_SERVICE_KAMU
```

7. Penting, konfigurasi **config-goboo.yaml** serta menjalankan **goboo** sebelum melanjutkan ke langkah berikutnya. Ulasan lengkap tentang **goboo** akan dibahas di section paling bawah.

```bash
./goboo config
```

8. Ubah **yaml** di direktori **deploy** dengan teliti juga ya.

9. Unggah repositori

```bash
git add .
git commit -m "Initial commit"
git push -u origin master
```