# Readme, please, biar kamu paham.

## NOTE

Penting, branch master merupakan branch yang merupakan tahap development dan mungkin memiliki bug bila digunakan untuk production. Silakan gunakan branch tag paling akhir.

## Instalasi golang

Ikuti instruksi ini untuk melakukan instalasi golang: [Tutorial instalasi golang](https://go.dev/doc/install)

Versi golang yang dipakai pada boilerplate ini: **go version go1.16.4 darwin/amd64**. Disarankan menggunakan versi golang yang sama.

Bila ingin memperbarui versi golang. Berikut cara memperbarui versi golang:

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

## Menghubungkan service dengan oracle secara lokal

Install sqlplus dan instantclient di MacBook kamu dengan cara sebagai berikut:

```bash
brew tap InstantClientTap/instantclient
brew install instantclient-basic
brew install instantclient-sqlplus
```

Bila kamu belum menggunakan brew, silakan install brew terlebih dahulu dengan [Mengikuti tutorial install brew ini.](https://brew.sh)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

## Konfigurasi deployment.yaml

Environtment pada deployment.yaml memiliki variabel yang sama dengan **run_example.sh**. Hanya saja, value-nya berbeda. Beberapa hal berikut ini penting untuk diperhatikan:

1. **GIN_MODE**. GIN_MODE merupakan mode server pada golang http. Terdapat dua value pada konfigurasi ini. **DEBUG** untuk dev dan staging. **RELEASE** untuk production.

2. **ORACLE_DB_SERVICE_NAME** dan **ORACLE_DB_HOST**. Bila kamu menggunakan koneksi basis data oracle, SERVICE_NAME merupakan kode supaya service kamu dapat terhubung dengan basis data yang tepat. Berikut beberapa SERVICE_NAME berserta HOST oracle pada UII:

**PRODUCTION**

- **SERVICE_NAME**: TRXAKADEMIK; **HOST**: 192.168.8.11;
- **SERVICE_NAME**: TRXPEGAWAI; **HOST**: 192.168.8.12;
- **SERVICE_NAME**: WHUTAMA; **HOST**: 192.168.8.13;
- **SERVICE_NAME**: TRXPMB; **HOST**: 192.168.8.17;
- **SERVICE_NAME**: TRXCBT; **HOST**: 192.168.8.51;

**DEV**

- **SERVICE_NAME**: TRANSDB1; **HOST**: 10.30.20.18;

**STAGING**

- **SERVICE_NAME**: TRANSDB1; **HOST**: 10.30.21.17;

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

7. Penting, konfigurasi **config-goboo.yaml** serta menjalankan **goboo** sebelum melanjutkan ke langkah berikutnya. Ulasan lengkap tentang **goboo** akan dibahas pada section **Bonus: goboo - An artisan framework** paling bawah yak.

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