package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func listMenu() {
	fmt.Println("1. Create New File")
	fmt.Println("2. Open File")
	fmt.Println("3. Write File")
	fmt.Println("4. Delete File")
	fmt.Println("0. Exit")
	fmt.Print("Pilihan [0-4]: ")
}

func createFile() {
	var fileName string
	fmt.Print("\n[!] Masukkan nama file yang akan dibuat: ")
	fmt.Scan(&fileName)

	file, err := os.Create(fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("=========================================\n")
	file.WriteString("Nama\t\tNpm\t\tKelas\t\t\n")
	file.WriteString("=========================================\n")

	defer file.Close()

	fileStat, err := os.OpenFile(fileName+".txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	//fileStat.WriteString("Alfin\t\t50421907\t2IA25\n")
	//fileStat.WriteString("\n")
	fileStat.Sync()

	fmt.Println("[!] Berhasil membuat file dengan Nama =>", fileName, ".txt")
}

func menu() {
	var choose int
	listMenu()
	for {
		fmt.Scan(&choose)
		if choose == 0 {
			fmt.Println("[!] Thank You, exiting.....")
			break
		} else {
			switch choose {
			case 1:
				createFile()
				listMenu()
			case 2:
				var fileName, pil string
				fmt.Print("[!] Nama File: ")
				fmt.Scan(&fileName)

				f, err := os.Open(fileName + ".txt")
				if err != nil {
					fmt.Println("\n[!] Nama file tersebut tidak tersedia..")
					fmt.Print("[?] Apakah anda ingin membuat file (y/!n)? ")
					fmt.Scan(&pil)

					if strings.ToLower(pil) == "y" {
						createFile()
					} else {
						menu()
					}
				}
				defer f.Close()

				s := bufio.NewScanner(f)
				for s.Scan() {
					fmt.Println(s.Text())
				}

				listMenu()
			case 3:
				var nama, npm, kelas, fileName, pil string

				fmt.Print("[!] Nama File: ")
				fmt.Scan(&fileName)

				fmt.Println("-------Input-------")
				fmt.Print("Nama\t: ")
				fmt.Scan(&nama)
				fmt.Print("NPM\t: ")
				fmt.Scan(&npm)
				fmt.Print("Kelas\t: ")
				fmt.Scan(&kelas)

				fileDataAPPENDS, err := os.OpenFile(fileName+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
					panic(err)
				}
				defer fileDataAPPENDS.Close()

				data := nama + "\t\t" + npm + "\t" + kelas + "\n"
				if _, err = fileDataAPPENDS.WriteString(data); err != nil {
					fmt.Println("[!] Nama file tersebut tidak tersedia..")
					fmt.Println("[?] Apakah anda ingin membuat file (y/!n)? ")
					fmt.Scan(&pil)

					if strings.ToLower(pil) == "y" {
						createFile()
					} else {
						menu()
					}
				}

				listMenu()
			case 4:
				var fileName string
				fmt.Print("[!] Nama File: ")
				fmt.Scan(&fileName)

				_, err := os.Stat(fileName + ".txt")

				if os.IsNotExist(err) {
					fmt.Println(err)
				}

				err = os.Remove(fileName + ".txt")

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("[!] Anda menghapus file dengan nama =>", fileName)
				listMenu()
			default:
				fmt.Println("[!] Silakan Pilih Secara Benar!")
				listMenu()
			}
		}
	}
}
func main() {
	menu()
}
