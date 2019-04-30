// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to takes slices of slices to create different
// views of and make changes to the underlying array.
package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1)

	// Take a slice of slice1. We want just indexes 2 and 3.
	// Parameters are [starting_index : (starting_index + length)]
	//  空きキャパシティは引き継がれる
	slice2 := slice1[2:4] // from 2 - (4-1) == from 2 - 3
	inspectSlice(slice2)

	// バグを呼び出す
	//slice2 = append(slice2, "CHERRY")
	//slice2 = append(slice2, "CHERRY2")

	fmt.Println("*************************")
	fmt.Println("*CHANGED*")

	// Change the value of the index 0 of slice2.
	slice2[0] = "CHANGED"
	// -> make a new slice

	// Display the change across all existing slices.
	inspectSlice(slice1)
	inspectSlice(slice2)

	fmt.Println("*************************")
	fmt.Println("*COPY*")
	// Make a new slice big enough to hold elements of slice 1 and copy the
	// values over using the builtin copy function.
	slice3 := make([]string, len(slice1))
	copy(slice3, slice1) // copy(dst, src) deep copy
	inspectSlice(slice3)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		//fmt.Printf("[%d] %p %s\n",
		//	i,
		//	&slice[i],
		//	s)
		fmt.Printf("[%d] s lives at %p /   s[i] lives at %p %s\n",
			i,
			&s,
			&slice[i],
			s)
	}
}

// Length[5] Capacity[8]
//[0] s lives at 0xc42000e1e0 /   s[i] lives at 0xc420098000 Apple
//[1] s lives at 0xc42000e1e0 /   s[i] lives at 0xc420098010 Orange
//[2] s lives at 0xc42000e1e0 /   s[i] lives at 0xc420098020 Banana
//[3] s lives at 0xc42000e1e0 /   s[i] lives at 0xc420098030 Grape
//[4] s lives at 0xc42000e1e0 /   s[i] lives at 0xc420098040 Plum
//Length[2] Capacity[6]
//[0] s lives at 0xc42000e240 /   s[i] lives at 0xc420098020 Banana
//[1] s lives at 0xc42000e240 /   s[i] lives at 0xc420098030 Grape
//*************************
//*CHANGED*
//Length[5] Capacity[8]
//[0] s lives at 0xc42000e270 /   s[i] lives at 0xc420098000 Apple
//[1] s lives at 0xc42000e270 /   s[i] lives at 0xc420098010 Orange
//[2] s lives at 0xc42000e270 /   s[i] lives at 0xc420098020 CHANGED
//[3] s lives at 0xc42000e270 /   s[i] lives at 0xc420098030 Grape
//[4] s lives at 0xc42000e270 /   s[i] lives at 0xc420098040 Plum
//Length[2] Capacity[6]
//[0] s lives at 0xc42000e2d0 /   s[i] lives at 0xc420098020 CHANGED
//[1] s lives at 0xc42000e2d0 /   s[i] lives at 0xc420098030 Grape
//*************************
//Length[5] Capacity[5]
//[0] s lives at 0xc42000e300 /   s[i] lives at 0xc4200940f0 Apple
//[1] s lives at 0xc42000e300 /   s[i] lives at 0xc420094100 Orange
//[2] s lives at 0xc42000e300 /   s[i] lives at 0xc420094110 CHANGED
//[3] s lives at 0xc42000e300 /   s[i] lives at 0xc420094120 Grape
//[4] s lives at 0xc42000e300 /   s[i] lives at 0xc420094130 Plum
