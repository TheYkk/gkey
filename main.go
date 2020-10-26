package main

import (
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var chars = []string{
	"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z",
	"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z",
	"1","2","3","4","5","6","7","8","9","0",
	"`","~","!","@","#","$","%","^","&","*","(",")","-","_","=","+",
	"[","{","]","}","|",";",":",",","<",".",">","?",
}

func main()  {

	password := flag.String("p","","-p Master Password")
	realm := flag.String("r","","-r Realm example github.com")
	length := flag.Int("l",16,"-l length of password")

	// ? Parse flags
	flag.Parse()

	fancyPassword,err := GenPass(*password,*realm,*length)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	// ? Print password and print new line to stderr
	// ? because we don't want to grap new line on pipeline

	fmt.Print(fancyPassword)
	_, _ = fmt.Fprint(os.Stderr, "\n")
}

func GenPass(password,realm string,length int) (string , error) {
	if password == ""{
		return "",errors.New("password can not be empty")
	}

	if realm == ""{
		return "",errors.New("realm can not be empty")
	}

	if length == 0{
		return "",errors.New("length can not be empty")
	}

	// ? Create sha256sum from string
	h := sha256.New()
	h.Write([]byte(password+"-"+realm))

	// ? Translate hex to base 94 password
	var str []string
	for _,hexa := range h.Sum(nil){
		pos := hexto94(hexa,byte(len(chars)))
		str = append(str, chars[pos])
	}

	// ? Shorten to a certain length
	return strings.Join(str[0:length],""),nil
}


// ? Translate base 255 code to base 94
func hexto94(base byte,max byte) byte  {
	return base % max
}