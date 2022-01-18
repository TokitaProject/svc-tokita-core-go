## Menjalankan service secara lokal

Jika kamu ingin menjalankan service ini secara lokal dengan database lokal, kamu tinggal copy *run_local.sh.example* ke *run_local.sh*.

```bash
cp run_local.sh.example run_local.sh
```

Ubah variabel dalam *run_local.sh*.

```bash
export DB_USERNAME="USERNAME_MYSQL_KAMU"
export DB_PASSWORD="PASSWORD_MYSQL_KAMU"
export DB_NAME="NAMA_DATABASE_MYSQL_KAMU"
```

Jalankan shell script untuk menjalankan service

```bash
sh ./run_local.sh
```

Jangan lupa, kamu perlu install **nodemon** terlebih dahulu. lihat [**"Tutorial instalasi nodemon"**] (https://www.npmjs.com/package/nodemon) ini bila kamu belum mengetahui cara instalasi nodemon.

```bash
npm install -g nodemon
```