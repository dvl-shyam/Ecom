Task: Create a web server that leverages templating to dynamically generate HTML content.
Utilize a templating engine like html/template to define reusable templates with placeholders for dynamic data.
Design routes to handle different HTTP requests and serve appropriate templates with data from your program or a database.
Implement basic functionalities like displaying a welcome message, user login form, or product listing page.

Challenge:

Integrate a database (e.g., MongoDB) to store and retrieve data for displaying on web pages.
Explore advanced templating features like loops, conditionals, and layouts for a more structured and dynamic user interface.


HomePage Displaying Welcome Message
```bash
curl --location --request GET 'http://localhost:{PORT}/
```

Get Poduct Page
```bash
curl --location --request GET 'http://localhost:{PORT}/products
```

Insert New Poduct
```bash
curl --location --request POST 'http://localhost:{PORT}/products
```

Go To Login Page
```bash
curl --location --request GET 'http://localhost:{PORT}/login
```