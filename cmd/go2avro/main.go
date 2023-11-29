// The go2avro command generates Avro schemas for Go types.
package main

import (
	"bytes"
	"encoding/json"
	stdflag "flag"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var flag = stdflag.NewFlagSet("", stdflag.ContinueOnError)

func main() {
	os.Exit(main1())
}

func main1() int {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `usage: go2avro [package.]type

This command prints the Avro schema for a given Go type on the
standard output.

If the package isn't specified, the current directory is used.
Current implementation restrictions mean that schemas can only
be generated for exported Go types.

For example:

	go2avro foo.com/bar/somepkg.Foo
`[1:])
	}
	if flag.Parse(os.Args[1:]) != nil {
		return 2
	}
	if len(flag.Args()) != 1 {
		flag.Usage()
		return 2
	}
	if err := main2(); err != nil {
		fmt.Fprintf(os.Stderr, "go2avro: %v\n", err)
		return 1
	}
	return 0
}

func main2() error {
	pkgType := flag.Arg(0)
	var p tmplParams
	if i := strings.LastIndex(pkgType, "."); i < 0 {
		var err error
		p.Package, err = currentPkg()
		if err != nil {
			return fmt.Errorf("cannot get current package: %v", err)
		}
		p.Type = pkgType
	} else {
		p.Package = pkgType[0:i]
		p.Type = pkgType[i+1:]
	}
	var codeBuf bytes.Buffer
	if err := tmpl.Execute(&codeBuf, p); err != nil {
		return fmt.Errorf("cannot execute template: %v", err)
	}
	code, err := format.Source(codeBuf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid temporary Go code:\n-------\n%s-----\n", codeBuf.String())
		return fmt.Errorf("invalid template code: %v", err)
	}
	// Build the binary before executing the code so we can
	// distinguish between build errors and Avro errors.
	exe, err := buildGo(code)
	if err != nil {
		return err
	}
	defer os.Remove(exe)

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	cmd := exec.Command(exe)
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		if errBuf.Len() > 0 {
			return fmt.Errorf("cannot get Avro type: %s", strings.TrimSpace(errBuf.String()))
		}
		return err
	}
	var indentJSON bytes.Buffer
	if err := json.Indent(&indentJSON, outBuf.Bytes(), "", "    "); err != nil {
		return fmt.Errorf("cannot indent JSON: %v", err)
	}
	fmt.Printf("%s", indentJSON.String())
	return nil
}

func buildGo(code []byte) (string, error) {
	// Create the Go file in the current directory so that we
	// take advantage of the current Go module.
	// TODO avoid the side-effect of adding the avro import, somehow.
	tmpFile, err := os.CreateTemp(".", "go2avro_temp_*.go")
	if err != nil {
		return "", fmt.Errorf("cannot generate temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write(code)
	tmpFile.Close()
	if err != nil {
		return "", fmt.Errorf("cannot write %q: %v", tmpFile.Name(), err)
	}
	tmpBinary, err := os.CreateTemp(".", "go2avro_temp_bin")
	if err != nil {
		return "", fmt.Errorf("cannot generate temp binary file: %v", err)
	}
	tmpBinary.Close()

	var errBuf bytes.Buffer
	cmd := exec.Command("go", "build", "-o", tmpBinary.Name(), tmpFile.Name())
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		defer os.Remove(tmpBinary.Name())
		return "", fmt.Errorf("cannot build: %v", errBuf.String())
	}
	// Use explicit "./" prefix because . isn't always in $PATH.
	return "." + string(filepath.Separator) + tmpBinary.Name(), nil
}

func currentPkg() (string, error) {
	var buf bytes.Buffer
	c := exec.Command("go", "list")
	c.Stderr = os.Stderr
	c.Stdout = &buf
	if err := c.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

type tmplParams struct {
	Package string
	Type    string
}

var tmpl = template.Must(template.New("").Parse(`
// Code generated by avrogen. DO NOT EDIT.

// This should be treated as a temporary file. Remove it if you find it.

// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/amedia/avro"

	pkg {{printf "%q" .Package}}
)

func main() {
	var x pkg.{{.Type}}
	t, err := avro.TypeOf(x)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get type: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(t)
}
`))
