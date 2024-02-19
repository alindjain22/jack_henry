# jack_henry
Programming assignment for Jack Henry

Steps:
1. Add your API key in line 26 of weatherService.go
2. Run the weatherService.go file from command line like this - go run weatherService.go
3. Check the stdout to see the port number the server is running on. The port will be 8080 unless specified otherwise in the stdout logs.
4. Open browser and type URL in the following format - `http://localhost:<port number>/weather?lat=<user latitude>&lon=<user longitude>`
5. Replace `<port number>` with the port number on which the HTTP server is running (e.g., 8080). Replace `<user latitude>` and `<user longitude>` with the desired geographical coordinates for which you want to fetch the weather data
6. Example URL to type in browser - http://localhost:8080/weather?lat=22.71&lon=75.85
7. Weather information will be displayed in the browser.
