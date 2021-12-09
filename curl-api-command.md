### **POST a JSON FILE**
```bash
curl -X POST -H 'Content-Type: application/json' -d @book.json http://127.0.0.1:3888/books
```
### **POST a JSON DATA**
```bash
curl -X POST -H 'Content-Type: application/json' -d '{"title":"Harry potter","sub_title":"Harry potter is a lewd movie and story","price":98000}' http://127.0.0.1:3888/books
```
