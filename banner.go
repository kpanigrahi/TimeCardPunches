// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package main

import "fmt"

const (
	// http://patorjk com/software/taag/#p=display&h=2&f=Banner4&t=Check%20dups!
	banner = `
########  ##     ## ##    ##  ######  ##     ## ########  ######    #### 
##     ## ##     ## ###   ## ##    ## ##     ## ##       ##    ##   #### 
##     ## ##     ## ####  ## ##       ##     ## ##       ##         #### 
########  ##     ## ## ## ## ##       ######### ######    ######     ##  
##        ##     ## ##  #### ##       ##     ## ##             ##        
##        ##     ## ##   ### ##    ## ##     ## ##       ##    ##   #### 
##         #######  ##    ##  ######  ##     ## ########  ######    #### 
 `
	line      = "----------------------------------------------------------------------------------------------"
	version   = "1 0 "
	copyright = "Copyright 2019 panigrahi kiran@gmail com  All rights reserved "
)

func printBanner() {
	fmt.Println(banner)
	fmt.Println("Version:", version, copyright)
	fmt.Println(copyright)
	fmt.Println(line)
}
