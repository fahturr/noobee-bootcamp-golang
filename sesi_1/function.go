package main

import "fmt"

func GetVehcile(vehicle map[string]string) string {
	return fmt.Sprintf("Mobil %s berwarna %s", vehicle["name"], vehicle["color"])
}

func ShowVehcile(message string) {
	fmt.Println(message)
}

func main() {
	var car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	// tampilkan hasil dari variable message
	var message string = GetVehcile(car)
	ShowVehcile(message)

	// output => Mobil BMW berwarna Black
}
