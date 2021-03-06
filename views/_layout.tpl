<!DOCTYPE html>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Last Chance Banking</title>
    <script src="//ajax.googleapis.com/ajax/libs/dojo/1.10.4/dojox/mobile/deviceTheme.js"></script>
    <link rel="stylesheet" href="/public/css/app.css">
    <script>
      var dojoConfig = {
        async: true,
        isDebug: true,
        parseOnLoad: false,
        packages: [{
          name: "app",
          location: "/public/js/app"
        }]
      };
    </script>
    <script src="//ajax.googleapis.com/ajax/libs/dojo/1.10.4/dojo/dojo.js"></script>
  </head>
  <body>
    {{block "content" . }}{{end}}
  </body>
</html>
