// this file was automatically generated using struct2ts -H -i -D -o frontend/src/logic/types.ts controller.AgendaUtilisateur views.OutIngredient views.OutIngredients views.OutRecette views.OutRecettes views.OutMenu views.OutMenus views.OutSejour views.OutAgenda views.OutUtilisateurs views.InResoudIngredients views.OutResoudIngredients controller.OutLoggin views.InLoggin
// +build ignore

package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/views"
	"github.com/benoitkugler/struct2ts"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var (
		out = flag.String("o", "-", "output")
		f   = os.Stdout
		err error
	)

	flag.Parse()
	if *out != "-" {
		if f, err = os.OpenFile(*out, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644); err != nil {
			panic(err)
		}
		defer f.Close()
	}
	if err = runStruct2TS(f); err != nil {
		panic(err)
	}
}

func runStruct2TS(w io.Writer) error {
	s := struct2ts.New(&struct2ts.Options{
		Indent: "	",

		NoAssignDefaults: false,
		InterfaceOnly:    true,

		NoConstructor: false,
		NoCapitalize:  false,
		MarkOptional:  false,
		NoToObject:    false,
		NoExports:     false,
		NoHelpers:     true,
		NoDate:        true,

		ES6: false,
	})

	s.Add(controller.AgendaUtilisateur{})
	s.Add(views.OutIngredient{})
	s.Add(views.OutIngredients{})
	s.Add(views.OutRecette{})
	s.Add(views.OutRecettes{})
	s.Add(views.OutMenu{})
	s.Add(views.OutMenus{})
	s.Add(views.OutSejour{})
	s.Add(views.OutAgenda{})
	s.Add(views.OutUtilisateurs{})
	s.Add(views.InResoudIngredients{})
	s.Add(views.OutResoudIngredients{})
	s.Add(controller.OutLoggin{})
	s.Add(views.InLoggin{})

	io.WriteString(w, "// this file was automatically generated, DO NOT EDIT\n")
	return s.RenderTo(w)
}
