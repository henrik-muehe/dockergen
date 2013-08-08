package main

import "text/template"
import "os/exec"
import "strings"
import "regexp"
import "os"

var tpl string=`{{range .}}
server {
	listen 80;
	server_name {{.Name}}.muehe.org *.{{.Name}}.muehe.org;

	location / {
    	proxy_pass http://localhost:{{.Port}};
	}
}

{{end}}`



type DockerEntry struct {
	Port string
	Name string
}


func main() {
	entries := make([]*DockerEntry,0,0)

	cmd := exec.Command("docker","ps")
	outBytes,_:=cmd.Output()

	out:=strings.Trim(string(outBytes),"\n")
	lines := strings.Split(string(out),"\n")

	splitRegexp,_ := regexp.Compile("[\\s\\t]+")

	for _,line := range(lines[1:]) {
		cols := splitRegexp.Split(line,-1)

		cmdPort := exec.Command("docker","port",cols[0],"8080")
		portBytes,_ := cmdPort.Output()
		port := strings.Trim(string(portBytes),"\n")

		name := cols[1][strings.Index(cols[1],"/")+1:]
		name = name[0:strings.Index(name,":")]

		if len(port) > 0 {
			entries = append(entries,&DockerEntry{Port: port, Name: name})
		}
	}

	// Run through template
	tplInst,_ := template.New("tpl").Parse(tpl)
	tplInst.Execute(os.Stdout,entries)
}
