package main

import "fmt"

/* 21.	Aplikasi Chatting
Deskripsi: Aplikasi digunakan untuk pengiriman pesan. Pengguna aplikasi adalah pengguna yang telah melakukan instalasi aplikasi dan admin aplikasi.
Spesifikasi:
a.	Pengguna bisa melakukan registrasi akun.
b.	Admin bisa melakukan persetujuan/penolakan registrasi akun dan mencetak daftar akun.
c.	Pengguna bisa mengirim pesan pribadi kepada akun lain
d.	Pengguna bisa membuat grup chatting dan menambahkan akun lain untuk masuk ke grup chattingnya.
e.	Pengguna bisa mengirim pesan kepada grup chatting.
f.	Pengguna bisa melihat peserta dalam grup.
*/

const NMAXpengguna = 100
const NMAXpesan = 100

type pengguna struct {
	username   string
	email      string
	noTelp     string
	isVerified string
}

type pesan struct {
	pengirim pengguna
	text     string
	penerima pengguna
}

type pesanGrup struct {
	pengirim pengguna
	text     string
}

type Grup struct {
	admin                pengguna
	anggota              [NMAXpengguna]pengguna
	nAnggota, nPesanGrup int
	pesan                [NMAXpesan]pesanGrup
	namaGrup             string
}

type tab [NMAXpengguna]pengguna
type tabChatpri [NMAXpesan]pesan
type tabGrup [1000]Grup

func main() {

	// inisialisasi a sebagai array user
	var a tab

	//insialisasi x sebagai masukan, n sebagai index jumlah pengguna
	var x, n int

	// inisialisasi chPri sebagai array untuk pesan pribadi
	var chPri tabChatpri

	//inisialisasi b sebagai index pengguna yang sedang login
	var b int

	//inisialisasi nPesan sebagai index jumlah pesan pribadi
	var nPesan int

	//inisialisasi nPesanGrup sebagai index jumlah pesan grup
	var maxPesanGrup int

	//inisialisasi grups sebagai array untuk grup
	var grups tabGrup

	//inisialisasi nGrups sebagai index jumlah grup
	var nGrups int

	// Data Dummy Pengguna
	a[0].username = "joshua"
	a[0].email = "shua"
	a[0].noTelp = "012346580"
	a[0].isVerified = "accepted"

	a[1].username = "wonwoo"
	a[1].email = "wonu"
	a[1].noTelp = "097625342"
	a[1].isVerified = "accepted"

	a[2].username = "jeonghan"
	a[2].email = "hani"
	a[2].noTelp = "012846580"
	a[2].isVerified = "accepted"

	a[3].username = "nawra"
	a[3].email = "nawrakirana"
	a[3].noTelp = "103022300054"
	a[3].isVerified = "accepted"

	a[4].username = "rahmah"
	a[4].email = "rahmahaisyah"
	a[4].noTelp = "103022300014"
	a[4].isVerified = "accepted"

	a[5].username = "luthfi"
	a[5].email = "luthfiramadhan"
	a[5].noTelp = "103022300147"
	a[5].isVerified = "accepted"

	n = 6

	// Data Dummy Grup
	grups[0].namaGrup = "SE-47-04"
	grups[0].admin = a[4]      // Rahmah sebagai Admin
	grups[0].anggota[0] = a[4] // Rahmah
	grups[0].anggota[1] = a[3] // Nawra
	grups[0].anggota[2] = a[5] // Luthfi
	grups[0].nAnggota = 3

	grups[1].namaGrup = "sebong"
	grups[1].admin = a[2]      // Jeonghan sebagai Admin
	grups[1].anggota[0] = a[2] // Jeonghan
	grups[1].anggota[1] = a[1] // Wonwoo
	grups[1].anggota[2] = a[0] // Joshua
	grups[1].nAnggota = 3

	nGrups = 2

	for x != 3 {
		fmt.Println("-------------------------")
		fmt.Println("    Welcome to ChatMe!    ")
		fmt.Println("Teman chatting setiamu ^^")
		fmt.Println("-------------------------")
		fmt.Println("Apakah kamu Pengguna atau Admin?")
		fmt.Println("1. Pengguna")
		fmt.Println("2. Admin")
		fmt.Println("3. Keluar dari program")
		fmt.Scan(&x)
		if x == 1 {
			userSelection(&a, &chPri, &grups, &n, &b, &nPesan, &nGrups, &maxPesanGrup)
		} else if x == 2 {
			admin(&a, n)
		}
	}
}

func userSelection(user *tab, chPri *tabChatpri, grups *tabGrup, n *int, b *int, nPesan *int, nGrups *int, maxPesanGrup *int) {
	/*
		IS. terdefinisi tipe bentukan array user, array chPri, array grups, variabel n, variabel b, variabel nPesan,
		 variabel nGrups, variabel maxPesanGrup. berisi masukan berupa variabel pilih
		FS. Pengguna dapat melanjutkan ke procedure selanjutnya (login, registrasi) berdasarkan masukan
	*/
	var pilih int
	for pilih != 3 {
		fmt.Println()
		fmt.Println("Apakah anda sudah memiliki akun sebelumnya?")
		fmt.Println("Jika sudah pilih Login, jika tidak pilih registrasi :")
		fmt.Println("1. Login")
		fmt.Println("2. Registrasi")
		fmt.Println("3. Kembali")
		fmt.Scan(&pilih)
		if pilih == 1 {
			login(user, chPri, grups, n, b, nPesan, nGrups, maxPesanGrup)
		} else if pilih == 2 {
			registrasi(user, n)
		}
	}
}

func registrasi(user *tab, n *int) {
	/*
	   IS. terdefinisi tipe bentukan array user dan n yang berupa jumlah data pengguna,
	   data pengguna baru masih kosong
	   FS. menambahkan masukan ke dalam array user berupa data pengguna, memperbarui n,
	   dan menginisialisasi isVerified sebagai "pending"
	*/
	if *n >= NMAXpengguna {
		fmt.Println("Tidak bisa registrasi, pengguna sudah penuh")
		return
	}

	fmt.Print("Masukkan username: ")
	fmt.Scan(&user[*n].username)
	fmt.Print("Masukkan email: ")
	fmt.Scan(&user[*n].email)
	fmt.Print("Masukkan nomor telepon: ")
	fmt.Scan(&user[*n].noTelp)
	user[*n].isVerified = "pending"
	fmt.Println()
	fmt.Println("Registrasi berhasil, silahkan tunggu persetujuan admin")
	*n++
}

func admin(user *tab, n int) {
	/*
	   IS. berisi menu admin, berupa masukan
	   FS. melanjutkan ke procedure selanjutnya berdasarkan masukan
	*/
	var pilih int
	for pilih != 3 {
		fmt.Println()
		fmt.Println("Selamat datang !")
		fmt.Println("Menu :")
		fmt.Println("1. Verifikasi Akun")
		fmt.Println("2. Cetak Daftar Akun")
		fmt.Println("3. Kembali")
		fmt.Scan(&pilih)
		if pilih == 1 {
			verifikasi(user, n)
		} else if pilih == 2 {
			cetakDaftar(user, &n)
		}
	}
}

func verifikasi(user *tab, n int) {
	/*
		       IS. terdefinisi tipe bentukan array user dan variabel n, menampilkan data array yang belum terverifikasi
					atau masih "pending", berisi masukan
		       FS. memperbarui data dalam array user, mengubah "pending" menjadi "accepted" atau "rejected",
			   mengeluarkan output "Semua akun sudah terverifikasi" apabila sudah terverifikasi semua.
	*/
	var username, email, notelp bool
	var pilih int

	fmt.Println()
	fmt.Println("Daftar akun yang belum terverifikasi :")
	for i := 0; i < n; i++ {
		if (*user)[i].isVerified == "pending" {
			fmt.Println("Akun no :", i+1)
			fmt.Println("Username :", (*user)[i].username)
			username = false
			email = false
			notelp = false

			for x := 0; x < n; x++ {
				if (*user)[i].username == (*user)[x].username && (*user)[x].isVerified == "accepted" {
					username = true
				}
			}
			if username {
				fmt.Println("Username sudah pernah didaftarkan")
				fmt.Println()
			} else {
				fmt.Println("Username boleh dipakai")
				fmt.Println()
			}

			fmt.Println("Email :", (*user)[i].email)
			for x := 0; x < n; x++ {
				if (*user)[i].email == (*user)[x].email && (*user)[x].isVerified == "accepted" {
					email = true
				}
			}
			if email {
				fmt.Println("Email sudah pernah didaftarkan")
				fmt.Println()
			} else {
				fmt.Println("Email boleh dipakai")
				fmt.Println()
			}

			fmt.Println("Nomor Telepon :", (*user)[i].noTelp)
			for x := 0; x < n; x++ {
				if (*user)[i].noTelp == (*user)[x].noTelp && (*user)[x].isVerified == "accepted" {
					notelp = true
				}
			}
			if notelp {
				fmt.Println("Nomor Telepon sudah pernah didaftarkan")
				fmt.Println()
			} else {
				fmt.Println("Nomor Telepon boleh dipakai")
				fmt.Println()
			}

			fmt.Println("Apakah anda ingin memverifikasi akun ini?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tolak")
			fmt.Println("0. Kembali")
			var x int
			fmt.Scan(&x)
			if x == 1 {
				(*user)[i].isVerified = "accepted"
				fmt.Println("Akun sudah terverifikasi")
			} else if x == 2 {
				(*user)[i].isVerified = "rejected"
				fmt.Println("Akun sudah ditolak")
			} else {
				fmt.Println("Masukan salah, silahkan masukkan kembali")
			}
			fmt.Println("")
		}
	}
	fmt.Println("Semua akun sudah terverifikasi")
	fmt.Println("------------------------------")
	for pilih != 1 {
		fmt.Println("Apakah anda ingin kembali?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Scan(&pilih)
	}
}

func hapusAkun(user *tab, n *int) {
	/*
	   IS. terdefinisi tipe bentukan array user dan variabel n
	   FS. memasukkan data, menghapus akun berdasarkan masukan username
	*/
	var username string
	var i, j int
	var ketemu bool
	var hapus int
	fmt.Println("Akun mana yang ingin dihapus?")
	fmt.Print("Masukkan username:")
	fmt.Scan(&username)

	ketemu = false
	i = 0
	for !ketemu && i < *n {
		if username == (*user)[i].username {
			ketemu = true
		} else {
			i++
		}
	}

	if !ketemu {
		fmt.Println("Username tidak ketemu")
	} else {
		fmt.Println("Apakah anda yakin ingin menghapus akun", username, "?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Scan(&hapus)
		if hapus == 1 {
			for j = i; j < *n; j++ {
				(*user)[j] = (*user)[j+1]
			}
			*n--
			fmt.Println("Akun berhasil dihapus")
		} else if hapus == 2 {
			fmt.Println("Akun tidak dihapus")
		} else {
			fmt.Println("Masukan salah, silahkan masukkan kembali")
		}
	}
}

func cetakDaftar(user *tab, n *int) {
	/*
		       IS. terdefinisi tipe bentukan array user dan variabel n
		       FS. mengurutkan akun menggunakan INSERTION SORT SECARA ASCENDING, menampilkan daftar pengguna,
			   memperbarui array user dan variabel n jika pengguna memilih hapus akun
	*/
	var kembali int
	for kembali != 2 {
		fmt.Println()
		fmt.Println("Daftar akun yang sudah terverifikasi :")
		var i, pass int
		var temp pengguna
		pass = 1
		for pass <= *n-1 {
			i = pass
			temp = user[pass]
			for i > 0 && temp.username < user[i-1].username {
				user[i] = user[i-1]
				i--
			}
			user[i] = temp
			pass++
		}
		terverifikasi := false
		j := 1
		fmt.Println("----------------------------------------------------")
		fmt.Printf("%-3s %10s %20s %15s\n", "No", "Username", "Email", "No Telpon")
		fmt.Println("----------------------------------------------------")
		for i := 0; i < *n; i++ {
			if user[i].isVerified == "accepted" {
				fmt.Printf("%-3d %10s %20s %15s\n", j, user[i].username, user[i].email, user[i].noTelp)
				terverifikasi = true
				j++
			}
		}
		if !terverifikasi {
			fmt.Println("belum ada akun yang terverifikasi")
		} else {
			fmt.Println()
			fmt.Println("Apa yang ingin anda lakukan?")
			fmt.Println("1. Hapus Akun")
			fmt.Println("2. Kembali")
			fmt.Scan(&kembali)
			if kembali == 1 {
				hapusAkun(user, n)
			}
		}
	}
}

func login(user *tab, chPri *tabChatpri, grups *tabGrup, n *int, b *int, nPesan *int, nGrups *int, indexGrup *int) {
	/*
		       IS. terdefinisi tipe bentukan array user, chPri, grups,
			    variabel n, b, nPesan, nGrups dan  indexGrup
		       FS. mengecek apakah username dan email yang dimasukkan sudah terdaftar atau belum,
			   jika sudah terdaftar maka akan masuk ke menu utama akun (chat), jika belum kembali meminta masukan
	*/
	var username, email string
	var autentifikasi bool
	var x int
	fmt.Println("Masukkan Username:")
	fmt.Scan(&username)
	fmt.Println("Masukkan Email:")
	fmt.Scan(&email)

	autusername := false
	autemail := false
	for i := 0; i < *n; i++ {
		if user[i].username == username && user[i].isVerified == "accepted" {
			autusername = true
			*b = i
		}
		if user[i].email == email && user[i].isVerified == "accepted" {
			autemail = true
			*b = i
		}
	}

	if autusername && autemail {
		fmt.Println()
		fmt.Println("Login berhasil!")
		autentifikasi = true
	} else if !autusername && autemail {
		fmt.Println()
		fmt.Println("Username salah.")
		autentifikasi = true
	} else if autusername && !autemail {
		fmt.Println()
		fmt.Println("Email salah.")
		autentifikasi = true
	} else {
		fmt.Println()
		fmt.Println("Akun belum terverifikasi.")
		fmt.Println("Apakah anda ingin kembali?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak, saya ingin login akun lain.")
		fmt.Scan(&x)
		if x == 2 {
			login(user, chPri, grups, n, b, nPesan, nGrups, indexGrup)
		}
	}

	if autentifikasi == true {
		chat(user, chPri, grups, n, *b, nPesan, nGrups, indexGrup)
	}
}

func chat(user *tab, chPri *tabChatpri, grups *tabGrup, n *int, b int, nPesan *int, nGrups *int, indexGrup *int) {
	/*
	   IS. terdefinisi tipe bentukan array user, chPri, grups,
	    variabel n, b, nPesan, nGrups dan  indexGrup.
	   FS. berisi menu pilihan masukan, mengarahkan ke procedure selanjutnya (chatPriv atau menuGroup) berdasarkan masukan
	*/
	var x int
	for x != 3 {
		fmt.Println("-------------------------")
		fmt.Println("Selamat datang", user[b].username, "!")
		fmt.Println("-------------------------")
		fmt.Println("Menu chat:")
		fmt.Println("1. Private Chat")
		fmt.Println("2. Grup Chat")
		fmt.Println("3. Kembali")
		fmt.Scan(&x)
		if x == 1 {
			chatPriv(user, chPri, n, b, nPesan)
		} else if x == 2 {
			menuGroup(user, grups, n, b, nGrups, indexGrup)
		}
	}
}

func chatPriv(user *tab, chPri *tabChatpri, n *int, b int, nPesan *int) {
	/*
		           IS. terdefinisi tipe bentukan array user, chPri, n, b, nPesan.
		           FS. berisi daftar kontak dan masukan, mengoutput pesan jika sudah ada,
				   memperbarui array chPri sesuai masukan, menambahkan nilai nPesan
	*/
	var i, pilih, j int
	var y int
	var pilihindex int
	var pesan string
	var pilihpesan int
	y = 1
	fmt.Println("Daftar Kontak :")
	for i = 0; i < *n; i++ {
		if i != b && user[i].isVerified == "accepted" {
			fmt.Println(y, ".", user[i].username)
			y++
		}
	}
	fmt.Println("Silahkan pilih kontak yang ingin anda chat :")
	fmt.Scan(&pilih)
	y = 1
	found := false

	for i := 0; i < *n && !found; i++ {
		if i != b && user[i].isVerified == "accepted" {
			if y == pilih {
				pilihindex = i
				found = true
			} else {
				y++
			}
		}
	}
	if pilih == y {
		fmt.Println("--------------------------------")
		fmt.Println(user[pilihindex].username)
		fmt.Println("ONLINE")
		fmt.Println("--------------------------------")
		fmt.Println()
	}
	for j = 0; j < *nPesan; j++ {
		if ((chPri[j].pengirim.username == user[b].username) && (chPri[j].penerima.username == user[pilihindex].username)) || ((chPri[j].pengirim.username == user[pilihindex].username) && (chPri[j].penerima.username == user[b].username)) {
			if chPri[j].pengirim.username == user[b].username {
				fmt.Println("Anda :")
				fmt.Println(chPri[j].text)
				fmt.Println()
			} else {
				fmt.Println(chPri[j].pengirim.username)
				fmt.Println(chPri[j].text)
				fmt.Println()
			}
		}
	}
	fmt.Println("Masukkan pesan anda...")
	fmt.Scan(&pesan)
	chPri[*nPesan].pengirim.username = user[b].username
	chPri[*nPesan].penerima.username = user[pilihindex].username
	chPri[*nPesan].text = pesan
	*nPesan++
	for pilihpesan != 2 {
		fmt.Println("Apakah anda ingin kembali mengirim pesan?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Scan(&pilihpesan)
		if pilihpesan == 1 {
			fmt.Println("Masukkan pesan anda...")
			fmt.Scan(&pesan)
			chPri[*nPesan].pengirim.username = user[b].username
			chPri[*nPesan].penerima.username = user[pilihindex].username
			chPri[*nPesan].text = pesan
			*nPesan++
		}
	}
}

func menuGroup(user *tab, grups *tabGrup, n *int, b int, nGrups *int, indexGrup *int) {
	/*
	   IS. terdefinisi tipe bentukan array user, grups, n, b, nGrups dan  indexGrup. berisi menu pilihan masukan
	   FS. mengarahkan ke procedure selanjutnya (buatGrup atau daftarGrup) berdasarkan masukan
	*/
	var pilih int
	fmt.Println()
	fmt.Println("Menu Grup :")
	fmt.Println("1.Buat Grup")
	fmt.Println("2.Daftar Grup")
	fmt.Scan(&pilih)
	if pilih == 1 {
		buatGrup(*user, grups, *n, b, nGrups)
	} else if pilih == 2 {
		daftarGrup(*user, grups, *n, b, nGrups, *indexGrup)
	} else {

	}
}

func buatGrup(user tab, grups *tabGrup, n int, b int, nGrups *int) {
	/*
	   IS. terdefinisi tipe bentukan array user, grups, n, b, nGrups. berisi masukan
	   FS. memperbarui array grups dan menambah nilai nGrups
	*/
	var pilih int
	var pilihindex int
	var pilihlagi int
	var y int
	fmt.Println("Nama Grup")
	fmt.Scan(&grups[*nGrups].namaGrup)
	grups[*nGrups].admin = user[b]
	grups[*nGrups].nAnggota = 1
	grups[*nGrups].anggota[0] = user[b]
	fmt.Println("Tambah Kontak:")
	y = 1
	for i := 0; i < n; i++ {
		if i != b && user[i].isVerified == "accepted" {
			fmt.Println(y, ".", user[i].username)
			y++
		}
	}

	// Pilih pengguna
	fmt.Println("Siapa yang ingin ditambahkan ke dalam grup?")
	fmt.Scan(&pilih)

	// Sesuaikan pilihan dengan indeks array
	pilihindex = pilih - 1

	y = 1
	found := false
	for i := 0; i < n && !found; i++ {
		if i != b && user[i].isVerified == "accepted" {
			if y == pilih {
				pilihindex = i
				found = true
			} else {
				y++
			}
		}
	}

	if found {
		grups[*nGrups].anggota[grups[*nGrups].nAnggota] = user[pilihindex]
		grups[*nGrups].nAnggota++
	}

	for pilihlagi != 2 {
		fmt.Println("Apakah anda ingin menambahkan kontak lain?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Scan(&pilihlagi)
		if pilihlagi == 1 {
			fmt.Println("Daftar Kontak :")
			y = 1
			for i := 0; i < n; i++ {
				if i != b && !isMember(grups[*nGrups].anggota, grups[*nGrups].nAnggota, user[i].username) && user[i].isVerified == "accepted" {
					fmt.Println(y, ".", user[i].username)
					y++
				}
			}

			fmt.Scan(&pilih)
			pilihindex = pilih - 1
			y = 1
			found = false
			for i := 0; i < n && !found; i++ {
				if i != b && !isMember(grups[*nGrups].anggota, grups[*nGrups].nAnggota, user[i].username) && user[i].isVerified == "accepted" {
					if y == pilih {
						pilihindex = i
						found = true
					} else {
						y++
					}
				}
			}

			if found {
				grups[*nGrups].anggota[grups[*nGrups].nAnggota] = user[pilihindex]
				grups[*nGrups].nAnggota++
			}
		}
	}
	*nGrups++
}

func isMember(anggota [NMAXpengguna]pengguna, nAnggota int, username string) bool {
	/*
	   IS. -
	   FS. mengembalikan nilai true jika username terdapat dalam array anggota, dan false jika tidak
	*/
	for i := 0; i < nAnggota; i++ {
		if anggota[i].username == username {
			return true
		}
	}
	return false
}

func daftarGrup(user tab, grups *tabGrup, n int, b int, nGrups *int, indexGrup int) {
	/*
		           IS. terdefinisi tipe bentukan array user, grups, n, b, nGrups dan  indexGrup.
				   FS. mengurutkan grup berdasarkan huruf SECARA ASCENDING MENGGUNAKAN SELECTION SORT,
				   menampilkan daftar grup yang dimiliki, mengarahkan ke procedure selanjutnya (chatGrup)
	*/
	var y int
	var i, j int
	var pass int
	var temp Grup
	var min int
	pass = *nGrups - 1
	for i = 0; i < pass; i++ {
		min = i
		for j = i + 1; j <= pass; j++ {
			if grups[j].namaGrup < grups[min].namaGrup {
				min = j
			}
		}
		temp = grups[i]
		grups[i] = grups[min]
		grups[min] = temp
	}
	y = 1
	fmt.Println("Daftar Grup :")
	for i := 0; i < *nGrups; i++ {
		for j := 0; j < grups[i].nAnggota; j++ {
			if grups[i].anggota[j].username == user[b].username {
				fmt.Println(y, ".", grups[i].namaGrup)
				y++
			}
		}
	}
	if y > 1 {
		chatGrup(user, grups, &n, b, nGrups, &indexGrup)
	} else {
		fmt.Println("Anda belum memiliki grup")
	}
}

func chatGrup(user tab, grups *tabGrup, n *int, b int, nGrups *int, indexGrup *int) {
	/*
		           IS. terdefinisi tipe bentukan array user, grups, n, b, nGrups dan  indexGrup.
		           berisi masukan pilih grup, output pesan yang tersimpan, pilihan(mengirim pesan,
				lihat anggota grup, tambah akun, hapus akun)
		           FS. memperbarui array grups (tambah pesan), mengarahkan ke procedure tambah akun atau
				   hapus akun sesuai pilihan
	*/
	var y int
	var pilih, pilihlagi, pilihpesan int
	var i, j int
	var temp pengguna
	var pilihindex int
	fmt.Println("Silahkan pilih grup yang ingin anda chat :")
	fmt.Scan(&pilih)
	y = 1
	found := false
	for i := 0; i < *nGrups && !found; i++ {
		for j := 0; j < grups[i].nAnggota; j++ {
			if grups[i].anggota[j].username == user[b].username {
				if y == pilih {
					pilihindex = i
					found = true
				} else {
					y++
				}
			}
		}
	}
	if found {
		fmt.Println("--------------------------------")
		fmt.Println(grups[pilihindex].namaGrup)
		fmt.Println(grups[pilihindex].nAnggota, "online")
		fmt.Println("--------------------------------")
		fmt.Println()
	}

	for j := 0; j < grups[pilihindex].nPesanGrup; j++ {
		if grups[pilihindex].pesan[j].pengirim.username == user[b].username {
			fmt.Println("Anda :")
			fmt.Println(grups[pilihindex].pesan[j].text)
			fmt.Println()
		} else {
			fmt.Println(grups[pilihindex].pesan[j].pengirim.username)
			fmt.Println(grups[pilihindex].pesan[j].text)
			fmt.Println()
		}
	}
	fmt.Println("1. Mengirim Pesan Grup")
	fmt.Println("2. Melihat anggota grup")
	fmt.Println("3. Kembali")
	fmt.Scan(&pilih)
	if pilih == 1 {
		fmt.Println("Masukkan pesan anda...")
		fmt.Scan(&grups[pilihindex].pesan[grups[pilihindex].nPesanGrup].text)
		grups[pilihindex].pesan[grups[pilihindex].nPesanGrup].pengirim.username = user[b].username
		grups[pilihindex].nPesanGrup++
		for pilihpesan != 2 {
			fmt.Println("Apakah anda ingin mengirim pesan lagi?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Scan(&pilihpesan)
			if pilihpesan == 1 {
				fmt.Println("Masukkan pesan anda...")
				fmt.Scan(&grups[pilihindex].pesan[grups[pilihindex].nPesanGrup].text)
				grups[pilihindex].pesan[grups[pilihindex].nPesanGrup].pengirim.username = user[b].username
				grups[pilihindex].nPesanGrup++
			}
		}
	} else if pilih == 2 {
		if pilihindex >= 0 && pilihindex < *nGrups {
			for i = 1; i < grups[pilihindex].nAnggota; i++ {
				temp = grups[pilihindex].anggota[i]
				j = i - 1
				for j >= 0 && grups[pilihindex].anggota[j].username > temp.username {
					grups[pilihindex].anggota[j+1] = grups[pilihindex].anggota[j]
					j = j - 1
				}
				grups[pilihindex].anggota[j+1] = temp
			}
			fmt.Println()
			fmt.Println("Daftar Anggota dalam Grup", grups[pilihindex].namaGrup, ":")
			j = 1
			for i = 0; i < grups[pilihindex].nAnggota; i++ {
				if grups[pilihindex].admin.username == grups[pilihindex].anggota[i].username {
					fmt.Println(j, ".", grups[pilihindex].anggota[i].username, "[Admin Grup]")
					j++
				}
			}
			for i = 0; i < grups[pilihindex].nAnggota; i++ {
				if grups[pilihindex].admin.username != grups[pilihindex].anggota[i].username {
					fmt.Println(j, ".", grups[pilihindex].anggota[i].username)
					j++
				}
			}
			fmt.Println()
			fmt.Println("Apa yang ingin anda lakukan?")
			fmt.Println("1. Tambah Anggota")
			fmt.Println("2. Hapus Anggota")
			fmt.Println("3. Kembali")
			fmt.Scan(&pilihlagi)
			if pilihlagi == 1 {
				*indexGrup = pilihindex
				tambahAnggota(&user, grups, *n, b, indexGrup)
			} else if pilihlagi == 2 {
				*indexGrup = pilihindex
				hapusAnggotaGrup(&user, grups, *n, b, nGrups, indexGrup)
			}
		}
	}
}

func tambahAnggota(user *tab, grups *tabGrup, n int, b int, indexGrup *int) {
	/*
	   IS. terdefinisi tipe bentukan array user, grups, n, b dan indexGrup.
	   berisi daftar pengguna yang ingin ditambahkan, masukan pilih pengguna.
	   FS. memperbarui array grups (anggota)
	*/

	var pilih, pilihindex, i, y, pilihlagi int
	fmt.Println("Daftar kontak")
	if grups[*indexGrup].admin.username == user[b].username {
		// Menampilkan pengguna yang bukan anggota grup
		y = 1
		for i = 0; i < n; i++ {
			if i != b && !isMember(grups[*indexGrup].anggota, grups[*indexGrup].nAnggota, user[i].username) && user[i].isVerified == "accepted" {
				fmt.Println(y, ". ", user[i].username)
				y++
			}
		}
		fmt.Println()
		fmt.Println("Pilih pengguna untuk ditambahkan ke grup:")
		fmt.Scan(&pilih)
		pilihindex = pilih - 1
		y = 1
		found := false
		for i = 0; i < n && !found; i++ {
			if i != b && !isMember(grups[*indexGrup].anggota, grups[*indexGrup].nAnggota, user[i].username) && user[i].isVerified == "accepted" {
				if y == pilih {
					pilihindex = i
					found = true
				} else {
					y++
				}
			}
		}

		if found {
			grups[*indexGrup].anggota[grups[*indexGrup].nAnggota] = user[pilihindex]
			grups[*indexGrup].nAnggota++
			fmt.Println("Pengguna berhasil ditambahkan ke grup.")
			for pilihlagi != 2 {
				fmt.Println("Apakah anda ingin menambahkan kontak lain?")
				fmt.Println("1. Ya")
				fmt.Println("2. Tidak")
				fmt.Scan(&pilihlagi)
				if pilihlagi == 1 {
					if grups[*indexGrup].admin.username == user[b].username {
						// Menampilkan pengguna yang bukan anggota grup
						y = 1
						for i = 0; i < n; i++ {
							if i != b && !isMember(grups[*indexGrup].anggota, grups[*indexGrup].nAnggota, user[i].username) {
								fmt.Println(y, ". ", user[i].username)
								y++
							}
						}
						fmt.Println()
						fmt.Println("Pilih pengguna untuk ditambahkan ke grup:")
						fmt.Scan(&pilih)
						pilihindex = pilih - 1
						y = 1
						found := false
						for i = 0; i < n && !found; i++ {
							if i != b && !isMember(grups[*indexGrup].anggota, grups[*indexGrup].nAnggota, user[i].username) {
								if y == pilih {
									pilihindex = i
									found = true
								} else {
									y++
								}
							}
						}

						if found {
							grups[*indexGrup].anggota[grups[*indexGrup].nAnggota] = user[pilihindex]
							grups[*indexGrup].nAnggota++
							fmt.Println("Pengguna berhasil ditambahkan ke grup.")
						}
					}
				}
			}
		} else {
			fmt.Println("Pengguna yang dipilih tidak valid.")
		}
	} else {
		fmt.Println("Anda bukan admin grup, tidak bisa menambah akun.")
	}
}

func hapusAnggotaGrup(user *tab, grups *tabGrup, n int, b int, nGrups *int, indexGrup *int) {
	/*
			    IS. terdefinisi tipe bentukan array user, grups, n, b, nGrups dan indexGrup.
		        berisi daftar pengguna yang ingin dihapus, masukan berupa pilih pengguna.
		        FS. mencari anggota yang ingin dihapus dari grup dalam array menggunakan
				BINARY SEARCH, memperbarui array grups (anggota)
	*/
	var kanan, kiri, tengah int
	var ketemu bool
	var username string
	var y, i int
	fmt.Println("Daftar Anggota dalam Grup", grups[*indexGrup].namaGrup, ":")
	if grups[*indexGrup].admin.username == user[b].username {
		y = 1
		for i = 0; i < n; i++ {
			if i != b && isMember(grups[*indexGrup].anggota, grups[*indexGrup].nAnggota, user[i].username) && user[i].isVerified == "accepted" {
				fmt.Println(y, ". ", user[i].username)
				y++
			}
		}
		fmt.Println()
		fmt.Println("masukkan username pengguna yang ingin dihapus dari grup:")
		if grups[*indexGrup].admin.username == user[b].username {
			fmt.Scan(&username)
			ketemu = false
			kiri = 0
			kanan = grups[*indexGrup].nAnggota - 1
			for kiri <= kanan && !ketemu {
				tengah = (kiri + kanan) / 2
				if username > grups[*indexGrup].anggota[tengah].username {
					kiri = tengah + 1
				} else if username < grups[*indexGrup].anggota[tengah].username {
					kanan = tengah - 1
				} else {
					ketemu = true
				}
			}
			if ketemu {
				for i := tengah; i < grups[*indexGrup].nAnggota; i++ {
					grups[*indexGrup].anggota[i] = grups[*indexGrup].anggota[i+1]
				}
				grups[*indexGrup].nAnggota--
			} else {
				fmt.Println("Username tidak ditemukan")
				fmt.Scan()
			}
		} else {
			fmt.Println("Anda bukan admin grup, tidak bisa menghapus akun.")
		}
	}
}
