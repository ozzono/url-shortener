# URL shortened code challenge

## Routes

- api routes:
  - PUT: path - `localhost:8000/api`
    - awaits source url as json data;
    - returns a url data as exampled below:

        ```json
        {
            "id":        "62bf0ff21d37e4cfa5608b0a",
            "source":    "https://go.dev",
            "shortened": "hkujzph",
            "count:      0,
        }
        ```
  - Get: path - `localhost:8000/api/:id`
    - awaits url id in path;
    - returns a url data as exampled below:

        ```json
        {
            "id":        "62bf0ff21d37e4cfa5608b0a",
            "source":    "https://go.dev",
            "shortened": "hkujzph",
            "count:      0,
        }
        ```
  - Del: path - `localhost:8000/api/:id`
    - awaits url id in path;
    - returns only request status code
- Redirect: path - `localhost:8000/<shortpath>`
  - redirects to stored shortened url; in this example:
    - shortpath: hkujzph
      - `localhost:8000/hkujzph`
    - redirects to <https://go.dev>
