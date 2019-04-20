//+build disabled

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ReadStringLine reads a set of characters before ENTER is hit from a standard input (terminal)
func ReadStringLine() (string, error) {
	stdinReader := bufio.NewReader(os.Stdin)

	if inputString, err := stdinReader.ReadString('\n'); err != nil {
		fmt.Println("Invalid input. Error: ", err)
		return "", err
	} else {
		// if exists, remove trailing Line Feed (added on Windows and on Unix/OSX)
		if strings.HasSuffix(inputString, "\n") {
			inputString = strings.TrimSuffix(inputString, "\n")
		}

		// if exist, remove trailing Carriage Return (added on Windows)
		if strings.HasSuffix(inputString, "\r") {
			inputString = strings.TrimSuffix(inputString, "\r")
		}

		inputString = strings.TrimSpace(inputString)

		// fmt.Println("inputString = ", inputString)
		return inputString, nil
	}
}

// ToIntegers takes a string, splits it in chunks separated by SPACE, converts each chunk into an integer and returns an
// array of integers.
func ToIntegers(str string) ([]int, error) {
	parts := strings.Split(str, " ")
	slice := make([]int, 0, len(parts))
	for _, v := range parts {
		if len(v) == 0 {
			continue
		}
		if n, err := strconv.Atoi(v); err != nil {
			fmt.Println("Error:", err)
			return nil, err
		} else {
			slice = append(slice, n)
		}
	}
	return slice, nil
}

// ReadIntegersLine reads characters before ENTER is hit from a standard input and returns an array of integers.
func ReadIntegersLine() ([]int, error) {
	if inputString, err := ReadStringLine(); err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	} else {
		// fmt.Println("inputString = ", inputString)
		return ToIntegers(inputString)
	}
}

func BubbleSort(slice []int, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				Swap(slice, i, j)
			}
		}
	}
}

func Swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func MergeTwoSorted(slice1 []int, slice2 []int, slice12 []int, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()
	i, j := 0, 0
	for k := 0; k < len(slice12); k++ {
		if i == len(slice1) || j == len(slice2) {
			if i == len(slice1) {
				slice12[k] = slice2[j]
				j++
			} else {
				slice12[k] = slice1[i]
				i++
			}
		} else if slice1[i] < slice2[j] {
			slice12[k] = slice1[i]
			i++
		} else {
			slice12[k] = slice2[j]
			j++
		}
	}
}

func main() {
	fmt.Println("Please enter a sequence of numbers to be sorted and then press ENTER:")
	if intSlice, err := ReadIntegersLine(); err != nil {
		fmt.Println("Error:", err)
	} else {
		elementsTotalCount := len(intSlice)
		if elementsTotalCount < 4 {
			BubbleSort(intSlice, nil)
			fmt.Println("All numbers sorted:", intSlice)
		} else {
			elementsChunkCount := elementsTotalCount / 4
			slice1 := intSlice[0:elementsChunkCount]
			slice2 := intSlice[elementsChunkCount : 2*elementsChunkCount]
			slice3 := intSlice[2*elementsChunkCount : 3*elementsChunkCount]
			slice4 := intSlice[3*elementsChunkCount : elementsTotalCount]
			var wg sync.WaitGroup
			wg.Add(4)
			go BubbleSort(slice1, &wg)
			go BubbleSort(slice2, &wg)
			go BubbleSort(slice3, &wg)
			go BubbleSort(slice4, &wg)
			wg.Wait()
			wg.Add(2)
			slice12 := make([]int, len(slice1)+len(slice2))
			go MergeTwoSorted(slice1, slice2, slice12, &wg)

			slice34 := make([]int, len(slice3)+len(slice4))
			go MergeTwoSorted(slice3, slice4, slice34, &wg)
			wg.Wait()

			slice1234 := make([]int, len(slice12)+len(slice34))
			MergeTwoSorted(slice12, slice34, slice1234, nil)

			fmt.Println("All numbers sorted:", slice1234)
		}
	}
}
