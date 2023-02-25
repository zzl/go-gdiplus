package gen

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Param struct {
	Name  string
	Type  string
	Stars int
}

type Fn struct {
	Cat        string
	Name       string
	ReturnType string
	Params     []Param
}

func ParseFn(sig string) Fn {
	par1Pos := strings.IndexByte(sig, '(')
	par2Pos := strings.IndexByte(sig, ')')
	sParams := sig[par1Pos+1 : par2Pos]

	paramParts := strings.Split(sParams, ",")
	var params []Param
	for n, paramPart := range paramParts {
		paramPart = strings.TrimSpace(paramPart)
		if paramPart == "" {
			continue
		}
		pos := strings.LastIndexByte(paramPart, ' ')
		if pos == -1 {
			pos = len(paramPart)
			paramPart = paramPart + " p" + strconv.Itoa(n)
		}
		name := paramPart[pos+1:]
		sType := strings.TrimSpace(paramPart[:pos])
		stars := 0
		//for strings.HasPrefix(name, "*") {
		//	name = name[1:]
		//	sType += "*"
		//}
		name1 := strings.ReplaceAll(name, "*", "")
		name1 = strings.ReplaceAll(name1, "&", "")
		stars += len(name) - len(name1)
		name = strings.TrimSpace(name1)

		sType1 := strings.ReplaceAll(sType, "*", "")
		sType1 = strings.ReplaceAll(sType1, "&", "")
		stars += len(sType) - len(sType1)
		sType = strings.TrimSpace(sType1)

		typeParts := strings.Split(sType, " ")
		for n, part := range typeParts {
			uPart := strings.ToUpper(part)
			if uPart != "GDIPCONST" && uPart != "CONST" &&
				uPart != "OUT" && uPart != "IN" {
				typeParts = typeParts[n:]
				break
			}
		}

		if name == "type" {
			name = "typ"
		} else if name == "map" {
			name = "map" + strconv.Itoa(n)
		}

		param := Param{
			Name:  name,
			Type:  strings.Join(typeParts, ""),
			Stars: stars,
		}
		//if param.Name == "pData16" {
		//	println("?")
		//}
		params = append(params, param)
	}

	sLeftParts := strings.TrimSpace(sig[:par1Pos])
	leftParts := strings.Split(sLeftParts, " ")
	fn := Fn{
		ReturnType: leftParts[0],
		Name:       leftParts[len(leftParts)-1],
		Params:     params,
	}
	if fn.ReturnType == "Status" {
		fn.ReturnType = "GpStatus"
	}
	return fn
}

func Main() {

	initTypeMappings()

	catFileMap := map[string]string{
		"adjustablearrowcap":  "pen",
		"cachedbitmap":        "bitmap",
		"customlinecap":       "pen",
		"fontfamily":          "font",
		"graphicspath":        "path",
		"hatchbrush":          "brush",
		"imageattributes":     "image",
		"lineargradientbrush": "brush",
		"memory":              "misc",
		"notification":        "misc",
		"pathgradientbrush":   "brush",
		"pathiterator":        "path",
		"solidbrush":          "brush",
		"stringformat":        "text",
		"texturebrush":        "brush",
	}

	fileFns := make(map[string][]Fn)

	bts, _ := ioutil.ReadFile("doc/gdip.txt")
	lines := strings.Split(string(bts), "\r\n")
	var cat string
	var file string
	typeSet := make(map[string]bool)
	for _, line := range lines {
		//println(line)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "/*@") {
			cat = line[3 : len(line)-2]
			println("@", cat)
			cg, ok := catFileMap[cat]
			if ok {
				file = cg
			} else {
				file = cat
			}
			continue
		}
		fn := ParseFn(line)
		for _, p := range fn.Params {
			typeSet[p.Type] = true
		}
		fn.Cat = cat
		print(fn.ReturnType, " ", fn.Name, "(")
		for n, p := range fn.Params {
			if n != 0 {
				print(", ")
			}
			print(p.Type)
			print(" ")
			for m := 0; m < p.Stars; m++ {
				print("*")
			}
			print(p.Name)
		}
		println(")")
		fns, ok := fileFns[file]
		if !ok {
			fns = make([]Fn, 0)
		}
		fns = append(fns, fn)
		fileFns[file] = fns
	}

	//println("-----------------")
	//for typ, _ := range typeSet {
	//	println(typ)
	//}
	//println("-----------------")

	for file, fns := range fileFns {
		writeFile(file, fns)
	}

}

func writeFile(name string, fns []Fn) {
	var text string
	text += `package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (
`
	var cat string
	cat = ""
	for _, fn := range fns {
		if cat != fn.Cat {
			cat = fn.Cat
			text += "\n\t// " + cat + "\n"
		}
		name := strings.ToLower(fn.Name[:1]) + fn.Name[1:]
		text += "\t" + name + "\t*windows.LazyProc\n"
	}
	text += ")\n"

	text += "\nfunc init() {\n"
	cat = ""
	for _, fn := range fns {
		if cat != fn.Cat {
			cat = fn.Cat
			text += "\n\t// " + cat + "\n"
		}
		name := strings.ToLower(fn.Name[:1]) + fn.Name[1:]
		text += "\t" + name + " = dll.NewProc(\"" + fn.Name + "\")\n"
	}
	text += "}\n"

	cat = ""
	for _, fn := range fns {
		if cat != fn.Cat {
			cat = fn.Cat
			text += "\n\t// " + cat + "\n"
		}
		text += "func " + fn.Name[4:] + "("
		for n, p := range fn.Params {
			if n != 0 {
				text += ", "
			}
			name := p.Name
			text += name + " "
			for n := 0; n < p.Stars; n++ {
				text += "*"
			}
			goType := mapGoType(p.Type)
			//p.Stars += len(strings.ReplaceAll(goType, "*", "")) - len(goType)
			text += goType
		}
		returnType := fn.ReturnType
		returnType = mapGoType(returnType)

		text += ") " + returnType + " "
		text += "{\n"
		//"\treturn Ok\n" +
		pCount := len(fn.Params)
		pCountCall := (pCount + 2) / 3 * 3
		_ = pCountCall
		text += "\tret, _, _ := syscall.SyscallN"
		//if pCountCall > 3 {
		//	text += strconv.Itoa(pCountCall)
		//}
		procName := strings.ToLower(fn.Name[:1]) + fn.Name[1:]
		//text += "(" + procName + ".Addr(), " + strconv.Itoa(pCount) + ",\n"
		text += "(" + procName + ".Addr(), \n"
		for n, p := range fn.Params {
			tm, ok := typeMappings[p.Type]
			var callConv string
			if !ok {
				//?
				callConv = p.Name
			} else {
				callConv = tm.CallConv
				if callConv == "" {
					callConv = p.Name
				} else {
					callConv = strings.ReplaceAll(callConv, "@", p.Name)
				}
			}

			if p.Stars > 0 || strings.ContainsRune(tm.GoType, '*') {
				callConv = "unsafe.Pointer(" + p.Name + ")"
			}

			text += "\t\tuintptr(" + callConv + ")"
			_ = n
			if n < pCount-1 {
				text += ",\n"
			} else {
				text += ")\n"
			}
		}
		//for n := pCount; n < pCountCall; n++ {
		//	text += "\t\t0"
		//	if n < pCountCall-1 {
		//		text += ",\n"
		//	} else {
		//		text += ")\n"
		//	}
		//}
		//text += "\t)\n"
		text += "\treturn " + returnType + "(ret)\n"
		text += "}\n\n"
		//name := strings.ToLower(fn.Name[:1]) + fn.Name[1:]
	}

	ioutil.WriteFile("gdip/"+name+".go", []byte(text), 0666)
}

func mapGoType(cType string) string {
	tm, ok := typeMappings[cType]
	if !ok {
		return cType
	}
	if tm.GdipType { //?
		return cType
	}
	return tm.GoType
}
