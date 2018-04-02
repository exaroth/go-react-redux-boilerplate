<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>{{ .appName }}</title>
    </head>
    <body>
        <div id="app"></div>
        {{ .appName }}
        {{ .developmentEnv }}
        {{ if not .developmentEnv }}
            <script src="/static/build/js/babel-polyfill.min.js"></script>
        {{ end }}

        <script src="/static/build/js/app.min.js"></script>
    </body>
</html>
