# Go RESTful API Design (Go Books)

## Starting the application

```bash
sudo docker compose up --build -d
```

## Testing the application

1. Install [Postman](https://www.postman.com/).
2. Import the Postman Collection `Go Books.postman_collection.json`.
3. Execute the HTTP request `Book Search` and play with the various pre-defined URL Query Parameters for filtering and sorting (see "Params" tab.), or add your own.

## Examples

Request:
```bash
GET http://localhost:8080/books?language=English&genre=self-help&genre=business&sort_by=published_at.desc
```

Response:
```json
{
    "Books": [
        {
            "BookID": 9,
            "Title": "Bad Blood: Secrets and Lies in a Silicon Valley Startup",
            "Language": "English",
            "ISBN": "9781524731656",
            "Pages": 352,
            "Genres": [
                "non-fiction",
                "business",
                "crime"
            ],
            "PublishedAt": "2018-05-21T00:00:00Z",
            "Authors": [
                {
                    "AuthorID": 2,
                    "Name": "John Carreyrou"
                }
            ],
            "Publisher": {
                "PublisherID": 2,
                "Name": "Alfred A. Knopf"
            }
        },
        {
            "BookID": 11,
            "Title": "The Subtle Art of Not Giving a Fuck",
            "Language": "English",
            "ISBN": "9780062457714",
            "Pages": 224,
            "Genres": [
                "fiction",
                "business",
                "self-help",
                "inspirational",
                "philosophy",
                "adult"
            ],
            "PublishedAt": "2016-09-13T00:00:00Z",
            "Authors": [
                {
                    "AuthorID": 4,
                    "Name": "Mark Manson"
                }
            ],
            "Publisher": {
                "PublisherID": 3,
                "Name": "HarperOne"
            }
        },
        {
            "BookID": 10,
            "Title": "The Alchemist",
            "Language": "English",
            "ISBN": "9780062315007",
            "Pages": 175,
            "Genres": [
                "fiction",
                "self-help",
                "inspirational"
            ],
            "PublishedAt": "1988-01-21T00:00:00Z",
            "Authors": [
                {
                    "AuthorID": 3,
                    "Name": "Paulo Coelho"
                }
            ],
            "Publisher": {
                "PublisherID": 3,
                "Name": "HarperOne"
            }
        }
    ]
}
```
