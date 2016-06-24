package template

import (
	"html/template"
)

var ROOT = template.Must(template.New("homepage").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{ .title }}</title>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/custom.css" rel="stylesheet">
  </head>
  <body>

    <div class="container">

      <!-- Static navbar -->
      <nav class="navbar navbar-default">
        <div class="container-fluid">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">{{ .app }}</a>
          </div>
          <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
              <li class="active"><a href="#">Home</a></li>
              <li><a href="#">About</a></li>
              <li><a href="#">Contact</a></li>
              <li class="dropdown">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
                <ul class="dropdown-menu">
                  <li><a href="#">Action</a></li>
                  <li><a href="#">Another action</a></li>
                  <li><a href="#">Something else here</a></li>
                  <li role="separator" class="divider"></li>
                  <li class="dropdown-header">Nav header</li>
                  <li><a href="#">Separated link</a></li>
                  <li><a href="#">One more separated link</a></li>
                </ul>
              </li>
            </ul>

            <!--ul class="nav navbar-nav navbar-right">
              <li class="active"><a href="./">Default <span class="sr-only">(current)</span></a></li>
              <li><a href="../navbar-static-top/">Static top</a></li>
              <li><a href="../navbar-fixed-top/">Fixed top</a></li>
            </ul-->
          </div><!--/.nav-collapse -->
        </div><!--/.container-fluid -->
      </nav>

      <!-- Main component for a primary marketing message or call to action -->
      <div class="jumbotron">
        <h1>{{ .title }}</h1>
        <p>This example is a quick exercise to illustrate how the default, static navbar and fixed to top navbar work.</p>
        <p>
          <a class="btn btn-lg btn-primary" href="../../components/#navbar" role="button">View navbar docs &raquo;</a>
        </p>
      </div>

    </div> <!-- /container -->

    
    <div class="container"> <!-- content -->

      <table class="table">
        <tr>
          <th>Firstname</th>
          <th>Lastname</th>
          <th>Points</th>
        </tr>
        <tr>
          <td>Eve</td>
          <td>Jackson</td>
          <td>94</td>
        </tr>
        {{ range .names }}
        <tr>
          <td><a href="/{{.}}">{{.}}</a></td>
          <td>other</td>
          <td>cool</td>
        </tr>
        {{ end }}
      </table>

    </div> <!-- /content -->

    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
  </body>
</html>
`))
