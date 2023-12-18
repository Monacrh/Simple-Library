package main

import (
	"fmt"
)

type Book struct {
	ISBN        int
	Title       string
	Author      string
	Genre       string
	IsAvailable bool
	Favorite    int
}

type Loan struct {
	ISBN       int
	Borrower   string
	BorrowedAt Borrowdate
	BorrowList int
	IsLoan     bool
}

type Borrowdate struct {
	Borrowday   int
	Borrowmonth int
	Borrowyear  int
}

type Library struct {
	Books [1000]Book
	Loans [1000]Loan
}

func main() {
	var L Library
	var ISBN, options int
	var title, author, genre string
	Index := 0
	Idx := 0
	for options != 10 {
		fmt.Println("Choose the number:")
		fmt.Println("1. ADD BOOK")
		fmt.Println("2. EDIT BOOK")
		fmt.Println("3. DELETE BOOK")
		fmt.Println("4. LIST OF BOOKS")
		fmt.Println("5. SEARCHING")
		fmt.Println("6. BORROWING")
		fmt.Println("7. RETURNED")
		fmt.Println("8. SHOW LOAN")
		fmt.Println("9. FAVORITE BOOK")
		fmt.Println("10. EXIT")

		fmt.Scan(&options)
		if options == 1 {
			add(&L, &Index, &ISBN, &title, &author, &genre)
		} else if options == 2 {
			edit(&L, &Index)
		} else if options == 3 {
			delete(&L, &Index, &ISBN)
		} else if options == 4 {
			showList(&L, Index)
		} else if options == 5 {
			searching(L, Index)
		} else if options == 6 {
			Borrow(&L, &Idx, &ISBN, &Index)
		} else if options == 7 {
			returned(&L, &Idx, &ISBN)
		} else if options == 8 {
			isLoans(L, Idx)
		} else if options == 9 {
			favorite(L, Index)
		} else if options == 10 {
			fmt.Println("Bye")
		} else {
			fmt.Println("Invalid option")
		}
	}
	// sortID(&L, &Index)
	// for i := 0; i < Index; i++ {
	// 	if L.Books[i].ISBN != 0 {
	// 		fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[i].ISBN, L.Books[i].Title, L.Books[i].Author, L.Books[i].Genre)
	// 	}
	// }
}

func validISBN(ISBN int) bool {
	count := 0
	temp := ISBN
	for temp != 0 {
		count++
		temp /= 10
	}
	return count == 13
}

func add(L *Library, Index *int, ISBN *int, title, author, genre *string) {
	isAdding := true
	for isAdding {
		fmt.Println("Enter Book ISBN (Enter -1 to stop adding):")
		fmt.Scan(ISBN)

		if *ISBN != -1 {
			if validISBN(*ISBN) {
				fmt.Println("Enter Book Title:")
				fmt.Scan(title)
				fmt.Println("Enter Book Author:")
				fmt.Scan(author)
				fmt.Println("Enter Book Genre:")
				fmt.Scan(genre)

				L.Books[*Index] = Book{
					ISBN:        *ISBN,
					Title:       *title,
					Author:      *author,
					Genre:       *genre,
					IsAvailable: true,
				}
				*Index++
			} else {
				fmt.Println("Invalid ISBN. The ISBN must be exactly 13 digits.")
			}
		} else {
			isAdding = false
		}
	}
}

func sortID(L *Library, Index *int) {
	for i := 1; i < *Index; i++ {
		key := L.Books[i]
		j := i - 1
		for j >= 0 && L.Books[j].ISBN > key.ISBN {
			L.Books[j+1] = L.Books[j]
			j--
		}
		L.Books[j+1] = key
	}
}

func SortSTR(L *Library, Index int, search int) {
	if search == 2 {
		for i := 0; i < Index-1; i++ {
			minIndex := i
			for j := i + 1; j < Index; j++ {
				if L.Books[j].Title < L.Books[minIndex].Title {
					minIndex = j
				}
			}
			temp := L.Books[minIndex]
			L.Books[minIndex] = L.Books[i]
			L.Books[i] = temp
		}
	} else if search == 3 {
		for i := 0; i < Index-1; i++ {
			minIndex := i
			for j := i + 1; j < Index; j++ {
				if L.Books[j].Author < L.Books[minIndex].Author {
					minIndex = j
				}
			}
			temp := L.Books[minIndex]
			L.Books[minIndex] = L.Books[i]
			L.Books[i] = temp
		}
	} else if search == 4 {
		for i := 0; i < Index-1; i++ {
			minIndex := i
			for j := i + 1; j < Index; j++ {
				if L.Books[j].Genre < L.Books[minIndex].Genre {
					minIndex = j
				}
			}
			temp := L.Books[minIndex]
			L.Books[minIndex] = L.Books[i]
			L.Books[i] = temp
		}
	} else if search == 10 {
		for i := 0; i < Index-1; i++ {
			minIndex := i
			for j := i + 1; j < Index; j++ {
				if L.Books[j].Favorite < L.Books[minIndex].Favorite {
					minIndex = j
				}
			}
			temp := L.Books[minIndex]
			L.Books[minIndex] = L.Books[i]
			L.Books[i] = temp
		}
	}
}

func showList(L *Library, Index int) {
	var options int
	isCanceled := false
	fmt.Println("Sorted by: (Enter the number (1, 2, 3), otherwise enter 0 to go back)")
	fmt.Println("1. Title")
	fmt.Println("2. Author")
	fmt.Println("3. Genre")
	fmt.Scan(&options)

	if options == 0 {
		fmt.Println("Canceled")
		isCanceled = true
	}
	if !isCanceled {
		SortSTR(L, Index, options+1)
	}
	for i := 0; i < Index; i++ {
		if L.Books[i].ISBN != 0 {
			fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[i].ISBN, L.Books[i].Title, L.Books[i].Author, L.Books[i].Genre)
		}
	}
}

func delete(L *Library, Index, ISBN *int) {
	sortID(&*L, &*Index)
	isCanceled := false
	deleteIndex := -1

	for i := 0; i < *Index; i++ {
		if L.Books[i].ISBN != 0 {
			fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[i].ISBN, L.Books[i].Title, L.Books[i].Author, L.Books[i].Genre)
		}
	}

	fmt.Println("Choose the ISBN book to delete (Enter -1 to cancel)")
	fmt.Scan(ISBN)

	if *ISBN == -1 {
		fmt.Println("Canceled")
		isCanceled = true
	}

	if !isCanceled {
		for i := 0; i < *Index; i++ {
			if L.Books[i].ISBN == *ISBN {
				deleteIndex = i
			}
		}

		if deleteIndex != -1 {
			for i := deleteIndex; i < *Index-1; i++ {
				L.Books[i] = L.Books[i+1]
			}
			L.Books[*Index-1] = Book{}
		}
	}
}

func edit(L *Library, Index *int) {
	sortID(&*L, &*Index)
	var ISBN int
	var editTitle, editAuthor, editCategory string
	isCanceled := false
	for i := 0; i < *Index; i++ {
		if L.Books[i].ISBN != 0 {
			fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[i].ISBN, L.Books[i].Title, L.Books[i].Author, L.Books[i].Genre)
		}
	}
	fmt.Println("Choose the ISBN book to edit (Enter -1 to cancel)")
	fmt.Scan(&ISBN)
	if ISBN == -1 {
		fmt.Println("Canceled")
		isCanceled = true
	}
	if !isCanceled {
		l := 0
		r := *Index - 1
		found := false
		for l <= r && !found {
			m := (l + r) / 2
			if L.Books[m].ISBN == ISBN {
				found = true
			} else if L.Books[m].ISBN < ISBN {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		if found {
			fmt.Println("Do you want to change the title? (y/n)")
			fmt.Scan(&editTitle)
			if editTitle == "y" {
				fmt.Println("Enter the new title")
				fmt.Scan(&L.Books[l].Title)
			}

			fmt.Println("Do you want to change the Author? (y/n)")
			fmt.Scan(&editAuthor)
			if editAuthor == "y" {
				fmt.Println("Enter the new Author")
				fmt.Scan(&L.Books[l].Author)
			}

			fmt.Println("Do you want to change the Genre? (y/n)")
			fmt.Scan(&editCategory)
			if editCategory == "y" {
				fmt.Println("Enter the new Genre")
				fmt.Scan(&L.Books[l].Genre)
			}
		}
	}
}

func Borrow(L *Library, Idx *int, ISBN *int, Index *int) {
	var found, bd, bm, by, options, lList int
	var name string
	sortID(&*L, &*Index)
	for i := 0; i < *Index; i++ {
		if L.Books[i].ISBN != 0 {
			fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[i].ISBN, L.Books[i].Title, L.Books[i].Author, L.Books[i].Genre)
		}
	}
	isAdding := true
	for isAdding {
		fmt.Println("Enter The Borrower Name:")
		fmt.Scan(&name)
		fmt.Println("Enter the ISBN of the book:")
		fmt.Scan(&found)
		for j := 0; j < *Index; j++ {
			if found == L.Books[j].ISBN && L.Books[j].IsAvailable == true {
				L.Books[j].IsAvailable = false
				L.Books[j].Favorite++
			}
		}
		fmt.Println("List of Price :")
		fmt.Println("1. 14 days (5.000)")
		fmt.Println("2. 30 days (11.000)")
		fmt.Println("3. 40 days (17.000)")
		fmt.Println("Late return Penalty : 1.000/day ")
		fmt.Print("\nSelect a loan period [1-3]: ")
		fmt.Scan(&options)
		if options == 1 {
			lList = 14
		} else if options == 2 {
			lList = 30
		} else if options == 3 {
			lList = 40
		}

		fmt.Println("Enter the borrowing date (dd mm yyyy):")
		fmt.Scan(&bd, &bm, &by)
		getDaysSinceStart(by, bm, bd)
		L.Loans[*Idx] = Loan{
			ISBN:       found,
			Borrower:   name,
			IsLoan:     true,
			BorrowList: lList,
		}
		L.Loans[*Idx].BorrowedAt = Borrowdate{
			Borrowday:   bd,
			Borrowmonth: bm,
			Borrowyear:  by,
		}
		*Idx++
		isAdding = false
		fmt.Println(L.Loans[0].BorrowList)
	}
}

func isLoans(L Library, Idx int) {
	for i := 0; i < Idx; i++ {
		if L.Loans[i].IsLoan == true {
			fmt.Printf("ISBN: %d, Name: %s, BorrowingDate: %d-%d-%d", L.Loans[i].ISBN, L.Loans[i].Borrower,
				L.Loans[i].BorrowedAt.Borrowday, L.Loans[i].BorrowedAt.Borrowmonth, L.Loans[i].BorrowedAt.Borrowyear)
			fmt.Println()
		}
	}
}

func returned(L *Library, Idx *int, ISBN *int) {
	var ID, rd, rm, ry, bd, bm, by, penalty, lList int
	var name string
	fmt.Println("Enter The Returner Name:")
	fmt.Scan(&name)
	fmt.Println("Enter the ISBN of the book:")
	fmt.Scan(&ID)

	returnedIdx := -1

	for i := 0; i < *Idx && returnedIdx == -1; i++ {
		if L.Loans[i].ISBN == ID && L.Loans[i].IsLoan == true && L.Loans[i].Borrower == name {
			bd = L.Loans[i].BorrowedAt.Borrowday
			bm = L.Loans[i].BorrowedAt.Borrowmonth
			by = L.Loans[i].BorrowedAt.Borrowyear
			lList = L.Loans[i].BorrowList
			returnedIdx = i
		}
	}
	fmt.Println("Enter the returning date (dd mm yyyy) :")
	fmt.Scan(&rd, &rm, &ry)

	borrowingDays := getDaysSinceStart(by, bm, bd)
	returnDays := getDaysSinceStart(ry, rm, rd)
	penalty = 0

	days := returnDays - borrowingDays

	if lList == 14 {
		if days >= lList {
			penalty = days - lList
		}
		fmt.Printf("Number of days: %d\n", days)
		fmt.Printf("The Borrowing Cost you need to pay : %d\n", 5000+penalty*1000)

	} else if lList == 30 {
		if days >= lList {
			penalty = days - lList
		}
		fmt.Printf("Number of days: %d\n", days)
		fmt.Printf("The Borrowing Cost you need to pay : %d\n", 11000+penalty*1000)

	} else if lList == 40 {
		if days >= lList {
			penalty = days - lList
		}
		fmt.Printf("Number of days: %d\n", days)
		fmt.Printf("The Borrowing Cost you need to pay : %d\n", 17000+penalty*1000)
	}

	if returnedIdx != -1 {
		for i := returnedIdx; i < *Idx-1; i++ {
			L.Loans[i] = L.Loans[i+1]
		}
		L.Loans[*Idx-1] = Loan{}
	} else {
		fmt.Println("Book not loaned")
	}
}

func leapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func getDaysInYear(year int) int {
	if leapYear(year) {
		return 366
	}
	return 365
}

func getDaysInMonth(month, year int) int {
	if month == 2 {
		if leapYear(year) {
			return 29
		} else {
			return 28
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else {
		return 31
	}
}

func getDaysSinceStart(year, month, day int) int {
	days := 0

	for y := 1; y < year; y++ {
		days += getDaysInYear(y)
	}

	for m := 1; m < month; m++ {
		days += getDaysInMonth(m, year)
	}

	days += day

	return days
}

func searching(L Library, Index int) {
	var options, SINT int
	var SSTRING string
	isCanceled := false
	fmt.Println("Search By: (Enter the number or -1 to exit)")
	fmt.Println("1. Book ISBN")
	fmt.Println("2. Book Title")
	fmt.Println("3. Book Author")
	fmt.Println("4. Book Genre")
	fmt.Scan(&options)
	if options == -1 {
		fmt.Println("Canceled")
		isCanceled = true
	}
	if !isCanceled {
		if options == 1 {
			fmt.Println("Enter the ISBN of the book you want to search for:")
			fmt.Scan(&SINT)
			sortID(&L, &Index)
			l := 0
			r := Index - 1
			found := false
			for l <= r && !found {
				m := (l + r) / 2
				if L.Books[m].ISBN == SINT {
					fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[m].ISBN, L.Books[m].Title, L.Books[m].Author, L.Books[m].Genre)
					found = true
				} else if L.Books[m].ISBN < SINT {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		} else if options == 2 {
			fmt.Println("Enter the Title of the book you want to search for:")
			fmt.Scan(&SSTRING)
			SortSTR(&L, Index, options)
			l := 0
			r := Index - 1
			found := false
			for l <= r && !found {
				m := (l + r) / 2
				if L.Books[m].Title == SSTRING {
					fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[m].ISBN, L.Books[m].Title, L.Books[m].Author, L.Books[m].Genre)
					found = true
				} else if L.Books[m].Title > SSTRING {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		} else if options == 3 {
			fmt.Println("Enter the Author of the book you want to search for:")
			fmt.Scan(&SSTRING)
			SortSTR(&L, Index, options)
			l := 0
			r := Index - 1
			found := false
			for l <= r && !found {
				m := (l + r) / 2
				if L.Books[m].Author == SSTRING {
					fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[m].ISBN, L.Books[m].Title, L.Books[m].Author, L.Books[m].Genre)
					found = true
				} else if L.Books[m].Author > SSTRING {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		} else if options == 4 {
			fmt.Println("Enter the Genre of the book you want to search for:")
			fmt.Scan(&SSTRING)
			SortSTR(&L, Index, options)
			l := 0
			r := Index - 1
			found := false
			for l <= r && !found {
				m := (l + r) / 2
				if L.Books[m].Genre == SSTRING {
					fmt.Printf("ISBN: %d, Title: %s, Author: %s, Genre: %s\n", L.Books[m].ISBN, L.Books[m].Title, L.Books[m].Author, L.Books[m].Genre)
					found = true
				} else if L.Books[m].Genre > SSTRING {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		}
	}
}

func favorite(L Library, Index int) {
	search := 10
	SortSTR(&L, Index, search)
	for i := 0; i < 5; i++ {
		if L.Books[i].ISBN != 0 {
			fmt.Printf("%d: ISBN: %d, Title: %s, Author: %s, Genre: %s, BorrowCount: %d\n", i+1, L.Books[i].ISBN, L.Books[i].Title, L.Books[i].Author, L.Books[i].Genre, L.Books[i].Favorite)
		}
	}
}
