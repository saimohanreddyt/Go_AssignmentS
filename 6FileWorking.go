//File_Reading
	package main

	import ("fmt"
		"os"
		"bufio"
		"io/ioutil"
)



func main ()  {
	data, err := ioutil.ReadFile("Assignment1.txt")
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))

	f,err := os.Open("Assignment1.txt")
	defer f.Close()

	if err != nil{
		fmt.Println(err)
	}
	reader  := bufio.NewReader(f)
	b1,err := reader.Peek(10)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b1))
}
// File_Writing
     package main

     import ("fmt"
		"os"
		"bufio"
		"io/ioutil"
)



func main ()  {
	message := []byte("Hello Friends")
	err := ioutil.WriteFile("new.txt", message,0644)
	if err != nil{
		fmt.Println(err)
	}

  // creating a file

  f,err := os.Create("Assignment1.txt")
  defer f.Close()
  if err != nil{
	fmt.Println(err)
  }
  f.WriteString("Hello friends !!!")
  f.Sync()

  w := bufio.NewWriter(f)

  w.WriteString("Welcome to Kloudone.")

  w.Flush()
}
