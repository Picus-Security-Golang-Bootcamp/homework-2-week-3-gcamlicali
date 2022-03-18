package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


type Books struct {
	Books 	[]Book	`json:"books"`	 
}

type AuthorS struct {
	Name	string 	`json:"AName"`
	ID		int		`json:"AID"`
}

type Book struct {
	ID 		int 	`json:"ID"`
	Name	string 	`json:"Name"`
	Pages	int 	`json:"Pages"`
	Stock	int 	`json:"Stock"`
	Price	int		`json:"Price"`
	StockID string	`json:"StockID"`
	ISBN 	int		`json:"ISBN"`
	Author 	AuthorS	`json:"Author"`
}

var List *Books

var (
	flagCommand 	=	flag.String("command","", "Booklist commands")
	flagName		=	flag.String("name", "", "Name of book parameter")
	flagID			=	flag.Int("ID", 0, "Book id parameter")
	flagQuantity	=	flag.Int("quantity", 0, "Book quantity paramter")
)

func main() {
	flag.Parse()
	OpenList()

	switch *flagCommand {
		case "list":
			GetBookList()
		case "get":
			ShowBookDetail(*flagID)
		case "search":
			FindBookByName(*flagName)
		case "delete":
			DeleteBookByID(*flagID)
		case "buy":
			BuyBookByID(*flagID, *flagQuantity)
		default:
			fmt.Println("\n Wrong Command !!!! \n")
			fmt.Println("You can use: -command <cmdParameter> <otherprmtr> <otherprmt> ")	
			fmt.Println("<cmdparamter> -> :")
			fmt.Println("\tlist")
			fmt.Println("\tsearch <name>")
			fmt.Println("\tdelete <id>")
			fmt.Println("\tbuy <id> <quantity>")
	}	
}
func GetBookList() {
	for _, book:= range List.Books{
		ShowBookDetail(book.ID) //Show all books
	}
}

func OpenList() {
	jsonFile, err := os.Open("books.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
	
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

	//Parse to List struct
	json.Unmarshal(byteValue, &List)
}

func UpdateList() {
	// Convert golang object back to byte
   byteValue, err := json.Marshal(List)
   if err != nil {
	   fmt.Println(err)
   }   
	// Write back to file
	err = ioutil.WriteFile("books.json", byteValue, 0644)	
	

	if err != nil {
	   fmt.Println(err)
   }  
}

func BuyBookByID(bookID int, quantity int) {
	isExist := false

	for i := 0; i < len(List.Books); i++ {
		if(bookID == List.Books[i].ID){	// Search book id in list
			if(List.Books[i].Stock >= quantity){	// if stock enough stock
				List.Books[i].Stock -= quantity
				UpdateList() // We need to change stock in json

				fmt.Println(strconv.Itoa(quantity) + " piece(s) "+ List.Books[i].Name +" sold. Remaining  : " + strconv.Itoa(List.Books[i].Stock))
			} else{
				fmt.Println("Not enough stock!!!")
			}
			isExist = true
			break			
		}
	}

	if(!isExist){
		fmt.Println("There is no ID in List")
	}
}

func ShowBookDetail(bookID int){
	isExist := false

	for _, book:= range List.Books{
		if(bookID == book.ID){
			fmt.Println("Book Name : \t\t" + book.Name)
			fmt.Println("\t ID: \t\t" + strconv.Itoa(book.ID))
			fmt.Println("\t Page count: \t" + strconv.Itoa(book.Pages))
			fmt.Println("\t Stock count: \t" + strconv.Itoa(book.Stock))
			fmt.Println("\t Price: \t" + strconv.Itoa(book.Price))
			fmt.Println("\t StockID: \t" + book.StockID)
			fmt.Println("\t ISBN: \t\t" + strconv.Itoa(book.ISBN))
			fmt.Println("\t Author: \t" + book.Author.Name)
			isExist = true
			break	
		}
	}
	if(!isExist){
		fmt.Println("İstenen kitap listede yok")
	}
}

func FindBookByName(bookName string){
	bookFound := 0

	bookName = strings.ToLower(bookName)
	for _, book:= range List.Books{
		// Compare two names. Is list contains that name?
		if(strings.Contains(strings.ToLower(book.Name), bookName)){
			ShowBookDetail(book.ID)
			bookFound++
		}
	}

	if(bookFound == 0){
		fmt.Println("There is no " + bookName +" named book in list")
	}else {
		fmt.Println("\n\n"+ strconv.Itoa(bookFound) + "	book(s) found")
	}
}

func DeleteBookByID(bookID int){
	isExist := false

	for i := 0; i < len(List.Books); i++ {
		if(bookID == List.Books[i].ID){
			//Delete from list index and update the list
			List.Books = append(List.Books[:i], List.Books[(i + 1):]...)
			
			UpdateList()
			fmt.Println("Book deleted")

			isExist = true
			break
		}
	}

	if(!isExist){
		fmt.Println("There is no ID in List")
	}
}