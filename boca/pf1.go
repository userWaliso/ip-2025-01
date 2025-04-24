package main
import (
	f"fmt"
)
func main() {
	var ano int
	var conceito string
	f.Scan(&ano)

	if ano == 2004 {
		conceito = "Macaco"
	} else if ano == 2005 {
		conceito = "Galo"
	} else if ano == 2006 {
		conceito = "Cao"
	} else if ano == 2007 {
		conceito = "Porco"
	} else if ano == 2008 {
		conceito = "Rato"
	} else if ano == 2009 {
		conceito = "Boi"
	} else if ano == 2010 {
		conceito = "Tigre"
	} else if ano == 2011 {
		conceito = "Coelho"
	} else if ano == 2012 {
		conceito = "Dragao"
	} else if ano == 2013 {
		conceito = "Serpente"
	} else if ano == 2014 {
		conceito = "Cavalo"
	} else if ano == 2015 {
		conceito = "Cabra"
	}
	f.Println(conceito)
}