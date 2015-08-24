package main

import (
	"html/template"
)

var (
	mainTpl *template.Template
)

const mainTplSrc = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Image to CSS</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
	<style>
	  	.img-container {
	  		width: {{.Width}}px;
	  		height: {{.Height}}px;
	  	}
	  	.img-container div {
	  		width: 1px;
	  		height: 1px;
	  		float: left;
	  	}
	</style>
</head>
<body>
	<div class="container">
	
	<div class="row">
		<div class="col-md-12">
			<h1>Image to CSS <small>Convert Pixels to divs</small></h1>
			<form action="/" method="POST" enctype="multipart/form-data" class="form-inline">
				<div class="form-group">
					<input type="file" name="image" accept="image/*" />
				</div>
				<button type="submit" class="btn btn-primary">Go!</button>
			</form>
		</div>
	</div>

	{{if .Pixels}}
	<div class="row">
		<div class="col-md-12">
			<h1>{{.PixelCount}} divs</h1>
			<div class="img-container">
	 		{{range .Pixels}}
	 			{{range .}}
	 				<div style="background: rgba({{.R}}, {{.G}}, {{.B}}, {{.A}});"></div>
	 			{{end}}
	 		{{end}}
	 		</div>
	 	</div>
 	</div>
 	{{end}}

 	</div>
 </body>
</html>
`

func init() {
	var err error
	mainTpl, err = template.New("gen").Parse(mainTplSrc)
	if err != nil {
		panic(err)
	}
}

type mainTplData struct {
	Width, Height int
	Pixels        [][]pixel
	PixelCount    int
}
