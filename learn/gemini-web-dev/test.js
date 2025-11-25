fetch('http://localhost:8080/api/search', {
	method:  'POST',
	headers:  {
		'Content-Type':  'application/json',
	},
	body:  JSON.stringify({
		question:  "what is the meaning of 42 (from a browser)?",
		limit: 3
	}),
})
.then(response => response.json())
.then(data => console.log(data))
	.catch(error => console.error('Error:', error));
