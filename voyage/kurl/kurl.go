package main

import (
	"github.com/lirlia/go-mypj/voyage/kurl/cmd"
)

func main() {
	cmd.Execute()
}

// func main() {

	kurl := &kurlObj{}

	flag.Parse()
	fqdn := flag.Args()[0]
	kurl.fqdn = fqdn

	err := kurl.sendRequest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := kurl.getResponse()
	fmt.Println(res.payload.stringer())
// }
