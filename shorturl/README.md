Exercise based on https://courses.calhoun.io/lessons/les_goph_04.

Provided short url will be redirected to destin page, much like URL shortener would.
All configuration is stored in yml or json file. For example:

yaml file
``` yaml
- path: /urlshort
  url: https://github.com/gophercises/urlshort
```
json

```json
[
  {
    "path": "/urlshort",
    "url": "https://github.com/gophercises/urlshort"
  }
]
```


To run use go run main.go.