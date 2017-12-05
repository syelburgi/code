package main

import (
	"fmt"
	"myalgorithms/hash"
	"myalgorithms/list"
)

//callback registartion
func myalgo(d interface{}) int {
	return 0
}

func main() {
	l, ok := lists.New()

	if ok == true {
		//l.Insert("Sandeep")
		//l.Insert("yelburgi")
		//l.Insert("acalvio")
		l.Insert(10)
		l.Insert(20)
		l.Delete("yelburgi")
		l.Delete(10)
		//l.Print()
	}

	hashb := hash.Newhashtable(2, myalgo)

	hashb.Insert(10)
	hashb.Insert(20)
	ok = hashb.Search(10)
	if ok == true {
		fmt.Println("Search 10 worked")
	} else {
		fmt.Println("search 10 failed")
	}

	hashb.Insert(30)
	hashb.Insert(40)
	hashb.Print()

	ok = hashb.Search(30)
	if ok == true {
		fmt.Println("Search 30 worked")
	} else {
		fmt.Println("Search 30 failed")
	}

	ok = hashb.Search(40)
	if ok == true {
		fmt.Println("Search 40 worked")
	} else {
		fmt.Println("Search 40 failed")
	}

	ok = hashb.Search(5)
	if ok == false {
		fmt.Println("Search 5 worked")
	} else {
		fmt.Println("Search 5 failed")
	}

	ok = hashb.Delete(10)
	ok = hashb.Search(10)
	if ok == false {
		fmt.Println("Search 10 worked")
	} else {
		fmt.Println("Search 10 failed")
	}

	ok = hashb.Delete(20)
	ok = hashb.Search(20)
	if ok == false {
		fmt.Println("Search 20 worked")
	} else {
		fmt.Println("Search 20 failed")
	}

	ok = hashb.Delete(10)
	ok = hashb.Search(10)
	if ok == false {
		fmt.Println("Search 10 worked")
	} else {
		fmt.Println("Search 10 failed")
	}

}
