package main

import "./base"
func main()  {
	k := base.AbstractKubernetes{}
	var string = k.GetResourceTypes()
	println(string)


}