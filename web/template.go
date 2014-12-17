package web

const templates = `
{{define "header"}}
<!DOCTYPE HTML>
<html>
<head>
    <meta charset="utf-8">
    <link rel="stylesheet" href="/static/css/main.css" type="text/css">
    <title>{{.PageTitle}}</title>
    <style type="text/css">
body, html  { margin: 5px; }
textarea    { width: 100%; padding: 10px; box-sizing: border-box; resize: vertical; }
label       { margin-right: 30px; }
.pull-right { float: right; }
.center     { text-align: center; margin: auto; }
    </style>
</head>
<body>
{{end}}



{{define "footer"}}
</body>
</html>
{{end}}


{{define "main-page"}}
{{template "header" .}}
<form action="." method="POST">
    <textarea name="content" id="content" placeholder="Paste your content here" required></textarea>

    <label for="host-required">
        <input id="host-required" type="checkbox" name="host-required"> Host required
    </label>
    <label for="one-use-only">
        <input id="one-use-only" type="checkbox" name="one-use-only"> One use only
    </label>
    <label for="expire-after">
        Expire after
        <select id="expire-after" name="expire-after">
            <option value="0">Never</option>
            <option value="300" selected>5 minutes</option>
            <option value="1800">30 minutes</option>
            <option value="3600">1 hour</option>
            <option value="21600">6 hours</option>
            <option value="86400">1 day</option>
        <select>
    </label>
    <button type="submit" class="pull-right">Create</button>
</form>
<script type="text/javascript">
(function () {
  window.addEventListener('load', function () {
      var t = document.getElementById('content');
      var resizeContentInput = function () {
        t.style.height = window.innerHeight - 60 + 'px';
      };

      window.addEventListener('resize', resizeContentInput);
      resizeContentInput();
  }, false);
}());
</script>
{{template "footer" .}}
{{end}}


{{define "paste-host"}}
{{template "header" .}}
<h1>
	You are hosting paste
	<small>
		<a href="/{{.PasteKey}}">{{.PasteKey}}</a>
	</small>
</h1>
<div>
	Closing browser tab will delete the paste.
</div>
<h3>Connected clients</h3>
<ul>
{{end}}


{{define "paste-client"}}
<li>
	{{.Client.Address}} - {{.Client.UserAgent}}
</li>
{{end}}


{{define "paste-client-end"}}
</ul>
<div>
    No more clients can connect.
</div>
{{template "footer" .}}
{{end}}


{{define "paste-one-use-created"}}
{{template "header" .}}
    <h1>Paste created</h1>
	<h2 class="center">
		<a href="/{{.PasteKey}}">{{.PasteUrl}}</a>
	</h2>
	<p>
		Once the following link will be accessed, paste will be deleted and no longer available.
	</p>
{{template "footer" .}}
{{end}}
`
