package main

import "fmt"

// driver function
func main(){
    arr := [4]int{2,7,11,15}
    target :=9
    fmt.Println(findIndex(arr[0:], target))

}

func findIndex(arr []int, target int) (int, int) {
    // creating map
    indexMap := make(map[int]int)

    // iterating over array
    for i := 0; i < len(arr); i++ {
        rem := target - arr[i]
        // checking if the value in present in the map or not
        if val, ok := indexMap[rem]; ok {
            return val, i 
        }
        indexMap[arr[i]] = i  // storing current value
    }
    return -1,-1 // or any sentinel value indicating "not found"
}
