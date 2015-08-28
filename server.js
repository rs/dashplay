var http = require('http');
var fs = require('fs');

var server = http.createServer(function(req, res) {
    fs.readFile(__dirname + '/index.html', function(err, data) {
        res.writeHead(200);
        res.end(data);
    });
});
server.listen(8080);
